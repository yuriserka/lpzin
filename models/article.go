package models

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Picture string `json:"picture_path"`
}

var articleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Article 1 body", Picture: "../static/img/art_1.png"},
	article{ID: 2, Title: "Article 2", Content: "Article 2 body", Picture: "../static/img/art_2.png"},
	article{ID: 3, Title: "Article 3", Content: "Higor Ensaboado", Picture: "../static/img/art_3.png"},
}

func GetAllArticles() []article {
	return articleList
}

func GetArticleByID(id int) (*article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
