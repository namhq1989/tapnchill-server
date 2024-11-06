package appjwt

import (
	"github.com/labstack/echo/v4"
	"github.com/namhq1989/go-utilities/appcontext"
)

func (j JWT) RequireLoggedIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Get("ctx").(*appcontext.AppContext)
		ctx.SetUserID("66d9bfc85efebfd68b22b13f")
		ctx.SetClientID("66d9c5c21cb6cb622e2be8dd")
		ctx.SetTimezone("Asia/Ho_Chi_Minh")
		return next(c)

		// var (
		// 	ctx   = c.Get("ctx").(*appcontext.AppContext)
		// 	token = strings.TrimPrefix(c.Request().Header.Get("Authorization"), "Bearer ")
		// )
		//
		// claims, err := j.ParseAccessToken(ctx, token)
		// if claims == nil || err != nil {
		// 	ctx.Logger().Error("check access token error", err, appcontext.Fields{"token": token})
		// 	return c.JSON(http.StatusUnauthorized, echo.Map{
		// 		"data":    nil,
		// 		"code":    "unauthorized",
		// 		"message": "Unauthorized",
		// 	})
		// }
		//
		// ctx.SetUserID(claims.UserID)
		// ctx.SetClientID(claims.ClientID)
		// ctx.SetTimezone(claims.Timezone)
		// return next(c)
	}
}
