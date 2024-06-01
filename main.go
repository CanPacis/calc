package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// file, _ := os.ReadFile("./data.calc")
	input := strings.Join(os.Args[1:], " ")
	parser := NewParser([]byte(input), "calc")
	expr, err := parser.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	compiler := NewCompiler(expr)
	_, err = compiler.Compile()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	vm := NewVmFromCompiler(compiler)
	result, err := vm.Run()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result)

	// js.Global().Set("build", build())
	// js.Global().Set("exec", exec())
	// <-make(chan struct{})
}

// type build_resp struct {
// 	Compiled []int  `json:"compiled"`
// 	Error    string `json:"error"`
// }

// func build_err(err string) build_resp {
// 	return build_resp{
// 		Compiled: []int{},
// 		Error:    err,
// 	}
// }

// func build() js.Func {
// 	return js.FuncOf(func(this js.Value, args []js.Value) any {
// 		if len(args) == 0 {
// 			m, _ := json.Marshal(build_err("no input"))
// 			return string(m)
// 		}

// 		input := args[0].String()
// 		parser := NewParser([]byte(input), "main.calc")
// 		expr, err := parser.Parse()
// 		if err != nil {
// 			m, _ := json.Marshal(build_err(err.Error()))
// 			return string(m)
// 		}
// 		compiler := NewCompiler(expr)
// 		compiled, err := compiler.Compile()
// 		if err != nil {
// 			m, _ := json.Marshal(build_err(err.Error()))
// 			return string(m)
// 		}
// 		result := []int{}

// 		for _, b := range compiled {
// 			result = append(result, int(b))
// 		}

// 		m, _ := json.Marshal(build_resp{
// 			Compiled: result,
// 			Error:    "",
// 		})
// 		return string(m)
// 	})
// }

// type exec_resp struct {
// 	Result float64 `json:"result"`
// 	Error  string  `json:"error"`
// }

// func exec_err(err string) exec_resp {
// 	return exec_resp{
// 		Result: 0,
// 		Error:  err,
// 	}
// }

// func exec() js.Func {
// 	return js.FuncOf(func(this js.Value, args []js.Value) any {
// 		if len(args) == 0 {
// 			m, _ := json.Marshal(exec_err("no input"))
// 			return string(m)
// 		}

// 		input := args[0].String()
// 		data := []byte{}
// 		json.Unmarshal([]byte(input), &data)

// 		vm, err := NewVm(data)
// 		if err != nil {
// 			m, _ := json.Marshal(exec_err(err.Error()))
// 			return string(m)
// 		}

// 		ret, err := vm.Run()
// 		if err != nil {
// 			m, _ := json.Marshal(exec_err(err.Error()))
// 			return string(m)
// 		}

// 		m, _ := json.Marshal(exec_resp{
// 			Result: ret,
// 			Error:  "",
// 		})
// 		return string(m)
// 	})
// }

// func check_err(err error) {
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// }

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("no action provided")
// 		os.Exit(1)
// 	}

// 	if len(os.Args) < 3 {
// 		fmt.Println("no input file provided")
// 		os.Exit(1)
// 	}

// 	switch os.Args[1] {
// 	case "build":
// 		p, err := filepath.Abs(os.Args[2])
// 		check_err(err)

// 		name := path.Base(p)
// 		name = strings.TrimSuffix(name, path.Ext(p))
// 		name = fmt.Sprintf("%s.cb", name)

// 		data, err := build(p)
// 		check_err(err)

// 		os.WriteFile(name, data, 0644)
// 	case "run":
// 		p, err := filepath.Abs(os.Args[2])
// 		check_err(err)

// 		ret, err := run(p)
// 		check_err(err)

// 		fmt.Println(ret)
// 	case "calc":
// 		p, err := filepath.Abs(os.Args[2])
// 		check_err(err)

// 		data, _ := os.ReadFile(p)
// 		check_err(err)
// 		parser := NewParser(data, p)
// 		expr, err := parser.Parse()
// 		check_err(err)

// 		compiler := NewCompiler(expr)
// 		_, err = compiler.Compile()
// 		check_err(err)

// 		vm := NewVmFromCompiler(compiler)

// 		ret, err := vm.Run()
// 		check_err(err)

// 		fmt.Println(ret)
// 	default:
// 		fmt.Println("unknown action", os.Args[1])
// 		os.Exit(1)
// 	}
// }

// func build(p string) ([]byte, error) {
// 	data, err := os.ReadFile(p)
// 	if err != nil {
// 		return []byte{}, err
// 	}
// 	parser := NewParser(data, p)

// 	expr, err := parser.Parse()
// 	fmt.Println(expr)
// 	if err != nil {
// 		return []byte{}, err
// 	}

// 	compiler := NewCompiler(expr)

// 	serialized, err := compiler.Compile()
// 	if err != nil {
// 		return []byte{}, err
// 	}

// 	return serialized, nil
// }

// func run(p string) (float64, error) {
// 	data, err := os.ReadFile(p)
// 	if err != nil {
// 		return 0, err
// 	}

// 	vm, err := NewVm(data)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return vm.Run()
// }
