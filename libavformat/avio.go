package libavformat

/*#cgo pkg-config: libavformat
#include <libavformat/avformat.h>
extern int AvRead(void *opaque, uint8_t *buf, int buf_size);
extern int AvWrite(void *opaque, uint8_t *buf, int buf_size);
extern int64_t AvSeek(void *opaque, int64_t pos, int whence);
*/
import "C"
import (
	"io"
	"unsafe"

	"github.com/jimyx17/goav/libavcodec"
	"github.com/jimyx17/goav/libavutil"
	gopointer "github.com/mattn/go-pointer"
)

// AvIOOpen Create and initialize a AVIOContext for accessing the resource indicated by url.
func AvIOOpen(url string, flags int) (ioctx *AvIOContext, err error) {
	cURL := C.CString(url)
	defer C.free(unsafe.Pointer(cURL))
	err = libavutil.ErrorFromCode(int(C.avio_open((**C.AVIOContext)(unsafe.Pointer(&ioctx)), cURL, C.int(flags))))
	return
}

// AvIOOpen Create and initialize a AVIOContext for accessing the resource indicated by url.
func AvIOReaderOpen(r io.Reader, bufferSize int64) (ioctx *AvIOContext, err error) {
	buf := libavutil.AvMalloc(bufferSize)

	var seeker *[0]byte
	if _, ok := r.(io.Seeker); ok {
		seeker = (*[0]byte)(C.AvSeek)
	}

	p := gopointer.Save(r)
	ctx := C.avio_alloc_context((*C.uchar)(buf), C.int(bufferSize),
		0, p, (*[0]byte)(C.AvRead), nil, (*[0]byte)(seeker))

	ioctx = (*AvIOContext)(ctx)
	return
}

// AvioClose Close the resource accessed by the AVIOContext s and free it.
// This function can only be used if s was opened by avio_open().
// The internal buffer is automatically flushed before closing the resource.
func (ctxt *AvIOContext) AvioBufClose() error {
	if ctxt.buffer != nil {
		libavutil.AvFree(unsafe.Pointer(ctxt.buffer))
	}
	libavutil.AvFree(unsafe.Pointer(ctxt))
	return nil
}

// AvioClose Close the resource accessed by the AVIOContext s and free it.
// This function can only be used if s was opened by avio_open().
// The internal buffer is automatically flushed before closing the resource.
func (ctxt *AvIOContext) AvioClose() error {
	return libavutil.ErrorFromCode(int(C.avio_close((*C.AVIOContext)(unsafe.Pointer(ctxt)))))
}

// AvioClosep the resource accessed by the AVIOContext *s, free it and set the pointer pointing to it to NULL.
// This function can only be used if s was opened by avio_open().
func AvioClosep(ctxt *AvIOContext) error {
	return libavutil.ErrorFromCode(int(C.avio_closep((**C.AVIOContext)(unsafe.Pointer(&ctxt)))))
}

// AvGetPacket Allocate and read the payload of a packet and initialize its fields with default values.
func (ctxt *AvIOContext) AvGetPacket(pkt *libavcodec.AvPacket, size int) int {
	return int(C.av_get_packet((*C.struct_AVIOContext)(ctxt),
		toCPacket(pkt), C.int(size)))
}

// AvAppendPacket Read data and append it to the current content of the Packet.
func (ctxt *AvIOContext) AvAppendPacket(pkt *libavcodec.AvPacket, size int) int {
	return int(C.av_append_packet((*C.struct_AVIOContext)(ctxt),
		toCPacket(pkt), C.int(size)))
}
