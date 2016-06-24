package main

import (
	"fmt"
	"time"

	"github.com/goiot/devices/grove/accel3xdigital"
	"golang.org/x/exp/io/i2c"
)

func main() {
	accel, err := accel3xdigital.Open(&i2c.Devfs{Dev: "/dev/i2c-1"})
	if err != nil {
		panic(err)
	}

	// Enabling tap switches from 32 samples/second to 120
	err = accel.EnableTap()
	if err != nil {
		panic(err)
	}

	defer accel.Close()

	for i := 0; i < 20; i++ {
		accel.Update()
		fmt.Println(accel.State)
		time.Sleep(500 * time.Millisecond)
	}

}
