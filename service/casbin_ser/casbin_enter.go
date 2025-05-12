package casbin_ser

import (
	"gin_web_frame/global"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"sync"
)

type CasbinService struct {
	Enforcer *casbin.Enforcer
	Adapte   *gormadapter.Adapter
}

// 角色组和用户增删改查
type User struct {
	Username string
	Roles    []string
}

// 获取所有用户以及关联的角色
func (c *CasbinService) GetUsers() (users []User, err error) {
	policy, err := c.Enforcer.GetGroupingPolicy()
	if err != nil {
		return nil, err
	}
	userMap := make(map[string]*User, 0)

	for _, _p := range policy {
		username, userGroup := _p[0], _p[1]
		if u, ok := userMap[username]; ok {
			u.Roles = append(u.Roles, userGroup)
		} else {
			userMap[username] = &User{
				Username: username,
				Roles:    []string{userGroup},
			}
		}
	}
	for _, user := range userMap {
		users = append(users, *user)
	}

	return
}

// 获取所有角色组
func (c *CasbinService) GetRoles() (roles []string, err error) {

	return c.Enforcer.GetAllRoles()
}

// 角色组添加用户，没有组默认添加
func (c *CasbinService) AddUserInRoles(username, role string) (err error) {
	_, err = c.Enforcer.AddRoleForUser(username, role)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy() // 保存策略 持久化到数据库

}

// 角色组删除用户
func (c *CasbinService) DeleteUserInRoles(username, role string) (err error) {
	_, err = c.Enforcer.DeleteRoleForUser(username, role)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy() // 保存策略 持久化到数据库

}

// 角色组权限的增删改查
type rolePolicy struct {
	RoleName string `gorm:"column:v0" json:"roleName"`
	Path     string `gorm:"column:v1" json:"path" `
	Method   string `gorm:"column:v2" json:"method"`
}

// 获取所有角色组的权限
func (c *CasbinService) GetRolesPolicy() (roles []rolePolicy, err error) {
	err = c.Adapte.GetDb().Model(&gormadapter.CasbinRule{}).Where("ptype = 'p'").Find(&roles).Error
	if err != nil {
		return nil, err
	}
	return
}

// 创建角色组权限, 已有的会忽略
func (c *CasbinService) CreateRolePolicy(role string, path string, method string) (err error) {

	err = c.Enforcer.LoadPolicy()
	if err != nil {
		return err
	}

	_, err = c.Enforcer.AddPolicy(role, path, method)
	if err != nil {
		return err
	}

	return c.Enforcer.SavePolicy() // 保存策略 持久化到数据库

}

// 修改角色组权限
func (c *CasbinService) UpdateRolePolicy(old, new rolePolicy) (err error) {
	_, err = c.Enforcer.UpdatePolicy([]string{old.RoleName, old.Path, old.Method}, []string{new.RoleName, new.Path, new.Method})
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy() // 保存策略 持久化到数据库

}

// 删除角色组权限
func (c *CasbinService) DeleteRolePolicy(rolePolicy rolePolicy) (err error) {
	_, err = c.Enforcer.RemovePolicy(rolePolicy.RoleName, rolePolicy.Path, rolePolicy.Method)
	if err != nil {
		return err
	}
	return c.Enforcer.SavePolicy() // 保存策略 持久化到数据库
}

// 验证用户权限
func (c *CasbinService) CanAccess(roleName, path, method string) (ok bool, err error) {
	return c.Enforcer.Enforce(roleName, path, method)
}

var (
	CasbinSrv *CasbinService
	once      sync.Once
)

func (c *CasbinService) CasbinInit() (casbinService *CasbinService) {
	once.Do(func() {
		adapter, err := gormadapter.NewAdapterByDB(global.DB)
		if err != nil {
			panic(err)
		}
		CasbinSrv.Adapte = adapter

		// 加载模型
		modelStr := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(modelStr)
		if err != nil {
			zap.L().Error("字符串加载模型失败:", zap.Error(err))
		}
		e, err := casbin.NewEnforcer(m, adapter)
		CasbinSrv.Enforcer = e

	})
	return CasbinSrv
}
