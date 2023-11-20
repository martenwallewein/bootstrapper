module github.com/netsec-ethz/bootstrapper

require (
	github.com/grandcat/zeroconf v1.0.0
	github.com/inconshreveable/log15 v0.0.0-20201112154412-8562bdadbbac
	github.com/mdlayher/ndp v0.10.0
	github.com/miekg/dns v1.1.27
	github.com/pelletier/go-toml v1.8.1-0.20200708110244-34de94e6a887
	golang.org/x/net v0.1.0
	golang.org/x/sys v0.1.0
)

require golang.org/x/sync v0.0.0-20220722155255-886fb9371eb4 // indirect

require (
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/mattn/go-colorable v0.1.8 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	// github.com/u-root/u-root v7.0.0+incompatible // indirect
	gitlab.com/golang-commonmark/puny v0.0.0-20191124015043-9f83538fa04f // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
)

replace github.com/insomniacslk/dhcp => github.com/stapelberg/dhcp v0.0.0-20190429172946-5244c0daddf0

go 1.18
