module github.com/netsec-ethz/bootstrapper

require (
	github.com/grandcat/zeroconf v1.0.0
	github.com/insomniacslk/dhcp v0.0.0-20200922210017-67c425063dca
	github.com/miekg/dns v1.1.27
	github.com/scionproto/scion v0.6.0
	github.com/stretchr/testify v1.6.1 // indirect
	github.com/u-root/u-root v7.0.0+incompatible // indirect
	golang.org/x/net v0.0.0-20200927032502-5d4f70055728
	golang.org/x/sys v0.0.0-20200916030750-2334cc1a136f
)

replace github.com/insomniacslk/dhcp => github.com/stapelberg/dhcp v0.0.0-20190429172946-5244c0daddf0

replace github.com/scionproto/scion => github.com/scionproto/scion v0.0.0-20201109090843-8b06407464f6

go 1.14