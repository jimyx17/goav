package libavcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

const (
	AV_FIELD_UNKNOWN     = 0
	AV_FIELD_PROGRESSIVE = 1
	AV_FIELD_TT          = 2 //< Top coded_first, top displayed first
	AV_FIELD_BB          = 3 //< Bottom coded first, bottom displayed first
	AV_FIELD_TB          = 4 //< Top coded first, bottom displayed first
	AV_FIELD_BT          = 5
)

// CodecType Return codec_type
func (cp *AvCodecParameters) CodecType() AvMediaType {
	return AvMediaType(cp.codec_type)
}

// CodecID Return codec_id
func (cp *AvCodecParameters) CodecID() AvCodecID {
	return AvCodecID(cp.codec_id)
}

// Format Return format
func (cp *AvCodecParameters) Format() AvSampleFormat {
	return AvSampleFormat(cp.format)
}

// SetFormat Set format
func (cp *AvCodecParameters) SetFormat(f int) {
	cp.format = C.int(f)
}

// Width Return width
func (cp *AvCodecParameters) Width() int {
	return int(cp.width)
}

// SetWidth Set width
func (cp *AvCodecParameters) SetWidth(w int) {
	cp.width = C.int(w)
}

// Height Return height
func (cp *AvCodecParameters) Height() int {
	return int(cp.height)
}

// SetHeight Set height
func (cp *AvCodecParameters) SetHeight(h int) {
	cp.height = C.int(h)
}

// SampleAspectRatio Return sample_aspect_ratio
func (cp *AvCodecParameters) SampleAspectRatio() AvRational {
	return NewAvRational(int(cp.sample_aspect_ratio.num), int(cp.sample_aspect_ratio.den))
}

// SetSampleAspectRatio Set sample_aspect_ratio
func (cp *AvCodecParameters) SetSampleAspectRatio(sampleAspectRatio AvRational) {
	cp.sample_aspect_ratio.num = C.int(sampleAspectRatio.Num())
	cp.sample_aspect_ratio.den = C.int(sampleAspectRatio.Den())
}

// ChannelLayout Return channel_layout
func (cp *AvCodecParameters) ChannelLayout() uint64 {
	return uint64(cp.channel_layout)
}

// SetChannelLayout Set channel_layout
func (cp *AvCodecParameters) SetChannelLayout(cl uint64) {
	cp.channel_layout = C.uint64_t(cl)
}

// Channels Return channels
func (cp *AvCodecParameters) Channels() int {
	return int(cp.channels)
}

// SetChannels Set channels
func (cp *AvCodecParameters) SetChannels(nc int) {
	cp.channels = C.int(nc)
}

// SampleRate Return sample_rate
func (cp *AvCodecParameters) SampleRate() int {
	return int(cp.sample_rate)
}

// SetSampleRate Set sample_rate
func (cp *AvCodecParameters) SetSampleRate(sr int) {
	cp.sample_rate = C.int(sr)
}

// SetSampleRate Set sample_rate
func (cp *AvCodecParameters) BitRate() int {
	return int(cp.bit_rate)
}

// Profile Get profile name
func (cp *AvCodecParameters) Profile() string {
	return C.GoString(C.avcodec_profile_name(cp.codec_id, cp.profile))
}

// CodecTagString return codec tag string
func (cp *AvCodecParameters) CodecTagString() string {
	b := [32]C.char{}
	return C.GoString(C.av_fourcc_make_string((*C.char)(&b[0]), cp.codec_tag))
}

// CodecTag return codec tag
func (cp *AvCodecParameters) CodecTag() uint32 {
	return uint32(cp.codec_tag)
}

// CodecTag return codec tag
func (cp *AvCodecParameters) VideoDelay() int {
	return int(cp.video_delay)
}

func (cp *AvCodecParameters) Level() int {
	return int(cp.level)
}

func (cp *AvCodecParameters) ColorRange() AvColorRange {
	return AvColorRange(cp.color_range)
}

func (cp *AvCodecParameters) ColorSpace() AvColorSpace {
	return AvColorSpace(cp.color_space)
}

func (cp *AvCodecParameters) ColorPrimaries() AvColorPrimaries {
	return AvColorPrimaries(cp.color_primaries)
}

func (cp *AvCodecParameters) ColorTrc() AvColorTransferCharacteristic {
	return AvColorTransferCharacteristic(cp.color_trc)
}

func (cp *AvCodecParameters) ChromaLocation() AvChromaLocation {
	return AvChromaLocation(cp.chroma_location)
}

func (cp *AvCodecParameters) FieldOrder() AvFieldOrder {
	return AvFieldOrder(cp.field_order)
}
