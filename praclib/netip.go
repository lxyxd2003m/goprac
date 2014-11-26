package goprac

import (
	"net"
)

func SubnetPool(networkString string) *net.IPNet {

	_, subnetPool := NetworkParms(networkString)
}

func NetworkParms(networkString string) (net.IP, *net.IPNet) {

	containerIP, subnet, err := net.ParseCIDR(networkString)
	return containerIP, subnet
}
