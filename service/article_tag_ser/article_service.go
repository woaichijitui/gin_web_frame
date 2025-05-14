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

func (a *ArticleTagService) ArticleCreateAndAppendTags1(article *models.Article, tags []models.Tag) error {
	return global.DB.Transaction(func(tx *gorm.DB) error {
		// 1. 提取所有标签名称
		tagNames := make([]string, len(tags))
		for i, t := range tags {
			tagNames[i] = t.TagName
		}

		// 2. 批量查询已存在的标签 (1次查询)
		var existingTags []models.Tag
		if err := tx.Model(&models.Tag{}).
			Select("id, tag_name").
			Where("tag_name IN ?", tagNames).
			Find(&existingTags).Error; err != nil {
			return err
		}

		// 3. 构建需要创建的标签 (利用唯一性保证)
		existingTagMap := make(map[string]struct{}, len(existingTags))
		for _, t := range existingTags {
			existingTagMap[t.TagName] = struct{}{}
		}

		newTags := make([]models.Tag, 0, len(tags)-len(existingTags))
		for _, t := range tags {
			if _, exists := existingTagMap[t.TagName]; !exists {
				newTags = append(newTags, models.Tag{
					TagName: t.TagName,
					TagDesc: t.TagDesc, // 确保传入的 tags 包含有效 TagDesc
				})
			}
		}

		// 4. 批量创建新标签 (1次查询，冲突时忽略)
		if len(newTags) > 0 {
			if err := tx.Clauses(clause.OnConflict{
				Columns:   []clause.Column{{Name: "tag_name"}},
				DoNothing: true, // 冲突时直接跳过，保持已有数据
			}).Create(&newTags).Error; err != nil {
				return err
			}
		}

		// 5. 合并全部标签ID (1次查询获取新标签ID)
		var allTags []models.Tag
		if len(newTags) > 0 {
			if err := tx.Model(&models.Tag{}).
				Select("id").
				Where("tag_name IN ?", tagNames).
				Find(&allTags).Error; err != nil {
				return err
			}
		} else {
			allTags = existingTags
		}

		// 6. 创建文章 (1次查询)
		if err := tx.Omit("Tags").Create(article).Error; err != nil { // 明确排除关联字段
			return err
		}

		// 7. 批量关联标签 (1次查询，使用 Replace 更高效)
		if len(allTags) > 0 {
			return tx.Model(article).Association("Tags").Replace(allTags)
		}
		return nil
	})
}

// 这个性能最好
func (a *ArticleTagService) ArticleCreateWithTags(article *models.Article, tags []models.Tag) error {
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
		return tx.Omit("Tags.*").Create(&article).Association("Tags").Append(tags)
	})
}

// 这个性能最好
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
