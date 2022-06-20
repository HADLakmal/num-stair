package number_stairs

import "github.com/senpathi/gofloat"

type Step struct {
	ID             uint64
	Inputs         []Block
	Next, Previous *Step
	Handrail       *Handrail
}

type Block struct {
	ID string
	Options
	Value gofloat.Float
}

func NewBlock(ID string, value gofloat.Float) Block {
	return Block{
		ID:    ID,
		Value: value,
	}
}

type Handrail struct {
	Height gofloat.Float
}

type Stair struct {
	End   *Step
	Steps map[uint64]*Step
	StairOptions
}

func NewStair(options ...StairOption) *Stair {
	stair := &Stair{
		Steps: make(map[uint64]*Step),
	}
	stair.StairOptions.apply(options...)

	return stair
}

func (s *Stair) AddStep(name uint64) bool {
	step := &Step{
		ID:       name,
		Inputs:   *new([]Block),
		Handrail: new(Handrail),
	}
	if _, ok := s.Steps[name]; ok || name <= 0 {
		return false
	}
	s.Steps[name] = step
	if s.End == nil {
		s.End = step
	} else {
		if s.End.ID > name {
			return false
		}
		s.End.Previous = step
		step.Next = s.End
		s.End = step
	}
	return true
}

func (s *Stair) AddBlock(stepName uint64, block Block) bool {
	if st, ok := s.Steps[stepName]; !ok {
		return ok
	} else {
		if val := st.Handrail.Height.Add(block.Value); val.Float64() < s.margin && block.Value.Float64() < 0 {
			return false
		} else {
			st.Inputs = append(st.Inputs, block)
			st.Handrail.Height = st.Handrail.Height.Add(block.Value)
			step := st.Previous
			for step != nil {
				step.Handrail.Height = step.Handrail.Height.Add(block.Value)
				step = step.Previous
			}
			return ok
		}
	}
}

func (s *Stair) PositionBlock(block Block, options ...Option) bool {
	block.Options.apply(options...)
	if s.End.Handrail.Height.Add(block.Value).Float64() < s.margin {
		return false
	}
	fitBlock(s.End, block, s.margin)
	return true
}

func fitBlock(step *Step, block Block, margin float64) {
	stepUpdate := func(step *Step) {
		step.Inputs = append(step.Inputs, block)
		step.Handrail.Height = step.Handrail.Height.Add(block.Value)
		s := step.Previous
		for s != nil {
			s.Handrail.Height = s.Handrail.Height.Add(block.Value)
			s = s.Previous
		}
	}
	if step.Next == nil ||
		step.Next.Handrail.Height.Add(block.Value).Float64() < margin ||
		step.ID <= block.offset {
		stepUpdate(step)
		return
	}
	fitBlock(step.Next, block, margin)
}

type Options struct {
	offset uint64
}
type Option func(*Options)

func Offset(offset uint64) Option {
	return func(options *Options) {
		options.offset = offset
	}
}

func (opt *Options) apply(options ...Option) {
	for _, option := range options {
		option(opt)
	}
}

type StairOptions struct {
	margin float64
}
type StairOption func(*StairOptions)

func Margin(offset float64) StairOption {
	return func(options *StairOptions) {
		options.margin = offset
	}
}

func (opt *StairOptions) apply(options ...StairOption) {
	for _, option := range options {
		option(opt)
	}
}
