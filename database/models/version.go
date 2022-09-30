package models

type Version struct {
	Id         string `json:"id"`
	Commit     string `json:"commit,omitempty"`
	ReleaseTag string `json:"version,omitempty"`
}
