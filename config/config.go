// Package config contains the configuration of bootstrapper.
package config

import (
	"github.com/scionproto/scion/go/lib/common"
	"github.com/scionproto/scion/go/lib/config"
	"github.com/scionproto/scion/go/lib/env"
	"io"
)

var _ config.Config = (*Config)(nil)

type Config struct {
	Interface string

	SciondDirectory string

	Mechanisms Mechanisms

	Logging env.Logging
}

func (cfg *Config) InitDefaults() {
	config.InitAll(
		&cfg.Mechanisms,
		&cfg.Logging,
	)

	if cfg.SciondDirectory == "" {
		cfg.SciondDirectory = "."
	}
}

func (cfg *Config) Validate() error {
	if cfg.Interface == "" {
		return common.NewBasicError("Interface must be set", nil)
	}

	return config.ValidateAll(
		&cfg.Logging,
	)
}

func (cfg *Config) Sample(dst io.Writer, path config.Path, _ config.CtxMap) {
	config.WriteString(dst, bootstrapperSample)
	config.WriteSample(dst, path, config.CtxMap{config.ID: idSample},
		&cfg.Mechanisms,
		&cfg.Logging,
	)
}

func (cfg *Config) ConfigName() string {
	return "bootstrapper_config"
}

var _ config.Config = (*Mechanisms)(nil)

type Mechanisms struct {
	DHCP bool

	MDNS bool

	DNSSD bool

	DNSNAPTR bool
}

func (cfg *Mechanisms) InitDefaults() {
	cfg.DHCP = true
	cfg.MDNS = true
	cfg.DNSSD = true
	cfg.DNSNAPTR = true
}

func (cfg *Mechanisms) Validate() error {
	return nil
}

func (cfg *Mechanisms) Sample(dst io.Writer, path config.Path, ctx config.CtxMap) {
	config.WriteString(dst, mechanismsSample)
}

func (cfg *Mechanisms) ConfigName() string {
	return "mechanisms"
}
