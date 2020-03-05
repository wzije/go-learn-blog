package article

import (
	"github.com/golang/learn-blog/domains/repository"
)

//buat interface add article usecase
type FindArticleUsecase interface {
	Find(id int) (interface{}, error)
}

//buat struct
type findArticleUsecase struct {
	articleRepo repository.ArticleRepository
}

func (f findArticleUsecase) Find(id int) (interface{}, error) {
	return f.articleRepo.Find(id)
}

func NewFindArticleUsecase(articleRepository repository.ArticleRepository) FindArticleUsecase {
	return &findArticleUsecase{articleRepository}
}
