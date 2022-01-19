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
	if err := c.ShouldBindJSON(&recipe); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"error": err.Error()})

		return

	}

	recipe.SetId(xid.New().String())

	recipe.SetPublishedAt(time.Now())

	recipes = append(recipes, recipe)

	c.JSON(http.StatusOK, recipe)

}

var recipes []model.Recipe

func init() {
	recipes = make([]model.Recipe, 0)
}
