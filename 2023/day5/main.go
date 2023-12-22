package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var seeds = map[int]int{}
var seedtosoil = map[int]int{}
var soiltofert = map[int]int{}
var ferttowater = map[int]int{}
var watertolight = map[int]int{}
var lighttotemp = map[int]int{}
var temptohumidity = map[int]int{}
var humidifytolocation = map[int]int{}

func main() {
	addSeeds()
	addSeedToSoil()
	addSoilToFert()
	addFertToWater()
	addWaterToLight()
	addLightToTemp()
	addTempToHumidity()
	addHumidityToLocation()
}
func addSeeds(){
    dat, _ := os.ReadFile("seeds.txt")
    splited := strings.Split((strings.Split(string(dat),"\n")[0])," ")
    for _,value := range splited{
        if value == ""{
            continue
        }
        convertedValue , _ := strconv.Atoi(value)
        seeds[convertedValue] = convertedValue
    }
}
func addSeedToSoil(){
    dat, _ := os.ReadFile("seedtosoil.txt")
    splited := strings.Split(string(dat),"\n")
}
func addSoilToFert(){
}
func addFertToWater(){
}
func addWaterToLight(){
}
func addLightToTemp(){
}
func addTempToHumidity(){
}
func addHumidityToLocation(){
}
