/**
 * User: Jackong
 * Date: 13-7-21
 * Time: 下午6:40
 */
package i

import (
	"net/http"
	"githud.com/Jackong/web/common/registry"
)

type Input struct {
	Req *http.Request
	registry.Registry
}
func New(req *http.Request) *Input{
	return &Input{Req: req, Registry: registry.Registry{}}
}

func (this *Input) Get(key string) interface {} {
	if value := this.Registry.Get(key); value != nil {
		return value
	}
	return this.Req.FormValue(key)
}
