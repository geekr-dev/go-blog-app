package service

import (
	"github.com/geekr-dev/go-blog-app/internal/dao"
	"github.com/geekr-dev/go-blog-app/internal/model"
	"github.com/geekr-dev/go-blog-app/pkg/app"
)

type ArticleRequest struct {
	ID    uint32 `form:"id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type ArticleListRequest struct {
	TagID uint32 `form:"tag_id" binding:"required,gte=1"`
	State uint8  `form:"state,default=1" binding:"oneof=0 1"`
}

type CreateArticleRequest struct {
	Title        string `form:"title" binding:"required,min=2,max=100"`
	Desc         string `form:"desc" binding:"min=2,max=255"`
	Content      string `form:"content" binding:"required"`
	FeatureImage string `form:"feature_image" binding:"required,max=255"`
	State        uint8  `form:"state,default=1" binding:"oneof=0 1"`
	TagID        uint32 `form:"tag_id" binding:"required,gte=1"`
	CreatedBy    string `form:"created_by" binding:"required,min=3,max=100"`
}

type UpdateArticleRequest struct {
	ID           uint32 `form:"id" binding:"required,gte=1"`
	Title        string `form:"title" binding:"min=2,max=100"`
	Desc         string `form:"desc" binding:"min=2,max=255"`
	Content      string `form:"content"`
	FeatureImage string `form:"feature_image" binding:"max=255"`
	State        uint8  `form:"state,default=1" binding:"oneof=0 1"`
	TagID        uint32 `form:"tag_id" binding:"gte=1"`
	UpdatedBy    string `form:"updated_by" binding:"required,min=3,max=100"`
}

type DeleteArticleRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

type Article struct {
	ID           uint32     `json:"id"`
	Title        string     `json:"title"`
	Desc         string     `json:"desc"`
	Content      string     `json:"content"`
	FeatureImage string     `json:"feature_image"`
	State        uint8      `json:"state"`
	Tag          *model.Tag `json:"tag"`
}

func (svc *Service) GetArticle(param *ArticleRequest) (*Article, error) {
	article, err := svc.dao.GetArticle(param.ID, param.State)
	if err != nil {
		return nil, err
	}

	if article.Model == nil {
		return nil, nil
	}

	articleTag, err := svc.dao.GetArticleTagByAID(article.ID)
	if err != nil {
		return nil, err
	}

	tag, err := svc.dao.GetTag(articleTag.TagID, model.STATE_OPEN)
	if err != nil {
		return nil, err
	}

	return &Article{
		ID:           article.ID,
		Title:        article.Title,
		Desc:         article.Desc,
		Content:      article.Content,
		FeatureImage: article.FeatureImage,
		State:        article.State,
		Tag:          &tag,
	}, nil
}

func (svc *Service) GetArticleList(param *ArticleListRequest, pager *app.Pager) ([]*Article, int64, error) {
	articleCount, err := svc.dao.CountArticleListByTagID(param.TagID, param.State)
	if err != nil {
		return nil, 0, err
	}

	articles, err := svc.dao.GetArticleListByTagID(param.TagID, param.State, pager.Page, pager.PageSize)
	if err != nil {
		return nil, 0, err
	}

	var articleList []*Article
	for _, article := range articles {
		articleList = append(articleList, &Article{
			ID:           article.ArticleID,
			Title:        article.ArticleTitle,
			Desc:         article.ArticleDesc,
			Content:      article.Content,
			FeatureImage: article.FeatureImage,
			Tag:          &model.Tag{Model: &model.Model{ID: article.TagID}, Name: article.TagName},
		})
	}

	return articleList, articleCount, nil
}

func (svc *Service) CreateArticle(param *CreateArticleRequest) error {
	article, err := svc.dao.CreateArticle(&dao.Article{
		Title:        param.Title,
		Desc:         param.Desc,
		Content:      param.Content,
		FeatureImage: param.FeatureImage,
		State:        param.State,
		CreatedBy:    param.CreatedBy,
	})
	if err != nil {
		return err
	}

	err = svc.dao.CreateArticleTag(article.ID, param.TagID)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) UpdateArticle(param *UpdateArticleRequest) error {
	err := svc.dao.UpdateArticle(&dao.Article{
		ID:           param.ID,
		Title:        param.Title,
		Desc:         param.Desc,
		Content:      param.Content,
		FeatureImage: param.FeatureImage,
		State:        param.State,
		UpdatedBy:    param.UpdatedBy,
	})
	if err != nil {
		return err
	}

	err = svc.dao.UpdateArticleTag(param.ID, param.TagID)
	if err != nil {
		return err
	}

	return nil
}

func (svc *Service) DeleteArticle(param *DeleteArticleRequest) error {
	err := svc.dao.DeleteArticle(param.ID)
	if err != nil {
		return err
	}

	err = svc.dao.DeleteArticleTag(param.ID)
	if err != nil {
		return err
	}

	return nil
}
