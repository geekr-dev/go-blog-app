package errcode

var (
	ErrorGetArticleFail    = NewError(20020001, "获取文章失败")
	ErrorGetArticlesFail   = NewError(20020002, "获取文章列表失败")
	ErrorCreateArticleFail = NewError(20020003, "创建文章失败")
	ErrorUpdateArticleFail = NewError(20020004, "更新文章失败")
	ErrorDeleteArticleFail = NewError(20020005, "删除文章失败")
	ErrorCountArticleFail  = NewError(20020006, "统计文章失败")
)
