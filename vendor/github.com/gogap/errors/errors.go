package errors

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"
)

type Params map[string]interface{}

const (
	ErrcodeNamespace      = "ERRCODE"
	DefaultErrorNamespace = "ERR"
)

const (
	ErrcodeParseTmplError = 1
	ErrcodeExecTmpleError = 2
)

var (
	errorTemplate  = make(map[string]*ErrCodeTemplate)
	errCodeDefined = make(map[string]bool)
)

type ErrCode interface {
	Id() string
	Code() uint64
	Namespace() string
	Error() string
	StackTrace() string
	Context() ErrorContext
	FullError() error
	Append(err ...interface{}) ErrCode
	WithContext(k string, v interface{}) ErrCode
	Marshal() ([]byte, error)
}

var (
	_ ErrCode = (*errorCode)(nil)
)

type Error struct {
	ID        string `json:"id"`
	Namespace string `json:"namespace"`
	Code      uint64 `json:"code"`
	Message   string `json:"message"`
}

type errorCode struct {
	err        Error
	stackTrace string
	context    map[string]interface{}
	errors     []string
}

func NewErrorCode(id string, code uint64, namespace string, message string, stackTrace string, context map[string]interface{}) ErrCode {
	e := &errorCode{
		err:        Error{ID: id, Namespace: namespace, Code: code, Message: message},
		stackTrace: stackTrace,
		context:    context,
	}

	if e.context == nil {
		e.context = make(map[string]interface{})
	}

	e.err.ID = id
	e.err.Code = code
	e.err.Message = message
	e.err.Namespace = namespace

	return e
}

func (p *errorCode) Id() string {
	return p.err.ID
}

func (p *errorCode) Code() uint64 {
	return p.err.Code
}

func (p *errorCode) Namespace() string {
	return p.err.Namespace
}

func (p *errorCode) Error() string {
	msg := p.err.Message

	if len(p.errors) > 0 {
		if msg != "" {
			msg += "; "
		}

		msg += strings.Join(p.errors, "; ") + "."
	}

	return msg
}

func (p *errorCode) FullError() error {
	errLines := make([]string, 1)

	errLines[0] = fmt.Sprintf("Id: %s#%d:%s", p.Namespace(), p.Code(), p.Id())

	errLines = append(errLines, "Error:")
	errLines = append(errLines, p.Error())
	errLines = append(errLines, "Context:")
	errLines = append(errLines, p.Context().String())
	errLines = append(errLines, "StackTrace:")
	errLines = append(errLines, p.stackTrace)
	return New(strings.Join(errLines, "\n"))
}

func Unmarshal(data []byte) ErrCode {
	decoder := json.NewDecoder(bytes.NewBuffer(data))
	decoder.UseNumber()
	v := errorCode{}
	if decoder.Decode(&v.err) != nil {
		return nil
	}

	return &v
}

func (p *errorCode) Context() ErrorContext {
	return p.context
}

func (p *errorCode) StackTrace() string {
	return p.stackTrace
}

func (p *errorCode) Append(err ...interface{}) ErrCode {
	if err != nil {
		for _, e := range err {
			switch ev := e.(type) {
			case ErrCode:
				{
					str := fmt.Sprintf("(%s#%d:%s) %s", ev.Namespace(), ev.Code(), ev.Id(), ev.Error())
					p.errors = append(p.errors, str)
				}
			case error:
				{
					p.errors = append(p.errors, ev.Error())
				}
			default:
				p.errors = append(p.errors, fmt.Sprintf("%v", e))
			}
		}
	}
	return p
}

func (p *errorCode) WithContext(key string, value interface{}) ErrCode {
	p.context[key] = value
	return p
}

func (p *errorCode) Marshal() (data []byte, err error) {

	e := Error{
		ID:        p.err.ID,
		Code:      p.err.Code,
		Message:   p.Error(),
		Namespace: p.err.Namespace,
	}

	return json.Marshal(e)
}

func IsErrCode(err error) bool {
	_, ok := err.(ErrCode)
	return ok
}

func (p errorCode) MarshalJSON() ([]byte, error) {
	str := "\"" + fmt.Sprintf("(%s#%d:%s) %s", p.err.Namespace, p.err.Code, p.err.ID, p.err.Message) + "\""
	return []byte(str), nil
}

func (p errorCode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	str := fmt.Sprintf("(%s#%d:%s) %s", p.err.Namespace, p.err.Code, p.err.ID, p.err.Message)
	e.EncodeElement(str, start)
	return nil
}
