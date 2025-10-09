package Extractor

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/FrakenboK/asnx/internal/logger"
	"github.com/fatih/color"
)

type Extractor struct {
	log            *logger.Logger
	name           string
	regexValidator *regexp.Regexp
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

	// -i 127.0.0.0,8.8.8.8,ips.txt
	// -d example.com,domains.txt
	if strings.Contains(value, ",") {
		return e.ExtractValues(strings.Split(value, ","))
	}

	if e.validateValue(value) {
		return []string{value}
	}

	e.log.Fail(fmt.Sprintf("Invalid %s found %s", e.name, color.RedString(value)))
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

func (e *Extractor) validateValue(ip string) bool {
	if e.regexValidator == nil {
		return true
	}
	return e.regexValidator.MatchString(ip)
}

func New(
	logger *logger.Logger,
	name string,
	regexValidator *regexp.Regexp,
) *Extractor {
	return &Extractor{
		log:            logger,
		name:           name,
		regexValidator: regexValidator,
	}
}
