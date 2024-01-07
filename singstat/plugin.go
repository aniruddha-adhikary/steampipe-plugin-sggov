package singstat

import (
	"context"
	"encoding/json"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"net/http"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	return &plugin.Plugin{
		Name: "steampipe-plugin-singstat",
		TableMap: map[string]*plugin.Table{
			"singstat": tableSingStat(ctx),
		},
	}
}

func tableSingStat(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "singstat",
		Description: "Search for statistical tables in the SingStat Table Builder.",
		List: &plugin.ListConfig{
			Hydrate: listSingStatTables,
		},
		Columns: []*plugin.Column{
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The ID of the statistical table."},
			{Name: "table_type", Type: proto.ColumnType_STRING, Description: "The type of the statistical table."},
			{Name: "title", Type: proto.ColumnType_STRING, Description: "The title of the statistical table."},
		},
	}
}

func listSingStatTables(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	// Retrieve the keyword and searchOption from the query qualifiers
	quals := d.EqualsQuals
	keyword := quals["keyword"].GetStringValue()
	searchOption := quals["searchOption"].GetStringValue()

	// Build the URL for the API request
	url := "https://tablebuilder.singstat.gov.sg/api/table/resourceid?keyword=" + keyword + "&searchOption=" + searchOption

	// Create a new HTTP client and execute the request
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		plugin.Logger(ctx).Error("listSingStatTables", "request_creation_error", err)
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		plugin.Logger(ctx).Error("listSingStatTables", "request_execution_error", err)
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the response body into the appropriate structure
	var responseData struct {
		Data struct {
			GeneratedBy   string `json:"generatedBy"`
			DateGenerated string `json:"dateGenerated"`
			Total         int    `json:"total"`
			Records       []struct {
				ID        string `json:"id"`
				TableType string `json:"tableType"`
				Title     string `json:"title"`
			} `json:"records"`
		} `json:"Data"`
		DataCount  int    `json:"DataCount"`
		StatusCode int    `json:"StatusCode"`
		Message    string `json:"Message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		plugin.Logger(ctx).Error("listSingStatTables", "response_decode_error", err)
		return nil, err
	}

	// Stream the data from the response
	for _, record := range responseData.Data.Records {
		d.StreamListItem(ctx, record)
	}

	return nil, nil
}

