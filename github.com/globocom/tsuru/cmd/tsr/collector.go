// Copyright 2013 tsuru authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/globocom/tsuru/cmd"
	"github.com/globocom/tsuru/collector"
	"launchpad.net/gnuflag"
)

type collectorCmd struct {
	fs     *gnuflag.FlagSet
	config string
	dry    bool
}

func (c *collectorCmd) Run(context *cmd.Context, client *cmd.Client) error {
	flags := map[string]interface{}{}
	flags["dry"] = c.dry
	flags["config"] = c.config
	collector.Run(flags)
	return nil
}

func (collectorCmd) Info() *cmd.Info {
	return &cmd.Info{
		Name:    "collector",
		Usage:   "collector",
		Desc:    "Starts the tsuru collector.",
		MinArgs: 0,
	}
}

func (c *collectorCmd) Flags() *gnuflag.FlagSet {
	if c.fs == nil {
		c.fs = gnuflag.NewFlagSet("api", gnuflag.ExitOnError)
		c.fs.BoolVar(&c.dry, "dry", false, "dry-run: does not run the collector (for testing purpose)")
		c.fs.BoolVar(&c.dry, "d", false, "dry-run: does not run the collector (for testing purpose)")
		c.fs.StringVar(&c.config, "config", "/etc/tsuru/tsuru.conf", "tsr collector config file.")
		c.fs.StringVar(&c.config, "c", "/etc/tsuru/tsuru.conf", "tsr collector config file.")
	}
	return c.fs
}
