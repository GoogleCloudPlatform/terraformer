package ali_mns

import (
	"bytes"
	"encoding/xml"
	"io"
)

type MNSDecoder interface {
	Decode(reader io.Reader, v interface{}) (err error)
	DecodeError(bodyBytes []byte, resource string) (decodedError error, err error)

	Test() bool
}

type aliMNSDecoder struct {
}

type batchOpDecoder struct {
	v interface{}
}

func NewAliMNSDecoder() MNSDecoder {
	return &aliMNSDecoder{}
}

func (p *aliMNSDecoder) Test() bool {
	return false
}

func (p *batchOpDecoder) Test() bool {
	return true
}

func (p *aliMNSDecoder) Decode(reader io.Reader, v interface{}) (err error) {
	decoder := xml.NewDecoder(reader)
	err = decoder.Decode(&v)

	return
}

func (p *aliMNSDecoder) DecodeError(bodyBytes []byte, resource string) (decodedError error, err error) {
	bodyReader := bytes.NewReader(bodyBytes)
	errResp := ErrorResponse{}

	decoder := xml.NewDecoder(bodyReader)
	err = decoder.Decode(&errResp)
	if err == nil {
		decodedError = ParseError(errResp, resource)
	}
	return
}

func NewBatchOpDecoder(v interface{}) MNSDecoder {
	return &batchOpDecoder{v: v}
}

func (p *batchOpDecoder) Decode(reader io.Reader, v interface{}) (err error) {
	decoder := xml.NewDecoder(reader)
	err = decoder.Decode(&v)

	if err == io.EOF {
		err = nil
	}

	return
}

func (p *batchOpDecoder) DecodeError(bodyBytes []byte, resource string) (decodedError error, err error) {
	bodyReader := bytes.NewReader(bodyBytes)

	decoder := xml.NewDecoder(bodyReader)
	err = decoder.Decode(&p.v)
	if err != nil {
		bodyReader.Seek(0, 0)
		errResp := ErrorResponse{}
		err = decoder.Decode(&errResp)
		if err == nil {
			decodedError = ParseError(errResp, resource)
		}
	} else {
		decodedError = ERR_MNS_BATCH_OP_FAIL.New()
	}
	return
}
