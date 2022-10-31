package number_stairs

type StairCase interface {
	AddStep(name uint64) bool
	AddBlock(stepName uint64, block Block, options ...Option) bool
	PositionBlock(block Block, options ...Option) (stepID uint64, ok bool)
	PositionBlockCheck(block Block, options ...Option) (stepID uint64, ok bool)
	GetSteps() map[uint64]*Step
	StairCaseOption
}

type StairCaseOption interface {
	GetMargin() float64
	Apply(...StairOption)
}
