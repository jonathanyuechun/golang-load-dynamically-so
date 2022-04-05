package main

import (
	"fmt"
)

type ElasticSourceSink struct {
}

func (es *ElasticSourceSink) Open() {
	fmt.Println("Opening")
}

func (es *ElasticSourceSink) Start() {
	fmt.Println("Starting")
}

func main() {
	// for testing purpose
	mySource := ElasticSourceSink{}
	mySource.Open()
	mySource.Start()
}

// exported
var ElasticSource ElasticSourceSink
