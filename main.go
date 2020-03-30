package main

import (
    "net/http"
    "strings"
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    mid "github.com/yellia1989/tex-web/backend/middleware"
    "github.com/yellia1989/tex-web/backend/api"
)

func httpErrorHandler(err error, c echo.Context) {
    he, ok := err.(*echo.HTTPError)
    if ok {
        if he.Internal != nil {
            if herr, ok := he.Internal.(*echo.HTTPError); ok {
                he = herr
            }
        }
    } else {
        he = &echo.HTTPError{
            Code:    http.StatusInternalServerError,
            Message: http.StatusText(http.StatusInternalServerError),
        }
    }

    code := he.Code
    message := he.Message

    c.Logger().Error("error:"+he.Error())

    // Send response
    if !c.Response().Committed {
        if c.Request().Method == http.MethodHead { // Issue #608
            err = c.NoContent(he.Code)
        } else {
            // 格式化错误消息
            path := c.Request().URL.Path
            if strings.HasPrefix(path, "/api") {
                err = c.JSON(http.StatusOK, map[string]interface{}{
                    "code": code,
                    "msg": message,
                })
            } else {
                // 没有登陆的话重定位到登陆
                if mid.GetUserId(c) == 0 {
                    err = c.Redirect(http.StatusMovedPermanently, "/login.html")
                } else {
                    redirect := "/500.html"
                    switch code {
                    case http.StatusForbidden:redirect = "/403.html"
                    case http.StatusNotFound:redirect = "/404.html"
                    }
                    err = c.Redirect(http.StatusMovedPermanently, redirect)
                }
            }
        }
        if err != nil {
            c.Logger().Error(err)
        }
    }
}

func main() {
    // Echo instance
    e := echo.New()
    e.HTTPErrorHandler = httpErrorHandler

    // Middleware
    e.Pre(middleware.RemoveTrailingSlash())
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.Use(mid.NewContext(), mid.RequireLogin(), mid.RequireAuth())

    e.Static("/", "front/pages")
    e.Static("/lib", "front/lib")
    e.Static("/css", "front/css")
    e.Static("/js", "front/js")
    e.Static("/images", "front/images")

    api.RegisterHandler(e.Group("/api"))

    // Start server
    e.Logger.Fatal(e.Start(":8080"))
}
