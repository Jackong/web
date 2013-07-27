/**
 * User: Jackong
 * Date: 13-7-21
 * Time: 下午4:36
 */
package test

import (
	"githud.com/Jackong/web/ctrl"
	"githud.com/Jackong/web/mapper"
	"githud.com/Jackong/web/io"
	"githud.com/Jackong/web/method"
)

type test struct {
	ctrl.Ctrl
}

func init() {
	xa := &test{}
	xa.Init()
	xa.Rules[method.GET] = ctrl.Rule{Param: "ok", Required: true}
	mapper.Set("/test", xa)
}

func (this *test) Read(ctx *io.Context) {
	ctx.Output.Set("test", ctx.Input.Get("ok"))
}