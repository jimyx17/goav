# goav
Golang binding for FFmpeg

A binding to the ffmpeg video/audio manipulation library.


[![GoDoc](https://godoc.org/github.com/giorgisio/goav?status.svg)](https://godoc.org/github.com/giorgisio/goav)


## Jimyx17 library version

This fork is only to play with ffmpeg bindings, nothing else. 
Don't use it!
It really needs a full library refactor and some further adaptions.

## Usage

`````go

import "github.com/jimyx17/goav/libavformat"

func main() {

	filename := "sample.mp4"

	// Register all formats and codecs
	libavformat.AvRegisterAll() // deprecated, no longer needed

	ctx := libavformat.AvformatAllocContext()

	// Open video file
	if libavformat.AvformatOpenInput(&ctx, filename, nil, nil) != 0 {
		log.Println("Error: Couldn't open file.")
		return
	}

	// Retrieve stream information
	if ctx.AvformatFindStreamInfo(nil) < 0 {
		log.Println("Error: Couldn't find stream information.")

		// Close input file and free context
		ctx.AvformatCloseInput()
		return
	}

	//...

}
`````

## Libraries

* `avcodec` corresponds to the ffmpeg library: libavcodec [provides implementation of a wider range of codecs]
* `avformat` corresponds to the ffmpeg library: libavformat [implements streaming protocols, container formats and basic I/O access]
* `avutil` corresponds to the ffmpeg library: libavutil [includes hashers, decompressors and miscellaneous utility functions]
* `avfilter` corresponds to the ffmpeg library: libavfilter [provides a mean to alter decoded Audio and Video through chain of filters]
* `avdevice` corresponds to the ffmpeg library: libavdevice [provides an abstraction to access capture and playback devices]
* `swresample` corresponds to the ffmpeg library: libswresample [implements audio mixing and resampling routines]
* `swscale` corresponds to the ffmpeg library: libswscale [implements color conversion and scaling routines]


## Installation

[FFMPEG INSTALL INSTRUCTIONS](https://github.com/FFmpeg/FFmpeg/blob/master/INSTALL.md)

``` sh
sudo apt-get -y install autoconf automake build-essential libass-dev libfreetype6-dev libsdl1.2-dev libtheora-dev libtool libva-dev libvdpau-dev libvorbis-dev libxcb1-dev libxcb-shm0-dev libxcb-xfixes0-dev pkg-config texi2html zlib1g-dev

sudo apt install -y libavdevice-dev libavfilter-dev libswscale-dev libavcodec-dev libavformat-dev libswresample-dev libavutil-dev

sudo apt-get install yasm

export FFMPEG_ROOT=$HOME/ffmpeg
export CGO_LDFLAGS="-L$FFMPEG_ROOT/lib/ -lavcodec -lavformat -lavutil -lswscale -lswresample -lavdevice -lavfilter"
export CGO_CFLAGS="-I$FFMPEG_ROOT/include"
export LD_LIBRARY_PATH=$HOME/ffmpeg/lib
``` 

``` 
go get github.com/jimyx17/goav

``` 

## More Examples

Coding examples are available in the examples/ directory.

## Note
- Function names in Go are consistent with that of the libraries to help with easy search
- [cgo: Extending Go with C](http://blog.giorgis.io/cgo-examples)
- goav comes with absolutely no warranty.

## Contribute
- Fork this repo and create your own feature branch.
- Follow standard Go conventions
- Test your code.
- Create pull request


## TODO

- [ ] Refactor original goav version
- [ ] Adapt C interfaces to more standard Go interfaces
- [ ] Bump up ffmpeg version to 4.3.X
- [ ] Garbage Collection
- [ ] Go Tests
- [ ] Review included/excluded functions from each library
- [ ] Examples

## License
This library is under the [MIT License](http://opensource.org/licenses/MIT)
