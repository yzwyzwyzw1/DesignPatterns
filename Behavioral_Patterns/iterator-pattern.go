package Behavioral_Patterns

import "fmt"

type Iterator struct {
	index int
	Container
}
//迭代器方法
func (i *Iterator) Next() Visitor1 {
	fmt.Println(i.index)
	visitor := i.list[i.index]
	i.index += 1
	return visitor
}

func (i *Iterator) HasNext() bool {
	if i.index >= len(i.list) {
		return false
	}
	return true
}



//创建容器
//Visitor1类型
type Container struct {
	list []Visitor1
}

func (c *Container) Add(visitor Visitor1) {
	c.list = append(c.list, visitor)
}

func (c *Container) Remove(index int) {
	if index < 0 || index > len(c.list) {
		return
	}
	c.list = append(c.list[:index], c.list[index+1:]...)
}



//创建Visitor接口
type Visitor1 interface {
	Visit()
}

//创建具体的visitor对象
//Teacher 接口实现类
type Teacher struct {}
func (t *Teacher) Visit() {
	fmt.Println("this is teacher visitor")
}

//Analysis 接口实现类
type Analysis struct {}
func (a *Analysis) Visit() {
	fmt.Println("this is analysis visitor")
}


//工厂方法创建迭代器
func NewIterator() *Iterator {
	return &Iterator{
		index: 0,
		Container: Container{},
	}
}
