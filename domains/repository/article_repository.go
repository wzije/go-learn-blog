package repository

import (
	"database/sql"
	"fmt"

	"github.com/golang/learn-blog/domains"
)

type ArticleRepository interface {
	Fetch() ([]*domains.Article, error)
	Find(ID int) (*domains.Article, error)
	Save(article *domains.Article) error
	Update(article *domains.Article) error
	Delete(id int) error
}

type articleService struct {
	db *sql.DB
}

func (a articleService) Fetch() ([]*domains.Article, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM articles"

	rows, err := a.db.Query(query)
	if err != nil {
		return nil, err
	}

	var articles []*domains.Article

	for rows.Next() {
		article := new(domains.Article)
		err := rows.Scan(
			&article.ID,
			&article.Name,
			&article.Description,
			&article.CreatedAt,
			&article.UpdatedAt,
		)

		if err != nil {
			fmt.Printf("error query", err)
		}

		articles = append(articles, article)
	}

	return articles, nil
}

func (a articleService) Find(ID int) (*domains.Article, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM articles WHERE id=?"

	row := a.db.QueryRow(query, ID)

	article := new(domains.Article)
	err := row.Scan(
		&article.ID,
		&article.Name,
		&article.Description,
		&article.CreatedAt,
		&article.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return article, nil
}

func (a articleService) Save(article *domains.Article) error {
	query := "INSERT INTO articles (name, description) values (?,?)"

	ps, err := a.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = ps.Exec(article.Name, article.Description)

	return err
}

func (a articleService) Update(article *domains.Article) error {
	query := "UPDATE articles SET name=?, description=? WHERE id=?"

	ps, err := a.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = ps.Exec(article.Name, article.Description, article.ID)

	return err
}

func (a articleService) Delete(id int) error {
	query := "DELETE FROM articles WHERE id=?"

	ps, err := a.db.Prepare(query)

	if err != nil {
		return err
	}

	_, err = ps.Exec(id)

	return err
}

func NewArticleRepository(db *sql.DB) ArticleRepository {
	return &articleService{db}
}
