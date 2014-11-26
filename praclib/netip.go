package goprac

import (
	"log"
	"net"
)

func SubnetPool(networkString string) *net.IPNet {

	_, subnetPool := NetworkParms(networkString)
	return subnetPool
}

func NetworkParms(networkString string) (net.IP, *net.IPNet) {

	containerIP, subnet, err := net.ParseCIDR(networkString)
	if err != nil {

		log.Fatal("parse ip error", err)
	}

	return containerIP, subnet
}
