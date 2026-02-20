package runner

import (
	"github.com/FrakenboK/asnx/internal/logger"
	"github.com/FrakenboK/asnx/internal/options"
	"github.com/FrakenboK/asnx/internal/resolver"
	extractor "github.com/FrakenboK/asnx/internal/runner/value-extractor"
	"github.com/FrakenboK/asnx/internal/ui"

	"github.com/spf13/cobra"
)

type Runner struct {
	opts *options.Options
	rdap *resolver.Resolver
	log  *logger.Logger
}

func (r *Runner) Start(cmd *cobra.Command, args []string) {
	if r.opts.Version || len(r.opts.Hosts) == 0 { // TODO: all options
		cmd.Help()
		return
	}

	if !r.opts.BannerDisabled {
		r.log.RawLog(ui.GetBanner())
	}

	// TODO: all options
	r.log.Note("Searching for ASNs...")
	ext := extractor.New(r.log)
	ips := ext.ExtractValues(r.opts.Hosts)
	r.rdap.HandleIPs(ips, r.opts.IPRangeFile)
}

func New(opts *options.Options) *Runner {
	logger := &logger.Logger{}

	rdap := resolver.NewClient(logger)

	return &Runner{
		opts: opts,
		rdap: rdap,
		log:  logger,
	}
}
