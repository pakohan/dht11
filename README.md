# dht

CGo client library for the DHT11/22 sensors on a Raspberry Pi 1 via GPIO Edit

```
go get -d -u github.com/pakohan/dht
go generate github.com/pakohan/dht
```

The code is copied over from [adafruit](https://github.com/adafruit/Adafruit_Python_DHT).
I removed the Python part and moved the conversion from bytes from the C library to Go
(mainly for learning about the library a little bit while changing it).

Since there is a [mmap library](golang.org/x/exp/mmap) it seems that this could
be written entirely in Go, but I think the garbage collector prevents us from getting
correct results here.
What is not supported yet is the RPi2 and BeagleBone Black, both have drivers in the
repo mentioned above. What I could not test as well was the DHT22 sensor, but I tried
to copy over that conversion part as well.

I think the more of the library is moved to Go, the more similarities will be between
the different platforms, which means less C code to maintain. Maybe opening the right
device file can be moved to Go as well.

## usage

`go get` will throw an error when it is retrieved without the `-d` flag since it
tries to build the library after downloading. This is not possible, because before we need
some C compilation to do. Everything is done via `go generate` and `make`, so when you just
want to use the library, running `go generate` will do the trick.

## development

`make` is a wonderful build tool that just builds what has been changed. Just change any file,
run `go generate` or `make` (whatever you prefer) and only the files that will be changed
will get recompiled. First it compiles the .c files to .o files, then it will combine
these object files to one static library called dht.a which then can be catched by the cgo
tool to build this library.
