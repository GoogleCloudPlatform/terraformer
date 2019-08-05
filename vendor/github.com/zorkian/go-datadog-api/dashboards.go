/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

import (
	"encoding/json"
	"fmt"
)

// GraphDefinitionRequestStyle represents the graph style attributes
type GraphDefinitionRequestStyle struct {
	Palette *string `json:"palette,omitempty"`
	Width   *string `json:"width,omitempty"`
	Type    *string `json:"type,omitempty"`
}

// GraphDefinitionRequest represents the requests passed into each graph.
type GraphDefinitionRequest struct {
	Stacked            *bool                        `json:"stacked,omitempty"`
	Aggregator         *string                      `json:"aggregator,omitempty"`
	ConditionalFormats []DashboardConditionalFormat `json:"conditional_formats,omitempty"`
	Type               *string                      `json:"type,omitempty"`
	Style              *GraphDefinitionRequestStyle `json:"style,omitempty"`

	// For change type graphs
	ChangeType     *string                            `json:"change_type,omitempty"`
	OrderDirection *string                            `json:"order_dir,omitempty"`
	CompareTo      *string                            `json:"compare_to,omitempty"`
	IncreaseGood   *bool                              `json:"increase_good,omitempty"`
	OrderBy        *string                            `json:"order_by,omitempty"`
	ExtraCol       *string                            `json:"extra_col,omitempty"`
	Metadata       map[string]GraphDefinitionMetadata `json:"metadata,omitempty"`

	// A Graph can only have one of these types of query.
	Query        *string             `json:"q,omitempty"`
	LogQuery     *GraphApmOrLogQuery `json:"log_query,omitempty"`
	ApmQuery     *GraphApmOrLogQuery `json:"apm_query,omitempty"`
	ProcessQuery *GraphProcessQuery  `json:"process_query,omitempty"`
}

// GraphApmOrLogQuery represents an APM or a Log query
type GraphApmOrLogQuery struct {
	Index   *string                     `json:"index"`
	Compute *GraphApmOrLogQueryCompute  `json:"compute"`
	Search  *GraphApmOrLogQuerySearch   `json:"search,omitempty"`
	GroupBy []GraphApmOrLogQueryGroupBy `json:"groupBy,omitempty"`
}

type GraphApmOrLogQueryCompute struct {
	Aggregation *string `json:"aggregation"`
	Facet       *string `json:"facet,omitempty"`
	Interval    *int    `json:"interval,omitempty"`
}

type GraphApmOrLogQuerySearch struct {
	Query *string `json:"query"`
}

type GraphApmOrLogQueryGroupBy struct {
	Facet *string                        `json:"facet"`
	Limit *int                           `json:"limit,omitempty"`
	Sort  *GraphApmOrLogQueryGroupBySort `json:"sort,omitempty"`
}

type GraphApmOrLogQueryGroupBySort struct {
	Aggregation *string `json:"aggregation"`
	Order       *string `json:"order"`
	Facet       *string `json:"facet,omitempty"`
}

type GraphProcessQuery struct {
	Metric   *string  `json:"metric"`
	SearchBy *string  `json:"search_by,omitempty"`
	FilterBy []string `json:"filter_by,omitempty"`
	Limit    *int     `json:"limit,omitempty"`
}

type GraphDefinitionMetadata TileDefMetadata

type GraphDefinitionMarker struct {
	Type  *string      `json:"type,omitempty"`
	Value *string      `json:"value,omitempty"`
	Label *string      `json:"label,omitempty"`
	Val   *json.Number `json:"val,omitempty"`
	Min   *json.Number `json:"min,omitempty"`
	Max   *json.Number `json:"max,omitempty"`
}

type GraphEvent struct {
	Query *string `json:"q,omitempty"`
}

type Yaxis struct {
	Min          *float64 `json:"min,omitempty"`
	AutoMin      bool     `json:"-"`
	Max          *float64 `json:"max,omitempty"`
	AutoMax      bool     `json:"-"`
	Scale        *string  `json:"scale,omitempty"`
	IncludeZero  *bool    `json:"includeZero,omitempty"`
	IncludeUnits *bool    `json:"units,omitempty"`
}

// UnmarshalJSON is a Custom Unmarshal for Yaxis.Min/Yaxis.Max. If the datadog API
// returns "auto" for min or max, then we should set Yaxis.min or Yaxis.max to nil,
// respectively.
func (y *Yaxis) UnmarshalJSON(data []byte) error {
	type Alias Yaxis
	wrapper := &struct {
		Min *json.Number `json:"min,omitempty"`
		Max *json.Number `json:"max,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(y),
	}

	if err := json.Unmarshal(data, &wrapper); err != nil {
		return err
	}

	if wrapper.Min != nil {
		if *wrapper.Min == "auto" {
			y.AutoMin = true
			y.Min = nil
		} else {
			f, err := wrapper.Min.Float64()
			if err != nil {
				return err
			}
			y.Min = &f
		}
	}

	if wrapper.Max != nil {
		if *wrapper.Max == "auto" {
			y.AutoMax = true
			y.Max = nil
		} else {
			f, err := wrapper.Max.Float64()
			if err != nil {
				return err
			}
			y.Max = &f
		}
	}
	return nil
}

type Style struct {
	Palette     *string      `json:"palette,omitempty"`
	PaletteFlip *bool        `json:"paletteFlip,omitempty"`
	FillMin     *json.Number `json:"fillMin,omitempty"`
	FillMax     *json.Number `json:"fillMax,omitempty"`
}

type GraphDefinition struct {
	Viz      *string                  `json:"viz,omitempty"`
	Requests []GraphDefinitionRequest `json:"requests,omitempty"`
	Events   []GraphEvent             `json:"events,omitempty"`
	Markers  []GraphDefinitionMarker  `json:"markers,omitempty"`

	// For timeseries type graphs
	Yaxis Yaxis `json:"yaxis,omitempty"`

	// For query value type graphs
	Autoscale  *bool       `json:"autoscale,omitempty"`
	TextAlign  *string     `json:"text_align,omitempty"`
	Precision  *PrecisionT `json:"precision,omitempty"`
	CustomUnit *string     `json:"custom_unit,omitempty"`

	// For hostmaps
	Style                 *Style   `json:"style,omitempty"`
	Groups                []string `json:"group,omitempty"`
	IncludeNoMetricHosts  *bool    `json:"noMetricHosts,omitempty"`
	Scopes                []string `json:"scope,omitempty"`
	IncludeUngroupedHosts *bool    `json:"noGroupHosts,omitempty"`
	NodeType              *string  `json:"nodeType,omitempty"`
}

// Graph represents a graph that might exist on a dashboard.
type Graph struct {
	Title      *string          `json:"title,omitempty"`
	Definition *GraphDefinition `json:"definition"`
}

// Template variable represents a template variable that might exist on a dashboard
type TemplateVariable struct {
	Name    *string `json:"name,omitempty"`
	Prefix  *string `json:"prefix,omitempty"`
	Default *string `json:"default,omitempty"`
}

// Dashboard represents a user created dashboard. This is the full dashboard
// struct when we load a dashboard in detail.
type Dashboard struct {
	Id                *int               `json:"id,omitempty"`
	NewId             *string            `json:"new_id,omitempty"`
	Description       *string            `json:"description,omitempty"`
	Title             *string            `json:"title,omitempty"`
	Graphs            []Graph            `json:"graphs,omitempty"`
	TemplateVariables []TemplateVariable `json:"template_variables,omitempty"`
	ReadOnly          *bool              `json:"read_only,omitempty"`
}

// DashboardLite represents a user created dashboard. This is the mini
// struct when we load the summaries.
type DashboardLite struct {
	Id          *int       `json:"id,string,omitempty"` // TODO: Remove ',string'.
	Resource    *string    `json:"resource,omitempty"`
	Description *string    `json:"description,omitempty"`
	Title       *string    `json:"title,omitempty"`
	ReadOnly    *bool      `json:"read_only,omitempty"`
	Created     *string    `json:"created,omitempty"`
	Modified    *string    `json:"modified,omitempty"`
	CreatedBy   *CreatedBy `json:"created_by,omitempty"`
}

// CreatedBy represents a field from DashboardLite.
type CreatedBy struct {
	Disabled   *bool   `json:"disabled,omitempty"`
	Handle     *string `json:"handle,omitempty"`
	Name       *string `json:"name,omitempty"`
	IsAdmin    *bool   `json:"is_admin,omitempty"`
	Role       *string `json:"role,omitempty"`
	AccessRole *string `json:"access_role,omitempty"`
	Verified   *bool   `json:"verified,omitempty"`
	Email      *string `json:"email,omitempty"`
	Icon       *string `json:"icon,omitempty"`
}

// reqGetDashboards from /api/v1/dash
type reqGetDashboards struct {
	Dashboards []DashboardLite `json:"dashes,omitempty"`
}

// reqGetDashboard from /api/v1/dash/:dashboard_id
type reqGetDashboard struct {
	Resource  *string    `json:"resource,omitempty"`
	Url       *string    `json:"url,omitempty"`
	Dashboard *Dashboard `json:"dash,omitempty"`
}

type DashboardConditionalFormat struct {
	Palette        *string      `json:"palette,omitempty"`
	Comparator     *string      `json:"comparator,omitempty"`
	CustomBgColor  *string      `json:"custom_bg_color,omitempty"`
	Value          *json.Number `json:"value,omitempty"`
	Inverted       *bool        `json:"invert,omitempty"`
	CustomFgColor  *string      `json:"custom_fg_color,omitempty"`
	CustomImageUrl *string      `json:"custom_image,omitempty"`
}

// GetDashboard returns a single dashboard created on this account.
func (client *Client) GetDashboard(id interface{}) (*Dashboard, error) {

	stringId, err := GetStringId(id)
	if err != nil {
		return nil, err
	}

	var out reqGetDashboard
	if err := client.doJsonRequest("GET", fmt.Sprintf("/v1/dash/%s", stringId), nil, &out); err != nil {
		return nil, err
	}
	return out.Dashboard, nil
}

// GetDashboards returns a list of all dashboards created on this account.
func (client *Client) GetDashboards() ([]DashboardLite, error) {
	var out reqGetDashboards
	if err := client.doJsonRequest("GET", "/v1/dash", nil, &out); err != nil {
		return nil, err
	}
	return out.Dashboards, nil
}

// DeleteDashboard deletes a dashboard by the identifier.
func (client *Client) DeleteDashboard(id int) error {
	return client.doJsonRequest("DELETE", fmt.Sprintf("/v1/dash/%d", id), nil, nil)
}

// CreateDashboard creates a new dashboard when given a Dashboard struct. Note
// that the Id, Resource, Url and similar elements are not used in creation.
func (client *Client) CreateDashboard(dash *Dashboard) (*Dashboard, error) {
	var out reqGetDashboard
	if err := client.doJsonRequest("POST", "/v1/dash", dash, &out); err != nil {
		return nil, err
	}
	return out.Dashboard, nil
}

// UpdateDashboard in essence takes a Dashboard struct and persists it back to
// the server. Use this if you've updated your local and need to push it back.
func (client *Client) UpdateDashboard(dash *Dashboard) error {
	return client.doJsonRequest("PUT", fmt.Sprintf("/v1/dash/%d", *dash.Id),
		dash, nil)
}
