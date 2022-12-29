package model

import "gorm.io/gorm"

type ArticleTag struct {
	*gorm.Model
	TagID     uint32 `json:"tag_id"`
	ArticleID uint32 `json:"article_id"`
}

func (a ArticleTag) TableName() string {
	return "article_tags"
}

func (a ArticleTag) GetByAID(db *gorm.DB) (ArticleTag, error) {
	var articleTag ArticleTag
	err := db.Where("article_id = ? AND deleted_at IS NULL", a.ArticleID).First(&articleTag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return articleTag, err
	}
	return articleTag, nil
}

func (a ArticleTag) ListByTID(db *gorm.DB) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	if err := db.Where("tag_id = ? AND deleted_at IS NULL", a.TagID).Find(&articleTags).Error; err != nil {
		return nil, err
	}
	return articleTags, nil
}

func (a ArticleTag) ListByAIDs(db *gorm.DB, articleIDs []uint32) ([]*ArticleTag, error) {
	var articleTags []*ArticleTag
	err := db.Where("article_id IN (?) AND deleted_at IS NULL", articleIDs).Find(&articleTags).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return articleTags, nil
}

func (a ArticleTag) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a ArticleTag) UpdateOne(db *gorm.DB, values interface{}) error {
	return db.Model(&a).Where("article_id = ? AND deleted_at IS NULL", a.ArticleID).Limit(1).Updates(values).Error
}

func (a ArticleTag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND deleted_at IS NULL", a.Model.ID).Delete(&a).Error
}

func (a ArticleTag) DeleteOne(db *gorm.DB) error {
	return db.Where("article_id = ? AND deleted_at IS NULL", a.ArticleID).Delete(&a).Limit(1).Error
}
