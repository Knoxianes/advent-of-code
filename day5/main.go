package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var minLocation = -1
var numberOfCores = 8
var waitGroup sync.WaitGroup

func main() {
	dat, _ := os.ReadFile("tests.txt")
	splited := strings.Split(string(dat), "\n")
	runtime.GOMAXPROCS(numberOfCores)
	for _, value := range splited {
		if value == "" {
			continue
		}
		value_splitted := strings.Split(value, " ")
		start_seed, _ := strconv.Atoi(value_splitted[0])
		len, _ := strconv.Atoi(value_splitted[1])
		divided := len / numberOfCores
	
		for i := 0; i < numberOfCores; i++ {
			waitGroup.Add(1)
			go Task(start_seed+divided*i, divided)
		}
		waitGroup.Wait()
	}
	fmt.Println(minLocation)
}
func Task(start_seed int, len int) {
	defer waitGroup.Done()
	var convertedValue int
	for i := 0; i < len; i++ {
		convertedValue = start_seed + i
		location := SeedToSoil(convertedValue);
		if minLocation == -1 {
			minLocation = location
			continue
		}
		if location < minLocation {
			minLocation = location
		}
	}
}
func SeedToSoil(seedNumber int) int {
	//dat, _ := os.ReadFile("seedtosoil.txt")
	dat, _ := os.ReadFile("ststest.txt")
	splited := strings.Split(string(dat), "\n")
	for _, value := range splited {
		if value == "" {
			continue
		}
		value_splited := strings.Split(value, " ")
		desc_start, _ := strconv.Atoi(value_splited[0])
		src_start, _ := strconv.Atoi(value_splited[1])
		len, _ := strconv.Atoi(value_splited[2])
		if src_start > seedNumber || src_start+len <= seedNumber {
			continue
		}
		return SoilToFert((seedNumber - src_start) + desc_start)

	}
	return SoilToFert(seedNumber)
}
func SoilToFert(soilNumber int) int {
	//dat, _ := os.ReadFile("soiltofertil.txt")
	dat, _ := os.ReadFile("stftest.txt")
	splited := strings.Split(string(dat), "\n")
	for _, value := range splited {
		if value == "" {
			continue
		}
		value_splited := strings.Split(value, " ")
		desc_start, _ := strconv.Atoi(value_splited[0])
		src_start, _ := strconv.Atoi(value_splited[1])
		len, _ := strconv.Atoi(value_splited[2])
		if src_start > soilNumber || src_start+len <= soilNumber {
			continue
		}
		return FertToWater((soilNumber - src_start) + desc_start)

	}
	return FertToWater(soilNumber)
}
func FertToWater(fertNumber int) int {
	//dat, _ := os.ReadFile("ferttowater.txt")
	dat, _ := os.ReadFile("ftwtest.txt")
	splited := strings.Split(string(dat), "\n")
	for _, value := range splited {
		if value == "" {
			continue
		}
		value_splited := strings.Split(value, " ")
		desc_start, _ := strconv.Atoi(value_splited[0])
		src_start, _ := strconv.Atoi(value_splited[1])
		len, _ := strconv.Atoi(value_splited[2])
		if src_start > fertNumber || src_start+len <= fertNumber {
			continue
		}
		return WaterToLight((fertNumber - src_start) + desc_start)

	}
	return WaterToLight(fertNumber)
}
func WaterToLight(waterNumber int) int {
	//dat, _ := os.ReadFile("watertolight.txt")
	dat, _ := os.ReadFile("wtltest.txt")
	splited := strings.Split(string(dat), "\n")
	for _, value := range splited {
		if value == "" {
			continue
		}
		value_splited := strings.Split(value, " ")
		desc_start, _ := strconv.Atoi(value_splited[0])
		src_start, _ := strconv.Atoi(value_splited[1])
		len, _ := strconv.Atoi(value_splited[2])
		if src_start > waterNumber || src_start+len <= waterNumber {
			continue
		}
		return LightToTemp((waterNumber - src_start) + desc_start)

	}
	return waterNumber
}
func LightToTemp(lightNumber int) int {
	//dat, _ := os.ReadFile("lighttotemp.txt")
	dat, _ := os.ReadFile("ltttest.txt")
	splited := strings.Split(string(dat), "\n")
	for _, value := range splited {
		if value == "" {
			continue
		}
		value_splited := strings.Split(value, " ")
		desc_start, _ := strconv.Atoi(value_splited[0])
		src_start, _ := strconv.Atoi(value_splited[1])
		len, _ := strconv.Atoi(value_splited[2])
		if src_start > lightNumber || src_start+len <= lightNumber {
			continue
		}
		return TempToHumidity((lightNumber - src_start) + desc_start)

	}
	return TempToHumidity(lightNumber)
}
func TempToHumidity(tempNumber int) int {
	//dat, _ := os.ReadFile("temptohumidity.txt")
	dat, _ := os.ReadFile("tthtest.txt")
	splited := strings.Split(string(dat), "\n")
	for _, value := range splited {
		if value == "" {
			continue
		}
		value_splited := strings.Split(value, " ")
		desc_start, _ := strconv.Atoi(value_splited[0])
		src_start, _ := strconv.Atoi(value_splited[1])
		len, _ := strconv.Atoi(value_splited[2])
		if src_start > tempNumber || src_start+len <= tempNumber {
			continue
		}
		return HumidityToLocation((tempNumber - src_start) + desc_start)

	}
	return HumidityToLocation(tempNumber)
}
func HumidityToLocation(humidityNumber int) int {
	//dat, _ := os.ReadFile("humiditytolocation.txt")
	dat, _ := os.ReadFile("htltest.txt")
	splited := strings.Split(string(dat), "\n")
	for _, value := range splited {
		if value == "" {
			continue
		}
		value_splited := strings.Split(value, " ")
		desc_start, _ := strconv.Atoi(value_splited[0])
		src_start, _ := strconv.Atoi(value_splited[1])
		len, _ := strconv.Atoi(value_splited[2])
		if src_start > humidityNumber || src_start+len <= humidityNumber {
			continue
		}
		return (humidityNumber - src_start) + desc_start

	}
	return humidityNumber
}
