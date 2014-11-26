package main

import (
	"fmt"
	"github.com/lxyxd2003m/goprac/praclib"
	"net"
)

func main() {
	var defaultSubnetPool *net.IPNet

	defaultSubnetPool = goprac.SubnetPool("10.2.3.0/32")
	fmt.Printf("net----->%s", defaultSubnetPool.Network())

}
