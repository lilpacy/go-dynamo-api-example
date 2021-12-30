package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/guregu/dynamo"
	"github.com/lilpacy/gin-nginx-next-docker-min/adapter/gateway"
	"github.com/lilpacy/gin-nginx-next-docker-min/adapter/interfaces"
	"github.com/lilpacy/gin-nginx-next-docker-min/domain"
	"github.com/lilpacy/gin-nginx-next-docker-min/usecase"
)

type PostController struct {
	Interactor usecase.PostInteractor
}

func NewPostController(table dynamo.Table) *PostController {
	return &PostController{
		Interactor: usecase.PostInteractor{
			PostRepository: &gateway.PostRepository{Table: table},
		},
	}
}

func (controller *PostController) Create(c interfaces.Context) {
	postParams := domain.PostDTO{}
	if err := c.Bind(&postParams); err != nil { panic(err) }
	post := domain.New(postParams)
	_, err := controller.Interactor.Add(post)
	if err != nil {
		fmt.Printf("Failed to get post[%v]\n", err)
		panic(err)
	}
	c.JSON(200, gin.H{
		"message": "ok",
	})
}

func (controller *PostController) MockCreate(c interfaces.Context) {
	postParams := domain.PostDTO{}
	if err := c.Bind(&postParams); err != nil { panic(err) }
	post := domain.New(postParams)
	c.JSON(200, gin.H{
		"post": post,
	})
	//_, err := controller.Interactor.Add(domain.MockPost)
	//if err != nil {
	//	fmt.Printf("Failed to get post[%v]\n", err)
	//	panic(err)
	//}
	//c.JSON(200, gin.H{
	//	"message": "ok",
	//})
}

func (controller *PostController) GetAll(c interfaces.Context) {
	posts, err := controller.Interactor.FindAll()
	if err != nil {
		fmt.Printf("Failed to get post[%v]\n", err)
	}
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func (controller *PostController) Get(c interfaces.Context) {
	post, err := controller.Interactor.Find(c.Query("id"))
	if err != nil {
		fmt.Printf("Failed to get post[%v]\n", err)
	}
	c.JSON(200, gin.H{
		"post": post,
	})
}
