package main

import (
	"github.com/samarqandi/blog-gin/api/controller"
	"github.com/samarqandi/blog-gin/api/repository"
	"github.com/samarqandi/blog-gin/api/routes"
	"github.com/samarqandi/blog-gin/api/service"
	"github.com/samarqandi/blog-gin/infrastructure"
	"github.com/samarqandi/blog-gin/models"
)

func init() {
	infrastructure.LoadEnv()
}

func main() {

	router := infrastructure.NewGinRouter()
	db := infrastructure.NewDatabase()
	postRepository := repository.NewPostRepository(db)
	postService := service.NewPostService(postRepository)
	postController := controller.NewPostController(postService)
	postRoute := routes.NewPostRoute(postController, router)
	postRoute.Setup()

	db.DB.AutoMigrate(&models.Post{})
	router.Gin.Run(":8000")

}
