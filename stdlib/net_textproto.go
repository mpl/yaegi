package stdlib

// Generated by 'goexports net/textproto'. Do not edit!

import (
	"net/textproto"
	"reflect"
)

func init() {
	Value["net/textproto"] = map[string]reflect.Value{
		"CanonicalMIMEHeaderKey": reflect.ValueOf(textproto.CanonicalMIMEHeaderKey),
		"Dial":                   reflect.ValueOf(textproto.Dial),
		"NewConn":                reflect.ValueOf(textproto.NewConn),
		"NewReader":              reflect.ValueOf(textproto.NewReader),
		"NewWriter":              reflect.ValueOf(textproto.NewWriter),
		"TrimBytes":              reflect.ValueOf(textproto.TrimBytes),
		"TrimString":             reflect.ValueOf(textproto.TrimString),
	}
	Type["net/textproto"] = map[string]reflect.Type{
		"Conn":          reflect.TypeOf((*textproto.Conn)(nil)).Elem(),
		"Error":         reflect.TypeOf((*textproto.Error)(nil)).Elem(),
		"MIMEHeader":    reflect.TypeOf((*textproto.MIMEHeader)(nil)).Elem(),
		"Pipeline":      reflect.TypeOf((*textproto.Pipeline)(nil)).Elem(),
		"ProtocolError": reflect.TypeOf((*textproto.ProtocolError)(nil)).Elem(),
		"Reader":        reflect.TypeOf((*textproto.Reader)(nil)).Elem(),
		"Writer":        reflect.TypeOf((*textproto.Writer)(nil)).Elem(),
	}
}