package model

import "github.com/geekr-dev/go-blog-app/pkg/app"

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "tags"
}

type TagSwagger struct {
	List  []*Tag
	Pager *app.Pager
}
