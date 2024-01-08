package singstat

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	return &plugin.Plugin{
		Name: "steampipe-plugin-sggov",
		TableMap: map[string]*plugin.Table{
			"singstat": tableSingStat(ctx),
		},
	}
}
