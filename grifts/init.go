package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/masato25/my104job_board/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
