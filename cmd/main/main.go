package main

import (
	"fiber-template/config"
	"fiber-template/database"
	//_ "fiber-template/docs"
	"fiber-template/routing"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

// @title           flygo API
// @version         1.0
// @description     This is flygo backend.
// @termsOfService  http://swagger.io/terms/

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      192.168.5.34:8330
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  flygo API
// @externalDocs.url          http://192.168.5.34:8330/swagger/index.html
func main() {
	database.ConnectDB()

	fiberApp := fiber.New(
		fiber.Config{
			BodyLimit: 40 * 1024 * 1024, // 40MB
		})
	fiberApp.Use(func(c *fiber.Ctx) error {
		// 允许所有域名进行跨域请求
		c.Set("Access-Control-Allow-Origin", "*")
		// 允许 GET、POST、PUT、DELETE 和 OPTIONS 方法进行跨域请求
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许客户端发送的请求头
		c.Set("Access-Control-Allow-Headers", "Content-Type, Authorization, token")
		// 在响应中添加 CORS 头
		if c.Method() == "OPTIONS" {
			return c.SendStatus(fiber.StatusOK)
		} else {
			return c.Next()
		}
	})
	database.InitDBData()

	fiberApp.Use(cors.New())
	fiberApp.Get("/swagger/*", fiberSwagger.WrapHandler)

	// Register routing
	routing.SetupApp(fiberApp)

	err := fiberApp.Listen(fmt.Sprintf(":%v", config.Config.Port))
	if err != nil {
		config.Log.Fatalf("Listen error: %v", err)
	}
}
