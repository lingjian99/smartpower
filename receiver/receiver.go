package main

import (
	"flag"
	"fmt"

	"smartpower/receiver/internal/config"
	"smartpower/receiver/internal/kqserver"
	"smartpower/receiver/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/receiver.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	server := kqserver.MustNewKqServer(ctx)  //.MustNewServer(ctx)
	defer server.Stop()
	fmt.Printf("Starting Server...\n")
	server.Start()
}
