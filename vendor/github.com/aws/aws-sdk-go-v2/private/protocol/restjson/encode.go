package restjson

import (
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/json"
	"github.com/aws/aws-sdk-go-v2/private/protocol/rest"
)

// An Encoder provides encoding of the AWS RESTJSON protocol. This encoder combindes
// the JSON and REST encoders deligating to them for their associated targets.
//
// It is invalid to set a JSON and stream payload on the same encoder.
type Encoder struct {
	method      string
	reqEncoder  *rest.Encoder
	bodyEncoder *json.Encoder
	t           protocol.Target

	err error
}

// NewEncoder creates a new encoder for encoding the AWS RESTJSON protocol.
// The request passed in will be the base the path, query, and headers encoded
// will be set on top of.
func NewEncoder(req *http.Request) *Encoder {
	e := &Encoder{
		method:      req.Method,
		reqEncoder:  rest.NewEncoder(req),
		bodyEncoder: json.NewEncoder(),
	}

	return e
}

// Encode returns the encoded request, and body payload. If no payload body was
// set nil will be returned.  If an error occurred while encoding the API an
// error will be returned.
func (e *Encoder) Encode() (*http.Request, io.ReadSeeker, error) {
	req, payloadBody, err := e.reqEncoder.Encode()
	if err != nil {
		return nil, nil, err
	}

	jsonBody, err := e.bodyEncoder.Encode()
	if err != nil {
		return nil, nil, err
	}

	havePayload := payloadBody != nil
	haveJSON := jsonBody != nil

	if havePayload == haveJSON && haveJSON {
		return nil, nil, fmt.Errorf("unexpected JSON body and request payload for AWSMarshaler")
	}

	body := payloadBody
	if body == nil {
		body = jsonBody
	}

	return req, body, err
}

// SetValue will set a value to the header, path, query, or body.
//
// If the request's method is GET all BodyTarget values will be written to
// the query string.
func (e *Encoder) SetValue(t protocol.Target, k string, v protocol.ValueMarshaler, meta protocol.Metadata) {
	if e.err != nil {
		return
	}

	switch t {
	case protocol.PathTarget:
		fallthrough
	case protocol.QueryTarget:
		fallthrough
	case protocol.HeaderTarget:
		e.reqEncoder.SetValue(t, k, v, meta)
	case protocol.BodyTarget:
		fallthrough
	case protocol.PayloadTarget:
		if e.method == "GET" {
			e.reqEncoder.SetValue(t, k, v, meta)
		} else {
			e.bodyEncoder.SetValue(t, k, v, meta)
		}
	default:
		e.err = fmt.Errorf("unknown SetValue restjson encode target, %s, %s", t, k)
	}
}

// SetStream will set the stream to the payload of the request.
func (e *Encoder) SetStream(t protocol.Target, k string, v protocol.StreamMarshaler, meta protocol.Metadata) {
	if e.err != nil {
		return
	}

	switch t {
	case protocol.PayloadTarget:
		e.reqEncoder.SetStream(t, k, v, meta)
	default:
		e.err = fmt.Errorf("invalid target %s, for SetStream, must be PayloadTarget", t)
	}
}

// List will return the proper list encoder based on the given protocol.Target.
func (e *Encoder) List(t protocol.Target, k string, meta protocol.Metadata) protocol.ListEncoder {
	if e.err != nil {
		return nil
	}
	e.t = t

	switch t {
	case protocol.HeaderTarget:
		fallthrough
	case protocol.QueryTarget:
		return e.reqEncoder.List(t, k, meta)
	case protocol.BodyTarget:
		return e.bodyEncoder.List(t, k, meta)
	default:
		e.err = fmt.Errorf("unknown SetList restjson encode target, %s, %s", t, k)
		return nil
	}
}

// Map will return the proper map encoder based on the given protocol Target.
func (e *Encoder) Map(t protocol.Target, k string, meta protocol.Metadata) protocol.MapEncoder {
	if e.err != nil {
		return nil
	}

	e.t = t
	switch t {
	case protocol.QueryTarget:
		fallthrough
	case protocol.HeadersTarget:
		return e.reqEncoder.Map(t, k, meta)
	case protocol.BodyTarget:
		return e.bodyEncoder.Map(t, k, meta)
	default:
		e.err = fmt.Errorf("unknown SetMap restjson encode target, %s, %s", t, k)
		return nil
	}
}

// SetFields will set the nested type's fields to the body.
func (e *Encoder) SetFields(t protocol.Target, k string, m protocol.FieldMarshaler, meta protocol.Metadata) {
	if e.err != nil {
		return
	}

	switch t {
	case protocol.PayloadTarget:
		fallthrough
	case protocol.BodyTarget:
		e.bodyEncoder.SetFields(t, k, m, meta)
	default:
		e.err = fmt.Errorf("unknown SetMarshaler restjson encode target, %s, %s", t, k)
	}
}
