package main

import (
	admin_models "goblog/admin/models"
	"goblog/config"
	"net/http"
)

func main() {
	admin_models.Post{}.Migrate()
	admin_models.User{}.Migrate()
	http.ListenAndServe(":8080", config.Routes())
}
