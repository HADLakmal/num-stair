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
			if _, ok := st.PositionBlock(Block{
				value: gofloat.ToFloat(test.value, 2),
			}); ok == test.invalid {
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
	if _, ok := st.PositionBlock(NewBlock(`position`, gofloat.ToFloat(-10, 2)), StepFunction(func(step *Step) {
		for i := range step.Inputs {
			step.Inputs[i].ID = `position-fixed`
		}
	})); ok {
		// print position-fixed
		fmt.Println(st.Steps[step].Inputs[0].ID)
	}
}

func TestStair_PositionBlockCheck(t *testing.T) {
	tests := map[string]struct {
		steps      []uint64
		stepValues []float64
		value      float64
		output     uint64
	}{
		`valid block`: {
			steps:      []uint64{1},
			stepValues: []float64{20},
			value:      10,
			output:     1,
		},
		`two steps`: {
			steps:      []uint64{1, 2},
			stepValues: []float64{5, 10},
			value:      -5,
			output:     1,
		},
		`first step with lower value`: {
			steps:      []uint64{1, 2},
			stepValues: []float64{3, 3},
			value:      -5,
			output:     2,
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			st := NewStair()
			for _, sp := range test.steps {
				if !st.AddStep(sp) {
					t.Errorf("step add fail")
				}
			}
			for i, sp := range test.steps {
				st.AddBlock(sp, NewBlock(fmt.Sprintf(`%d`, i), gofloat.ToFloat(test.stepValues[i], 2)))
			}

			if val, _ := st.PositionBlockCheck(Block{
				value: gofloat.ToFloat(test.value, 2),
			}); val != test.output {
				t.Errorf("value got : %v, but expect : %v", val, test.output)
			}
		})
	}
}
