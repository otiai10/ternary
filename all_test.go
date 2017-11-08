package ternary

import (
	"os"
	"testing"

	. "github.com/otiai10/mint"
)

func TestIf(t *testing.T) {
	Expect(t, If(true)).TypeOf("ternary.Condition")
}

func TestCondition_String(t *testing.T) {

	message := If(true).String("Yes!", "No...")
	Expect(t, message).ToBe("Yes!")

	message = If(false).String("Yes!", "No...")
	Expect(t, message).ToBe("No...")

}

func TestCondition_Int(t *testing.T) {

	status := If(true).Int(200, 500)
	Expect(t, status).ToBe(200)

	status = If(false).Int(200, 500)
	Expect(t, status).ToBe(500)

}

func TestCondition_Interface(t *testing.T) {

	status := 200
	auth := If(status == 200).Interface("ok", nil)
	Expect(t, auth).ToBe("ok")

	status = 403
	auth = If(status == 200).Interface("ok", nil)
	Expect(t, auth).ToBe(nil)

}

func TestCondition_Place(t *testing.T) {

	response := map[string]interface{}{
		"message": "Hello, btw!",
	}

	status := 200
	If(status != 200).Put("error", "Failed")(response)
	_, exists := response["error"]
	Expect(t, exists).ToBe(false)

	status = 500
	If(status != 200).Put("error", "Failed")(response)
	value, exists := response["error"]
	Expect(t, exists).ToBe(true)
	Expect(t, value).ToBe("Failed")
}

func TestZero(t *testing.T) {
	port := Zero("").String("9090", "8080")
	Expect(t, port).ToBe("9090")
}

func TestDefault(t *testing.T) {
	port := Default("7777").String(os.Getenv("SPECIFIED_PORT"))
	Expect(t, port).ToBe("7777")
	port = Default(100).String(os.Getenv("SPECIFIED_PORT"))
	Expect(t, port).ToBe("")
}
