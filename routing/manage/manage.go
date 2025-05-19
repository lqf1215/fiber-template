package manage

import (
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fiber-template/config"
	"fiber-template/database"
	"fiber-template/model"
	"fiber-template/pkg"
	"fiber-template/routing/types"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/pbkdf2"
	"gorm.io/gorm"
)

// ManagerLogin
// @Tags 管理端
// @Summary 【管理】密码登录
// @Accept json
// @Description 密码登录(用户名) API
// @Param param body types.ManagerLoginReq true "请求参数"
// @Success 200 {object} types.ManagerLoginResp "成功响应"
// @Router /manage/login [post]
func ManagerLogin(c *fiber.Ctx) error {
	reqParams := types.ManagerLoginReq{}
	db := database.DB
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "request params error", "请求参数错误"))
	}
	if reqParams.Username == "" && strings.TrimSpace(reqParams.Password) == "" {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "Please input a password", "请输入用户名或密码"))
	}
	manager, err := model.GetManagerByUsername(db, reqParams.Username)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "The manager id does not exist", "No,用户名或者密码错误"))
		}
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "get manager by username err", "Err,用户名或者密码错误"))
	}
	salt := []byte(manager.Username) // 自定义盐
	// 对支付密码进行hash并使用自定义盐
	hashedPassword := pbkdf2.Key([]byte(reqParams.Password), salt, 4096, 32, sha256.New)
	// 将二进制哈希密码转换为Base64编码，便于存储和传输
	hashedPasswordBase64 := base64.StdEncoding.EncodeToString(hashedPassword)

	if manager.Password != hashedPasswordBase64 {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "The password is incorrect", "001,用户名或者密码错误"))
	}
	returnT := pkg.RandomString(64)
	token := returnT + ":" + strconv.FormatInt(time.Now().Unix(), 10)
	if manager.ID != 0 {
		manager.Token = token
		if err = model.ManagerRefreshToken(db, int64(manager.ID), token); err != nil {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "db UserRefreshAppToken err", "002,登录失败"))
		}
	}
	return c.JSON(pkg.SuccessResponse(types.ManagerLoginResp{
		Token: returnT,
	}))
}
