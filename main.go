package main

import "fmt"

func main() {
	fmt.Println("Hello, mock data generator")

	var e *Email = new(Email)

	e.value = "hi"
	e.prefix = "bnb1083"
	e.domain = "google.com"

	fmt.Println(e.getValue())

}

func (b *Basic) getValue() interface{} {
	return b.value.(string)
}

func (e *Email) getValue() interface{} {
	e.value = e.prefix + "@" + e.domain
	return e.value.(string)
}

/*
config file format
json：
{
	row: (integer),
	columns:{
		"id" : {
			"type": "integer",
			"unique": true,
			""
		}
		"name":"string",
		""
	}
}

*/
