package model

import (
	"github.com/geekr-dev/go-blog-app/pkg/app"
	"gorm.io/gorm"
)

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "tags"
}

func (t Tag) Get(db *gorm.DB) (Tag, error) {
	var tag Tag
	err := db.Where("id = ? AND state = ? AND deleted_at IS NULL", t.Model.ID, t.State).First(&tag).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return tag, err
	}
	return tag, nil
}

func (t Tag) Count(db *gorm.DB) (int64, error) {
	var count int64
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err := db.Model(&t).Where("deleted_at is NULL").Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

func (t Tag) List(db *gorm.DB, pageOffset, pageSize int) ([]*Tag, error) {
	var tags []*Tag
	var err error
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}
	if t.Name != "" {
		db = db.Where("name = ?", t.Name)
	}
	db = db.Where("state = ?", t.State)
	if err = db.Where("deleted_at is NULL").Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (t Tag) Create(db *gorm.DB) error {
	return db.Create(&t).Error
}

func (t Tag) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&t).Where("id = ? AND deleted_at is NULL", t.ID).Updates(values).Error
}

func (t Tag) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND deleted_at is NULL", t.ID).Delete(&t).Error
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
