package enhance

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestEnhance(t *testing.T) {
	expected := Output{
		Name: "Constantin is Awesome!",
	}

	check := Enhance("Constantin")

	assert.Equal(t, check, expected)
}
