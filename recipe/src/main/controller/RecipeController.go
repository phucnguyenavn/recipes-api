package controller

import (
	"encoding/json"
	"ex.recipes/recipe/src/main/model"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"io/ioutil"
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

func UpdateRecipeHandler(c *gin.Context) {
	var recipe model.Recipe
	id := c.Param("id")
	err := c.ShouldBindJSON(&recipe)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{

			"error": err.Error()})

		return

	}

	index := -1

	for i := 0; i < len(recipes); i++ {

		if recipes[i].ID == id {

			index = i

		}

	}

	if index == -1 {

		c.JSON(http.StatusNotFound, gin.H{

			"error": "Recipe not found"})

		return

	}

	recipes[index] = recipe

	c.JSON(http.StatusOK, recipe)

}

func DeleteRecipeHandler(c *gin.Context) {

	id := c.Param("id")

	index := -1

	for i := 0; i < len(recipes); i++ {

		if recipes[i].ID == id {

			index = i

		}

	}

	if index == -1 {

		c.JSON(http.StatusNotFound, gin.H{

			"error": "Recipe not found"})

		return

	}

	recipes = append(recipes[:index], recipes[index+1:]...)

	c.JSON(http.StatusOK, gin.H{

		"message": "Recipe has been deleted",
	})
}

var recipes []model.Recipe

func init() {
	recipes = make([]model.Recipe, 0)
	file, _ := ioutil.ReadFile("recipe/src/main/recipes.json")

	_ = json.Unmarshal(file, &recipes)
}
