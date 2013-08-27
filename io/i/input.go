/**
 * User: Jackong
 * Date: 13-7-21
 * Time: 下午6:40
 */
package i

import (
	"net/http"
	"regexp"
)

type Input struct {
	Req *http.Request
}

func (this *Input) Get(name, pattern string, defo interface {}) interface {} {
	value := this.Req.FormValue(name)
	if value == "" {
		this.Req.ParseForm()
		value = this.Req.Form.Get(name)
	}
	if value == "" {
		if defo != nil {
			return defo
		}
		panic(InputError("Invalid param: " + name))
	}
	if pattern == "" {
		return value
	}
	if match, _ := regexp.MatchString(pattern, value); !match {
		panic(InputError("Invalid param: " + name))
	}
	return value
}

func (this *Input) Default(name string, defo interface {}) interface {} {
	return this.Get(name, "", defo)
}

func (this *Input) Pattern(name, pattern string) interface {} {
	return this.Get(name, pattern, nil)
}

func (this *Input) Required(name string) interface {}{
	return this.Pattern(name, "")
}
