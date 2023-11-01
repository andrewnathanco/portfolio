package index

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Link struct {
	Name string
	URL  string
}

func GetIndex(c echo.Context) error {
	links := []Link{
		{
			Name: "Music",
			URL:  "/music",
		},
		{
			Name: "Projects",
			URL:  "/projects",
		},
		{
			Name: "Blog",
			URL:  "/blog",
		},
		{
			Name: "LinkedIn",
			URL:  "https://www.linkedin.com/in/andrew-cohen-089972156/",
		},
		{
			Name: "Instagram",
			URL:  "https://www.instagram.com/andrewnathanco/",
		},
		{
			Name: "Github",
			URL:  "https://github.com/andrewnathanco",
		},
	}

	return c.Render(http.StatusOK, "index.html", links)
}
