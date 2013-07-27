/**
 * User: Jackong
 * Date: 13-7-27
 * Time: 上午10:19
 */
package io

import (
	"githud.com/Jackong/web/io/i"
	"githud.com/Jackong/web/io/o"
)

type Context struct {
	*i.Input
	*o.Output
}

func New(input *i.Input, output *o.Output) *Context {
	return &Context{Input: input, Output: output}
}
