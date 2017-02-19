package circular

import (
	"io"
)

const testVersion = 4

// Implement a circular buffer of bytes supporting both overflow-checked writes
// and unconditional, possibly overwriting, writes.

// We chose the below API so that Buffer implements io.ByteReader
// and io.ByteWriter and can be used (size permitting) as a drop in
// replacement for anything using that interface.

type Buffer struct {
	// current pointer for read and write
	rpt      int
	wpt      int
	length   int
	capacity int
	content  []byte
}

// Create a circular buffer
func NewBuffer(size int) *Buffer {
	return &Buffer{0, 0, 0, size, make([]byte, size)}
}

// Read the next byte. If read is ahead of write, then return read fail.
// Once read, the data is consumed.
func (buf *Buffer) ReadByte() (b byte, e error) {
	if buf.length > 0 {
		b = buf.content[buf.rpt]
		buf.rpt++
		buf.rpt = buf.rpt % buf.capacity
		buf.length--
		return b, nil
	}
	return '\x00', io.EOF

}

// Write the next byte, if content is not read before write,
// then it means the buffer full, return error
func (buf *Buffer) WriteByte(c byte) error {
	if buf.length == buf.capacity {
		return io.ErrShortBuffer
	} else {
		buf.content[buf.wpt] = c
		buf.wpt++
		buf.wpt = buf.wpt % buf.capacity
		buf.length++
		return nil
	}

}

// Write the next byte even if it is full
// Reset W/R count in case the counter overflows
func (buf *Buffer) Overwrite(c byte) {
	// if current read pointer is same as write pointer, then read pointer need
	// to move to the next unconsumed byte
	if buf.length == buf.capacity { // an actual overwrite, read also need to move
		buf.rpt++
		buf.rpt = buf.rpt % buf.capacity
	}else{
		buf.length++
	}
	buf.content[buf.wpt] = c
	buf.wpt++
	buf.wpt = buf.wpt % buf.capacity
}

// put buffer in an empty state
func (buf *Buffer) Reset() {
	buf.wpt, buf.rpt, buf.length = 0, 0, 0

}
