package api

import (
	"errors"
	"fiber-template/config"
	"fiber-template/database"
	"fiber-template/model"
	"fiber-template/pkg"
	"fiber-template/routing/types"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Register
// @Tags Api端
// @Summary 注册
// @Accept json
// @Description 注册(手机号或邮箱) API
// @Param param body types.RegisterReq true "请求参数"
// @Success 200 string  "成功响应"
// @Router /api/register [post]
func Register(c *fiber.Ctx) error {
	reqParams := types.RegisterReq{}
	err := c.BodyParser(&reqParams)
	db := database.DB
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "request params error", "请求参数错误"))
	}
	if strings.TrimSpace(reqParams.LoginPwd) == "" {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "Please input a password", "请输入密码"))
	}
	if len(reqParams.LoginPwd) < 6 {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "Password length must be greater than or equal to 6", "密码长度必须大于等于6"))
	}
	sender := strings.TrimSpace(reqParams.Email)
	phoneSender := ""

	// 邮箱
	if reqParams.LoginType == "2" {
		if strings.TrimSpace(reqParams.Email) == "" {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "email is empty", "邮箱不能为空"))
		}
		// 先检测邮箱是否正确的格式
		if !pkg.ValidateEmail(reqParams.Email) {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "email format is not correct", "邮箱格式不正确"))
		}
		reqParams.Area = ""
		reqParams.Phone = ""

	} else {
		if strings.TrimSpace(reqParams.Phone) == "" {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "phone is empty", "手机号不能为空"))
		}
		sender = fmt.Sprintf("%v%v", reqParams.Area, strings.TrimSpace(reqParams.Phone))
		phoneSender = fmt.Sprintf("%v%v", reqParams.Area, strings.TrimSpace(reqParams.Phone))
		reqParams.Email = ""
	}

	if reqParams.LoginPwd != reqParams.ConfirmLoginPwd {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "login password is not equal to confirm login password", "登录密码不一致"))
	}

	// v, err := model.GetValidCodeBySender(db, sender)
	// if err != nil {
	// 	config.Log.Errorf("[Register] query valid code error, sender: %s\n", sender)
	// 	return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "Verification code error", "验证码错误"))
	// }
	// if v.Code != reqParams.Code {
	// 	return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "Please enter the correct verification code", "请输入正确的验证码"))
	// }

	// if time.Now().Unix()-v.UpdatedAt.Unix() > config.ValidCodeOverTime {
	// 	v.State = "2" // 验证码过期
	// 	err := v.UpdateValidCode(db)
	// 	if err != nil {
	// 		fmt.Printf("[Register] valid code fail, Phone: %s, Email: %s, Code: %s\n", phoneSender, reqParams.Email, reqParams.Code)
	// 		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "valid code error", "验证码失败"))
	// 	}
	// 	config.Log.Errorf("[Register] code time out, Phone: %s, Email: %s, Code: %s\n", phoneSender, reqParams.Email, reqParams.Code)
	// 	return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "verification code has out of date", "验证码已超时"))
	// }
	var user *model.User
	if reqParams.LoginType == "2" {
		user, err = model.SelectUserByEmail(db, reqParams.Email)
	} else {
		user, err = model.SelectUserByPhone(db, phoneSender)
	}

	if err != nil {
		config.Log.Errorf("[Register] query user error, Phone: %s, Email: %s \n", phoneSender, reqParams.Email)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "register fail22", "注册失败2"))
		}
	}
	if user != nil {
		config.Log.Errorf("[Register] exit, sender: %s ", sender)
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "user has already registered", "该用户已注册"))
	}
	encryptPwd, err := pkg.EncryptData(config.LOGIN_CRT_PATH, strings.TrimSpace(reqParams.LoginPwd))
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "encryptPwd error", "密码加密失败"))
	}
	// 生成token
	returnT := pkg.RandomString(64)
	token := returnT + ":" + strconv.FormatInt(time.Now().Unix(), 10)
	newUser := model.User{
		Username: sender,
		LoginPwd: encryptPwd,
		Email:    reqParams.Email,
		Phone:    phoneSender,
		Flag:     "1",
		//AvatarUrl:       config.UserDefaultAvatar,
		Token: token,
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		err = newUser.Create(tx)
		if err != nil {
			config.Log.Errorf("[Register] create user error, Sender=%v err=%v \n", sender, err)
			return err
		}
		return nil
	})
	if err != nil {
		config.Log.Errorf("[Register] tx err, Sender=%v err=%v \n", sender, err)
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "register fail", "注册失败"))
	}

	//// 最后更新验证码状态
	//v.State = "1"
	//err = v.UpdateValidCode(db)
	//if err != nil {
	//	config.Log.Errorf("[Register] update valid code error, Sender=%v err=%v \n", sender, err)
	//	return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "Verification code verification failed", "验证码验证失败"))
	//}

	return c.JSON(pkg.SuccessResponse("success"))
}

// Login
// @Tags Api端
// @Summary 密码登录
// @Accept json
// @Description 密码登录(手机号或邮箱) API
// @Param param body types.LoginReq true "请求参数"
// @Success 200 {object} types.LoginResp "成功响应"
// @Router /api/login [post]
func Login(c *fiber.Ctx) error {
	reqParams := types.LoginReq{}
	db := database.DB
	err := c.BodyParser(&reqParams)
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "request params error", "请求参数错误"))
	}
	if strings.TrimSpace(reqParams.LoginPwd) == "" {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "Please input a password", "请输入密码"))
	}
	var user *model.User
	phone := fmt.Sprintf("%v%v", reqParams.Area, strings.TrimSpace(reqParams.Phone))
	if reqParams.LoginType == "2" {
		if strings.TrimSpace(reqParams.Email) == "" {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "email is empty", "邮箱不能为空"))
		}
		user, err = model.SelectUserByEmail(db, strings.TrimSpace(reqParams.Email))
	} else {
		if strings.TrimSpace(reqParams.Phone) == "" {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "phone is empty", "手机号不能为空"))
		}

		user, err = model.SelectUserByPhone(db, phone)
	}

	if err != nil {
		config.Log.Errorf("[Login] query user error, Phone: %s, Email: %s,\n", phone, reqParams.Email)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "login fail", "登录失败"))
		}
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "please register", "首次使用请注册"))
	}

	if user == nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "please register", "首次使用请注册"))
	}

	deLoginPwd, err := pkg.DecryptData(user.LoginPwd, config.LOGIN_FILE_KEY, config.LOGIN_P12_PATH)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "query user not exist", "请先设置登录密码"))
		}
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "login password error", ""))
	}
	// 确认密码正确
	if deLoginPwd != strings.TrimSpace(reqParams.LoginPwd) {
		//if deLoginPwd != reqParams.LoginPwd {
		config.Log.Errorf("pwd error, Phone: %s, Email: %s, DePwd: %s, Pwd: %s\n", reqParams.Phone, reqParams.Email, deLoginPwd, reqParams.LoginPwd)
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "login password error", "登录密码错误"))
	}
	// 更新token
	returnT := pkg.RandomString(64)
	token := returnT + ":" + strconv.FormatInt(time.Now().Unix(), 10)
	user.Token = token
	if err = model.UserRefreshToken(db, int64(user.ID), token); err != nil {
		return c.JSON(pkg.MessageResponse(config.TOKEN_FAIL, "db UserRefreshAppToken err", "更新token失败"))
	}

	return c.JSON(pkg.SuccessResponse(types.LoginResp{
		Token: returnT,
	}))
}

// UserInfo
// @Tags Api端
// @Summary 查询账户信息
// @Accept json
// @Description 查账户信息 API
// @Param token header string true "token" default(test_hash_token_user_id_1)
// @Success 200 {object} types.UserInfoResp "成功响应"
// @Router /api/user/info [post]
func UserInfo(c *fiber.Ctx) error {
	userId := c.Locals(config.LOCAL_USERID_INT64).(int64)

	db := database.DB
	//user, err := model.SelectUserByUserId(db.Preload("Balance"), userId)
	user, err := model.SelectUserByUserId(db, userId)
	if err != nil {
		config.Log.Errorf("[UserInfo] get user info error=%v \n", err.Error())
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "Get user exceptions", "获取用户异常"))
	}

	userInfoResp := types.UserInfoResp{
		Username: user.Username,
		Email:    user.Email,
		Phone:    user.Phone,
	}

	return c.JSON(pkg.SuccessResponse(userInfoResp))
}
