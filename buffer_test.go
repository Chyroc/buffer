package buffer_test

import (
	"github.com/Chyroc/buffer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuffer(t *testing.T) {
	as := assert.New(t)

	buf := buffer.New()

	buf.AppendInt(-1)
	buf.AppendInt8(-2)
	buf.AppendInt16(-3)
	buf.AppendInt32(-4)
	buf.AppendInt64(-5)
	as.Equal("-1-2-3-4-5", buf.String())

	buf.Reset()
	buf.AppendUint(1)
	buf.AppendUint8(2)
	buf.AppendUint16(3)
	buf.AppendUint32(4)
	buf.AppendUint64(5)
	as.Equal("12345", buf.String())

	buf.Reset()
	buf.AppendBool(true)
	buf.AppendBool(false)
	as.Equal("truefalse", buf.String())

	buf.Reset()
	buf.AppendString("test")
	buf.AppendQuote("1234")
	buf.AppendByte('=')
	as.Equal(`test"1234"=`, buf.String())
}
