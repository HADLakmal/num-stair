package number_stairs

import (
	"github.com/senpathi/gofloat"
	"testing"
)

func TestNewStair(t *testing.T) {
	st := NewStair()
	if !st.AddStep(1) {
		t.Errorf("step add fail")
	}
}

func TestNewStairMargin(t *testing.T) {
	margin := 10.0
	st := NewStair(Margin(margin))
	if !st.AddStep(1) && st.margin != margin {
		t.Errorf("step add fail")
	}
}

func TestStair_AddBlock(t *testing.T) {
	tests := map[string]struct {
		step    int64
		value   float64
		invalid bool
	}{
		`valid block`: {
			step:  1,
			value: 10,
		},
		`below the margin`: {
			step:    1,
			value:   -5,
			invalid: true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			st := NewStair()
			if !st.AddStep(test.step) {
				t.Errorf("step add fail")
			}
			if st.AddBlock(test.step, Block{
				Value: gofloat.ToFloat(test.value, 2),
			}) == test.invalid {
				t.Errorf("block add fail")
			}
		})
	}
}

func TestStair_PositionBlock(t *testing.T) {
	tests := map[string]struct {
		step    int64
		value   float64
		invalid bool
	}{
		`valid block`: {
			step:  1,
			value: 10,
		},
		`below the margin`: {
			step:    1,
			value:   -5,
			invalid: true,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			st := NewStair()
			if !st.AddStep(test.step) {
				t.Errorf("step add fail")
			}
			if st.PositionBlock(Block{
				Value: gofloat.ToFloat(test.value, 2),
			}) == test.invalid {
				t.Errorf("block position fail")
			}
		})
	}
}
