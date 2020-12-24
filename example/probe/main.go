package main

import (
	"log"
	"os"

	"github.com/jimyx17/goav/libavcodec"
	"github.com/jimyx17/goav/libavformat"
	"github.com/jimyx17/goav/libavutil"
)

func main() {

	if len(os.Args) != 2 {
		log.Fatal("probe <filename>")
	}
	fd, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("probe <filename>")
	}

	ctx, err := libavformat.AvIOReaderOpen(fd, 4096)
	if err != nil {
		log.Fatal("could not open context")
	}

	fmt := libavformat.AvformatAllocContext()
	fmt.SetPb(ctx)
	ret := libavformat.AvformatOpenInput(&fmt, "", nil, nil)
	if ret < 0 {
		log.Printf("error opening input")
	}

	ret = fmt.AvformatFindStreamInfo(nil)
	if ret < 0 {
		log.Printf("error finding streams")
	}

	defer fmt.AvformatCloseInput()
	defer ctx.AvioBufClose()

	log.Printf("filename: %v", fmt.Url())
	log.Printf("NB Streams: %v", fmt.NbStreams())
	log.Printf("NB Programs: %v", fmt.NbPrograms())
	log.Printf("Format Name: %v", fmt.Iformat().Name())
	log.Printf("Format Long Name: %v", fmt.Iformat().LongName())
	log.Printf("Start Time: %v", fmt.StartTime())
	log.Printf("Duration: %v", fmt.Duration())
	log.Printf("Size: %v", "TBF")
	log.Printf("BitRate: %v", fmt.BitRate())
	log.Printf("Probe Score: %v", fmt.ProbeScore())

	d := fmt.Metadata()
	var entry *libavutil.AvDictionaryEntry
	for i := 0; i < d.AvDictCount(); i++ {
		entry = d.AvDictGet("", entry, libavutil.AvDictIgnoreSuffix)
		log.Printf("%v : %v", entry.Key(), entry.Value())
	}

	for _, chpt := range fmt.Chapters() {
		log.Printf("Chapter ID: %v", chpt.ID())
		log.Printf("Chapter TimeBase: %v", chpt.TimeBase().String())
		log.Printf("Chapter Start: %v", chpt.Start())
		log.Printf("Chapter Start Time: %v", float64(chpt.Start())*chpt.TimeBase().Q2d())
		log.Printf("Chapter End: %v", chpt.End())
		log.Printf("Chapter End Time: %v", float64(chpt.End())*chpt.TimeBase().Q2d())

		d := chpt.MetaData()
		var entry *libavutil.AvDictionaryEntry
		for i := 0; i < d.AvDictCount(); i++ {
			entry = d.AvDictGet("", entry, libavutil.AvDictIgnoreSuffix)
			log.Printf("%v : %v", entry.Key(), entry.Value())
		}
	}

	for _, str := range fmt.Streams() {
		cid := libavcodec.AvcodecDescriptorGet(str.CodecParameters().CodecID())
		codecCtx := str.Codec()
		params := str.CodecParameters()
		log.Printf("Stream Index: %v", str.Index())
		str.
			log.Printf("Codec Name: %v", cid.Name())
		log.Printf("Codec Long Name: %v", cid.LongName())
		log.Printf("Codec Profile: %v", params.Profile())
		log.Printf("Codec Type: %v", libavutil.AvGetMediaTypeString(libavutil.AvMediaType(params.CodecType())))
		log.Printf("Codec Time Base: %v", codecCtx.TimeBase().String())
		log.Printf("Codec Tag String: %v", params.CodecTagString())
		log.Printf("Codec Tag: 0x%04x", params.CodecTag())
		switch params.CodecType() {
		case libavutil.AvmediaTypeVideo:
			log.Printf("width: %v", params.Width())
			log.Printf("height: %v", params.Height())
			log.Printf("Coded width: %v", codecCtx.CodedWidth())
			log.Printf("Coded height: %v", codecCtx.CodedHeight())
			log.Printf("Closed captions: %v", codecCtx.Properties()&0x00000002)
			log.Printf("has_b_frames: %v", params.VideoDelay())
			sar := libavformat.AvformatAllocContext().AvGuessSampleAspectRatio(str, nil)
			log.Printf("Sample Aspect Ratio: %v", sar.String())
			log.Printf("Display Aspect Ratio: %v", libavcodec.AvReduce(int64(params.Width()*sar.Num()),
				int64(params.Height()*sar.Den()), 1024*1024))
			log.Printf("Pixel format: %v", libavutil.AvGetPixFmtName(params.Format()))
			log.Printf("Level: %v", params.Level())
			log.Printf("Color Range: %v", libavcodec.AvColorRangeName(params.ColorRange()))
			log.Printf("Color Space: %v", libavcodec.AvColorSpaceName(params.ColorSpace()))
			log.Printf("Color Primaries: %v", libavcodec.AvColorPrimariesName(params.ColorPrimaries()))
			log.Printf("Color Transfer: %v", libavcodec.AvColorTrcName(params.ColorTrc()))
			log.Printf("Chroma Location: %v", libavcodec.AvChromaLocationName(params.ChromaLocation()))
			switch params.FieldOrder() {
			case libavcodec.AV_FIELD_PROGRESSIVE:
				log.Printf("Field order: progressive")
			case libavcodec.AV_FIELD_TT:
				log.Printf("Field order: tt")
			case libavcodec.AV_FIELD_BB:
				log.Printf("Field order: bb")
			case libavcodec.AV_FIELD_TB:
				log.Printf("Field order: tb")
			case libavcodec.AV_FIELD_BT:
				log.Printf("Field order: bt")
			default:
				log.Printf("Field order: unknown")
			}
			log.Printf("Timecode: %v", libavutil.AvTimecodeMakeMpegTCString(codecCtx.TimecodeFrameStart()))
		case libavutil.AvmediaTypeAudio:
			log.Printf("Sample Format: %v", libavcodec.AvGetSampleFmtName(params.Format()))
			log.Printf("Sample rate: %v", params.SampleRate())
			log.Printf("Channels: %v", params.Channels())
			log.Printf("Channels layout: %v", libavcodec.AvChannelLayoutStr(params.Channels(), params.ChannelLayout()))
			log.Printf("Bits per sample: %v", libavcodec.AvGetBitsPerSample(params.CodecID()))
		case libavutil.AvmediaTypeSubtitle:
			log.Printf("width: %v", params.Width())
			log.Printf("height: %v", params.Height())
		}

		log.Printf("R Frame Rate: %v", str.RFrameRate())
		log.Printf("Average Frame Rate: %v", str.AvgFrameRate().String())
		log.Printf("Time Base: %v", str.TimeBase().String())
		log.Printf("Start PTS: %v", str.StartTime())
		log.Printf("Start Time: %v", float64(str.StartTime())*str.TimeBase().Q2d())
		log.Printf("Duration ts: %v", str.Duration())
		log.Printf("Duration: %v", float64(str.Duration())*str.TimeBase().Q2d())
		log.Printf("BitRate: %v", params.BitRate())
		log.Printf("Max Bit Rate: %v", codecCtx.RcMaxRate())
		log.Printf("Bits per raw sample: %v", codecCtx.BitsPerRawSample())
		log.Printf("NB Frames: %v", str.NbFrames())

		log.Printf("Default: %v", str.Disposition()&libavformat.AvDispositionDefault)
		log.Printf("Dub: %v", str.Disposition()&libavformat.AvDispositionDub)
		log.Printf("Original: %v", str.Disposition()&libavformat.AvDispositionOriginal)
		log.Printf("Comment: %v", str.Disposition()&libavformat.AvDispositionComment)
		log.Printf("Lyrics: %v", str.Disposition()&libavformat.AvDispositionLyrics)
		log.Printf("Karaoke: %v", str.Disposition()&libavformat.AvDispositionKaraoke)
		log.Printf("Forced: %v", str.Disposition()&libavformat.AvDispositionForced)
		log.Printf("Hearing Impaired: %v", str.Disposition()&libavformat.AvDispositionHearingImpaired)
		log.Printf("Visual Impaired: %v", str.Disposition()&libavformat.AvDispositionVisualImpaired)
		log.Printf("Clean Effects: %v", str.Disposition()&libavformat.AvDispositionCleanEffects)
		log.Printf("Attached Pic: %v", str.Disposition()&libavformat.AvDispositionAttachedPic)
		log.Printf("Timed Thumbnails: %v", str.Disposition()&libavformat.AvDispositionTimedThumbnails)

		d := str.Metadata()
		var entry *libavutil.AvDictionaryEntry
		for i := 0; i < d.AvDictCount(); i++ {
			entry = d.AvDictGet("", entry, libavutil.AvDictIgnoreSuffix)
			log.Printf("%v : %v", entry.Key(), entry.Value())
		}

	}
}
