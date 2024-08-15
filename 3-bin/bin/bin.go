package bin

import (
	"errors"
	"time"
)

type BinList struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

func newBin(id, name string, private bool) (*Bin, error) {
	if id == "" {
		return nil, errors.New("invalid ID")
	}
	if name == "" {
		return nil, errors.New("invalid Name")
	}

	bin := &Bin{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}

	return bin, nil

}

func newBinList(id, name string, private bool) (*BinList, error) {
	if id == "" {
		return nil, errors.New("invalid ID")
	}
	if name == "" {
		return nil, errors.New("invalid Name")
	}

	binList := &BinList{
		id:        id,
		private:   private,
		createdAt: time.Now(),
		name:      name,
	}

	return binList, nil

}
