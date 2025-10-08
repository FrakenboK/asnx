package options

type Options struct {
	IPs             []string
	FullNetworkEnum bool
	Version         bool
}

func New() *Options {
	return &Options{}
}
