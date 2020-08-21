package collection

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringSliceToSet(t *testing.T) {
	tests := []struct {
		givenSlice []string
		wantSet    map[string]interface{}
	}{
		{
			[]string{"a", "b", "c"},
			map[string]interface{}{
				"a": struct{}{},
				"b": struct{}{},
				"c": struct{}{},
			},
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.wantSet, StringSliceToSet(tt.givenSlice))
	}
}

func TestStringSliceFilter(t *testing.T) {
	tests := []struct {
		givenSlice  []string
		givenFilter func(string) bool
		wantSlice   []string
	}{
		{
			[]string{"a", "b", "c"},
			func(ele string) bool {
				return ele != "a"
			},
			[]string{"b", "c"},
		},
	}

	for _, tt := range tests {
		assert.Equal(t, tt.wantSlice, StringSliceFilter(tt.givenSlice, tt.givenFilter))
	}
}

func TestStringSliceForEach(t *testing.T) {
	type item struct {
		i   int
		ele string
		raw []string
	}

	tests := []struct {
		givenSlice []string
		givenDo    func(int, string, []string)
	}{
		{
			[]string{"a", "b", "c"},
			func(i int, ele string, raw []string) {
				items := []item{
					{0, "a", []string{"a", "b", "c"}},
					{1, "b", []string{"a", "b", "c"}},
					{2, "c", []string{"a", "b", "c"}},
				}
				assert.Equal(t, i, items[i].i)
				assert.Equal(t, ele, items[i].ele)
				assert.Equal(t, raw, items[i].raw)
			},
		},
	}

	for _, tt := range tests {
		StringSliceForEach(tt.givenSlice, tt.givenDo)
	}
}
