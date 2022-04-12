package put

import (
	"encoding/xml"
	"errors"
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

func (data *SStore) Set(name string, city string) {
	data.m.Lock()
	defer data.m.Unlock()

	if data.db == nil {
		data.db = make(map[string]string)
	}
	data.db[name] = city
}

func (data *SStore) Get(name string) (string, error) {
	data.m.Lock()
	defer data.m.Unlock()

	okStatus := data.db
	st := data.db[name]

	if okStatus == nil {
		return "Error: ", errors.New("not Found")
	}

	return st, nil
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
