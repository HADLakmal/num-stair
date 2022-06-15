package number_stairs

import "github.com/senpathi/gofloat"

type Step struct {
	ID             int64
	Inputs         []Block
	Next, Previous *Step
	Handrail       *Handrail
}

type Block struct {
	ID string
	Options
	Value gofloat.Float
}

type Handrail struct {
	Height gofloat.Float
}

type Stair struct {
	End   *Step
	Steps map[int64]*Step
}

func NewStair() *Stair {
	return &Stair{
		Steps: make(map[int64]*Step),
	}
}

func (s *Stair) AddStep(name int64) bool {
	step := &Step{
		ID:       name,
		Inputs:   *new([]Block),
		Handrail: new(Handrail),
	}
	if _, ok := s.Steps[name]; ok {
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

func (s *Stair) AddBlock(stepName int64, block Block) bool {
	if st, ok := s.Steps[stepName]; !ok {
		return ok
	} else {
		if val := st.Handrail.Height.Add(block.Value); val.Float64() < 0 {
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
	if s.End.Handrail.Height.Add(block.Value).Float64() < 0 {
		return false
	}
	fitBlock(s.End, block)
	return true
}

func fitBlock(step *Step, block Block) {
	stepUpdate := func(step *Step) {
		step.Inputs = append(step.Inputs, block)
		step.Handrail.Height = step.Handrail.Height.Add(block.Value)
		s := step.Previous
		for s != nil {
			s.Handrail.Height = s.Handrail.Height.Add(block.Value)
			s = s.Previous
		}
	}
	if step.Next == nil {
		stepUpdate(step)
		return
	}
	if step.Next.Handrail.Height.Add(block.Value).Float64() < 0 {
		stepUpdate(step)
		return
	} else if step.ID <= block.offset {
		stepUpdate(step)
		return
	}
	fitBlock(step.Next, block)
}

type Options struct {
	offset int64
}
type Option func(*Options)

func Offset(offset int64) Option {
	return func(options *Options) {
		options.offset = offset
	}
}

func (opt *Options) apply(options ...Option) {
	for _, option := range options {
		option(opt)
	}
}
