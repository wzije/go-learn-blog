package article

import (
	"github.com/golang/learn-blog/domains"
	"github.com/golang/learn-blog/domains/repository"
)

//buat interface add article usecase
type AddArticleUsecase interface {
	Add(article *domains.Article) (interface{}, error)
}

//buat struct
type addArticleUsecase struct {
	articleRepo repository.ArticleRepository
}

func (a addArticleUsecase) Add(article *domains.Article) (interface{}, error) {
	return article, a.articleRepo.Save(article)
}

func NewAddArticleUsecase(articleRepository repository.ArticleRepository) AddArticleUsecase {
	return &addArticleUsecase{articleRepository}
}
