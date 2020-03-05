package article

import (
	"encoding/json"
	"net/http"

	"github.com/golang/learn-blog/domains"

	"github.com/golang/learn-blog/usecases/article"
)

type AddHandler interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type addHandler struct {
	addUsecase article.AddArticleUsecase
}

func (a addHandler) Handle(w http.ResponseWriter, r *http.Request) {
	req := domains.Article{}

	//mapping request to object
	err := json.NewDecoder(r.Body).Decode(req)

	if err != nil {
		w.Header().Set("content-type", "Application/Json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Request Error")
		return
	}

	response, err := a.addUsecase.Add(&req)

	if err != nil {
		w.Header().Set("content-type", "Application/Json")
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Request Error")
		return
	}

	w.Header().Set("content-type", "Application/Json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(response)

}

func NewAddHandler(usecase article.AddArticleUsecase) AddHandler {
	return &addHandler{usecase}
}
