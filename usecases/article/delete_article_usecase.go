package article

import (
	"github.com/golang/learn-blog/domains/repository"
)

//buat interface add article usecase
type DeleteArticleUsecase interface {
	Delete(id int) (interface{}, error)
}

//buat struct
type deleteArticleUsecase struct {
	articleRepo repository.ArticleRepository
}

func (f deleteArticleUsecase) Delete(id int) (interface{}, error) {
	return id, f.articleRepo.Delete(id)
}

func NewDeleteArticleUsecase(articleRepository repository.ArticleRepository) DeleteArticleUsecase {
	return &deleteArticleUsecase{articleRepository}
}
