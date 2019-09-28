package sls

import (
	"strings"
)

// GetHistogramsResponse defines response from GetHistograms call
type SingleHistogram struct {
	Progress string `json:"progress"`
	Count    int64  `json:"count"`
	From     int64  `json:"from"`
	To       int64  `json:"to"`
}

type GetHistogramsResponse struct {
	Progress   string            `json:"progress"`
	Count      int64             `json:"count"`
	Histograms []SingleHistogram `json:"histograms"`
}

func (resp *GetHistogramsResponse) IsComplete() bool {
	return strings.ToLower(resp.Progress) == "complete"
}

// GetLogsResponse defines response from GetLogs call
type GetLogsResponse struct {
	Progress string              `json:"progress"`
	Count    int64               `json:"count"`
	Logs     []map[string]string `json:"logs"`
}

func (resp *GetLogsResponse) IsComplete() bool {
	return strings.ToLower(resp.Progress) == "complete"
}

// IndexKey ...
type IndexKey struct {
	Token         []string `json:"token"` // tokens that split the log line.
	CaseSensitive bool     `json:"caseSensitive"`
	Type          string   `json:"type"` // text, long, double
	DocValue      bool     `json:"doc_value,omitempty"`
	Alias         string   `json:"alias,omitempty"`
	Chn           bool     `json:"chn"` // parse chinese or not
}

type IndexLine struct {
	Token         []string `json:"token"`
	CaseSensitive bool     `json:"caseSensitive"`
	IncludeKeys   []string `json:"include_keys,omitempty"`
	ExcludeKeys   []string `json:"exclude_keys,omitempty"`
	Chn           bool     `json:"chn"` // parse chinese or not
}

// Index is an index config for a log store.
type Index struct {
	Keys map[string]IndexKey `json:"keys,omitempty"`
	Line *IndexLine          `json:"line,omitempty"`
}

// CreateDefaultIndex return a full text index config
func CreateDefaultIndex() *Index {
	return &Index{
		Line: &IndexLine{
			Token:         []string{" ", "\n", "\t", "\r", ",", ";", "[", "]", "{", "}", "(", ")", "&", "^", "*", "#", "@", "~", "=", "<", ">", "/", "\\", "?", ":", "'", "\""},
			CaseSensitive: false,
		},
	}
}
