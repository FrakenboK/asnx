package runner

import (
	"github.com/FrakenboK/asnx/internal/logger"
	"github.com/FrakenboK/asnx/internal/options"
	"github.com/FrakenboK/asnx/internal/resolver"
	"github.com/FrakenboK/asnx/internal/runner/validator"
	extractor "github.com/FrakenboK/asnx/internal/runner/value-extractor"

	"github.com/spf13/cobra"
)

type Runner struct {
	opts *options.Options
	rdap *resolver.Resolver
	log  *logger.Logger
}

func (r *Runner) Start(cmd *cobra.Command, args []string) {
	if r.opts.Version {
		cmd.Help()
		return
	}

	r.log.Note("Searching for ASNs...")

	if len(r.opts.IPs) > 0 {
		ext := extractor.New(r.log, "IP", validator.IpRegex)
		ips := ext.ExtractValues(r.opts.IPs)
		r.rdap.HandleIPs(ips, r.opts.IPRangeFile)
		return
	}
	cmd.Help()
}

func New(opts *options.Options) *Runner {
	logger := logger.New()

	rdap := resolver.NewClient(logger)

	return &Runner{
		opts: opts,
		rdap: rdap,
		log:  logger,
	}
}
