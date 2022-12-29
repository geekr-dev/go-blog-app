package model

import "github.com/geekr-dev/go-blog-app/pkg/app"

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

type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}
