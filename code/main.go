package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	ErrCouldNotParseTempaltes = fmt.Errorf("could not parse templates")
)

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	templates, err := template.ParseGlob("view/*.html")
	if err != nil {
		log.Panic("could not parse templates: %w", err)
	}

	for _, t := range templates.Templates() {
		fmt.Println(t.Name())
	}

	e.Renderer = &TemplateRenderer{
		templates: templates,
	}

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/dist", "./dist")

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	e.GET("/music", func(c echo.Context) error {
		return c.Render(http.StatusOK, "music.html", nil)
	})

	e.Logger.Fatal(e.Start(":1323"))
}
