package Behavioral_Patterns

import "fmt"

//上下文 保存状态值和状态接口的实现类对象
type Context1 struct {
	State 		ActionState
	HealthValue int
}

//账号的行为
func (a *Context1) View() {
	a.State.View()
}
func (a *Context1) Comment() {
	a.State.Comment()
}
func (a *Context1) Create() {
	a.State.Create()
}

func (a *Context1) SetHealth(value int) {
	a.HealthValue = value
	a.changestate()
}

//根据不同的状态值，获取不同的状态
func (a *Context1) changestate() {
	if a.HealthValue < 0 {
		a.State = &ClosedState{}
	} else if a.HealthValue > 10 {
		a.State = &NormalState{}
	} else if a.HealthValue <10 && a.HealthValue > 0 {
		a.State = &RestrictedState{}
	}
}
//简单工厂方式创建
func NewContext(health int) *Context1 {
	con := &Context1{HealthValue: health}
	con.changestate()
	return con
}



//因为含有多种状态，状态类型不确定，这里需要声明State接口，多态原则，在golang中通过接口实现多态性
type ActionState interface {
	View()
	Comment()
	Create()
}
/*
具体的状态，继承接口
*/
//Normal state
type NormalState struct {}

func (n *NormalState) View() {
	fmt.Println("view normal")
}
func (n *NormalState) Comment() {
	fmt.Println("comment normal")
}
func (n *NormalState) Create() {
	fmt.Println("create normal")
}
//Restricted State
type RestrictedState struct {}

func (n *RestrictedState) View() {
	fmt.Println("view Restricted")
}
func (n *RestrictedState) Comment() {
	fmt.Println("comment Restricted")
}
func (n *RestrictedState) Create() {
	fmt.Println("create Restricted")
}
//Closed State
type ClosedState struct {}

func (n *ClosedState) View() {
	fmt.Println("view closed")
}
func (n *ClosedState) Comment() {
	fmt.Println("comment closed")
}
func (n *ClosedState) Create() {
	fmt.Println("create closed")
}

