package main

import (
	"flag"

	"github.com/PoacherBro/kafka-gateway/server/global"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	configFile = flag.String("c", "", "config file path")
)

func assertError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	log.SetLevel(log.InfoLevel)

	flag.Parse()
	if len(*configFile) == 0 {
		panic("config file is empty")
	}

	cfg, err := global.ParseConfig(*configFile)
	assertError(err)

	r := gin.Default()
	r.Run(cfg.HTTPAddr)
}
