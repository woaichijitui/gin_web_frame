package common

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// 自定义泛型
type intType interface {
	int | int8 | int16 | int32 | int64
}

// CheckFileSizeOutOfLimit
//
//	Description: 校验文件大小是否超出配置限制大小
//	param fileSize 单位b
//	param LimitSize 单位MB
//	return error
func CheckFileSizeOutOfLimit[T intType](fileSize int64, limitSize T) error {
	sizeMB := float64(fileSize) / (float64(1024 * 1024)) //将字节转换为MB
	//将泛型全部转换成float64
	lsize := float64(limitSize)
	if sizeMB >= lsize {
		return errors.New(fmt.Sprintf("校验文件大小超出配置限制大小，文件大小为： %.2f MB,系统文件上传阈值为： %d MB", sizeMB, limitSize))
	}
	return nil
}

// CheckFileSuffixIsRight
//
//	Description: 校验文件后缀是否符合配置
//	param fileName
//	return suffix
//	return err
/*func CheckFileSuffixIsRight(fileName string) (suffix string, err error) {
	filenameSplitList := strings.Split(fileName, ".")
	suffix = filenameSplitList[len(filenameSplitList)-1]

	if exit := InList(global.CONFIG.Upload.Suffix, suffix); exit {
		return suffix, nil
	}
	return "", errors.New(fmt.Sprintf("不允许上传的%s文件", suffix))
}*/

// GetValidMsg
//
//	Description: 获取结构体中msg标签信息
//	param err
//	param obj
//	return string
func GetValidMsg(err error, obj any) string {

	typeObj := reflect.TypeOf(obj)

	//将err接口断言成具体类型
	if errors, ok := err.(validator.ValidationErrors); ok {
		//断言成功
		for _, e := range errors {
			//循环每个错误信息
			//根据报错字段名，获取结构体的具体字段
			//	e.Field() 返回字段名字 json标签的名字优先
			if filedObj, exist := typeObj.Elem().FieldByName(e.Field()); exist {
				//若有tag标签 获取tag标签里的值
				msg := filedObj.Tag.Get("msg")
				return msg
			}
		}
	}

	return err.Error()
}

// 自定义验证器函数
func RequiredBool(fl validator.FieldLevel) bool {
	field := fl.Field()
	if field.Kind() == reflect.Bool {
		return true
	}
	return false
}
