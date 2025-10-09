package ui

import (
	"fmt"

	"github.com/FrakenboK/asnx/internal/version"
	"github.com/fatih/color"
)

var (
	// r = rand.New(rand.NewSource(time.Now().UnixNano()))
	// randomColor = color.RGB(r.Intn(255), r.Intn(255), r.Intn(255))

	// 209, 59, 107
	randomColor = color.RGB(177, 70, 102)

	FailPrefix  = color.RedString("[-]")
	InfoPrefix  = color.GreenString("[+]")
	VersionDesc = fmt.Sprintf("\t\t  You are using asnx version %s!\n", randomColor.Sprintf(version.Version))

	desc = "\tA tool for obtaining information about ASN hosts using RDAP"

	Banner = `
		▄█████▄  ▄▄█████▄  ██▄████▄  ▀██  ██▀ 
		▀ ▄▄▄██  ██▄▄▄▄ ▀  ██▀   ██    ████   
		▄██▀▀▀██   ▀▀▀▀██▄  ██    ██    ▄██▄   
		██▄▄▄███  █▄▄▄▄▄██  ██    ██   ▄█▀▀█▄  
		▀▀▀▀ ▀▀   ▀▀▀▀▀▀   ▀▀    ▀▀  ▀▀▀  ▀▀▀ 
`
)

func GetBanner() string {
	return fmt.Sprintf("%s\n%s", randomColor.Sprint(Banner), VersionDesc)
}

func GetDesc() string {
	return fmt.Sprintf("%s\n%s", GetBanner(), desc)
}
