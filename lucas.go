package main

/**
	lucas is a etcd v3 key/value browser implemented in go
	etcd-browser works very well in etcd v2 api
	but there isn't a  mature solution to etcd v3 api

	Then lucas is born for this.
	lucas is from the movie <Lucas the Spider>
	Hope you love it.

	2017/12/15
	21:43:36
	zhongwei.lzw@alibaba-inc.com
 */

import (
	"github.com/ringtail/lucas/cmd"
	"github.com/ringtail/lucas/backend/types"
	"flag"
	"fmt"
	"os"
)

const COMMAND_DESC = `
NAME:
   lucas - A etcd v3 key/value browser implemented in go

USAGE:
   lucas [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   run  run lucas web server

GLOBAL OPTIONS:
   --debug                          set debug mode to lucas
   --endpoints value                machine addresses in the cluster (default: "http://127.0.0.1:2379,http://127.0.0.1:4001")
   --cert-file value                identify HTTPS client using this SSL certificate file
   --key-file value                 identify HTTPS client using this SSL key file
   --ca-file value                  verify certificates of HTTPS-enabled servers using this CA bundle
   --help                           show help
   --version                        print the version
`

const VERSION = "0.0.1"

var (
	debug_mode bool
	endpoints  string
	cert_file  string
	key_file   string
	ca_file    string
	help       bool
	version    bool
)

func init() {
	flag.BoolVar(&debug_mode, "debug", false, "set debug mode to lucas")
	flag.StringVar(&endpoints, "endpoints", "http://127.0.0.1:2379", "machine addresses in the cluster (default: \"http://127.0.0.1:2379,http://127.0.0.1:4001\")", )
	flag.StringVar(&ca_file, "ca-file", "", "verify certificates of HTTPS-enabled servers using this CA bundle")
	flag.StringVar(&key_file, "key-file", "", "identify HTTPS client using this SSL key file")
	flag.StringVar(&cert_file, "cert-file", "", "identify HTTPS client using this SSL certificate file")
	flag.BoolVar(&help, "help", false, "show help")
	flag.BoolVar(&version, "version", false, "print the version")
}

func main() {
	flag.Parse()
	if help == true {
		fmt.Println(COMMAND_DESC)
		return
	} else if version == true {
		fmt.Println(VERSION)
		return
	}
	if os.Getenv("CA_FILE") != "" || os.Getenv("KEY_FILE") != "" || os.Getenv("CERT_FILE") != "" {
		ca_file = os.Getenv("CA_FILE")
		key_file = os.Getenv("KEY_FILE")
		cert_file = os.Getenv("CERT_FILE")
	}

	if os.Getenv("ENDPOINTS") != "" {
		endpoints = os.Getenv("ENDPOINTS")
	}

	args := flag.Args()
	if args != nil && len(args) != 0 && args[0] == "run" {
		cl := &cmd.CommandLine{
			Opts: &types.Opts{
				DebugMode: debug_mode,
				Cert:      cert_file,
				Key:       key_file,
				Ca:        ca_file,
				Endpoints: endpoints,
			},
		}
		cl.Run()
	} else {
		fmt.Printf("unknown command: %v\n", args)
	}

}
