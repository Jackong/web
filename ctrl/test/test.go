/**
 * User: Jackong
 * Date: 13-7-21
 * Time: 下午4:36
 */
package test

import (
	"github.com/Jackong/web/ctrl"
	"github.com/Jackong/web/mapper"
	"github.com/Jackong/web/io"
)

type test struct {
	ctrl.Ctrl
}

func init() {
	xa := &test{}
	mapper.Set("/test", xa)
}

func (this *test) Read(ctx *io.Context) {
	ctx.Output.Set("test", ctx.Input.Required("ok"))
}
