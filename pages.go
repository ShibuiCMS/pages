package main

import (
	"log"

	pages "github.com/ShibuiCMS/pages/lib"
	"github.com/vroomy/plugins"
)

var p *pages.Pages

func init() {
	var err error
	if p, err = pages.New("./data"); err != nil {
		log.Fatal(err)
	}
}

// Backend returns the underlying backend to the plugin
func Backend() interface{} {
	return p
}

// OnInit will be called by Vroomie on initialization
func OnInit(p *plugins.Plugins, flags, env map[string]string) (err error) {
	return
}
