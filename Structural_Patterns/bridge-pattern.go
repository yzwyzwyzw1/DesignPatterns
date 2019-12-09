package Structural_Patterns

import "fmt"

//创建桥接接口
type SoftWare interface {
	Run()
}

//创建实现run方法的结构体
type Cpu struct {

}
func (c *Cpu) Run() {
	fmt.Println("this is cpu run")
}

type Storage struct {

}
func (s *Storage) Run() {
	fmt.Println("this is storage run")
}

//创建Shape struct（抽象部分）
type Shape struct {
	software SoftWare
}
//赋值实现桥接接口的具体结构体
func (s *Shape) SetSoftWare(soft SoftWare) {
	s.software = soft
}

//组合模式创建具体的struct
type Phone struct {
	shape Shape
}

func (p *Phone) SetShape(soft SoftWare) {
	p.shape.SetSoftWare(soft)
}

func (p *Phone) Print() {
	p.shape.software.Run()
}
