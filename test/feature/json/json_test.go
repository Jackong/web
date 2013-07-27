/**
 * User: Jackong
 * Date: 13-7-20
 * Time: 下午10:50
 */
package json

import (
	"testing"
	"encoding/json"
)

func TestObj(t *testing.T) {
	b, err := json.Marshal(struct {Obj string} {Obj: "str"})
	if err != nil {
		t.Error(err)
		return
	}
	bs := string(b)
	if bs != `{"Obj":"str"}` {
		t.Error(bs)
	}
	var v interface{}
	if err = json.Unmarshal(b, &v); err != nil {
		t.Error(err)
	}
	m := v.(map[string]interface {})
	if m["Obj"] != "str" {
		t.Error(m)
	}
}

func TestSlice(t *testing.T) {
	b, err := json.Marshal(struct{Slice []int} {Slice: []int{1, 2, 4, 5}})
	if err != nil {
		t.Error(err)
		return
	}
	bs := string(b)
	if bs != `{"Slice":[1,2,4,5]}` {
		t.Error(bs)
	}
}

func TestMap(t *testing.T) {
	b, err := json.Marshal(map[string]struct {
		Bool interface {}
		Float float32
		Null interface {}
	}{"A":{Bool: false, Float: 1.264, Null: nil}, "B":{Bool: true, Float: -1.254, Null: "null"}, "C":{Bool: nil, Float: 0.00, Null: "123"}})
	if err != nil {
		t.Error(err)
		return
	}
	bs := string(b)
	if bs != `{"A":{"Bool":false,"Float":1.264,"Null":null},"B":{"Bool":true,"Float":-1.254,"Null":"null"},"C":{"Bool":null,"Float":0,"Null":"123"}}` {
		t.Error(bs)
	}
}
