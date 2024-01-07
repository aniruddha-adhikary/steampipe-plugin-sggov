package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"steampipe-plugin-singstat/singstat"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: singstat.Plugin,
	})
}
