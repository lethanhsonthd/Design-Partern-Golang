package main

import "fmt"

type Strategy interface {
	doOperation(num1, num2 int) int
}
type OperationAdd struct {
}

func (op *OperationAdd) doOperation(num1, num2 int) int {
	return num1 + num2
}

type OperationSub struct {
}

func (op *OperationSub) doOperation(num1, num2 int) int {
	return num1 - num2
}

type OperationMul struct {
}

func (op *OperationMul) doOperation(num1, num2 int) int {
	return num1 * num2
}

type Context struct {
	Strategy Strategy
}

func (c *Context) Context(s Strategy) {
	c.Strategy = s
}
func (c *Context) executeStrategy(num1, num2 int) int {
	return c.Strategy.doOperation(num1, num2)
}
func main() {
	tmp2 := Context{&OperationMul{}}
	fmt.Println(tmp2.executeStrategy(2, 3))
}
