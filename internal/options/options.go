package options

type Options struct {
	IPs     []string
	Domains []string

	IPRangeFile string
	OutputFile  string

	Version        bool
	BannerDisabled bool
}
