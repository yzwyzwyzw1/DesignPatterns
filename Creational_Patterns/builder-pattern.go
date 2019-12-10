package Creational_Patterns


type Wheel int

type Vehicle struct{
	Wheels Wheel
	Seats int
	Structure string
}

//定义建造者接口
type Builder interface{
	SetWheels() Builder
	SetSeats()  Builder
	SetStructure() Builder
	GetVehicle() Vehicle
}

//定义结构体
type Director struct{
	builder Builder
}

func (director *Director)Construct(){
	director.builder.SetWheels().SetSeats().SetStructure()//一次链式调用
}

func (director *Director)SetBuilder(builder Builder){
	director.builder=builder
}

//定义car对象用来实现接口,作为接口的一个实现类
type Car struct{
	vehicle Vehicle
}

//实现接口
func (car *Car)SetWheels() Builder{
	car.vehicle.Wheels=4
	return car
}

func (car *Car)SetSeats() Builder{
	car.vehicle.Seats=4
	return car
}

func (car *Car)SetStructure() Builder{
	car.vehicle.Structure="Car"
	return car
}

func (car *Car)GetVehicle() Vehicle{
	return car.vehicle
}