package steampipe-plugin-singstat

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/your_github_username/steampipe-plugin-singstat/singstat"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		PluginFunc: singstat.Plugin,
	})
}
