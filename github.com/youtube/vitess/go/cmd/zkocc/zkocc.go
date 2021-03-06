// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"
	"os"

	"github.com/youtube/vitess/go/proc"
	"github.com/youtube/vitess/go/vt/servenv"
	"github.com/youtube/vitess/go/vt/topo"
	"github.com/youtube/vitess/go/zk"
	"github.com/youtube/vitess/go/zk/zkocc"
)

var usage = `Cache open zookeeper connections and allow cheap read requests
through a lightweight RPC interface.  The optional parameters are cell
names to try to connect to at startup, versus waiting for the first
request to connect.`

var (
	resolveLocal = flag.Bool("resolve-local", false, "if specified, will try to resolve /zk/local/ paths. If not set, they will fail.")
	port         = flag.Int("port", 14850, "port for the server")
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, usage)
	}
}

// zkocc: a proxy for zk
func main() {
	flag.Parse()
	servenv.Init()
	defer servenv.Close()

	servenv.ServeRPC()
	zkr := zkocc.NewZkReader(*resolveLocal, flag.Args())
	zk.RegisterZkReader(zkr)

	topo.RegisterTopoReader(&TopoReader{zkr: zkr})

	proc.ListenAndServe(fmt.Sprintf("%v", *port))
}
