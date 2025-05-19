package routing

import (
	intcpt "fiber-template/intercept"
	"fiber-template/routing/api"
	"github.com/gofiber/fiber/v2"
)

// 注册路由
func SetupApp(f *fiber.App) {
	appApi := f.Group("/api")

	appApi.Post("/register", api.Register) // 注册
	appApi.Post("/login", api.Login)       // 登录
	appApi.Post("/upload", api.Upload)     // 上传图片

	appApi.Post("/user/info", intcpt.AuthApp(), api.UserInfo) // 用户信息
}
