package options

type Options struct {
	Hosts []string

	IPRangeFile string
	OutputFile  string

	Version        bool
	BannerDisabled bool
}
