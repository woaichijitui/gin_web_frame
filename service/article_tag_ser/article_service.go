package article_tag_ser

import (
	"gin_web_frame/global"
	models "gin_web_frame/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ArticleTagService struct {
}

func (a *ArticleTagService) DeleteArticlesWithTag(articleIds []uint) error {
	if len(articleIds) == 0 {
		return nil
	}

	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 批量删除中间表关联
		if err := tx.Exec(
			"DELETE FROM article_tag_relations WHERE article_id IN (?)",
			articleIds,
		).Error; err != nil {
			return err
		}

		// 批量删除文章（物理删除）
		//return tx.Unscoped().Where("id IN (?)", articleIds).Delete(&models.Article{}).Error

		// 如果使用软删除：
		return tx.Where("id IN (?)", articleIds).Delete(&models.Article{}).Error
	})
}

func (a *ArticleTagService) ArticleCreateAndAppendTags(article *models.Article, tags []models.Tag) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		var tagsToAppend []models.Tag

		// 处理每个传入的标签
		for _, t := range tags {
			var tag models.Tag
			// 使用原子操作查找或创建标签
			err := tx.Where(models.Tag{TagName: t.TagName}).FirstOrCreate(&tag).Error
			if err != nil {
				return err
			}
			tagsToAppend = append(tagsToAppend, tag)
		}

		// 创建文章
		if err := tx.Create(article).Error; err != nil {
			return err
		}

		// 关联标签
		if err := tx.Model(article).Association("Tags").Append(tagsToAppend); err != nil {
			return err
		}

		return nil
	})
}

// ArticleCreateWithTags 这个性能最好
func (a *ArticleTagService) ArticleCreateWithTags(article *models.Article, tags []models.Tag) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {

		// 提取标签名用于后续查询
		tagNames := make([]string, len(tags))
		for i, tag := range tags {
			tagNames[i] = tag.TagName
		}

		// 批量处理标签 (原子操作)
		if len(tags) > 0 {
			// 使用 Upsert 语法批量创建/跳过标签
			if err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "tag_name"}},
				DoNothing: true,
			}).Create(tags).Error; err != nil {
				return err
			}
		}

		// 重新查询完整标签数据（获取 ID）
		if err := tx.Where("tag_name IN ?", tagNames).Find(&tags).Error; err != nil {
			return err
		}

		// 2. 创建文章并直接关联标签 (单次操作)
		return tx.Omit("Tags.*").Create(&article).Association("Tags").Append(tags)
	})
}

func (a *ArticleTagService) ArticleUpdateWithTags(article *models.Article, mp *map[string]interface{}, tags []models.Tag) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 批量处理标签 (原子操作)
		if len(tags) > 0 {
			// 使用 Upsert 语法批量创建/跳过标签
			if err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "tag_name"}},
				DoNothing: true,
			}).Create(&tags).Error; err != nil {
				return err
			}
		}

		// 2. 创建文章并直接关联标签 (单次操作)
		// 2. 更新文章主体
		if err := tx.Model(article).Updates(mp).Error; err != nil {
			return err
		}
		// 3. 替换关联（GORM 仍然会操作中间表）
		return tx.Model(article).Association("Tags").Replace(tags)

	})
}

func (a *ArticleTagService) GetArticlesByTagId(tagID uint, page models.PageInfo) (
	tag models.Tag,
	articles []models.Article,
	count int64,
	err error,
) {
	db := global.DB

	// 1. 获取标签基础信息
	if err = db.First(&tag, tagID).Error; err != nil {
		return models.Tag{}, nil, 0, err
	}

	// 2. 使用 JOIN 直接分页查询
	query := db.
		Joins("JOIN article_tag_relations ON article.id = article_tag_relations.article_id").
		Where("article_tag_relations.tag_id = ?", tagID)

	//3. 获取总数
	if err = query.Model(&models.Article{}).Count(&count).Error; err != nil {
		return tag, nil, 0, err
	} else if count == 0 {
		return tag, []models.Article{}, 0, nil
	}

	// 4. 分页查询
	if err = query.
		Order("article.created_at DESC").
		Offset((page.Page - 1) * page.Limit).
		Limit(page.Limit).
		Find(&articles).Error; err != nil {
		return tag, nil, 0, err
	}

	return tag, articles, count, nil
}

func (a *ArticleTagService) CreateTags(tag *models.Tag) error {
	return global.DB.Create(tag).Error
}

// FindAllTags 获取所有的tag
func (a *ArticleTagService) FindAllTags(tags []*models.Tag, db *gorm.DB) (err error) {

	// 查询所有的 tag 记录
	err = db.Find(tags).Error
	if err != nil {
		return err
	}

	return nil
}
