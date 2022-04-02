package main

import (
	"fmt"
	"net/http"
	"simple-article/article"
	"simple-article/config"
)

var Ctrl *article.ControllerImpl

func init() {
	config.ProvideDB()
	ProvideController()
	registerRoutes()
}

func main() {
	var address = "localhost:9000"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func registerRoutes() {
	fmt.Println("-------register routes-------")
	http.HandleFunc("/list", Ctrl.GetListOfArticles)
	http.HandleFunc("/create", Ctrl.PostArticleHandler)
}

func ProvideController() {
	repo := article.ProvideRepository(config.GetGormClient().DB)
	svc := article.ProvideService(repo)
	iCtrl := article.Provide(svc)

	Ctrl = iCtrl.(*article.ControllerImpl)
}
