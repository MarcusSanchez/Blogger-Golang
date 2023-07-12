package handlers

import (
	"github.com/gofiber/fiber/v2"
	"strings"
)

// global items

type Post struct {
	Title   string
	Content string
}

var posts []Post

//functions

func GetRoot(c *fiber.Ctx) error {
	return c.Render("home", fiber.Map{
		"posts": posts,
	})
}

func GetCreate(c *fiber.Ctx) error {
	return c.SendFile("public/create.html")
}

func PostCreate(c *fiber.Ctx) error {
	currPost := Post{
		strings.Clone(c.FormValue("title")),
		strings.Clone(c.FormValue("content")),
	}
	posts = append(posts, currPost)
	return c.Redirect("/")
}
