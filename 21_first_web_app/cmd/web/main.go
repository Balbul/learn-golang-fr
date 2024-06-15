package main

import (
	"fmt"
	"net/http"

	"github.com/Balbul/firstwebapp/config"
	"github.com/Balbul/firstwebapp/internal/handlers"
)

func main() {
	var appConfig config.Config

	templateCache, err := handlers.CreateTemplateCache()

	if err != nil {
		panic(err)
	}

	appConfig.TemplateCache = templateCache
	appConfig.Port = ":8080"

	handlers.CreateTemplate(&appConfig)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/contact", handlers.Contact)

	fmt.Printf("(http://localhost%v) - server is runing.", appConfig.Port)
	http.ListenAndServe(appConfig.Port, nil)
}
