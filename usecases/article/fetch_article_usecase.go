package article

import (
	"github.com/golang/learn-blog/domains/repository"
)

//buat interface add article usecase
type FetchArticleUsecase interface {
	Fetch() (interface{}, error)
}

//buat struct
type fetchArticleUsecase struct {
	articleRepo repository.ArticleRepository
}

func (a fetchArticleUsecase) Fetch() (interface{}, error) {
	return a.articleRepo.Fetch()
}

func NewFetchArticleUsecase(articleRepository repository.ArticleRepository) FetchArticleUsecase {
	return &fetchArticleUsecase{articleRepository}
}
