package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"taskOne/models"
	"taskOne/service"
)

var controllerOnce sync.Once
var controller CategoryController

type CategoryController struct {
	service service.Service
}

func NewController(categoryService service.Service) CategoryController {

	controllerOnce.Do(func() {
		controller = CategoryController{
			service: categoryService,
		}
	})
	return controller
	//categorySvc, err := cat.NewService()
	//if err != nil {
	//	return CategoryController{}, err
	//}
	//return CategoryController{
	//	service: categoryService,
	//}
}

func (ctrl *CategoryController) CreateEntity(c *gin.Context) {
	var entity *models.Entity
	if err := c.ShouldBindJSON(entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	entity, err := ctrl.service.CreateEntity(c, entity)
	if err != nil {
		fmt.Printf("create error %v", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, entity)
}

func (ctrl *CategoryController) CreateEntityHandler(context *gin.Context) {
	
}

//
//func ReadEntityHandler(c *gin.Context) {
//
//	idStr := c.Param("id")
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
//		return
//	}
//
//	entity, err := service.ReadEntity(id)
//	if err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": "Entity not found"})
//		return
//	}
//
//	c.JSON(http.StatusOK, entity)
//}
//
//func UpdateEntityHandler(c *gin.Context) {
//	idStr := c.Param("id")
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
//		return
//	}
//
//	var updatedEntity models.Entity
//	if err := c.ShouldBindJSON(&updatedEntity); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	fmt.Println(updatedEntity)
//	if err := service.UpdateEntity(id, &updatedEntity); err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, updatedEntity)
//}
//
//func DeleteEntityHandler(c *gin.Context) {
//	idStr := c.Param("id")
//	id, err := strconv.Atoi(idStr)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
//		return
//	}
//	if err := service.DeleteEntity(id); err != nil {
//		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
//		return
//	}
//	c.JSON(http.StatusOK, gin.H{"message": "Entity deleted successfully"})
//}
