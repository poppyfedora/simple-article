package article

import (
	"github.com/go-chi/render"
	"net/http"
	"simple-article/utils"
	"time"
)

type ControllerImpl struct {
	svc Service
}

func Provide(svc Service) IController {
	return &ControllerImpl{
		svc: svc,
	}
}

func (c ControllerImpl) PostArticleHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	errMsg := "success"
	statusCode := http.StatusOK
	var data interface{}

	defer func() {
		utils.ReturnResponse(w, startTime, data, errMsg, statusCode)
	}()

	request := new(Article)
	if err := render.DecodeJSON(r.Body, &request); err != nil {
		errMsg = err.Error()
		statusCode = http.StatusBadRequest
		return
	}

	err := c.svc.AddNew(*request)
	if err != nil {
		errMsg = err.Error()
		statusCode = http.StatusInternalServerError
		return
	}
}

func (c ControllerImpl) GetListOfArticles(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	errMsg := "success"
	statusCode := http.StatusOK
	var data interface{}

	defer func() {
		utils.ReturnResponse(w, startTime, data, errMsg, statusCode)
	}()

	keyword := r.URL.Query().Get("keyword")
	author := r.URL.Query().Get("author")

	list, err := c.svc.GetLists(keyword, author)
	if err != nil {
		errMsg = err.Error()
		statusCode = http.StatusInternalServerError
	}
	data = list
}
