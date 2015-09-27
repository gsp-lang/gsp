package core

import (
	"fmt"
	"reflect"
)

type Any interface{}

func MOD(a, b Any) int {
	var n, m int

	if isInt(a) {
		n = a.(int)
	} else if isFloat(a) {
		n = int(a.(float64))
	} else {
		panic("need int/float argument to mod!")
	}

	if isInt(b) {
		m = b.(int)
	} else if isFloat(a) {
		m = int(b.(float64))
	} else {
		panic("need int/float argument to mod!")
	}

	return n % m
}

func ADD(args ...Any) float64 {
	var sum float64 = 0

	for i := 0; i < len(args); i++ {
		switch n := args[i]; {
		case isInt(n):
			sum += float64(n.(int))
		case isFloat(n):
			sum += n.(float64)
		}
	}

	return sum
}

func SUB(args ...Any) float64 {
	var result float64
	if isInt(args[0]) {
		result = float64(args[0].(int))
	} else if isFloat(args[0]) {
		result = args[0].(float64)
	} else {
		panic("need int/float for SUB")
	}

	for i := 1; i < len(args); i++ {
		switch n := args[i]; {
		case isInt(n):
			result -= float64(n.(int))
		case isFloat(n):
			result -= n.(float64)
		}
	}

	return result
}

func MUL(args ...Any) float64 {
	var prod float64 = 1

	for i := 0; i < len(args); i++ {
		switch n := args[i]; {
		case isInt(n):
			prod *= float64(n.(int))
		case isFloat(n):
			prod *= n.(float64)
		}
	}

	return prod
}
func DIV() {}

// TODO: can only compare ints and slice lens for now.
func LT(args ...Any) bool {
	if len(args) < 2 {
		panic("can't compare less than 2 values!")
	}

	for i := 0; i < len(args)-1; i++ {
		var n float64
		if isInt(args[i]) {
			n = float64(args[i].(int))
		} else if isFloat(args[i]) {
			n = args[i].(float64)
		} else {
			panic("you can't compare that!")
		}

		var m float64
		if isInt(args[i+1]) {
			m = float64(args[i+1].(int))
		} else if isFloat(args[i+1]) {
			m = args[i+1].(float64)
		} else {
			panic("you can't compare that!")
		}

		if n >= m {
			return false
		}
	}

	return true
}

// TODO: can only compare ints and slice lens for now.
func GT(args ...Any) bool {
	if len(args) < 2 {
		panic("can't compare less than 2 values!")
	}

	for i := 0; i < len(args)-1; i++ {
		var n float64
		if isInt(args[i]) {
			n = float64(args[i].(int))
		} else if isFloat(args[i]) {
			n = args[i].(float64)
		} else {
			panic("you can't compare that!")
		}

		var m float64
		if isInt(args[i+1]) {
			m = float64(args[i+1].(int))
		} else if isFloat(args[i+1]) {
			m = args[i+1].(float64)
		} else {
			panic("you can't compare that!")
		}

		if n <= m {
			return false
		}
	}

	return true
}

func EQ(args ...Any) bool {
	if len(args) < 2 {
		panic("can't compare less than 2 values!")
	}

	for i := 0; i < len(args)-1; i++ {
		var n float64
		var nNil bool
		var nBool bool
		if isInt(args[i]) {
			n = float64(args[i].(int))
		} else if isFloat(args[i]) {
			n = args[i].(float64)
		} else if isBool(args[i]) {
			nBool = true
			nBoolVal := args[i].(bool)
			n = 0
			if nBoolVal {
				n = 1
			}
		} else if isNil(args[i]) {
			nNil = true
		} else {
			panic("you can't compare that!")
		}

		var m float64
		var mNil bool
		var mBool bool
		if isInt(args[i+1]) {
			m = float64(args[i+1].(int))
		} else if isFloat(args[i+1]) {
			m = args[i+1].(float64)
		} else if isBool(args[i+1]) {
			mBool = true
			mBoolVal := args[i+1].(bool)
			m = 0
			if mBoolVal {
				m = 1
			}
		} else if isNil(args[i+1]) {
			mNil = true
		} else {
			panic("you can't compare that!")
		}

		if mNil || nNil {
			return mNil && nNil
		}

		if mBool || nBool {
			mBoolVal := m != 0
			nBoolVal := n != 0
			return mBoolVal == nBoolVal
		}

		if n != m {
			return false
		}
	}

	return true
}

// greater than or equal
func GTEQ(args ...Any) bool {
	if GT(args...) || EQ(args...) {
		return true
	}

	return false
}

// less than or equal
func LTEQ(args ...Any) bool {
	if LT(args...) || EQ(args...) {
		return true
	}

	return false
}

// Source: http://stackoverflow.com/q/13476349/1507139
func isNil(a interface{}) bool {
	defer func() { recover() }()
	return a == nil || reflect.ValueOf(a).IsNil()
}

func isFloat(n Any) bool {
	_, ok := n.(float64)
	return ok
}

func isInt(n Any) bool {
	_, ok := n.(int)
	return ok
}

func isBool(n Any) bool {
	_, ok := n.(bool)
	return ok
}

func Get(args ...Any) Any {
	if len(args) != 2 && len(args) != 3 {
		panic(fmt.Sprintf("get needs 2 or 3 arguments %d given.", len(args)))
	}

	if len(args) == 2 {
		if a, ok := args[1].([]Any); ok {
			return a[args[0].(int)]
		} else if a, ok := args[1].(string); ok {
			return a[args[0].(int)]
		} else {
			panic("arguments to get must include slice/vector/string")
		}
	} else {
		if a, ok := args[2].([]Any); ok {
			if args[1].(int) == -1 {
				return a[args[0].(int):]
			}

			return a[args[0].(int):args[1].(int)]
		} else if a, ok := args[2].(string); ok {
			if args[1].(int) == -1 {
				return a[args[0].(int):]
			}

			return a[args[0].(int):args[1].(int)]
		} else {
			panic("arguments to get must include slice/vector/string")
		}
	}
}

func Len(arg Any) Any {
	if a, ok := arg.([]Any); ok {
		return len(a)
	} else if a, ok := arg.(string); ok {
		return len(a)
	} else {
		panic("argument to let must include slice/vector/string")
	}
}
