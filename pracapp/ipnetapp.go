package main

import (
	"fmt"
	"github.com/lxyxd2003m/goprac/praclib"
	"net"
)

func main() {
	var defaultSubnetPool *net.IPNet

	defaultSubnetPool = goprac.SubnetPool("10.2.3.0/32")
	fmt.Printf("ip----->%s      mask----->%s\n", defaultSubnetPool.IP, defaultSubnetPool.Mask)

	ip, _ := goprac.NetworkParms("192.168.1.1/24")

	fmt.Printf("ip----->%s      mask----->%s", ip, ip.DefaultMask())

}
