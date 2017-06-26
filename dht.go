package dht

//go:generate make

// #cgo CFLAGS: -I./Raspberry_Pi -I./Raspberry_Pi_2 -I.
// #cgo LDFLAGS: ${SRCDIR}/dht.a
// #include "Raspberry_Pi/pi_dht_read.h"
// #include "Raspberry_Pi_2/pi_2_dht_read.h"
import "C"
import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
)

const (
	SensorDHT11 = iota
	SensorDHT22
	ModelUnknown = 0
	ModelRPi     = iota
	ModelRPi2
	ModelFile = "/sys/firmware/devicetree/base/model"
)

// GetPlatformModel detects the Raspberry Pi model
// https://www.raspberrypi.org/forums/viewtopic.php?f=31&t=120385
func GetPlatformModel() int {
	model, err := ioutil.ReadFile(ModelFile)
	if err != nil {
		return ModelUnknown
	}
	switch {
	case bytes.HasPrefix(model, []byte("Raspberry Pi 3")):
		return ModelRPi2
	case bytes.HasPrefix(model, []byte("Raspberry Pi 2")):
		return ModelRPi2
	case bytes.HasPrefix(model, []byte("Raspberry Pi")):
		return ModelRPi
	default:
		return ModelUnknown
	}
}

// GetSensorData calls ReadSensor for the given pin and returns the converted
// humidty and temparature depending on the sensor model.
func GetSensorData(mtype, stype, pin int) (humidity, temperature float32, err error) {
	if stype != SensorDHT11 && stype != SensorDHT22 {
		err = fmt.Errorf("sensor type must be either %d or %d", SensorDHT11, SensorDHT22)
		return
	}

	var data [5]byte
	data, err = ReadSensor(mtype, pin)
	if err != nil {
		return
	}

	if stype == SensorDHT11 {
		humidity = float32(data[0])
		temperature = float32(data[2])
	} else {
		humidity = float32(int(data[0])*256+int(data[1])) / 10.0
		temperature = float32(int(data[2]&0x7F)*256+int(data[3])) / 10.0
		if data[2]&0x80 > 0 {
			temperature *= -1.0
		}
	}
	return
}

// ReadSensor returns the raw bit sequence read from the GPIO pin attached to the
// data pin of the DHT sensors.
func ReadSensor(model, pin int) (data [5]byte, err error) {
	var res C.int
	if model == 1 {
		res = C.pi_dht_read(C.int(pin), (*C.uint8_t)(&data[0]))
	} else {
		res = C.pi_2_dht_read(C.int(pin), (*C.uint8_t)(&data[0]))
	}
	if res == C.DHT_ERROR_GPIO {
		err = errors.New("could not open gpio device")
		return
	} else if res == C.DHT_ERROR_TIMEOUT {
		err = errors.New("got timeout while reading from sensor")
		return
	} else if int(data[4]) != ((int(data[0]) + int(data[1]) + int(data[2]) + int(data[3])) & 0xFF) {
		err = errors.New("checksum does not match")
		return
	}

	return
}
