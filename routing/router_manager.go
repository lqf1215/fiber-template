package routing

import (
	"fiber-template/routing/manage"
	"github.com/gofiber/fiber/v2"
)

// SetupManager manager 接口
func SetupManager(f *fiber.App) {

	mangerApi := f.Group("/manager")
	mangerApi.Post("/login", manage.ManagerLogin) //登录

}
