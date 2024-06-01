package main

import (
	"fmt"
	"math"
)

var builtin_consts = map[string]Float64Object{
	"e":        {math.E},
	"pi":       {math.Pi},
	"phi":      {math.Phi},
	"sqrt_2":   {math.Sqrt2},
	"sqrt_e":   {math.SqrtE},
	"sqrt_pi":  {math.SqrtPi},
	"sqrt_phi": {math.SqrtPhi},
	"ln_2":     {math.Ln2},
	"ln_10":    {math.Ln10},
}

type BuiltinFn func(args ...Float64Object) (Float64Object, error)

type BuiltinFnDescriptor struct {
	Pointer int
	Fn      BuiltinFn
}

var zero = Float64Object{0}

func arg_len_err(name string, expected, found int) error {
	if expected != found {
		return fmt.Errorf("function '%s' expects exactly %d arguments but got %d", name, expected, found)
	}

	return nil
}

type BuiltinFnList map[string]BuiltinFnDescriptor

func (list BuiltinFnList) GetPointer(pointer int) *BuiltinFn {
	for _, descriptor := range list {
		if descriptor.Pointer == pointer {
			return &descriptor.Fn
		}
	}

	return nil
}

var builtin_fns = BuiltinFnList{
	"abs": {
		Pointer: 0,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("abs", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Abs(args[0].Value)}, nil
		},
	},
	"acos": {
		Pointer: 1,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("acos", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Acos(args[0].Value)}, nil
		},
	},
	"acosh": {
		Pointer: 2,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("acosh", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Acosh(args[0].Value)}, nil
		},
	},
	"asin": {
		Pointer: 3,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("asin", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Asin(args[0].Value)}, nil
		},
	},
	"asinh": {
		Pointer: 4,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("asinh", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Asinh(args[0].Value)}, nil
		},
	},
	"atan": {
		Pointer: 5,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("atan", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Atan(args[0].Value)}, nil
		},
	},
	"atanh": {
		Pointer: 6,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("atanh", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Atanh(args[0].Value)}, nil
		},
	},
	"cbrt": {
		Pointer: 7,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("cbrt", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Cbrt(args[0].Value)}, nil
		},
	},
	"ceil": {
		Pointer: 8,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("ceil", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Ceil(args[0].Value)}, nil
		},
	},
	"cos": {
		Pointer: 9,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("cos", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Cos(args[0].Value)}, nil
		},
	},
	"cosh": {
		Pointer: 10,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("cosh", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Cosh(args[0].Value)}, nil
		},
	},
	"exp": {
		Pointer: 11,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("exp", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Exp(args[0].Value)}, nil
		},
	},
	"expm1": {
		Pointer: 12,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("expm1", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Expm1(args[0].Value)}, nil
		},
	},
	"floor": {
		Pointer: 13,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("floor", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Floor(args[0].Value)}, nil
		},
	},
	"log": {
		Pointer: 14,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("log", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Log(args[0].Value)}, nil
		},
	},
	"log10": {
		Pointer: 15,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("log10", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Log10(args[0].Value)}, nil
		},
	},
	"log1p": {
		Pointer: 16,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("log1p", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Log1p(args[0].Value)}, nil
		},
	},
	"log2": {
		Pointer: 17,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("log2", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Log2(args[0].Value)}, nil
		},
	},
	"round": {
		Pointer: 18,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("round", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Round(args[0].Value)}, nil
		},
	},
	"sin": {
		Pointer: 19,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("sin", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Sin(args[0].Value)}, nil
		},
	},
	"sinh": {
		Pointer: 20,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("sinh", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Sinh(args[0].Value)}, nil
		},
	},
	"sqrt": {
		Pointer: 21,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("sqrt", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Sqrt(args[0].Value)}, nil
		},
	},
	"tan": {
		Pointer: 22,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("tan", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Tan(args[0].Value)}, nil
		},
	},
	"tanh": {
		Pointer: 23,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("tanh", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Tanh(args[0].Value)}, nil
		},
	},
	"trunc": {
		Pointer: 24,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("trunc", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{math.Trunc(args[0].Value)}, nil
		},
	},
	"rad": {
		Pointer: 25,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("rad", 1, len(args)); err != nil {
				return zero, err
			}

			rad := args[0].Value * (math.Pi / 180)
			return Float64Object{rad}, nil
		},
	},
	"deg": {
		Pointer: 26,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("deg", 1, len(args)); err != nil {
				return zero, err
			}

			deg := args[0].Value / (math.Pi / 180)
			return Float64Object{deg}, nil
		},
	},
	"neg": {
		Pointer: 27,
		Fn: func(args ...Float64Object) (Float64Object, error) {
			if err := arg_len_err("deg", 1, len(args)); err != nil {
				return zero, err
			}

			return Float64Object{0 - args[0].Value}, nil
		},
	},
}
