package Creational_Patterns

import (
	"fmt"
	"testing"
)

func TestExample_Clone(t *testing.T) {
	origin := pNew("origin object")
	current := origin.Clone()
	fmt.Println(current.Description)
	//	订制clone对象的属性
	current.Description = "current object"
	fmt.Println(current.Description)
}

