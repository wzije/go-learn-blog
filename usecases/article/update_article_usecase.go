package article

import (
	"github.com/golang/learn-blog/domains"
	"github.com/golang/learn-blog/domains/repository"
)

//buat interface add article usecase
type UpdateArticleUsecase interface {
	Update(article *domains.Article) (interface{}, error)
}

//buat struct
type updateArticleUsecase struct {
	articleRepo repository.ArticleRepository
}

func (f updateArticleUsecase) Update(article *domains.Article) (interface{}, error) {
	return article, f.articleRepo.Update(article)
}

func NewUpdateArticleUsecase(articleRepository repository.ArticleRepository) UpdateArticleUsecase {
	return &updateArticleUsecase{articleRepository}
}
