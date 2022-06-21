package number_stairs

import (
	"fmt"
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
		step    uint64
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
				value: gofloat.ToFloat(test.value, 2),
			}) == test.invalid {
				t.Errorf("block add fail")
			}
		})
	}
}

func TestStair_PositionBlock(t *testing.T) {
	tests := map[string]struct {
		step    uint64
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
				value: gofloat.ToFloat(test.value, 2),
			}) == test.invalid {
				t.Errorf("block position fail")
			}
		})
	}
}

func Test_BlockFunction(t *testing.T) {
	st := NewStair()
	var step uint64 = 1
	// add three steps
	st.AddStep(step)

	st.AddBlock(step, NewBlock(`block-1`, gofloat.ToFloat(20, 2)))

	// block position in block-2
	if st.PositionBlock(NewBlock(`position`, gofloat.ToFloat(-10, 2)), StepFunction(func(step *Step) {
		for i := range step.Inputs {
			step.Inputs[i].ID = `position-fixed`
		}
	})) {
		// print position-fixed
		fmt.Println(st.Steps[step].Inputs[0].ID)
	}
}
