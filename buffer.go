package buffer

import "strconv"

type Builder struct {
	bs []byte
}

func (b *Builder) AppendInt64(i int64) {
	b.bs = strconv.AppendInt(b.bs, i, 10)
}

func (b *Builder) AppendInt32(i int32) {
	b.bs = strconv.AppendInt(b.bs, int64(i), 10)
}

func (b *Builder) AppendInt16(i int16) {
	b.bs = strconv.AppendInt(b.bs, int64(i), 10)
}

func (b *Builder) AppendInt8(i int8) {
	b.bs = strconv.AppendInt(b.bs, int64(i), 10)
}

func (b *Builder) AppendInt(i int) {
	b.bs = strconv.AppendInt(b.bs, int64(i), 10)
}

func (b *Builder) AppendUint64(i uint64) {
	b.bs = strconv.AppendUint(b.bs, i, 10)
}

func (b *Builder) AppendUint32(i uint32) {
	b.bs = strconv.AppendUint(b.bs, uint64(i), 10)
}

func (b *Builder) AppendUint16(i uint16) {
	b.bs = strconv.AppendUint(b.bs, uint64(i), 10)
}

func (b *Builder) AppendUint8(i uint8) {
	b.bs = strconv.AppendUint(b.bs, uint64(i), 10)
}

func (b *Builder) AppendUint(i uint) {
	b.bs = strconv.AppendUint(b.bs, uint64(i), 10)
}

func (b *Builder) AppendBool(i bool) {
	b.bs = strconv.AppendBool(b.bs, i)
}

func (b *Builder) AppendFloat64(f float64) {
	b.bs = strconv.AppendFloat(b.bs, f, 'f', -1, 64)
}

func (b *Builder) AppendFloat32(f float32) {
	b.bs = strconv.AppendFloat(b.bs, float64(f), 'f', -1, 64)
}

func (b *Builder) AppendString(s string) {
	b.bs = append(b.bs, s...)
}

func (b *Builder) AppendQuote(s string) {
	b.bs = strconv.AppendQuote(b.bs, s)
}

func (b *Builder) AppendByte(s byte) {
	b.bs = append(b.bs, s)
}

func (b *Builder) Len() int {
	return len(b.bs)
}

func (b *Builder) Cap() int {
	return cap(b.bs)
}

func (b *Builder) String() string {
	return string(b.bs)
}

func (b *Builder) Bytes() []byte {
	return b.bs
}

func (b *Builder) Reset() {
	b.bs = b.bs[:0]
}
