package servers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

type Servers struct {
}

func New() *Servers {
	return &Servers{}
}

func (s *Servers) GetRblServers() (*RBLServers, error) {
	file, err := os.Open(RBLServersFile)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalln(err)
		}
	}()

	bFile, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var rbls RBLServers

	if err := json.Unmarshal(bFile, &rbls); err != nil {
		return nil, err
	}

	return &rbls, nil
}
