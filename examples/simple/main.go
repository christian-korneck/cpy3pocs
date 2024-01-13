package main

import (
	python3 "github.com/go-python/cpy3"
)

var pycode = `
import sys
print(sys.version)
`

func main() {
	defer python3.Py_Finalize()
	python3.Py_Initialize()
	// python3.PyRun_SimpleString("print('hello world')")
	python3.PyRun_SimpleString(pycode)

}
