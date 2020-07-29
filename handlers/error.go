package handlers

import "github.com/gkzy/gow"

func NotFoundHandler(c *gow.Context) {
	data := gow.H{
		"title": "404 page not found",
	}
	c.HTML("error.html", data)
}
