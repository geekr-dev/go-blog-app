package model

import (
	"github.com/geekr-dev/go-blog-app/pkg/app"
	"gorm.io/gorm"
)

type Article struct {
	*Model
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	Content      string `json:"content"`
	FeatureImage string `json:"feature_image"`
	State        uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "articles"
}

func (a Article) Create(db *gorm.DB) (*Article, error) {
	if err := db.Create(&a).Error; err != nil {
		return nil, err
	}
	return &a, nil
}

func (a Article) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&a).Where("id = ? AND deleted_at is NULL", a.ID).Updates(values).Error
}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ? AND deleted_at is NULL", a.ID).Delete(&a).Error
}

func (a Article) Get(db *gorm.DB) (Article, error) {
	var article Article
	db = db.Where("id = ? AND state = ? AND deleted_at IS NULL", a.ID, a.State)
	err := db.First(&article).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return article, err
	}
	return article, nil
}

// 根据标签获取文章列表(关联查询)
func (a Article) ListByTagID(db *gorm.DB, tagID uint32, pageOffset, pageSize int) ([]*ArticleRow, error) {
	fields := []string{"ar.id AS article_id", "ar.title AS article_title", "ar.desc AS article_desc", "ar.feature_image", "ar.content"}
	fields = append(fields, []string{"t.id AS tag_id", "t.name AS tag_name"}...)

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	rows, err := db.Select(fields).Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").
		Where("at.tag_id = ? AND ar.state = ? AND ar.deleted_at IS NULL", tagID, a.State).
		Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []*ArticleRow
	for rows.Next() {
		r := &ArticleRow{}
		if err := rows.Scan(&r.ArticleID, &r.ArticleTitle, &r.ArticleDesc, &r.FeatureImage, &r.Content, &r.TagID, &r.TagName); err != nil {
			return nil, err
		}
		articles = append(articles, r)
	}

	return articles, nil
}

func (a Article) CountByTagID(db *gorm.DB, tagID uint32) (int64, error) {
	var count int64
	err := db.Table(ArticleTag{}.TableName()+" AS at").
		Joins("LEFT JOIN `"+Tag{}.TableName()+"` AS t ON at.tag_id = t.id").
		Joins("LEFT JOIN `"+Article{}.TableName()+"` AS ar ON at.article_id = ar.id").
		Where("at.tag_id = ? AND ar.state = ? AND ar.deleted_at IS NULL", tagID, a.State).
		Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

type ArticleRow struct {
	ArticleID    uint32
	TagID        uint32
	TagName      string
	ArticleTitle string
	ArticleDesc  string
	FeatureImage string
	Content      string
}

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}
