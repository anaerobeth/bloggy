package grifts

import (
	"github.com/anaerobeth/bloggy/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
