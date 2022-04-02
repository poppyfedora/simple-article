package article

import "net/http"

type IController interface {
	PostArticleHandler(w http.ResponseWriter, r *http.Request)
	GetListOfArticles(w http.ResponseWriter, r *http.Request)
}
