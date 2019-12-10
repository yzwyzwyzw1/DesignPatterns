package Structural_Patterns

import "fmt"


//IObject接口
type IObject interface {
	ObjDo(action string)
}

// IObject的接口实现类Object
type Object struct {
	action string
}
func (obj *Object) ObjDo(action string) {
	fmt.Printf("I can %s", action)
}


//IObject的接口实现类ProxyObject,该类封装了接口实现类Object
type ProxyObject struct {
	object *Object
}
func (p *ProxyObject) ObjDo(action string) {
	if p.object == nil {
		p.object = new(Object)
	}
	if action == "run" {
		p.object.ObjDo(action)
	}
}
