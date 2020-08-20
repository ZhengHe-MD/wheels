package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringSliceToSet(t *testing.T) {
	tests := []struct{
		givenSlice []string
		wantSet map[string]interface{}
	} {
		{
			[]string{"a", "b", "c"},
			map[string]interface{}{
				"a": struct {}{},
				"b": struct {}{},
				"c": struct {}{},
			},
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.wantSet, StringSliceToSet(tt.givenSlice))
	}
}

func TestFilterStringSlice(t *testing.T) {
	tests := []struct{
		givenSlice []string
		givenFilter func(string) bool
		wantSlice []string
	} {
		{
			[]string{"a", "b", "c"},
			func(ele string) bool {
				return ele != "a";
			},
			[]string{"b", "c"},
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.wantSlice, FilterStringSlice(tt.givenSlice, tt.givenFilter))
	}
}
