package main

type Basic struct {
	unique bool
	value  interface{}
}

type StringValue struct {
	Basic
	maxLength int32
	minLength int32
}

type NickName struct {
	StringValue
}

type FirstName struct {
	StringValue
}

type LastName struct {
	StringValue
}

type Email struct {
	Basic
	prefix string
	domain string
}

type IntegerValue struct {
	Basic
	autoInc bool
	start   int
	end     int
}

type FloatValue struct {
}
