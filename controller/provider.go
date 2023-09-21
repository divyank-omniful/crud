package controller

import (
	"github.com/google/wire"
	"taskOne/interfaces"
	"taskOne/repository"
	"taskOne/service"
)

var CategoryProviderSet = wire.NewSet(
	NewController,
	service.NewService,
	repository.NewCategoryRepo,
	wire.Bind(new(interfaces.CategoryController), new(*CategoryController)),
	wire.Bind(new(interfaces.CategoryService), new(*service.Service)),
	wire.Bind(new(interfaces.CategoryRepository), new(*repository.Repository)),
)
