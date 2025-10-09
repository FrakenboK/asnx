package options

type Options struct {
	IPs         []string
	IPRangeFile string
	Version     bool
}

func New() *Options {
	return &Options{}
}
