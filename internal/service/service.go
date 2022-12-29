package service

import (
	"context"

	"github.com/geekr-dev/go-blog-app/global"
	"github.com/geekr-dev/go-blog-app/internal/dao"
)

type Service struct {
	ctx context.Context
	dao *dao.Dao
}

func New(ctx context.Context) Service {
	svc := Service{ctx: ctx}
	svc.dao = dao.New(global.DBEngine)
	return svc
}
