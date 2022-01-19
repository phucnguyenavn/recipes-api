package controller

import (
	"ex.recipes/recipe/src/main/model"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"net/http"
	"time"
)

func NewRecipeHandler(c *gin.Context) {
	var recipe model.Recipe
	err := c.ShouldBindJSON(&recipe)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"error": err.Error()})

		return

	}

	recipe.ID = xid.New().String()
	recipe.PublishedAt = time.Now()

	recipes = append(recipes, recipe)

	c.JSON(http.StatusOK, recipe)

}

func ListRecipesHandler(c *gin.Context) {

	c.JSON(http.StatusOK, recipes)

}

var recipes []model.Recipe

func init() {
	recipes = make([]model.Recipe, 0)
}
