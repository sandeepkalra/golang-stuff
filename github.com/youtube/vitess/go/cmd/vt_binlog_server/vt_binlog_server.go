// Copyright 2012, Google Inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// vt binlog server: Serves binlog for out of band replication.
package main

import (
	"flag"
	"fmt"
	_ "net/http/pprof"

	log "github.com/golang/glog"
	"github.com/youtube/vitess/go/proc"
	"github.com/youtube/vitess/go/rpcwrap"
	"github.com/youtube/vitess/go/vt/mysqlctl"
	"github.com/youtube/vitess/go/vt/mysqlctl/proto"
	"github.com/youtube/vitess/go/vt/servenv"
)

var (
	port      = flag.Int("port", 6614, "port for the server")
	dbname    = flag.String("dbname", "", "database name")
	mycnfFile = flag.String("mycnf-file", "", "path of mycnf file")
)

func main() {
	flag.Parse()
	servenv.Init()
	defer servenv.Close()

	if *mycnfFile == "" {
		log.Fatalf("Please specify the path for mycnf file.")
	}
	mycnf, err := mysqlctl.ReadMycnf(*mycnfFile)
	if err != nil {
		log.Fatalf("Error reading mycnf file %v", *mycnfFile)
	}

	binlogServer := mysqlctl.NewBinlogServer(mycnf)
	mysqlctl.EnableBinlogServerService(binlogServer, *dbname)

	proto.RegisterBinlogServer(binlogServer)
	rpcwrap.RegisterAuthenticated(binlogServer)

	servenv.ServeRPC()

	proc.ListenAndServe(fmt.Sprintf("%v", *port))
	mysqlctl.DisableBinlogServerService(binlogServer)
}
