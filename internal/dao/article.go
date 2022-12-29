package dao

import (
	"github.com/geekr-dev/go-blog-app/internal/model"
	"github.com/geekr-dev/go-blog-app/pkg/app"
)

type Article struct {
	ID           uint32 `json:"id"`
	TagID        uint32 `json:"tag_id"`
	Title        string `json:"title"`
	Desc         string `json:"desc"`
	Content      string `json:"content"`
	FeatureImage string `json:"feature_image"`
	CreatedBy    string `json:"created_by"`
	UpdatedBy    string `json:"updated_by"`
	State        uint8  `json:"state"`
}

func (d *Dao) CreateArticle(param *Article) (*model.Article, error) {
	article := model.Article{
		Title:        param.Title,
		Desc:         param.Desc,
		Content:      param.Content,
		FeatureImage: param.FeatureImage,
		State:        param.State,
		Model:        &model.Model{CreatedBy: param.CreatedBy},
	}
	return article.Create(d.engine)
}

func (d *Dao) UpdateArticle(param *Article) error {
	article := model.Article{Model: &model.Model{ID: param.ID}}
	values := map[string]interface{}{
		"updated_by": param.UpdatedBy,
		"state":      param.State,
	}
	if param.Title != "" {
		values["title"] = param.Title
	}
	if param.FeatureImage != "" {
		values["feature_image"] = param.FeatureImage
	}
	if param.Desc != "" {
		values["desc"] = param.Desc
	}
	if param.Content != "" {
		values["content"] = param.Content
	}
	return article.Update(d.engine, values)
}

func (d *Dao) GetArticle(id uint32, state uint8) (model.Article, error) {
	article := model.Article{Model: &model.Model{ID: id}, State: state}
	return article.Get(d.engine)
}

func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{Model: &model.Model{ID: id}}
	return article.Delete(d.engine)
}

func (d *Dao) CountArticleListByTagID(id uint32, state uint8) (int64, error) {
	article := model.Article{State: state}
	return article.CountByTagID(d.engine, id)
}

func (d *Dao) GetArticleListByTagID(id uint32, state uint8, page, pageSize int) ([]*model.ArticleRow, error) {
	article := model.Article{State: state}
	return article.ListByTagID(d.engine, id, app.GetPageOffset(page, pageSize), pageSize)
}
