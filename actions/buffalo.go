package actions

import "github.com/gobuffalo/buffalo"

func BuffaloHandler(c buffalo.Context) error {
	name := c.Param("name")
	if name == "" {
		name = "World"
	}
	c.Set("name", name)
	return c.Render(200, r.HTML("hello.html"))
}
