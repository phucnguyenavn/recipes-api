package main

import (
	"ex.recipes/recipe/src/main/controller"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.POST("/recipes", controller.NewRecipeHandler)
	router.Run()
}