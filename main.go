package main

import (
	"github.com/algorand/go-algorand/logging"
	"github.com/algorand/go-algorand/tools/network"
	"strings"
)

var log = logging.Base()

func printAddress(addrs []string) {
	log.Infof("Found %d SRV records: %v", len(addrs), strings.Join(addrs, ","))
}

func testBootstrap(fallbackDNSResolverAddress string, secure bool) {
	log.Infof("Trying: fallback=\"%v\", secure=%v", fallbackDNSResolverAddress, secure)

	addrs, err := network.ReadFromSRV(
		"algobootstrap",
		"tcp",
		"mainnet.algorand.network",
		"",
		true,
	)
	if err != nil {
		log.Error(err)
	}
	printAddress(addrs)
}

func main() {
	log.SetLevel(logging.Debug)

	testBootstrap("", true)
	testBootstrap("8.8.8.8", true)
	testBootstrap("", false)
	testBootstrap("8.8.8.8", false)
}
