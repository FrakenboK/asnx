/*
Copyright © 2025 Kobyak Mikhail FrakenboK@cR4.sh
*/
package cmd

import (
	"github.com/FrakenboK/asnx/internal/options"
	"github.com/FrakenboK/asnx/internal/runner"
	"github.com/FrakenboK/asnx/internal/ui"

	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	opts := &options.Options{}
	runner := runner.New(opts)

	var cmd = &cobra.Command{
		Use:   "asnx",
		Short: ui.GetDesc(),
		Run:   runner.Start,
	}

	// Search by IP
	cmd.Flags().StringArrayVarP(&opts.Hosts, "hosts", "H", []string{}, "Hosts (usage: -h ya.ru,8.8.8.8 or --hosts hosts.txt)")

	cmd.Flags().StringVarP(&opts.IPRangeFile, "ip-range-file", "f", "", "File to save all ip ranges handeled from ASN")
	cmd.Flags().StringVarP(&opts.OutputFile, "output", "o", "", "[IN DEVELOPMENT] File to save output")

	// Search by Domain
	cmd.Flags().BoolVarP(&opts.Version, "version", "v", false, "Shows asnx version")
	cmd.Flags().BoolVarP(&opts.BannerDisabled, "no-banner", "", false, "Disables banner")

	return cmd
}
