package sls

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// CreateETLJob creates a new ETL job in SLS.
func (p *LogProject) CreateETLJob(j *ETLJob) error {
	body, err := json.Marshal(j)
	if err != nil {
		return NewClientError(err)
	}
	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}
	r, err := request(p, "POST", "/etljobs", h, body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}

// GetETLJob returns ETL job according to job name.
func (p *LogProject) GetETLJob(name string) (*ETLJob, error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	r, err := request(p, "GET", "/etljobs/"+name, h, nil)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	job := &ETLJob{}
	err = json.Unmarshal(buf, job)
	if err != nil {
		return nil, err
	}
	return job, nil
}

// UpdateETLJob updates an ETL job according to job name,
// Not all fields of ETLJob can be updated
func (p *LogProject) UpdateETLJob(name string, job *ETLJob) error {
	body, err := json.Marshal(job)
	if err != nil {
		return NewClientError(err)
	}
	h := map[string]string{
		"x-log-bodyrawsize": fmt.Sprintf("%v", len(body)),
		"Content-Type":      "application/json",
		"Accept-Encoding":   "deflate", // TODO: support lz4
	}
	r, err := request(p, "PUT", "/etljobs/"+name, h, body)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}

// DeleteETLJob deletes a job according to job name.
func (p *LogProject) DeleteETLJob(name string) error {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	r, err := request(p, "DELETE", "/etljobs/"+name, h, nil)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return nil
}

// ListETLJobs returns all job names of project.
func (p *LogProject) ListETLJobs() ([]string, error) {
	h := map[string]string{
		"x-log-bodyrawsize": "0",
	}
	uri := fmt.Sprintf("/etljobs")
	r, err := request(p, "GET", uri, h, nil)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	type Body struct {
		Count   int
		ETLJobs []string `json:"etlJobNameList"`
		Total   int
	}
	body := &Body{}
	err = json.Unmarshal(buf, body)
	if err != nil {
		return nil, err
	}
	return body.ETLJobs, nil
}
