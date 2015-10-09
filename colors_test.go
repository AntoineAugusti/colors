package colors

import (
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

func TestisHexadecimalCharacter(t *testing.T) {
	assert.True(t, isHexadecimalCharacter("A"))
	assert.True(t, isHexadecimalCharacter("B"))
	assert.True(t, isHexadecimalCharacter("C"))
	assert.True(t, isHexadecimalCharacter("F"))
	assert.True(t, isHexadecimalCharacter("5"))
	assert.True(t, isHexadecimalCharacter("4"))
}

func TestisHexadecimal(t *testing.T) {
	assert.True(t, isHexadecimal("ffffff"))
	assert.True(t, isHexadecimal("FFaa65"))
	assert.True(t, isHexadecimal("c769fe"))

	assert.False(t, isHexadecimal("fffff"))
	assert.False(t, isHexadecimal("c769fea"))
	assert.False(t, isHexadecimal("c769fg"))
}
