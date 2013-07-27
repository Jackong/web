/**
 * User: Jackong
 * Date: 13-7-21
 * Time: 下午6:41
 */
package o

import "net/http"

type Output struct {
	Writer http.ResponseWriter
	name string
	output interface {}
}

func (this *Output) Get() (string, interface {}) {
	return this.name, this.output
}

func (this *Output) Set(name string, output interface {}) {
	this.name = name
	this.output = output
}
