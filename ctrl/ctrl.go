/**
 * User: Jackong
 * Date: 13-7-14
 * Time: 下午2:57
 */
package ctrl

import (
	"githud.com/Jackong/web/io/i"
	"githud.com/Jackong/web/io"
)

type ICtrl interface {
	Check(*i.Input) (bool, string)
	Create(*io.Context)
	Read(*io.Context)
	Update(*io.Context)
	Delete(*io.Context)
}

type Ctrl struct {
	Rules map[string] Rule
}

func (this *Ctrl) Init() {
	this.Rules = make(map[string] Rule)
}

func (this *Ctrl) Check(input *i.Input) (bool, string) {
	if rule, ok := this.Rules[input.Req.Method]; ok {
		value := rule.Value(input.Req)
		if value == nil {
			if rule.Tips == "" {
				rule.Tips = "Error param: " + rule.Param
			}
			return false, rule.Tips
		}
		input.Set(rule.Param, value)
	}
	return true, ""
}

func (this *Ctrl) Create(*io.Context) {
}

func (this *Ctrl) Read(*io.Context) {
}

func (this *Ctrl) Update(*io.Context) {
}

func (this *Ctrl) Delete(*io.Context) {
}
