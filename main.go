package main

import (
    "fmt"
    "reflect"
)

type myStruct struct {
    FirstSlice []string `yaml:"foo,flow"`
    SecondSlice []string `yaml:"bar,flow"`
}

func sliceDisplayer(s []string) {
    fmt.Println("Now in sliceDisplayer()")
    for _, v := range s {
        fmt.Println(v)
    }
}

func main() {
    firstStruct := myStruct{}
    firstStruct.FirstSlice = []string{"FirstSliceVal1", "FirstSliceVal2"}
    firstStruct.SecondSlice = []string{"SecondSliceVal1", "SecondSliceVal2"}
    
    v := reflect.ValueOf(firstStruct)
    typeOfV := v.Type()
    for i := 0; i < v.NumField(); i++ {
        fmt.Printf("Field: %v, Value: %v\n", typeOfV.Field(i).Name, v.Field(i).Interface())
        values := v.Field(i).Interface().([]string)
        fmt.Printf("values: %v, values type: %T\n", values, values)
        sliceDisplayer(values) 
    }
}
