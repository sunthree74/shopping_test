package handler

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/sunthree74/shopping_test/helper"
	"github.com/sunthree74/shopping_test/interfaces"
	"github.com/sunthree74/shopping_test/model"
	"github.com/sunthree74/shopping_test/structs"
	"github.com/sunthree74/shopping_test/structs/request"
	"net/http"
	"time"
)

type middleware struct {
	jwtHandler *jwt.GinJWTMiddleware
}

var identityKey string = "Email"

func InitializeMiddleware(userUsecase interfaces.UserUsecase) (*middleware, error) {
	jwtHandler, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:          "SIMPLESHOPPINGTEST",
		Key:            []byte("SIMPLESHOPPINGTEST"),
		SendCookie:     true,
		CookieHTTPOnly: true,
		MaxRefresh:     time.Hour,
		Timeout:        time.Hour * 24 * 30,
		IdentityKey:    identityKey,
		TokenLookup:    "header: Authorization, query: token, cookie: jwt",
		TokenHeadName:  "SIMPLESHOPPINGTEST",
		TimeFunc:       time.Now,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*structs.CachedUser); ok {
				return jwt.MapClaims{
					identityKey: v.Email,
					"Email":     v.Email,
					"Name":      v.Name,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			if claims["Name"] == nil {
				claims["Name"] = ""
			}

			return &structs.CachedUser{
				Email: claims["Email"].(string),
				Name:  claims["Name"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var (
				user    model.User
				message error
				success bool
			)

			method := c.Param("method")
			success = true

			if method == "register" {
				var registerForm request.ValidateRegister
				if err := c.ShouldBind(&registerForm); err != nil {
					return "", jwt.ErrMissingLoginValues
				}

				user.Email = registerForm.Email
				user.Name = registerForm.Name
				user.Password = registerForm.Password
				err := userUsecase.Create(c, &user)
				if err != nil {
					message = jwt.ErrFailedAuthentication
					success = false
				}
			} else if method == "login" {
				var loginForm request.ValidateLogin
				if err := c.ShouldBind(&loginForm); err != nil {
					return "", jwt.ErrMissingLoginValues
				}

				var err error
				user, err = userUsecase.Login(c, loginForm.Email, loginForm.Password)
				if err != nil {
					message = jwt.ErrFailedAuthentication
					success = false
				}
			}

			if success == true {
				return &structs.CachedUser{
					Email: user.Email,
					Name:  user.Name,
				}, nil
			} else {
				message = jwt.ErrMissingLoginValues
			}

			return nil, message
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			v, ok := data.(*structs.CachedUser)

			usr, _ := userUsecase.FindByEmail(c.Request.Context(), v.Email)
			var id = usr.ID

			if ok && id != 0 {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			if message == jwt.ErrMissingLoginValues.Error() {
				code = http.StatusBadRequest
				message = "Lengkapi email/password anda."
			} else if message == jwt.ErrFailedAuthentication.Error() {
				code = http.StatusUnauthorized
				message = "email/password tidak sesuai."
			} else {
				code = http.StatusUnauthorized
				message = "Silahkan login kembali untuk melanjutkan aktivitas anda."
			}

			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	})

	if err != nil {
		return nil, err
	}

	return &middleware{jwtHandler: jwtHandler}, nil
}

func (m *middleware) JwtAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := m.jwtHandler.ParseToken(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Mohon login untuk melanjutkan aktivitas"})
			return
		}

		claims := jwt.ExtractClaimsFromToken(token)
		exp := claims["exp"].(float64)
		expUnix := int64(exp)

		expireDate := time.Unix(expUnix, 0)
		if expireDate.Before(time.Now()) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "Mohon login untuk melanjutkan aktivitas"})
			return
		}

		email := claims["Email"].(string)
		ctx.Set("Email", email)
		ctx.Set("JWT_PAYLOAD", claims)
		ctx.Next()
	}
}

func (m *middleware) UserAuth() *jwt.GinJWTMiddleware {
	return m.jwtHandler
}

func VerifyHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		authBearer := c.Request.Header.Get("Authorization")
		if authBearer == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Authorization header is required"})
			return
		}

		if !helper.IsAuthTokenValid(authBearer) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"success": false, "message": "Invalid authorization header"})
			return
		}

		c.Next()
	}
}
