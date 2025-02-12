package coldp_test

import (
	"testing"

	"github.com/gnames/coldp/ent/coldp"
	"github.com/stretchr/testify/assert"
)

func TestToBool(t *testing.T) {
	assert := assert.New(t)
	tests := []struct {
		input         string
		valid, output bool
	}{
		{"smth", false, false},
		{"", false, false},
		{"true", true, true},
		{"TRUE", true, true},
		{"tRUe", true, true},
		{"T", true, true},
		{"yes", true, true},
		{"y", true, true},
		{"1", true, true},
		{"0", true, false},
		{"01", true, false},
		{"false", true, false},
		{"FAlse", true, false},
		{"no", true, false},
		{"N", true, false},
	}

	for _, v := range tests {
		res := coldp.ToBool(v.input)
		assert.Equal(v.output, res.Bool, v.input+"-bool")
		assert.Equal(v.valid, res.Valid, v.input+"-valid")
	}
}
