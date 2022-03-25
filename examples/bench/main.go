package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	sypexgeo "github.com/tenrok/go-sxgeo"
)

func main() {
	geo := sypexgeo.New("../SxGeoCity.dat")

	t := time.Now()
	for i := 0; i < 1000000; i++ {
		ip := strconv.Itoa(rand.Intn(222)) + "." + strconv.Itoa(rand.Intn(222)) + "." + strconv.Itoa(rand.Intn(222)) + "." + strconv.Itoa(rand.Intn(222))
		geo.Info(ip)
	}
	fmt.Printf("The call took %v to run.\n", time.Now().Sub(t))
}
