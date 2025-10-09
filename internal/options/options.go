package options

type Options struct {
	IPs     []string
	Domains []string

	IPRangeFile string
	Version     bool
}

func New() *Options {
	return &Options{}
}
