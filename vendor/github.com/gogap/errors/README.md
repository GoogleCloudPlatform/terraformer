errors
======
For replacement offical package of `errors`

for more detials, see the example


#### `main.go`
```go
package main

import (
    "fmt"

    "github.com/gogap/errors"
)

var (
    ErrParseTest  = errors.TN("GOGAP", 10001, "hello {{.param1}}")
    ErrParseTest2 = errors.TN("GOGAP", 10002, "test error")
    ErrStackTest  = errors.TN("GOGAP", 10003, "call stack test")
)

func main() {

    e1 := ErrParseTest.New(errors.Params{"param1": "world"}).WithContext("key", "value")

    e1.Append("I am append errors")

    fmt.Println("always equal while errors append:", ErrParseTest.IsEqual(e1))

    fmt.Println("ErrParseTest = ErrParseTest2 :", ErrParseTest2.IsEqual(e1))

    data, _ := e1.(errors.ErrCode).Marshal()

    fmt.Println(string(data))

    fmt.Println(e1.Error())

    stack3Error := call_3()

    fmt.Println(stack3Error.(errors.ErrCode).StackTrace())

    fmt.Println(e1.(errors.ErrCode).FullError())
}

func call_1() error {
    return call_2()
}
func call_2() error {
    return call_3()
}
func call_3() error {
    return ErrStackTest.New()
}

```

#### example output
```bash
$ go run main.go

always equal while errors append: true
ErrParseTest = ErrParseTest2 : false
{"id":"6F93654","namespace":"GOGAP","code":10001,"message":"hello world; I am append errors."}

hello world; I am append errors.

github.com/gogap/errors/example/main.go:45 call_3
github.com/gogap/errors/example/main.go:31 main
/usr/local/go/src/runtime/proc.go:188      main
/usr/local/go/src/runtime/asm_amd64.s:1998 goexit

Id: GOGAP#10001:6F93654
Error:
hello world; I am append errors.
Context:
{"key":"value"}
StackTrace:
github.com/gogap/errors/example/main.go:17 main
/usr/local/go/src/runtime/proc.go:188      main
/usr/local/go/src/runtime/asm_amd64.s:1998 goexit
```