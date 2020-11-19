package hinting

import (
	"context"
	"github.com/grandcat/zeroconf"
	"github.com/scionproto/scion/go/lib/log"
	"net"
	"time"
)

const (
	resolverTimeout = 2 * time.Second
)

type MDNSHintGeneratorConf struct {
	Enable bool `toml:"enable"`
}

var _ HintGenerator = (*MDNSSDHintGenerator)(nil)

// Multicast Domain Name System Service Discovery
type MDNSSDHintGenerator struct {
	cfg   *MDNSHintGeneratorConf
	iface *net.Interface
}

func NewMDNSHintGenerator(cfg *MDNSHintGeneratorConf, iface *net.Interface) *MDNSSDHintGenerator {
	return &MDNSSDHintGenerator{cfg, iface}
}

func (g *MDNSSDHintGenerator) Generate(ipHintsChan chan<- net.IP) {
	if !g.cfg.Enable {
		return
	}
	resolver, err := zeroconf.NewResolver(zeroconf.SelectIfaces([]net.Interface{*g.iface}))
	if err != nil {
		log.Error("mDNS could not construct dns resolver", "err", err)
		return
	}
	entriesChan := make(chan *zeroconf.ServiceEntry)
	go func() {
		defer log.HandlePanic()
		handleEntries(entriesChan, ipHintsChan)
	}()
	discoverEntries(resolver, entriesChan)
}

func handleEntries(entriesChan <-chan *zeroconf.ServiceEntry, ipHintsChan chan<- net.IP) {
	for entry := range entriesChan {
		for _, address := range entry.AddrIPv4 {
			log.Info("mDNS hint", "IP", address.String())
			ipHintsChan <- address
		}
		for _, address := range entry.AddrIPv6 {
			log.Info("mDNS hint", "IP", address.String())
			ipHintsChan <- address
		}
	}
	log.Info("mDNS hinting done")
}

func discoverEntries(resolver *zeroconf.Resolver, entriesChan chan *zeroconf.ServiceEntry) {
	ctx, cancel := context.WithTimeout(context.Background(), resolverTimeout)
	defer cancel()
	err := resolver.Browse(ctx, "_sciondiscovery._tcp", "local.", entriesChan)
	if err != nil {
		log.Error("mDNS could not lookup", "err", err)
		return
	}
	<-ctx.Done()
}