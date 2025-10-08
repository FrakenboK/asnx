package logger

import (
	"fmt"

	"github.com/FrakenboK/asnx/internal/version"
	"github.com/fatih/color"
)

var (
	failString    = color.RedString("[-]")
	infoString    = color.GreenString("[+]")
	versionString = fmt.Sprintf("\n\t\tYou are using asnx version %s!\n", color.MagentaString(version.Version))
)
