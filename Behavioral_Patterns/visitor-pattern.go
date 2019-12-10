package Behavioral_Patterns

import "fmt"

//visitor 接口
type Visitor interface {
	Visit()
}
//具体Visitor对象

//访问者A 的接口实现类
type ConcreteVisitorA struct {
	Name string
}
func (conV *ConcreteVisitorA) Visit() {
	fmt.Println("this is visitor A")
}

//访问者B 的接口实现类
type ConcreteVisitorB struct {
	Name string
}
func (conV *ConcreteVisitorB) Visit() {
	fmt.Println("this is visitor B")
}


//创建元素接口
type Element interface {
	Accept(visitor Visitor)
}


//ElementA实现Element的接口实现类
type ElementA struct {}

func (e *ElementA) Accept(visitor Visitor)  {
	visitor.Visit()
}



//Element容器
type ElementContainer struct {
	list []Element
}

//实现容器的元素的添加和移除
func (container *ElementContainer) Add(element Element)  {
	if container == nil || element == nil {
		return
	}
	container.list = append(container.list, element)
}

func (container *ElementContainer) Delete(element Element) {
	for i, val := range container.list {
		if val == element {
			container.list = append(container.list[:i], container.list[i+1:]...)
			break
		}
	}
}
