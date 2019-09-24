package testlib

import (
	"log"
	"reflect"
	"strconv"
)

type TestValidator struct {
	Response TestResponse
}

func (tv TestValidator) ExpectResponseStatus(status string) bool {
	s := strconv.Itoa(tv.Response.StatusCode)
	if string(status[0]) == "!" {
		return tv.ExpectValueNotEqual(status[1:], s)
	} else {
		return tv.ExpectValueEqual(status, s)
	}
}

func (tv TestValidator) ExpectBody(body map[string]interface{}) bool {
	for k, v := range body {
		val := tv.Response.RetrieveValueFromBody(k, body)
		if val != nil && reflect.TypeOf(v) == reflect.TypeOf(val) {
			return tv.ExpectValueEqual(v, val)
		}
	}
	return true
}

func (tv TestValidator) ExpectValueEqual(expected interface{}, actual interface{}) bool {
	return expected == actual
}

func (tv TestValidator) ExpectValueNotEqual(expected interface{}, actual interface{}) bool {
	return expected != actual
}

func (tv TestValidator) ExpectNotNilAndSave(key string) interface{} {
	body, err := tv.Response.ParseBody()
	if err != nil {
		log.Fatalf("Failed to encode response body: %v\n", err)
	}
	value := tv.Response.RetrieveValueFromBody(key, body)
	if value != "" {
		return value
	}
	return nil
}
