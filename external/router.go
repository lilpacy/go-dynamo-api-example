package external

import (
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
	"github.com/lilpacy/gin-nginx-next-docker-min/adapter/controllers"
)

func Setup(table dynamo.Table) *gin.Engine {
	router := gin.Default()

	postController := controllers.NewPostController(table)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/v1/posts", func(c *gin.Context) {
		postController.GetAll(c)
	})
	router.GET("/v1/post", func(c *gin.Context) {
		postController.Get(c)
	})
	router.POST("/v1/post", func(c *gin.Context) {
		postController.Create(c)
	})
	router.POST("/v1/mockpost", func(c *gin.Context) {
		postController.MockCreate(c)
	})
	return router
}
