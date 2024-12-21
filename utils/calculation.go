package utils

import (
	"errors"
	"strconv"
	"strings"
)

func Calc(expression string) (float64, error) {
	expression = strings.ReplaceAll(expression, " ", "")
	if expression == "" {
		return 0, errors.New("empty expression")
	}

	opePriorities := map[rune]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	opeStack := []rune{}
	numStack := []float64{}

	applyOperation := func() error {
		if len(opeStack) == 0 || len(numStack) < 2 {
			return errors.New("not enough data for calculation")
		}

		b := numStack[len(numStack)-1]
		a := numStack[len(numStack)-2]
		numStack = numStack[:len(numStack)-2]

		op := opeStack[len(opeStack)-1]
		opeStack = opeStack[:len(opeStack)-1]
		var result float64

		switch op {
		case '+':
			result = a + b
		case '-':
			result = a - b
		case '*':
			result = a * b
		case '/':
			if b == 0 {
				return errors.New("can't divide by zero")
			}

			result = a / b
		default:
			return errors.New("unknown operation")
		}

		numStack = append(numStack, result)

		return nil
	}

	i := 0
	for i < len(expression) {

		if expression[i] >= '0' && expression[i] <= '9' {
			start := i
			for i < len(expression) && (expression[i] >= '0' && expression[i] <= '9' || expression[i] == '.') {
				i++
			}

			num, err := strconv.ParseFloat(expression[start:i], 64)
			if err != nil {
				return 0, errors.New("wrong number in expression")
			}
			numStack = append(numStack, num)

			continue
		}

		if expression[i] == '(' {
			opeStack = append(opeStack, '(')
		} else if expression[i] == ')' {
			for len(opeStack) > 0 && opeStack[len(opeStack) - 1] != '(' {
				if err := applyOperation(); err != nil {
					return 0, err
				}
			}

			if len(opeStack) == 0 {
				return 0, errors.New("unbalanced brackets")
			}

			opeStack = opeStack[:len(opeStack)-1]
		} else {
			for len(opeStack) > 0 && opePriorities[opeStack[len(opeStack)-1]] >= opePriorities[rune(expression[i])] {
				if err := applyOperation(); err != nil {
					return 0, err
				}
			}
			opeStack = append(opeStack, rune(expression[i]))
		}

		i++
	}

	for len(opeStack) > 0 {
		if err := applyOperation(); err != nil {
			return 0, err
		}
	}

	if len(numStack) != 1 {
		return 0, errors.New("expressions failure")
	}

	return numStack[0], nil
}