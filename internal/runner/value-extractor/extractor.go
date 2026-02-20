package extractor

import (
	"errors"
	"fmt"
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/FrakenboK/asnx/internal/logger"
	"github.com/fatih/color"
)

type Extractor struct {
	log                  *logger.Logger
	regexIPValidator     *regexp.Regexp
	regexDomainValidator *regexp.Regexp
}

func (e *Extractor) ExtractValues(
	values []string,
) []string {
	resultValues := []string{}
	for _, value := range values {
		values := e.ProcessValue(value)
		// logger with name
		resultValues = append(resultValues, values...)
	}
	return resultValues
}

func (e *Extractor) ProcessValue(
	value string,
) []string {
	value = strings.TrimSpace(value)

	if value == "" {
		return []string{}
	}

	_, err := os.Stat(value)
	// extract values from file
	if err == nil {
		return e.extractFileValues(value)
	}

	if !errors.Is(err, os.ErrNotExist) {
		return []string{}
	}

	// EXAMPLE: ya.ru,8.8.8.8,ips.txt
	if strings.Contains(value, ",") {
		return e.ExtractValues(strings.Split(value, ","))
	}

	if e.isIPAddress(value) {
		return []string{value}
	}

	if e.isDomainName(value) {
		resolvedIPs, err := net.LookupIP(value)
		if err == nil {
			var ips []string
			for _, ip := range resolvedIPs {
				ips = append(ips, ip.String())
			}
			return ips
		}
	}

	e.log.Fail(fmt.Sprintf("Faied to get ASN info for host %s", color.RedString(value)))
	return []string{}
}

func (e *Extractor) extractFileValues(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		e.log.Fail(fmt.Sprintf("Failed to open file %s", filename))
		return []string{}
	}

	values := strings.Split(string(content), "\n")
	return e.ExtractValues(values)
}

func (e *Extractor) isIPAddress(value string) bool {
	return e.regexIPValidator.MatchString(value)
}

func (e *Extractor) isDomainName(value string) bool {
	return e.regexDomainValidator.MatchString(value)
}

func New(
	logger *logger.Logger,
) *Extractor {
	return &Extractor{
		log:                  logger,
		regexIPValidator:     ipRegex,
		regexDomainValidator: domainRegex,
	}
}
