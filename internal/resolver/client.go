package resolver

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/FrakenboK/asnx/internal/logger"
	"github.com/fatih/color"
	"github.com/openrdap/rdap"
)

type Resolver struct {
	client *rdap.Client
	log    *logger.Logger
}

type Response struct {
	IP     string
	Domain string

	StartAddress string
	EndAddress   string

	Info SpecialInfo
}

type SpecialInfo struct {
	Name      string
	Author    string
	Telephone string
	Country   string
	Email     string
	Address   string
}

func (r *Resolver) HandleIPs(ips []string, ipRangeFilename string) {
	asnIPs := []string{}

	for _, ip := range ips {
		info, err := r.client.QueryIP(ip)
		if err != nil {
			r.log.Fail(
				fmt.Sprintf("Failed to handle IP info: %s, [%s]", ip, err.Error()),
			)
			continue
		}
		response := r.processIPResponse(ip, info)
		r.log.Info(r.fmtResponse(response))

		responseAsnAddrs, err := getIPRange(response.StartAddress, response.EndAddress)
		if err != nil {
			r.log.Fail(
				fmt.Sprintf("Failed to extract IP-range: %s - %s; %s",
					response.StartAddress,
					response.EndAddress,
					err.Error(),
				),
			)
		}

		asnIPs = append(asnIPs, responseAsnAddrs...)
	}

	if ipRangeFilename == "" {
		r.log.Note("Complete!")
		return
	}
	r.log.Note(fmt.Sprintf("Saving IP ranges to file %s", ipRangeFilename))
	r.saveIPRange(asnIPs, ipRangeFilename)
}

func (r *Resolver) saveIPRange(ipRange []string, filename string) error {
	text := strings.Join(ipRange, "\n")

	return os.WriteFile(
		filename,
		[]byte(text),
		0644,
	)
}

func (r *Resolver) processIPResponse(ip string, info *rdap.IPNetwork) *Response {
	parsedInfo := r.processRDAPEntities(info.Entities)

	response := &Response{
		Info: parsedInfo,
	}

	response.Info.Name = info.Name
	response.StartAddress = info.StartAddress
	response.EndAddress = info.EndAddress
	response.IP = ip

	return response
}

func (r *Resolver) processRDAPEntities(entities []rdap.Entity) SpecialInfo {

	for _, entity := range entities {
		if entity.VCard.Tel() == "" && entity.VCard.Email() == "" {
			continue
		}
		return SpecialInfo{
			Author:    entity.VCard.Name(),
			Email:     entity.VCard.Email(),
			Telephone: entity.VCard.Tel(),
			Country:   entity.VCard.Country(),
			Address:   entity.VCard.ExtendedAddress(),
		}
	}
	return SpecialInfo{}
}

func (r *Resolver) fmtResponse(resp *Response) string {
	var printable string
	if resp.IP != "" {
		printable = fmt.Sprintf("IP: %24s => ", color.YellowString(resp.IP))
	} else {
		printable = fmt.Sprintf("Domain: %24s => ", color.YellowString(resp.Domain))
	}

	printable = fmt.Sprintf("%s %s", printable, fmtSpecialInfo(resp.Info))
	printable = fmt.Sprintf("%s IP_range=\"%s - %s\"", printable, resp.StartAddress, resp.EndAddress)

	return printable
}

func fmtSpecialInfo(info SpecialInfo) string {
	printable := ""
	reflectionValue := reflect.ValueOf(info)

	for i := 0; i < reflectionValue.NumField(); i++ {
		field := reflectionValue.Field(i).Interface()
		if field.(string) == "" {
			continue
		}
		fieldName := reflectionValue.Type().Field(i).Name
		printable = fmt.Sprintf("%s %s=\"%s\"", printable, fieldName, field)
	}

	return printable
}

func NewClient(
	log *logger.Logger,
) *Resolver {
	client := &rdap.Client{}

	return &Resolver{
		client: client,
		log:    log,
	}
}
