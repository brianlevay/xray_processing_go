package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func readConfigToMap(filepath string) (map[string]float64, error) {
	cfgFlt := make(map[string]float64)
	var linePts []string
	var key, valueStr string
	var valueFlt float64
	var errConv error
	fileBytes, errRead := ioutil.ReadFile(filepath)
	if errRead != nil {
		return cfgFlt, errRead
	}
	fileString := string(fileBytes)
	fileLines := strings.Split(fileString, "\n")
	for k := 0; k < len(fileLines); k++ {
		if strings.Contains(fileLines[k], "#") == false {
			linePts = strings.Split(fileLines[k], ":")
			if len(linePts) == 2 {
				key = strings.Trim(linePts[0], " \t")
				valueStr = strings.Trim(linePts[1], " \t\r")
				valueFlt, errConv = strconv.ParseFloat(valueStr, 64)
				if errConv == nil {
					cfgFlt[key] = valueFlt
				}
			}
		}
	}
	return cfgFlt, nil
}