package number_stairs

import "github.com/senpathi/gofloat"

type Step struct {
	id             uint64
	Inputs         []Block
	Next, Previous *Step
	handrail       *Handrail
}

func (s *Step) ID() uint64 {
	return s.id
}

func (s *Step) Height() gofloat.Float {
	return s.handrail.Height
}

type Block struct {
	ID string
	Options
	value gofloat.Float
}

func NewBlock(ID string, value gofloat.Float) Block {
	return Block{
		ID:    ID,
		value: value,
	}
}

type Handrail struct {
	Height gofloat.Float
}

type Stair struct {
	End   *Step
	Steps map[uint64]*Step
	StairCaseOption
}

func NewStair(options ...StairOption) StairCase {
	stair := &Stair{
		Steps:           make(map[uint64]*Step),
		StairCaseOption: new(StairOptions),
	}
	stair.Apply(options...)

	return stair
}

func (s *Stair) AddStep(name uint64) bool {
	step := &Step{
		id:       name,
		Inputs:   *new([]Block),
		handrail: new(Handrail),
	}
	if _, ok := s.Steps[name]; ok || name <= 0 {
		return false
	}
	s.Steps[name] = step
	if s.End == nil {
		s.End = step
	} else {
		if s.End.id > name {
			return false
		}
		step.handrail.Height = s.End.Height()
		s.End.Previous = step
		step.Next = s.End
		s.End = step
	}
	return true
}

func (s *Stair) AddBlock(stepName uint64, block Block, options ...Option) bool {
	if st, ok := s.Steps[stepName]; !ok {
		return ok
	} else {
		if val := st.handrail.Height.Add(block.value); val.Float64() < s.GetMargin() && block.value.Float64() < 0 {
			return false
		} else {
			block.Options.apply(options...)
			st.Inputs = append(st.Inputs, block)
			st.handrail.Height = st.handrail.Height.Add(block.value)
			step := st.Previous
			if block.fn != nil {
				block.fn(st)
			}
			for step != nil {
				step.handrail.Height = step.handrail.Height.Add(block.value)
				step = step.Previous
			}
			return ok
		}
	}
}

func (s *Stair) PositionBlock(block Block, options ...Option) (stepID uint64, ok bool) {
	block.Options.apply(options...)
	if s.End.handrail.Height.Add(block.value).Float64() < s.GetMargin() {
		return 0, false
	}

	return fitBlock(s.End, block, s.GetMargin(), false), true
}

func (s *Stair) PositionBlockCheck(block Block, options ...Option) (stepID uint64, ok bool) {
	block.Options.apply(options...)
	if s.End.handrail.Height.Add(block.value).Float64() < s.GetMargin() {
		return 0, false
	}

	return fitBlock(s.End, block, s.GetMargin(), true), true
}

func fitBlock(step *Step, block Block, margin float64, debug bool) (stepID uint64) {
	stepUpdate := func(step *Step) {
		step.Inputs = append(step.Inputs, block)
		step.handrail.Height = step.handrail.Height.Add(block.value)
		s := step.Previous
		if block.fn != nil {
			block.fn(step)
		}
		for s != nil {
			s.handrail.Height = s.handrail.Height.Add(block.value)
			s = s.Previous
		}
	}
	if step.Next == nil ||
		step.Next.handrail.Height.Add(block.value).Float64() < margin ||
		step.id <= block.offset {
		if !debug {
			stepUpdate(step)
		}

		return step.id
	}
	return fitBlock(step.Next, block, margin, debug)
}

func (s *Stair) GetSteps() map[uint64]*Step {
	return s.Steps
}

func (s *Stair) GetEndStep() *Step {
	return s.End
}
