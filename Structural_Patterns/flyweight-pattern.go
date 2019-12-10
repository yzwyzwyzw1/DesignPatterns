package Structural_Patterns



//创建shape接口
type Shape1 interface {
	SetRadius(radius int)
	SetColor(color string)
}
//接口实现类Circle
type Circle struct {
	color string
	radius int
}

func (c *Circle) SetRadius(radius int) {
	c.radius = radius
}

func (c *Circle) SetColor(color string) {
	c.color = color
}

/*********************************************************/

//创建ShapeFactory
type ShapeFactory struct {
	circleMap map[string]Shape1  //用 map 存储这些对象
}

//GetCircle 对象不存在则创建
func (sh *ShapeFactory) GetCircle(color string) Shape1 {
	if sh.circleMap == nil {
		sh.circleMap = make(map[string]Shape1)
	}
	if shape, ok := sh.circleMap[color]; ok {
		return shape
	}
	circle := new(Circle)
	circle.SetColor(color)
	sh.circleMap[color] = circle
	return circle
}

