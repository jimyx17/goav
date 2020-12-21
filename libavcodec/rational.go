// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package libavcodec contains the codecs (decoders and encoders) provided by the libavcodec library
//Provides some generic global options, which can be set on all the encoders and decoders.
package libavcodec

//#cgo pkg-config: libavformat libavcodec libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
import "C"
import (
	"fmt"
	"unsafe"
)

// NewAvRational ...
func NewAvRational(num, den int) AvRational {
	return AvRational{
		num: C.int(num),
		den: C.int(den),
	}
}

// String ...
func (r AvRational) String() string {
	return fmt.Sprintf("%d/%d", int(r.num), int(r.den))
}

// Num Return num
func (r AvRational) Num() int {
	return int(r.num)
}

// Den Return den
func (r AvRational) Den() int {
	return int(r.den)
}

// Q2d
func (r AvRational) Q2d() float64 {
	return float64(float64(r.num) / (float64)(r.den))
}

// Assign ...
func (r *AvRational) Assign(o AvRational) {
	r.Set(o.Num(), o.Den())
}

// Set ...
func (r *AvRational) Set(num, den int) {
	r.num, r.den = C.int(num), C.int(den)
}

// AVRescaleQRnd The operation is mathematically equivalent to `a * bq / cq`.
func AVRescaleQRnd(a int64, bq, cq AvRational, rnd AvRounding) int64 {
	return int64(C.av_rescale_q_rnd(C.int64_t(a), C.struct_AVRational(bq),
		C.struct_AVRational(cq), C.enum_AVRounding(rnd)))
}

// AvInvQ Invert a rational.
func AvInvQ(q AvRational) AvRational {
	return NewAvRational(q.Den(), q.Num())
}

func AvReduce(width, height int64, max int) AvRational {
	var num int
	var den int
	C.av_reduce((*C.int)(unsafe.Pointer(&num)), (*C.int)(unsafe.Pointer(&den)), C.long(width), C.long(height), 1024*1024)
	return NewAvRational(num, den)
}
