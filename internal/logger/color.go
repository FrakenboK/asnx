package logger

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
	randomColor = color.RGB(r.Intn(255), r.Intn(255), r.Intn(255))

	failString    = color.RedString("[-]")
	infoString    = color.GreenString("[+]")
	versionString = fmt.Sprintf("\t\t  You are using asnx version %s!\n", randomColor.Sprintf(version.Version))

	banner = `								
												
		▄█████▄  ▄▄█████▄  ██▄████▄  ▀██  ██▀ 
		▀ ▄▄▄██  ██▄▄▄▄ ▀  ██▀   ██    ████   
		▄██▀▀▀██   ▀▀▀▀██▄  ██    ██    ▄██▄   
		██▄▄▄███  █▄▄▄▄▄██  ██    ██   ▄█▀▀█▄  
		▀▀▀▀ ▀▀   ▀▀▀▀▀▀   ▀▀    ▀▀  ▀▀▀  ▀▀▀ 
											
		`
)
