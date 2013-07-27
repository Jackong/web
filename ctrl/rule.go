/**
 * User: Jackong
 * Date: 13-7-14
 * Time: 下午9:49
 */
package ctrl

import (
	"net/http"
	"regexp"
)

/**
 * 1.Default means not Required
 * 2.Match may be Default(Not-Required)
 */
type Rule struct {
	Param string
	Required bool
	Match string
	Default interface {}
	Tips string
}

func (this *Rule) Value(req *http.Request) interface{} {
	value := req.FormValue(this.Param)
	if this.Match != "" {
		if match, _ := regexp.MatchString(this.Match, value); !match {
			if this.Default == nil {
				return nil
			}
			return this.Default
		}
		return value
	}

	if value == "" {
		if this.Required {
			return nil
		}
		if this.Default != nil {
			return this.Default
		}
	}
	return value
}
