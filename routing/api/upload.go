package api

import (
	"fiber-template/config"
	"fiber-template/pkg"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

// Upload
// @Tags Api端
// @Param token header string true "token" default(test_hash_token_user_id_1)
// @Summary 上传图片 返回url
// @Description 上传图片 返回url API file_type
// @Param file formData file true "文件" format(file)
// @Success 200 {string}  string "返回url"
// @Router /api/upload [post]
func Upload(c *fiber.Ctx) error {
	// 检查请求数据流的大小
	if c.Request().Header.ContentLength() > (40 * 1024 * 1024) { // 20MB
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "The maximum size of the image is 40mb", "图片不可大于40mb"))
	}
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "from filer err", ""))
	}
	open, err := file.Open()
	if err != nil {
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "Open filer err", ""))
	}
	defer open.Close()

	fileName := fmt.Sprintf("%v_%v", time.Now().Format("20060102150405"), file.Filename)

	url, err := pkg.AwsUpload(open, fileName)
	if err != nil {
		config.Log.Errorf("AwsUpload err: %v", err)
		return c.JSON(pkg.MessageResponse(config.MESSAGE_FAIL, "Upload exception", "上传异常"))
	}

	return c.JSON(pkg.SuccessResponse(map[string]interface{}{"url": url}))
}
