package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"portfolio/controller/index"

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
	echo.NotFoundHandler = func(c echo.Context) error {
		return c.Render(http.StatusNotFound, "404.html", nil)
	}

	e := echo.New()
	templates, err := template.ParseGlob("view/*.html")
	if err != nil {
		log.Panic("could not parse templates: %w", err)
	}

	e.Renderer = &TemplateRenderer{
		templates: templates,
	}

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/static", "./static")

	e.GET("/", index.GetIndex)
	e.GET("/music", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "https://open.spotify.com/artist/4CpIdORYoIiB04WjB6wTNq?si=G-EoIRtSTpyxBigjrmOYYA")
	})

	e.GET("/projects", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "https://mural.andrewnathan.net")
	})
	
	e.GET("/blog", func(c echo.Context) error {
		return c.Redirect(http.StatusPermanentRedirect, "https://joinpickup.com/blog")
	})

	e.Logger.Fatal(e.Start(":1324"))
}
