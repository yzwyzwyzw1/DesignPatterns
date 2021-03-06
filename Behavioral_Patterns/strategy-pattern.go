package Behavioral_Patterns



type Operator interface {
	Apply(int, int) int
}

//包装器
type Operation struct {
	operator Operator
}

func (op *Operation) Operate(left, right int) int {
	return op.operator.Apply(left, right)
}

//加法接口实现类
type Addition struct {}

func (add *Addition) Apply(left, right int) int {
	return left + right
}

//乘法接口实现类对象
type Multiplication struct {}

func (mu *Multiplication) Apply(left, right int) int {
	return left * right
}


//创建策略生成器，传入策略，获取具体的策略方法
func CreateOpration(operator Operator) Operation {
	return Operation{operator}
}
