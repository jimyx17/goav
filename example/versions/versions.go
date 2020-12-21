package main

import (
	"log"

	"github.com/jimyx17/goav/libavcodec"
	"github.com/jimyx17/goav/libavdevice"
	"github.com/jimyx17/goav/libavfilter"
	"github.com/jimyx17/goav/libavformat"
	"github.com/jimyx17/goav/libavutil"
	"github.com/jimyx17/goav/libswresample"
	"github.com/jimyx17/goav/libswscale"
)

func main() {
	log.Printf("AvCodec  Version:\t%v", libavcodec.AvcodecVersion())
	log.Printf("AvDevice Version:\t%v", libavdevice.AvdeviceVersion())
	log.Printf("AvFilter Version:\t%v", libavfilter.AvfilterVersion())
	log.Printf("AvFormat Version:\t%v", libavformat.AvformatVersion())
	log.Printf("AvUtil   Version:\t%v", libavutil.AvutilVersion())
	log.Printf("Resample Version:\t%v", libswresample.SwresampleLicense())
	log.Printf("SwScale  Version:\t%v", libswscale.SwscaleVersion())
}
