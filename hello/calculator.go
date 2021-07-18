package hello

import (
	"fmt"
	"strconv"
)

type Expression struct {
	op   operator
	a, b Expresser
}

type Expresser interface {
	calculate() int
}

type SimpleValueExpression struct {
	value int
}

type operator int

const plus operator = 0
const minus operator = 1

func calc(e Expresser) int {
	if e == nil {
		return 0
	}
	return e.calculate()
}

func (e *Expression) calculate() int {
	if e == nil {
		return 0
	}

	switch e.op {
	case plus:
		return calc(e.a) + calc(e.b)
	case minus:
		return calc(e.a) - calc(e.b)
	}
	return 0
}

func (e *Expression) consolidate() {
	e.a = &SimpleValueExpression{calc(e)}
	e.b = nil
	e.op = plus
}

func (e *SimpleValueExpression) calculate() int {
	if e == nil {
		return 0
	}
	return e.value
}

func seekForFullInt(in string, start int) (integerValue, endIndex int) {
	endIndex = start
	for ; endIndex < len(in); endIndex++ {
		if in[endIndex] < '0' || in[endIndex] > '9' {
			break
		}
	}

	integerValue, _ = strconv.Atoi(in[start:endIndex])
	endIndex--
	return
}

func stringToExpression(in string) *Expression {
	result := Expression{}
	stack := []*Expression{&result}
	index := 0
	for len(stack) > 0 {
		expr := stack[len(stack)-1]
		stack[0] = nil
		stack = stack[:len(stack)-1]
		popStack := false

		for ; index < len(in); index++ {
			switch in[index] {
			case '(':
				nextExpr := &Expression{}
				if expr.a == nil {
					expr.a = nextExpr
				} else {
					expr.b = nextExpr
				}
				stack = append(stack, expr)
				expr = nextExpr
				continue

			case ')':
				popStack = true

			case '+':
				if expr.a == nil || expr.b == nil {
					expr.op = plus
					continue
				}
				expr.consolidate()

			case '-':
				if expr.a == nil {
					nextExpr := &Expression{
						op: minus,
						a:  &SimpleValueExpression{},
						b:  nil,
					}
					expr.a = nextExpr
					stack = append(stack, expr)
					expr = nextExpr
					continue
				}
				expr.consolidate()
				expr.op = minus

			default:
				var integerValue int
				integerValue, index = seekForFullInt(in, index)
				if expr.a == nil {
					expr.a = &SimpleValueExpression{integerValue}
					continue
				}
				if expr.b == nil {
					expr.b = &SimpleValueExpression{integerValue}
					continue
				}
				// malformed?
			}
			if popStack {
				break
			}
		}
	}

	return &result
}

func Calculator(expr string) int {
	return stringToExpression(expr).calculate()
}

func TestCalculator() {
	fmt.Println(Calculator("-(3+(2-1))"))
	// {-, {+, 3, {-, 2, 1}}}
	fmt.Println(Calculator("-3+(2-1)"))
	fmt.Println(Calculator("(-3+(2-1))"))
	fmt.Println(Calculator("-3+(2-(1+10+12-2-2))"))
	// {+, {-, 3, 0}, {-, 2, 1}}
}
