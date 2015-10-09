package colors

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToHexColor(t *testing.T) {
	assert.Equal(t, "#ffffff", StringToHexColor("white"))
	assert.Equal(t, "transparent", StringToHexColor("transparent"))
	assert.Equal(t, "#ffffff", StringToHexColor("FFFFFF"))
	assert.Equal(t, "#ffffff", StringToHexColor("#FFFFFF"))

	assert.Equal(t, "#c00000", StringToHexColor("chucknorris"))
	assert.Equal(t, "#ad0e0e", StringToHexColor("adamlevine"))
}

func TestIsReserved(t *testing.T) {
	res, value := IsReserved("transparent")
	assert.True(t, res)
	assert.Equal(t, "transparent", value)

	res, value = IsReserved(" ffffff ")
	assert.True(t, res)
	assert.Equal(t, "#ffffff", value)

	res, value = IsReserved("#ffffff")
	assert.True(t, res)
	assert.Equal(t, "#ffffff", value)
	res, value = IsReserved(" RED ")
	assert.True(t, res)
	assert.Equal(t, "#ff0000", value)

	res, _ = IsReserved("foo")
	assert.False(t, res)
	res, _ = IsReserved("c769fea")
	assert.False(t, res)
}

func TestIsHexadecimalCharacter(t *testing.T) {
	assert.True(t, isHexadecimalCharacter("A"))
	assert.True(t, isHexadecimalCharacter("B"))
	assert.True(t, isHexadecimalCharacter("C"))
	assert.True(t, isHexadecimalCharacter("F"))
	assert.True(t, isHexadecimalCharacter("5"))
	assert.True(t, isHexadecimalCharacter("4"))
}

func TestIsHexadecimal(t *testing.T) {
	assert.True(t, isHexadecimal("ffffff"))
	assert.True(t, isHexadecimal("FFaa65"))
	assert.True(t, isHexadecimal("c769fe"))

	assert.False(t, isHexadecimal("fffff"))
	assert.False(t, isHexadecimal("c769fea"))
	assert.False(t, isHexadecimal("c769fg"))
}

func ExampleStringToHexColor_namedColor() {
	fmt.Println(StringToHexColor("white"))
	// Output: #ffffff
}

func ExampleStringToHexColor_hexadecimalString() {
	fmt.Println(StringToHexColor("FFFFFF"))
	// Output: #ffffff
}

func ExampleStringToHexColor_otherHexadecimalString() {
	fmt.Println(StringToHexColor("#FFFFFF"))
	// Output: #ffffff
}

func ExampleStringToHexColor_randomString() {
	fmt.Println(StringToHexColor("chucknorris"))
	// Output: #c00000
}

func ExampleStringToHexColor_otherNamedColor() {
	fmt.Println(StringToHexColor("limegreen"))
	// Output: #32cd32
}

func ExampleIsReserved_namedColor() {
	reserved, value := IsReserved("red")
	fmt.Fprintf(os.Stdout, "reserved: %t, value: %s", reserved, value)
	// Output: reserved: true, value: #ff0000
}

func ExampleIsReserved_hexadecimalColor() {
	reserved, value := IsReserved(" ffFFff  ")
	fmt.Fprintf(os.Stdout, "reserved: %t, value: %s", reserved, value)
	// Output: reserved: true, value: #ffffff
}

func ExampleIsReserved_unknowWord() {
	reserved, value := IsReserved("foo")
	fmt.Fprintf(os.Stdout, "reserved: %t, value: %s", reserved, value)
	// Output: reserved: false, value:
}
