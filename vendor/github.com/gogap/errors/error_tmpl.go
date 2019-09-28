package errors

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"text/template"
	"time"

	"github.com/gogap/stack"
)

type ErrCodeTemplate struct {
	namespace string
	code      uint64
	template  string
}

func T(code uint64, template string) ErrCodeTemplate {
	return TN(DefaultErrorNamespace, code, template)
}

func TN(namespace string, code uint64, template string) ErrCodeTemplate {
	key := fmt.Sprintf("%s:%d", namespace, code)
	if _, exist := errCodeDefined[key]; exist {
		strErr := fmt.Sprintf("error code %s already exist", key)
		panic(strErr)
	} else {
		errCodeDefined[key] = true
	}
	return ErrCodeTemplate{code: code, namespace: namespace, template: template}
}

func (p *ErrCodeTemplate) New(v ...Params) (err ErrCode) {
	params := Params{}
	if v != nil {
		for _, param := range v {
			for pn, pv := range param {
				params[pn] = pv
			}
		}
	}

	strCode := fmt.Sprintf("%s#%d", p.namespace, p.code)

	stack := stack.CallersDeepth(1, 5)

	errId := fmt.Sprintf("%s.%d.%s.%d", p.namespace, p.code, p.template, time.Now().UnixNano())

	crcErrId := crc32.ChecksumIEEE([]byte(errId))

	strCRCErrId := fmt.Sprintf("%0X", crcErrId)

	if len(strCRCErrId) > 7 {
		strCRCErrId = strCRCErrId[0:7]
	}

	var tpl *ErrCodeTemplate = p

	if t, e := template.New(strCode).Parse(tpl.template); e != nil {
		strErr := fmt.Sprintf("parser error template failed, namespace: %s, code: %d, error: %s", tpl.namespace, tpl.code, e.Error())
		err = &errorCode{
			err: Error{
				ID:        strCRCErrId,
				Namespace: ErrcodeNamespace,
				Code:      ErrcodeParseTmplError,
				Message:   strErr,
			},
			stackTrace: stack.String(),
			context:    make(map[string]interface{}),
		}
		return
	} else {
		var buf bytes.Buffer
		if e := t.Execute(&buf, params); e != nil {
			strErr := fmt.Sprintf("execute template failed, namespace: %s code: %d, error: %s", tpl.namespace, tpl.code, e.Error())
			err = &errorCode{
				err: Error{
					ID:        strCRCErrId,
					Namespace: ErrcodeNamespace,
					Code:      ErrcodeExecTmpleError,
					Message:   strErr,
				},
				stackTrace: stack.String(),
				context:    make(map[string]interface{})}
			return
		} else {
			err = &errorCode{
				err: Error{
					ID:        strCRCErrId,
					Namespace: tpl.namespace,
					Code:      tpl.code,
					Message:   buf.String(),
				},
				stackTrace: stack.String(),
				context:    make(map[string]interface{})}
			return
		}
	}
}

func (p *ErrCodeTemplate) IsEqual(err error) bool {
	if e, ok := err.(ErrCode); ok {
		if e.Code() == p.code && e.Namespace() == p.namespace {
			return true
		}
	}
	return false
}
