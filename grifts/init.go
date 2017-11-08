package grifts

import (
	"github.com/ECAllen/continuous_improvement/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
