//+build wireinject

package main

import (
	controllers2 "ginapp/auth/controllers"
	services2 "ginapp/auth/services"
	"ginapp/core/repositories"
	"ginapp/core/services"
	"ginapp/user/controllers"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func initUserController(db *gorm.DB) controllers.UserController {
	wire.Build(repositories.ProvideUserRepository, services.ProvideUserService, controllers.ProvideUserController)

	return controllers.UserController{}
}

func initAuthController(db *gorm.DB) controllers2.AuthController {
	wire.Build(repositories.ProvideUserRepository, services2.ProvideAuthService, services2.ProvideJwtService, controllers2.ProvideAuthController)

	return controllers2.AuthController{}
}
