package resolver

import (
	"fmt"
	"os"
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
	Name   string

	StartAddress string
	EndAddress   string

	Author  string
	Tel     string
	Country string
	Email   string
	Address string
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
	response := r.processRDAPEntities(info.Entities)
	response.Name = info.Name
	response.StartAddress = info.StartAddress
	response.EndAddress = info.EndAddress
	response.IP = ip
	return response
}

func (r *Resolver) processRDAPEntities(entities []rdap.Entity) *Response {

	for _, entity := range entities {
		if entity.VCard.Tel() == "" && entity.VCard.Email() == "" {
			continue
		}
		return &Response{
			Author:  entity.VCard.Name(),
			Email:   entity.VCard.Email(),
			Tel:     entity.VCard.Tel(),
			Country: entity.VCard.Country(),
			Address: entity.VCard.ExtendedAddress(),
		}
	}
	return &Response{}
}

func (r *Resolver) fmtResponse(resp *Response) string {
	var printable string
	if resp.IP != "" {
		printable = fmt.Sprintf("IP: %24s =>", color.YellowString(resp.IP))
	} else {
		printable = fmt.Sprintf("ASN info for Domain name %24s =>", color.YellowString(resp.Domain))
	}

	if resp.Name != "" {
		printable = fmt.Sprintf("%s Name=\"%s\"", printable, resp.Name)
	}
	if resp.Author != "" {
		printable = fmt.Sprintf("%s Person=\"%s\"", printable, resp.Author)
	}
	if resp.Tel != "" {
		printable = fmt.Sprintf("%s Telephone=\"%s\"", printable, resp.Tel)
	}
	if resp.Country != "" {
		printable = fmt.Sprintf("%s Country=\"%s\"", printable, resp.Country)
	}
	if resp.Email != "" {
		printable = fmt.Sprintf("%s Email=\"%s\"", printable, resp.Email)
	}
	if resp.Address != "" {
		printable = fmt.Sprintf("%s Address=\"%s\"", printable, resp.Address)
	}

	printable = fmt.Sprintf("%s IP_range=\"%s - %s\"", printable, resp.StartAddress, resp.EndAddress)

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
