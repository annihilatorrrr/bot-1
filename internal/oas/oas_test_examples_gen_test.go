// Code generated by ogen, DO NOT EDIT.

package oas

import (
	"github.com/go-faster/jx"

	std "encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateTelegramAccountOK_EncodeDecode(t *testing.T) {
	var typ CreateTelegramAccountOK
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateTelegramAccountOK
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestCreateTelegramAccountReq_EncodeDecode(t *testing.T) {
	var typ CreateTelegramAccountReq
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 CreateTelegramAccountReq
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestError_EncodeDecode(t *testing.T) {
	var typ Error
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Error
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestHealth_EncodeDecode(t *testing.T) {
	var typ Health
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 Health
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestSetTelegramAccountCodeOK_EncodeDecode(t *testing.T) {
	var typ SetTelegramAccountCodeOK
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 SetTelegramAccountCodeOK
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestSetTelegramAccountCodeReq_EncodeDecode(t *testing.T) {
	var typ SetTelegramAccountCodeReq
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 SetTelegramAccountCodeReq
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestSpanID_EncodeDecode(t *testing.T) {
	var typ SpanID
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 SpanID
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestTelegramAccountID_EncodeDecode(t *testing.T) {
	var typ TelegramAccountID
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 TelegramAccountID
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
func TestTraceID_EncodeDecode(t *testing.T) {
	var typ TraceID
	typ.SetFake()

	e := jx.Encoder{}
	typ.Encode(&e)
	data := e.Bytes()
	require.True(t, std.Valid(data), "Encoded: %s", data)

	var typ2 TraceID
	require.NoError(t, typ2.Decode(jx.DecodeBytes(data)))
}
