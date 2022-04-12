package put

import (
	"encoding/xml"
	"io/ioutil"
	"sync"
)

// helper
type IData interface {
	Set(
		name string,
		city string,
	)
}

type SData struct {
	name string
	city string
}

// store to db

type SStore struct {
	db map[string]string
	m  sync.Mutex
}

func (data *SStore) Save() error {

	body, err := xml.Marshal(data.db)

	if err != nil {
		return err
	}

	file, err := xml.MarshalIndent(body, "  ", "    ")

	if err != nil {
		return err
	}

	_ = ioutil.WriteFile("db.xml", file, 0644)

	return nil
}
