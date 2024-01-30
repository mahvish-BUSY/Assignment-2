/*Consider this as your input string. You have to iterate it and print the type and value of the entity.
If it is a data structure then go inside it and do the same. You have to keep on going nested until you reach data types.
(Do this using reflect concept)*/

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func printValueAndType(data interface{}) {
	dataVal := reflect.ValueOf(data)
	dataTyp := reflect.TypeOf(data)

	switch dataVal.Kind() {
	case reflect.Map:
		Keys := dataVal.MapKeys()
		for _, mapKey := range Keys {
			mapVal := dataVal.MapIndex(mapKey)
			printValueAndType(mapVal.Interface())
		}
	case reflect.Slice:
		for i := 0; i < dataVal.Len(); i++ {
			printValueAndType(dataVal.Index(i).Interface())
		}

	case reflect.Struct:
		for i := 0; i < dataVal.NumField(); i++ {
			printValueAndType(dataVal.Field(i).Interface())
		}
	}
	//print the value and type of current data
	fmt.Printf("Type : %s   Value : %v\n", dataTyp, dataVal)

}
func main() {

	input := `{
        "name": "Tolexo Online Pvt. Ltd",
        "age_in_years": 8.5,
        "origin": "Noida",
        "head_office": "Noida, Uttar Pradesh",
		"address": [
	    {
	        "street": "91 Springboard",
	        "landmark": "Axis Bank",
	        "city": "Noida",
	        "pincode": 201301,
	        "state": "Uttar Pradesh"
	    },
	    {
	        "street": "91 Springboard",
	        "landmark": "Axis Bank",
	        "city": "Noida",
	        "pincode": 201301,
	        "state": "Uttar Pradesh"
	    }
	],
        "sponsors": {
     	 "name": "One"
        },
        "revenue": "19.8 million$",
        "no_of_employee": 630,
        "str_text": ["one", "two"],
        "int_text": [1, 3, 4]
    }`

	myMap := make(map[string]interface{})

	err := json.Unmarshal([]byte(input), &myMap)

	if err == nil {

		printValueAndType(myMap)
	} else {
		return
	}

}
