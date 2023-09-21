//go:build wireinject
// +build wireinject

package controller

import (
	"context"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
)

func Wire(database mongo.Database) CategoryController {
	panic(wire.Build(CategoryProviderSet))
}
