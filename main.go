package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/dandimuzaki/project-app-portfolio-golang/database"
	"github.com/dandimuzaki/project-app-portfolio-golang/handler"
	"github.com/dandimuzaki/project-app-portfolio-golang/repository"
	"github.com/dandimuzaki/project-app-portfolio-golang/router"
	"github.com/dandimuzaki/project-app-portfolio-golang/service"
	"github.com/dandimuzaki/project-app-portfolio-golang/utils"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	var templates = template.Must(template.New("").ParseGlob("views/**/*.html"))

	logger, err := utils.InitLogger("./logs/app", true)

	repo := repository.NewRepository(db, logger)
	service := service.NewService(repo, logger)
	handler := handler.NewHandler(service, templates, logger)

	r := router.NewRouter(service, handler, logger)
	
	// fmt.Println(service.ProjectService.GetAllProjects(nil, nil))

	fmt.Println("server running on port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("error server")
	}
}
