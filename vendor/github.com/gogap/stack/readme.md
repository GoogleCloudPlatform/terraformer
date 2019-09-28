# stack
--
    import "github.com/facebookgo/stack"

Package stack provides utilities to capture and pass around stack traces.

This is useful for building errors that know where they originated from, to
track where a certain log event occured and so on.

The package provides stack.Multi which represents a sequence of stack traces.
Since in Go we return errors they don't necessarily end up with a single useful
stack trace. For example an error may be going thru a channel across goroutines,
in which case we may want to capture a stack trace in both (or many) goroutines.
stack.Multi in turn is made up of stack.Stack, which is a set of stack.Frames.
Each stack.Frame contains the File/Line/Name (function name). All these types
implement a pretty human readable String() function.

The GOPATH is stripped from the File location. Look at the StripGOPATH function
on instructions for how to embed to GOPATH into the binary for when deploying to
production and the GOPATH environment variable may not be set. The package name
is stripped from the Name of the function since it included in the File
location.

## Usage

#### func  StripGOPATH

```go
func StripGOPATH(f string) string
```
StripGOPATH strips the GOPATH prefix from the file path f. In development, this
will be done using the GOPATH environment variable. For production builds, where
the GOPATH environment will not be set, the GOPATH can be included in the binary
by passing ldflags, for example:

    GO_LDFLAGS="$GO_LDFLAGS -X github.com/facebookgo/stack.gopath $GOPATH"
    go install "-ldflags=$GO_LDFLAGS" my/pkg

#### func  StripPackage

```go
func StripPackage(n string) string
```
StripPackage strips the package name from the given Func.Name.

#### type Frame

```go
type Frame struct {
	File string
	Line int
	Name string
}
```

Frame identifies a file, line & function name in the stack.

#### func  Caller

```go
func Caller(skip int) Frame
```
Caller returns a single Frame for the caller. The argument skip is the number of
stack frames to ascend, with 0 identifying the caller of Callers.

#### func (Frame) String

```go
func (f Frame) String() string
```
String provides the standard file:line representation.

#### type Multi

```go
type Multi struct {
}
```

Multi represents a number of Stacks. This is useful to allow tracking a value as
it travels thru code.

#### func  CallersMulti

```go
func CallersMulti(skip int) *Multi
```
CallersMulti returns a Multi which includes one Stack for the current callers.
The argument skip is the number of stack frames to ascend, with 0 identifying
the caller of CallersMulti.

#### func (*Multi) Add

```go
func (m *Multi) Add(s Stack)
```
Add the given Stack to this Multi.

#### func (*Multi) AddCallers

```go
func (m *Multi) AddCallers(skip int)
```
AddCallers adds the Callers Stack to this Multi. The argument skip is the number
of stack frames to ascend, with 0 identifying the caller of Callers.

#### func (*Multi) Stacks

```go
func (m *Multi) Stacks() []Stack
```
Stacks returns the tracked Stacks.

#### func (*Multi) String

```go
func (m *Multi) String() string
```
String provides a human readable multi-line stack trace.

#### type Stack

```go
type Stack []Frame
```

Stack represents an ordered set of Frames.

#### func  Callers

```go
func Callers(skip int) Stack
```
Callers returns a Stack of Frames for the callers. The argument skip is the
number of stack frames to ascend, with 0 identifying the caller of Callers.

#### func (Stack) String

```go
func (s Stack) String() string
```
String provides the standard multi-line stack trace.
