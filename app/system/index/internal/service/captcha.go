package service

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
)

// 验证码管理服务
var (
	Captcha = captchaService{}
)

type captchaService struct{}

var (
	captchaStore  = base64Captcha.DefaultMemStore
	captchaDriver = newDriver()
)

func newDriver() *base64Captcha.DriverString {
	driver := &base64Captcha.DriverString{
		Height:          44,
		Width:           126,
		NoiseCount:      5,
		ShowLineOptions: base64Captcha.OptionShowSineLine | base64Captcha.OptionShowSlimeLine | base64Captcha.OptionShowHollowLine,
		Length:          4,
		Source:          "1234567890",
		Fonts:           []string{"wqy-microhei.ttc"},
	}
	return driver.ConvertFonts()
}

// 创建验证码，直接输出验证码图片内容到HTTP Response.
func (s *captchaService) NewAndStore(c *gin.Context, name string) error {
	captcha := base64Captcha.NewCaptcha(captchaDriver, captchaStore)
	// 生成Key，内容，答案
	captchaStoreKey, content, answer := captcha.Driver.GenerateIdQuestionAnswer()
	// 绘制二进制项
	item, _ := captcha.Driver.DrawCaptcha(content)
	// c.Session.Set(name, captchaStoreKey)
	// 设置验证码id的数字
	captcha.Store.Set(captchaStoreKey, answer)
	// 生成base64 图像字符串
	base64 := item.EncodeB64string()

	// 存储session
	session := sessions.DefaultMany(c, "Captcha")
	session.Set(name, captchaStoreKey)
	err := session.Save()

	c.JSON(http.StatusOK, base64)
	return err
}

// 校验验证码，并清空缓存的验证码信息
func (s *captchaService) VerifyAndClear(c *gin.Context, name string, value string) bool {
	session := sessions.DefaultMany(c, "Captcha")
	defer session.Delete(name)
	captchaStoreKey := fmt.Sprintf("%v", session.Get(name))
	return captchaStore.Verify(captchaStoreKey, value, true)
}
