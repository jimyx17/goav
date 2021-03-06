// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package libavformat

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
import "C"
import (
	"github.com/jimyx17/goav/libavcodec"
)

func newAvRational(r C.struct_AVRational) libavcodec.AvRational {
	return libavcodec.NewAvRational(int(r.num), int(r.den))
}
