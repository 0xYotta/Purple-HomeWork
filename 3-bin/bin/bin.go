package bin

import (
	"cloud-json/file"
	"encoding/json"
	"errors"
	"time"
)

//TODO Fix BinList and make all like in vault

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func newBin(id, name string, private bool) *Bin {
	// if id == "" {
	// 	return nil, errors.New("invalid ID")
	// }
	// if name == "" {
	// 	return nil, errors.New("invalid Name")
	// }

	bin := &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}

	return bin

}

type BinList struct {
	Bins      []Bin     `json:"binList"`
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func newBinList(id, name string, private bool) (*BinList, error) {
	if id == "" {
		return nil, errors.New("invalid ID")
	}
	if name == "" {
		return nil, errors.New("invalid Name")
	}
	var binList = BinList{} //empty but nil, safe from pointer to nil error
	if ok := binList.availableData(); ok {
		return &binList, nil
	}

	binList = BinList{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
		Bins:      []Bin{},
	}

	return &binList, nil

}

func (binList *BinList) save() (bool, error) {

	//Converting structure to a byte slice.
	bytes, err := json.Marshal(binList)
	if err != nil {
		return false, err
	}

	file.WriteFile(bytes, "data.json")
	return true, nil
}

func (binList *BinList) availableData() bool {
	data, err := file.ReadFile("data.json")
	if err != nil {
		return false
	}

	err = json.Unmarshal(data, binList)
	if err != nil {
		return false
	}

	return true
}
