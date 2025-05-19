package intcpt

import (
	"fiber-template/config"
	"fiber-template/database"
	"fiber-template/model"
	"fiber-template/pkg"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

// AuthApp Protected protect routes
func AuthApp() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var (
			userId int64
			token  = c.Get(config.LOCAL_TOKEN)
			db     = database.DB
			err    error
		)

		// 打印请求地址
		config.Log.Info("Request URL: ", c.Path())
		if token == "" || len(token) < 10 {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is null", ""))
		}

		user, err := model.UserSelectIdByToken(db, token)
		if err != nil {
			if err.Error() != "record not found" {
				fmt.Println(err.Error())
			}
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token invalid or network error", "token失效或网络故障"))
		}
		userId = int64(user.ID)

		if !pkg.CheckSpecialCharacters(&token) {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is invalid", ""))
		}
		//检查token 有效时间
		if !pkg.CheckTokenValidityTime(&user.Token) {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is exceed", ""))
		}

		//刷新token有效时间
		if err = model.UserRefreshToken(db, userId, user.Token); err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "database UserRefreshAppToken err", ""))
		}

		c.Locals(config.LOCAL_USERID_UINT, uint(userId))
		c.Locals(config.LOCAL_USERID_INT64, userId)
		_ = c.Next()

		return nil
	}
}

// AuthManagerApp Protected protect routes
func AuthManagerApp() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var (
			userId int64
			token  = c.Get(config.LOCAL_TOKEN)
			db     = database.DB
			err    error
		)

		// 打印请求地址
		config.Log.Info("Request URL: ", c.Path())
		if token == "" || len(token) < 10 {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is null", ""))
		}

		user, err := model.ManagerSelectIdByToken(db, token)
		if err != nil {
			if err.Error() != "record not found" {
				fmt.Println(err.Error())
			}
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token invalid or network error", "token失效或网络故障"))
		}
		userId = int64(user.ID)

		if !pkg.CheckSpecialCharacters(&token) {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is invalid", ""))
		}
		//检查token 有效时间
		if !pkg.CheckTokenValidityTime(&user.Token) {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "token is exceed", ""))
		}

		//刷新token有效时间
		if err = model.ManagerRefreshToken(db, userId, user.Token); err != nil {
			return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "database UserRefreshAppToken err", ""))
		}

		c.Locals(config.MANAGER_LOCAL_USERID_UINT, uint(userId))
		c.Locals(config.MANAGER_LOCAL_USERID_INT64, userId)
		_ = c.Next()
		return nil
	}
}
