/**
 * User: Jackong
 * Date: 13-7-26
 * Time: 下午11:02
 */
package registry

import (
	"testing"
)

type Registry map[string] interface{}

func (this Registry) Get(key string) (interface{}) {
	return this[key]
}

func (this Registry) Set(key string, value interface{}) {
	this[key] = value
}

type concrete struct {
	Registry
}

func (this concrete) Get(key string) (iSome) {
	if val := this.Registry.Get(key); val != nil {
		return val.(iSome)
	}
	return nil
}
type iSome interface {
	Op() bool
}

type some struct {

}
func (this *some) Op() bool{
	return true
}

func TestReg(t *testing.T) {
	c := concrete{Registry{}}
	s := &some{}
	c.Set("ww", s)
	if val := c.Get("ww");  val == nil || !val.Op() {
		t.Fail()
	}
	if v := c.Get("w"); v != nil {
		t.Fail()
	}
}
