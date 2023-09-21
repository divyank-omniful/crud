package interfaces

import (
	"context"
	"github.com/gin-gonic/gin"
	"taskOne/models"
)

type CategoryController interface {
	CreateEntity(c *gin.Context)
}

type CategoryService interface {
	CreateEntity(ctx context.Context, entity *models.Entity) (entity2 *models.Entity, err error)
}

type CategoryRepository interface {
	Create(ctx context.Context, entity *models.Entity) (err error)
}
