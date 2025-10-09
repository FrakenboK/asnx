package ui

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/FrakenboK/asnx/internal/version"
	"github.com/fatih/color"
)

var (
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	// 57 82 216
	RandomColor = color.RGB(r.Intn(255), r.Intn(255), r.Intn(255))

	FailPrefix  = color.RedString("[-]")
	InfoPrefix  = color.GreenString("[+]")
	VersionDesc = fmt.Sprintf("\t\t  You are using asnx version %s!\n", RandomColor.Sprintf(version.Version))

	Banner = `
		▄█████▄  ▄▄█████▄  ██▄████▄  ▀██  ██▀ 
		▀ ▄▄▄██  ██▄▄▄▄ ▀  ██▀   ██    ████   
		▄██▀▀▀██   ▀▀▀▀██▄  ██    ██    ▄██▄   
		██▄▄▄███  █▄▄▄▄▄██  ██    ██   ▄█▀▀█▄  
		▀▀▀▀ ▀▀   ▀▀▀▀▀▀   ▀▀    ▀▀  ▀▀▀  ▀▀▀ 
`
)
