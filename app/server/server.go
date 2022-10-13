package server

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"text/template"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"github.com/yamazaki-ko/koTutorial/app/db"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func RunServer() {
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "time=${time_custom}, method=${method}, uri=${uri}, status=${status}\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))

	e.Static("/css", "public/views/css")

	e.GET("/loginForm", LoginForm)
	e.POST("/login", Login)

	e.Logger.SetLevel(log.INFO)

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func LoginForm(c echo.Context) error {
	var u db.User
	return c.Render(http.StatusOK, "loginForm.html", u)
}

func Login(c echo.Context) error {
	LoginId := c.FormValue("loginId")
	Password := c.FormValue("password")

	ret := db.AuthLogin(LoginId, Password)

	if !ret {
		var u db.User
		u.LoginId = LoginId
		u.Password = Password
		return c.Render(http.StatusOK, "loginForm.html", u)
	}

	return c.String(http.StatusOK, "Success!")
}
