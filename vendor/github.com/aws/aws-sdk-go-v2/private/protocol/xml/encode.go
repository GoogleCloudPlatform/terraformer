package xml

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// An Encoder provides encoding of the AWS XML protocol. This encoder will will
// write all content to XML. Only supports body and payload targets.
type Encoder struct {
	encoder    *xml.Encoder
	encodedBuf *bytes.Buffer
	fieldBuf   protocol.FieldBuffer
	err        error
}

// NewEncoder creates a new encoder for encoding AWS XML protocol. Only encodes
// fields into the XML body, and error is returned if target is anything other
// than Body or Payload.
func NewEncoder() *Encoder {
	encodedBuf := bytes.NewBuffer(nil)
	return &Encoder{
		encodedBuf: encodedBuf,
		encoder:    xml.NewEncoder(encodedBuf),
	}
}

// Encode returns the encoded XMl reader. An error will be returned if one was
// encountered while building the XML body.
func (e *Encoder) Encode() (io.ReadSeeker, error) {
	if e.err != nil {
		return nil, e.err
	}

	if err := e.encoder.Flush(); err != nil {
		return nil, fmt.Errorf("unable to marshal XML, %v", err)
	}

	if e.encodedBuf.Len() == 0 {
		return nil, nil
	}

	return bytes.NewReader(e.encodedBuf.Bytes()), e.err
}

// SetValue sets an individual value to the XML body.
func (e *Encoder) SetValue(t protocol.Target, k string, v protocol.ValueMarshaler, meta protocol.Metadata) {
	if e.err != nil {
		return
	}
	if t != protocol.BodyTarget && t != protocol.PayloadTarget {
		e.err = fmt.Errorf(" invalid target %s for xml encoder SetValue, %s", t, k)
		return
	}

	e.err = addValueToken(e.encoder, &e.fieldBuf, k, v, meta)
}

// SetStream is not supported for XML protocol marshaling.
func (e *Encoder) SetStream(t protocol.Target, k string, v protocol.StreamMarshaler, meta protocol.Metadata) {
	e.err = fmt.Errorf("xml encoder SetStream not supported, %s, %s", t, k)
}

// List creates an XML list and calls the passed in fn callback with a list encoder.
func (e *Encoder) List(t protocol.Target, k string, meta protocol.Metadata) protocol.ListEncoder {
	if e.err != nil {
		return nil
	}
	if t != protocol.BodyTarget && t != protocol.PayloadTarget {
		e.err = fmt.Errorf(" invalid target %s for xml encoder SetValue, %s", t, k)
		return nil
	}

	if v := meta.ListLocationName; len(v) == 0 {
		if meta.Flatten {
			meta.ListLocationName = k
		} else {
			meta.ListLocationName = "member"
		}
	}

	return &ListEncoder{
		Base:     e,
		Key:      k,
		Metadata: meta,
	}
}

// Map creates an XML map and calls the passed in fn callback with a map encoder.
func (e *Encoder) Map(t protocol.Target, k string, meta protocol.Metadata) protocol.MapEncoder {
	if e.err != nil {
		return nil
	}
	if t != protocol.BodyTarget && t != protocol.PayloadTarget {
		e.err = fmt.Errorf(" invalid target %s for xml encoder SetValue, %s", t, k)
		return nil
	}

	me := MapEncoder{Base: e,
		// TODO: Get rid of these fields as we need the metadata structure now
		Flatten:   meta.Flatten,
		KeyName:   meta.MapLocationNameKey,
		ValueName: meta.MapLocationNameValue,
		Metadata:  meta,
		Key:       k,
	}

	return &me

}

// SetFields sets the nested fields to the XML body.
func (e *Encoder) SetFields(t protocol.Target, k string, m protocol.FieldMarshaler, meta protocol.Metadata) {
	if e.err != nil {
		return
	}
	if t != protocol.BodyTarget && t != protocol.PayloadTarget {
		e.err = fmt.Errorf(" invalid target %s for xml encoder SetFields, %s", t, k)
		return
	}

	tok, err := xmlStartElem(k, meta)
	if err != nil {
		e.err = err
		return
	}

	e.encoder.EncodeToken(tok)
	m.MarshalFields(e)
	e.encoder.EncodeToken(xml.EndElement{Name: tok.Name})
}

// A ListEncoder encodes elements within a list for the XML encoder.
type ListEncoder struct {
	Base     *Encoder
	Key      string
	Metadata protocol.Metadata
	Token    xml.StartElement
	err      error
}

// Map will return an error since nested collections are not support by this protocol.
func (e *ListEncoder) Map() protocol.MapEncoder {
	e.err = fmt.Errorf("xml list encoder ListSetMap not supported")
	return nil
}

// List will return an error since nested collections are not support by this protocol.
func (e *ListEncoder) List() protocol.ListEncoder {
	e.err = fmt.Errorf("xml list encoder ListSetList not supported")
	return nil
}

// Start will write the start element and set the token for closing
func (e *ListEncoder) Start() {
	var tok xml.StartElement
	var err error
	if !e.Metadata.Flatten {
		tok, err = xmlStartElem(e.Key, e.Metadata)
		if err != nil {
			e.err = err
			return
		}

		e.Base.encoder.EncodeToken(tok)
	}

	e.Token = tok
}

// End will write the end element if the list is not flat.
func (e *ListEncoder) End() {
	if !e.Metadata.Flatten {
		e.err = e.Base.encoder.EncodeToken(xml.EndElement{Name: e.Token.Name})
	}
}

// ListAddValue will add the value to the list.
func (e *ListEncoder) ListAddValue(v protocol.ValueMarshaler) {
	if e.err != nil {
		return
	}

	e.err = addValueToken(e.Base.encoder, &e.Base.fieldBuf, e.Metadata.ListLocationName, v, protocol.Metadata{})
}

// ListAddFields will set the nested type's fields to the list.
func (e *ListEncoder) ListAddFields(m protocol.FieldMarshaler) {
	if e.err != nil {
		return
	}

	var tok xml.StartElement
	tok, e.err = xmlStartElem(e.Metadata.ListLocationName, protocol.Metadata{})
	if e.err != nil {
		return
	}

	e.Base.encoder.EncodeToken(tok)
	m.MarshalFields(e.Base)
	e.Base.encoder.EncodeToken(xml.EndElement{Name: tok.Name})
}

// A MapEncoder encodes key values pair map values for the XML encoder.
type MapEncoder struct {
	Base      *Encoder
	Flatten   bool
	Key       string
	KeyName   string
	ValueName string
	err       error

	Token    xml.StartElement
	Metadata protocol.Metadata
}

// Start will open a new scope by creating a new XML start element tag.
func (e *MapEncoder) Start() {
	tok, err := xmlStartElem(e.Key, e.Metadata)
	if err != nil {
		e.err = err
		return
	}

	e.Token = tok
	e.Base.encoder.EncodeToken(tok)
}

// End will close the associated tag.
func (e *MapEncoder) End() {
	e.Base.encoder.EncodeToken(xml.EndElement{Name: e.Token.Name})
}

// Map will set err as nested collections are not supported in this protocol.
func (e *MapEncoder) Map(k string) protocol.MapEncoder {
	e.err = fmt.Errorf("xml map encoder MapSetList not supported, %s", k)
	return nil
}

// List will set err as nested collections are not supported in this protocol.
func (e *MapEncoder) List(k string) protocol.ListEncoder {
	e.err = fmt.Errorf("xml map encoder ListSetList not supported, %s", k)
	return nil
}

// MapSetValue sets a map value.
func (e *MapEncoder) MapSetValue(k string, v protocol.ValueMarshaler) {
	if e.err != nil {
		return
	}

	var tok xml.StartElement
	if !e.Flatten {
		tok, e.err = xmlStartElem("entry", protocol.Metadata{})
		if e.err != nil {
			return
		}
		e.Base.encoder.EncodeToken(tok)
	}

	keyName, valueName := e.KeyName, e.ValueName
	if len(keyName) == 0 {
		keyName = "key"
	}
	if len(valueName) == 0 {
		valueName = "value"
	}

	e.err = addValueToken(e.Base.encoder, &e.Base.fieldBuf, keyName, protocol.StringValue(k), protocol.Metadata{})
	if e.err != nil {
		return
	}

	e.err = addValueToken(e.Base.encoder, &e.Base.fieldBuf, valueName, v, protocol.Metadata{})
	if e.err != nil {
		return
	}

	if !e.Flatten {
		e.Base.encoder.EncodeToken(xml.EndElement{Name: tok.Name})
	}
}

// MapSetList is not supported.
func (e *MapEncoder) MapSetList(k string, fn func(le protocol.ListEncoder)) {
	e.err = fmt.Errorf("xml map encoder MapSetList not supported, %s", k)
}

// MapSetMap is not supported.
func (e *MapEncoder) MapSetMap(k string, fn func(me protocol.MapEncoder)) {
	e.err = fmt.Errorf("xml map encoder MapSetMap not supported, %s", k)
}

// MapSetFields will set the nested type's fields under the map.
func (e *MapEncoder) MapSetFields(k string, m protocol.FieldMarshaler) {
	if e.err != nil {
		return
	}

	var tok xml.StartElement
	if !e.Flatten {
		tok, e.err = xmlStartElem("entry", protocol.Metadata{})
		if e.err != nil {
			return
		}
		e.Base.encoder.EncodeToken(tok)
	}

	keyName, valueName := e.KeyName, e.ValueName
	if len(keyName) == 0 {
		keyName = "key"
	}
	if len(valueName) == 0 {
		valueName = "value"
	}

	e.err = addValueToken(e.Base.encoder, &e.Base.fieldBuf, keyName, protocol.StringValue(k), protocol.Metadata{})
	if e.err != nil {
		return
	}

	valTok, err := xmlStartElem(valueName, protocol.Metadata{})
	if err != nil {
		e.err = err
		return
	}
	e.Base.encoder.EncodeToken(valTok)

	m.MarshalFields(e.Base)

	e.Base.encoder.EncodeToken(xml.EndElement{Name: valTok.Name})

	if !e.Flatten {
		e.Base.encoder.EncodeToken(xml.EndElement{Name: tok.Name})
	}
}

func addValueToken(e *xml.Encoder, fieldBuf *protocol.FieldBuffer, k string, v protocol.ValueMarshaler, meta protocol.Metadata) error {
	b, err := fieldBuf.GetValue(v)
	if err != nil {
		return err
	}

	tok, err := xmlStartElem(k, meta)
	if err != nil {
		return err
	}

	e.EncodeToken(tok)
	e.EncodeToken(xml.CharData(b))
	e.EncodeToken(xml.EndElement{Name: tok.Name})

	return nil
}

func xmlStartElem(k string, meta protocol.Metadata) (xml.StartElement, error) {
	tok := xml.StartElement{Name: xmlName(k, meta)}
	attrs, err := buildAttributes(meta)
	if err != nil {
		return xml.StartElement{}, err
	}
	tok.Attr = attrs

	return tok, nil
}

func xmlName(k string, meta protocol.Metadata) xml.Name {
	name := xml.Name{Local: k}

	// TODO need to do something with namespace?
	//	if len(meta.XMLNamespacePrefix) > 0  && len(meta.XMLNamespaceURI) {
	//		name.Space = prefix
	//	}

	return name
}

func buildAttributes(meta protocol.Metadata) ([]xml.Attr, error) {
	n := len(meta.Attributes)
	if len(meta.XMLNamespaceURI) > 0 {
		n++
	}

	if n == 0 {
		return nil, nil
	}

	attrs := make([]xml.Attr, n)

	for _, a := range meta.Attributes {
		str, err := a.Value.MarshalValue()
		if err != nil {
			return nil, err
		}

		attrs = append(attrs, xml.Attr{Name: xmlName(a.Name, a.Meta), Value: str})
	}

	if uri := meta.XMLNamespaceURI; len(uri) > 0 {
		attr := xml.Attr{
			Name:  xml.Name{Local: "xmlns"},
			Value: uri,
		}
		if p := meta.XMLNamespacePrefix; len(p) > 0 {
			attr.Name.Local += ":" + p
		}
		attrs = append(attrs, attr)
	}

	return attrs, nil
}
