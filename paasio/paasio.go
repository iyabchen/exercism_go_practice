// Paas: Bill customers based on network and filesystem usage.
// Create a wrapper for network connections and files that can report IO
// statistics. The wrapper must report:
// - The total number of bytes read/written.
// - The total number of read/write operations.
package paasio

import (
	"io"
	"sync"
)

// testVersion identifies the API tested by the test program.
const testVersion = 3

// ++ is not atomic operation, need to protect it with
type MyReadWriteCounter struct {
	mutex sync.RWMutex

	reader  io.Reader
	readOps int
	readCnt int64

	writer   io.Writer
	writeOps int
	writeCnt int64
}

// implement interface ReadCounter
func (c *MyReadWriteCounter) ReadCount() (n int64, nops int) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.readCnt, c.readOps

}

// implement interface WriteCounter
func (c *MyReadWriteCounter) WriteCount() (n int64, nops int) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.writeCnt, c.writeOps

}

// implement interface io.Reader
func (c *MyReadWriteCounter) Read(p []byte) (n int, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.readOps++
	n, err = c.reader.Read(p)
	if err == nil {
		c.readCnt += int64(n)
	}
	return n, err
}

// implement interface io.Writer
func (c *MyReadWriteCounter) Write(p []byte) (n int, err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.writeOps++
	n, err = c.writer.Write(p)
	if err == nil {
		c.writeCnt += int64(n)
	}
	return n, err
}

func NewReadCounter(r io.Reader) ReadCounter {
	return &MyReadWriteCounter{reader: r}
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &MyReadWriteCounter{writer: w}
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &MyReadWriteCounter{reader: rw, writer: rw}
}
