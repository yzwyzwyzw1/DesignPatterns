package Creational_Patterns

import "testing"

func TestConCreteFactory_CreateProduct(t *testing.T) {
	conFactory := &ConCreteFactory{}  //创建工厂

	product := conFactory.CreateProduct()//工厂创建商品接口

	conProduct := product.(*ConcreteProduct)//指定接口的实现类

	if conProduct.Name != "KG" {
		t.Error("abstract factory can not create the concreate product")
	}
}
