package out

import (
	"github.com/frodenas/gcs-resource"
)

type OutRequest struct {
	Source gcsresource.Source `json:"source"`
	Params Params             `json:"params"`
}

type Params struct {
	File string `json:"file"`
}

func (params Params) IsValid() (bool, string) {
	if params.File == "" {
		return false, "please specify file"
	}

	return true, ""
}

type OutResponse struct {
	Version  gcsresource.Version        `json:"version"`
	Metadata []gcsresource.MetadataPair `json:"metadata"`
}