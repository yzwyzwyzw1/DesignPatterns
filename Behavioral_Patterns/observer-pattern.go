package Behavioral_Patterns

import "fmt"

type Event struct {
	Info string
}
//设计观察者和被观察者接口
type Observer interface {
	Receive(event Event)
}

type Notifier interface {
	Register(observer Observer)
	Remove(observer Observer)
	Notify(event Event)
}

//通过具体对象来实现接口
//投资人观察者
type InvestorObserver struct {
	Name string
}

func (invester *InvestorObserver) Receive(event Event) {
	fmt.Printf("%s 收到事件通知 %s\n", invester.Name, event.Info)
}
//股票被观察者
type ShareNotifier struct {
	Price float64
	oblist []Observer //注册链表
}

func (share *ShareNotifier) Register(observer Observer) {
	share.oblist = append(share.oblist, observer)
}
//移除观察者
func (share *ShareNotifier) Remove(observer Observer) {
	if len(share.oblist) == 0 {
		return
	}
	for i, ob := range share.oblist {
		if ob == observer {
			share.oblist = append(share.oblist[:i], share.oblist[i+1:]...)
		}
	}
}

func (share *ShareNotifier) Notify(event Event) {
	for _, ob := range share.oblist {
		ob.Receive(event)
	}
}

func NewEvent() Event {
	return Event{Info: "价格变动通知"}
}

func NewInvestorObserver(name string) *InvestorObserver {
	return &InvestorObserver{Name: name}
}

func NewShareNotifier(price float64) *ShareNotifier {
	return &ShareNotifier{Price: price}
}

