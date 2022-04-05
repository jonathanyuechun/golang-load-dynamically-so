package main

import (
	"fmt"
	"os"
	"plugin"
)

type Sourcer interface {
	Open()
	Start()
}

const (
	customSource string = "./mysource.so"
)

func main() {
	var pluginSource *plugin.Plugin
	var pluginSymbol plugin.Symbol
	var err error
	if pluginSource, err = plugin.Open(customSource); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if pluginSymbol, err = pluginSource.Lookup("ElasticSource"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var elasticSource Sourcer
	var ok bool
	if elasticSource, ok = pluginSymbol.(Sourcer); !ok {
		fmt.Println("missmatch type")
		os.Exit(1)
	}
	elasticSource.Open()
	elasticSource.Start()
}
