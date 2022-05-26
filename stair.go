package number_stairs

import "github.com/senpathi/gofloat"

type Step struct {
	ID             string
	Inputs         []Block
	Next, Previous *Step
	Handrail       *Handrail
}

type Block struct {
	ID    string
	Value gofloat.Float
}

type Handrail struct {
	Height gofloat.Float
}

type Stair struct {
	End   *Step
	Steps map[string]*Step
}

func NewStair() *Stair {
	return &Stair{
		Steps: make(map[string]*Step),
	}
}

func (s *Stair) AddStep(name string) {
	step := &Step{
		ID:       name,
		Inputs:   *new([]Block),
		Handrail: new(Handrail),
	}
	if _, ok := s.Steps[name]; ok {
		return
	}
	s.Steps[name] = step
	if s.End == nil {
		s.End = step
	} else {
		s.End.Previous = step
		step.Next = s.End
		s.End = step
	}
}

func (s *Stair) AddBlock(stepName string, block Block) bool {
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

func (s *Stair) PositionBlock(block Block) bool {
	if s.End.Handrail.Height.Add(block.Value).Float64() < 0 {
		return false
	}
	fitBlock(s.End, block)
	return true
}

func fitBlock(step *Step, block Block) {
	if step.Next == nil {
		step.Inputs = append(step.Inputs, block)
		step.Handrail.Height = step.Handrail.Height.Add(block.Value)
		s := step.Previous
		for s != nil {
			s.Handrail.Height = s.Handrail.Height.Add(block.Value)
			s = s.Previous
		}
		return
	}
	if step.Next.Handrail.Height.Add(block.Value).Float64() < 0 {
		if step.Handrail.Height.Float64() >= block.Value.Float64() {
			step.Inputs = append(step.Inputs, block)
			step.Handrail.Height = step.Handrail.Height.Add(block.Value)
			s := step.Previous
			for s != nil {
				s.Handrail.Height = s.Handrail.Height.Add(block.Value)
				s = s.Previous
			}
		}
		return
	}
	fitBlock(step.Next, block)
}
