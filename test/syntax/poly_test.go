/**
 * User: Jackong
 * Date: 13-7-14
 * Time: 下午6:44
 */
package syntax

import "testing"

type A struct {
	v string
}
func (this *A) K() {
	this.v = "xxx"
}
func (this *A) F() string {
	return this.v
}
type B struct {
	A
}
func (this *B) K() {
	this.v = "www"
}
func TestPrivate(t *testing.T) {
	c := new(B)
	c.K()
	if v := c.F(); v == "xxx" {
		t.Error("error")
	}
}
