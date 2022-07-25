package servers

import (
	"embed"
	"encoding/json"
)

//go:embed rbls.json
var reader embed.FS

type Servers struct {
}

func New() *Servers {
	return &Servers{}
}

func (s *Servers) GetRblServers() (*RBLServers, error) {
	bFile, err := reader.ReadFile("rbls.json")
	if err != nil {
		return nil, err
	}

	var rbls RBLServers

	if err := json.Unmarshal(bFile, &rbls); err != nil {
		return nil, err
	}

	return &rbls, nil
}
