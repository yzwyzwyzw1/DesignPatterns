package Creational_Patterns

import "fmt"

//定义工厂接口
type Factory interface {
	CreateProduct() Product
}

type Product interface {
	Describe()
}

//实现接口，成为接口的实现对象
type ConcreteProduct struct {
	Name string
}

func (conproduct *ConcreteProduct) Describe() {
	fmt.Println(conproduct.Name)
}


//具体工厂
type ConCreteFactory struct {}

//工厂类的方法： 返回产品接口类型的实现类
func (confactory *ConCreteFactory) CreateProduct() Product {
	return &ConcreteProduct{Name: "KG"}
}