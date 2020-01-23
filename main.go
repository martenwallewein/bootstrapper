package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/scionproto/scion/go/bootstrapper/config"
	"github.com/scionproto/scion/go/lib/env"
	"github.com/scionproto/scion/go/lib/fatal"
	"github.com/scionproto/scion/go/lib/infra/modules/itopo"
	"github.com/scionproto/scion/go/lib/log"
	"github.com/scionproto/scion/go/proto"
	_ "net/http/pprof"
	"os"
)

var (
	cfg config.Config
)

func init() {
	flag.Usage = env.Usage
}

func main() {
	os.Exit(realMain())
}

func realMain() int {
	fatal.Init()
	env.AddFlags()
	flag.Parse()
	if v, ok := env.CheckFlags(&cfg); !ok {
		return v
	}
	if _, err := toml.DecodeFile(env.ConfigFile(), &cfg); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	cfg.InitDefaults()
	if err := env.InitLogging(&cfg.Logging); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	defer log.Flush()
	defer env.LogAppStopped("bootstrapper", "")
	defer log.LogPanicAndExit()

	if err := cfg.Validate(); err != nil {
		log.Error("Unable to validate config", "err", err)
		return 1
	}
	itopo.Init("", proto.ServiceType_unset, itopo.Callbacks{})
	_, err := tryBootstrapping()
	if err != nil {
		log.Error("Unable to load topology", "err", err)
	}

	return 0
}
