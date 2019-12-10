package Creational_Patterns

type Example struct {
	Description string
}

//实现Clone
func (e *Example) Clone() *Example {
	res := *e
	return &res
}

func pNew(des string) *Example{
	return &Example{
		Description: des,
	}
}
