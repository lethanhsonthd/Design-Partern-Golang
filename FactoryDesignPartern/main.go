package main

import "fmt"

type Shape interface {
	Draw()
}

type Rectangle struct {
	Name string
}

func (r Rectangle) Draw() {
	fmt.Println("Draw rectangle")
}

type Circle struct {
	Name string
}

func (c Circle) Draw() {
	fmt.Println("Draw circle")
}

type ShapeFactory struct {
	Option string
}

func (sf ShapeFactory) NewShapeFactory() Shape {
	if sf.Option == "rectangle" {
		return Rectangle{Name: "asd"}
	}
	if sf.Option == "circle" {
		return Circle{Name: "circle"}
	}
	return nil
}

func main() {
	tmp := ShapeFactory{Option: "circle"}
	rec := tmp.NewShapeFactory()
	rec.Draw()
	tmp = ShapeFactory{Option: "rectangle"}
	rec = tmp.NewShapeFactory()
	rec.Draw()
}
