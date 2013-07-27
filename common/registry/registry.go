/**
 * User: Jackong
 * Date: 13-7-26
 * Time: 下午10:29
 */
package registry

type IRegistry interface {
	Get(key string) (interface{})
	Set(key string, value interface{})
}

type Registry (map[string] interface{})

func (this Registry) Get(key string) (interface{}) {
	return this[key]
}

func (this Registry) Set(key string, value interface{}) {
	this[key] = value
}
