package libavformat

/*#cgo pkg-config: libavformat
#include <libavformat/avformat.h>
#define AVERROR_EOF                FFERRTAG( 'E','O','F',' ')
*/
import "C"
import (
	"io"
	"unsafe"

	"github.com/jimyx17/goav/libavutil"
	gopointer "github.com/mattn/go-pointer"
)

type Sizer interface {
	Size() int64
}

//export AvRead
func AvRead(opaque unsafe.Pointer, buff unsafe.Pointer, size int) int {
	r, ok := gopointer.Restore(opaque).(io.Reader)
	if !ok {
		return -1
	}

	data := make([]byte, size)

	n, err := r.Read(data[:])
	if err != nil {
		return libavutil.AvErrorEOF
	}

	C.memcpy(buff, unsafe.Pointer(&data[0]), C.ulong(n))

	return n
}

//export AvWrite
func AvWrite(opaque unsafe.Pointer, buff unsafe.Pointer, size int) int {
	w, ok := gopointer.Restore(opaque).(io.Writer)
	if !ok {
		return -1
	}

	b := C.GoBytes(buff, C.int(size))

	n, err := w.Write((b)[:])
	if err != nil {
		return -1
	}

	return n
}

//export AvSeek
func AvSeek(opaque unsafe.Pointer, offset int64, whence int) int64 {
	s, ok := gopointer.Restore(opaque).(io.Seeker)
	if !ok {
		return -1
	}

	if whence == 0x10000 {
		if s, ok := gopointer.Restore(opaque).(Sizer); ok {
			ret := s.Size()
			return ret
		}
		return -1
	}

	n, err := s.Seek(offset, int(whence))
	if err != nil {
		return -1
	}

	return n
}
