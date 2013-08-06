/**
 * User: Jackong
 * Date: 13-8-4
 * Time: 下午1:23
 */
package ctrl

import (
	"fmt"
	"testing"
	"time"
	"reflect"
)

type mock struct {
}
func (this *mock) do(in *input) bool{
	in.get("test")
	fmt.Println("do something")
	return false
}

type input struct {
	inerr chan bool
}

func (this *input)get(name string) {
	fmt.Println("getting", name, "error")
	this.inerr <- true
	fmt.Println("over")
	/*
		//0. never close
		tmp := make(chan bool)
		defer close(tmp)
		<- tmp
	 */
	/*
		//1. will continue when close
		tmp := make(chan bool)
		time.Sleep(1 * time.Second)
		close(tmp)
		<- tmp
	*/
	/*
		//2. will continue when close
		time.Sleep(1 * time.Second)
		<- this.inerr
	*/
	fmt.Println("nerver")
}

func wrap(m *mock) {
	inerr := make(chan bool)
	defer close(inerr)
	in := &input{inerr}
	go m.do(in)
	stop := <- inerr
	if stop {
		fmt.Println("stop")
		return
	}
	fmt.Println("normal")
}


func TestNormal(t *testing.T) {
	m := &mock{}
	wrap(m)
	time.Sleep(time.Second)
}

func TestReflect(t *testing.T) {
	m := &mock{}
	f := reflect.TypeOf(m).Method(0)
	t.Log(f)
}

type inputErr string
func (this inputErr) Error() string {
	return string(this)
}
func pk() {
	panic(inputErr("err"))
}
func TestPanic(t *testing.T) {
	defer func() {
		if e := recover(); e != nil  {
			t.Log(e)
		}
	}()
	pk()
}
