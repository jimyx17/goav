// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

// Package libavutil is a utility library to aid portable multimedia programming.
// It contains safe portable string functions, random number generators, data structures,
// additional mathematics functions, cryptography and multimedia related functionality.
// Some generic features and utilities provided by the libavutil library
package libavutil

//#cgo pkg-config: libavutil
//#include <libavutil/avutil.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

type (
	Options       C.struct_AVOptions
	AvTree        C.struct_AVTree
	Rational      C.struct_AVRational
	MediaType     C.enum_AVMediaType
	AvPictureType C.enum_AVPictureType
	PixelFormat   C.enum_AVPixelFormat
	File          C.FILE
)

// AvutilVersion Return the LIBAvUTIL_VERSION_INT constant.
func AvutilVersion() uint {
	return uint(C.avutil_version())
}

// AvutilConfiguration Return the libavutil build-time configuration.
func AvutilConfiguration() string {
	return C.GoString(C.avutil_configuration())
}

// AvutilLicense Return the libavutil license.
func AvutilLicense() string {
	return C.GoString(C.avutil_license())
}

// AvGetMediaTypeString Return a string describing the media_type enum, NULL if media_type is unknown.
func AvGetMediaTypeString(mt MediaType) string {
	return C.GoString(C.av_get_media_type_string((C.enum_AVMediaType)(mt)))
}

// AvGetPictureTypeChar Return a single letter to describe the given picture type pict_type.
func AvGetPictureTypeChar(pt AvPictureType) string {
	return string(C.av_get_picture_type_char((C.enum_AVPictureType)(pt)))
}

// AvXIfNull Return x default pointer in case p is NULL.
func AvXIfNull(p, x int) {
	C.av_x_if_null(unsafe.Pointer(&p), unsafe.Pointer(&x))
}

// AvIntListLengthForSize Compute the length of an integer list.
func AvIntListLengthForSize(e uint, l int, t uint64) uint {
	return uint(C.av_int_list_length_for_size(C.uint(e), unsafe.Pointer(&l), (C.uint64_t)(t)))
}

// AvFopenUtf8 Open a file using a UTF-8 filename.
func AvFopenUtf8(p, m string) *File {
	f := C.av_fopen_utf8(C.CString(p), C.CString(m))
	return (*File)(f)
}

// AvGetTimeBaseQ Return the fractional representation of the internal time base.
func AvGetTimeBaseQ() Rational {
	return (Rational)(C.av_get_time_base_q())
}