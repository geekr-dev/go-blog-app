package middleware

import (
	"fmt"
	"time"

	"github.com/geekr-dev/go-blog-app/global"
	"github.com/geekr-dev/go-blog-app/pkg/app"
	"github.com/geekr-dev/go-blog-app/pkg/email"
	"github.com/geekr-dev/go-blog-app/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	defaultMailer := email.NewEmail(&email.SMTPInfo{
		Host:     global.EmailConfig.Host,
		Port:     global.EmailConfig.Port,
		IsSSL:    global.EmailConfig.IsSSL,
		UserName: global.EmailConfig.UserName,
		Password: global.EmailConfig.Password,
		From:     global.EmailConfig.From,
	})
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				global.Logger.WithCallersFrames().Errorf(c, "panic recover err: %v", err)
				//  出现异常给管理员发送邮件
				err := defaultMailer.SendMail(
					[]string{"geekr@hey.com"},
					fmt.Sprintf("异常抛出，发生时间：%d", time.Now().Unix()),
					fmt.Sprintf("错误信息: %v", err),
				)
				if err != nil {
					global.Logger.Panicf(c, "mail.SendMail err: %v", err)
				}

				app.NewResponse(c).ToErrorResponse(errcode.ServerError)
				c.Abort()
			}
		}()
		c.Next()
	}
}
