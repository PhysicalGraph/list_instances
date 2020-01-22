package main

import (
        //"fmt"
	"github.com/aws/aws-sdk-go/aws"
        "strings"
	"net"
	"os"
	"sort"

	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
)

type instance struct {
	*ec2.Instance
	name      string
	privateIP net.IP
        ipv6_addresses string
}

func newInstance(i *ec2.Instance) (ret instance) {
	ret.Instance = i
	ret.privateIP = net.ParseIP(*i.PrivateIpAddress)
        ipv6_addresses := []string{}
        for _, network := range i.NetworkInterfaces {
           for _, ipv6 := range network.Ipv6Addresses {
               //fmt.Println(*i.InstanceId, aws.StringValue(ipv6.Ipv6Address))
               ipv6_addresses = append(ipv6_addresses, aws.StringValue(ipv6.Ipv6Address))
           }
        }
	for _, t := range i.Tags {
		if *t.Key == "Name" {
			ret.name = *t.Value
		}
	}
        ret.ipv6_addresses = strings.Join(ipv6_addresses,", ")
        //fmt.Println(ret.ipv6_addresses)
	return ret
}

func (i *instance) toRow() []string {
	return []string{
		i.name,
		*i.InstanceId,
		stringify(i.PublicIpAddress),
		*i.PrivateIpAddress,
                i.ipv6_addresses,
		stringify(i.KeyName),
	}
}

func stringify(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

type instances []*instance

func (s instances) printTable() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Id", "PublicIP", "PrivateIP", "Ipv6", "Key"})
	for _, i := range s {
		table.Append(i.toRow())
	}
	table.Render()
}

func (s instances) sort() {
	sort.Sort(s)
}

// implement sort.Interface
func (s instances) Len() int      { return len(s) }
func (s instances) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

// Less sorts instances by name and then by private IP address
func (s instances) Less(i, j int) bool {
	if s[i].name < s[j].name {
		return true
	}
	if s[i].name > s[j].name {
		return false
	}
	for n, v := range s[i].privateIP {
		if v < s[j].privateIP[n] {
			return true
		}
		if v > s[j].privateIP[n] {
			return false
		}
	}
	return false
}
