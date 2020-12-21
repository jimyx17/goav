// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"unsafe"

	"github.com/jimyx17/goav/libavcodec"
	"github.com/jimyx17/goav/libavutil"
)

// ID Return ID
func (st *AvChapter) ID() int {
	return int(st.id)
}

// Start Return Start
func (st *AvChapter) Start() int64 {
	return int64(st.start)
}

// Index Return index
func (st *AvChapter) End() int64 {
	return int64(st.end)
}

// Index Return index
func (st *AvChapter) TimeBase() libavcodec.AvRational {
	return newAvRational(st.time_base)
}

// Index Return index
func (st *AvChapter) MetaData() *libavutil.AvDictionary {
	return (*libavutil.AvDictionary)(unsafe.Pointer(st.metadata))
}
