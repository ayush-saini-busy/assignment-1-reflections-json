package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func inspect(v reflect.Value) {
	switch v.Kind() {
	case reflect.Interface:
		inspect(v.Elem())
	case reflect.Map:
		fmt.Println("Map detected")
		for _, key := range v.MapKeys() {
			fmt.Printf("Key (%s): %v\n", key.Kind(), key.Interface())
			inspect(v.MapIndex(key))
		}
	case reflect.Slice:
		fmt.Println("Slice detected")
		for i := 0; i < v.Len(); i++ {
			inspect(v.Index(i))
		}

	default:
		fmt.Printf("Type: %s | Value: %v\n", v.Kind(), v.Interface())
	}

}

func main() {
	input := `{
	"name" : "Tolexo Online Pvt. Ltd",
	"age_in_years" : 8.5,
	"origin" : "Noida",
	"head_office" : "Noida, Uttar Pradesh",
	"address" : [
	{
	"street" : "91 Springboard",
	"landmark" : "Axis Bank",
	"city" : "Noida",
	"pincode" : 201301,
	"state" : "Uttar Pradesh"
	},
	{
	"street" : "91 Springboard",
	"landmark" : "Axis Bank",
	"city" : "Noida",
	"pincode" : 201301,
	"state" : "Uttar Pradesh"
	}
	],
	"sponsers" : {
	"name" : "One"
	},
	"revenue" : "19.8 million$",
	"no_of_employee" : 630,
	"str_text" : ["one","two"],
	"int_text" : [1,3,4]
	}`
	var data interface{}
	if err := json.Unmarshal([]byte(input), &data); err != nil {
		fmt.Println("An error occured while unmarshalling the file")
	}

	fmt.Println("*** Inspecting the JSON ***")
	inspect(reflect.ValueOf(data))
}
