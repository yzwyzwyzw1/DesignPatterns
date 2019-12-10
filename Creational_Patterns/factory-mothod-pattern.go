package Creational_Patterns

import "errors"

type Kind int


//定义Cash常量和 Credit常量
const (
	Cash Kind = 1 << iota
	Credit
)

type Payment interface {
	Pay(money float32) error
}

type CashPay struct {
	Balance float32
}
func (cash *CashPay) Pay(money float32)error{
    if cash.Balance <0 || cash.Balance<money{
    	return errors.New("balance not enouch")
	}
	cash.Balance -= money
	return nil
}

type CreditPay struct {
	Balance float32
}
func (credit *CreditPay)Pay(money float32) error{
	if credit.Balance<0|| credit.Balance<money{
		return errors.New("balance not enough")
	}
	credit.Balance -=money
	return nil
}

// 返回Payment接口，由于Cash和Credit对象都实现得了Payment都实现了该接口的Pay方法，因此，可以作为实现类返回
func GeneratePayment(k Kind,balance float32)(Payment,error){
	switch k {
	case Cash:
		cash := new(CashPay)
		cash.Balance = balance
		return cash, nil
	case Credit:
		credit :=new(CreditPay)
		credit.Balance= balance
		return credit,nil
	default:
		return  nil,errors.New("Payment do not support this")
	}
}
