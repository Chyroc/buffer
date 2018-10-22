package buffer

import "strconv"

type Builder interface {
	AppendInt64(i int64)
	AppendInt32(i int32)
	AppendInt16(i int16)
	AppendInt8(i int8)
	AppendInt(i int)
	AppendUint64(i uint64)
	AppendUint32(i uint32)
	AppendUint16(i uint16)
	AppendUint8(i uint8)
	AppendUint(i uint)
	AppendBool(i bool)
	AppendFloat64(f float64)
	AppendFloat32(f float32)
	AppendString(s string)
	AppendQuote(s string)
	AppendByte(s byte)
	Len() int
	Cap() int
	String() string
	Bytes() []byte
	Reset()
}

type builder struct {
	bs []byte
}

func New() Builder {
	return &builder{bs: make([]byte, 0, 100)}
}

func (b *builder) AppendInt64(i int64) {
	b.bs = strconv.AppendInt(b.bs, i, 10)
}

func (b *builder) AppendInt32(i int32) {
	b.bs = strconv.AppendInt(b.bs, int64(i), 10)
}

func (b *builder) AppendInt16(i int16) {
	b.bs = strconv.AppendInt(b.bs, int64(i), 10)
}

func (b *builder) AppendInt8(i int8) {
	b.bs = strconv.AppendInt(b.bs, int64(i), 10)
}

func (b *builder) AppendInt(i int) {
	b.bs = strconv.AppendInt(b.bs, int64(i), 10)
}

func (b *builder) AppendUint64(i uint64) {
	b.bs = strconv.AppendUint(b.bs, i, 10)
}

func (b *builder) AppendUint32(i uint32) {
	b.bs = strconv.AppendUint(b.bs, uint64(i), 10)
}

func (b *builder) AppendUint16(i uint16) {
	b.bs = strconv.AppendUint(b.bs, uint64(i), 10)
}

func (b *builder) AppendUint8(i uint8) {
	b.bs = strconv.AppendUint(b.bs, uint64(i), 10)
}

func (b *builder) AppendUint(i uint) {
	b.bs = strconv.AppendUint(b.bs, uint64(i), 10)
}

func (b *builder) AppendBool(i bool) {
	b.bs = strconv.AppendBool(b.bs, i)
}

func (b *builder) AppendFloat64(f float64) {
	b.bs = strconv.AppendFloat(b.bs, f, 'f', -1, 64)
}

func (b *builder) AppendFloat32(f float32) {
	b.bs = strconv.AppendFloat(b.bs, float64(f), 'f', -1, 64)
}

func (b *builder) AppendString(s string) {
	b.bs = append(b.bs, s...)
}

func (b *builder) AppendQuote(s string) {
	b.bs = strconv.AppendQuote(b.bs, s)
}

func (b *builder) AppendByte(s byte) {
	b.bs = append(b.bs, s)
}

func (b *builder) Len() int {
	return len(b.bs)
}

func (b *builder) Cap() int {
	return cap(b.bs)
}

func (b *builder) String() string {
	return string(b.bs)
}

func (b *builder) Bytes() []byte {
	return b.bs
}

func (b *builder) Reset() {
	b.bs = b.bs[:0]
}
