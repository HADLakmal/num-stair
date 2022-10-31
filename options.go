package number_stairs

type Options struct {
	offset uint64
	attach interface{}
	fn     func(step *Step)
}
type Option func(*Options)

func (opt *Options) GetAttachValue() interface{} {
	return opt.attach
}

func Offset(offset uint64) Option {
	return func(options *Options) {
		options.offset = offset
	}
}

func ValueAttach(attach interface{}) Option {
	return func(options *Options) {
		options.attach = attach
	}
}

func StepFunction(fn func(step *Step)) Option {
	return func(options *Options) {
		options.fn = fn
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

func (opt *StairOptions) Apply(options ...StairOption) {
	for _, option := range options {
		option(opt)
	}
}

func (opt *StairOptions) GetMargin() float64 {
	return opt.margin
}
