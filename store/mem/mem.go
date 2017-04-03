// Copyright 2017 Joan Llopis. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
package mem

import (
	"encoding/json"
	"os"

	"io"

	"github.com/jllopis/zbs/log"
	"github.com/jllopis/zbs/services"
	"github.com/jllopis/zbs/store"
)

// MemStore is a Storer implementation over PostgresSQL
type MemStore struct {
	Pool          map[int64]*services.JobDataMsg
	StoreFileName string
	StoreFile     *os.File
	Stat          int
}

var _ store.Storer = (*MemStore)(nil)

// New is a default Storer implementation based upon PostgresSQL
func New() (*MemStore, error) {
	return &MemStore{}, nil
}

// Dial perform the connection to the underlying database server
func (d *MemStore) Dial(options store.Options) error {
	// read the config file (if not provided, start anew)

	d.StoreFileName = options["host"].(string) + "/" + options["dbname"].(string)
	if err := d.initFromFile(); err != nil {
		return err
	}

	d.Stat = store.CONNECTED

	log.Info("created mem store", "mc", len(d.Pool))

	return nil
}

// Status return the current status of the underlying database
func (d *MemStore) Status() (int, string) {
	return d.Stat, store.StatusStr[d.Stat]
}

// Close effectively close de database connection
func (d *MemStore) Close() error {
	log.Info("GoPg store CLOSING")
	// save changes if any
	d.StoreFile.Sync()
	d.StoreFile.Close()
	d.Pool = nil
	d.Stat = store.DISCONNECTED
	return nil
}

func (d *MemStore) initFromFile() error {
	f, err := os.OpenFile(d.StoreFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	d.StoreFile = f

	// Load Jobs from file
	dec := json.NewDecoder(f)
	jobarr := &services.JobArrayMsg{}
	err = dec.Decode(jobarr)
	if err != nil {
		if err != io.EOF {
			return err
		}
	}
	if jobarr.Total > 0 {
		d.Pool = make(map[int64]*services.JobDataMsg, jobarr.Total)
		for _, j := range jobarr.Jobs {
			d.Pool[j.GetId()] = j.GetJobData()
		}
	} else {
		// Initialize map
		d.Pool = make(map[int64]*services.JobDataMsg, 0)
	}
	return nil
}

func (d *MemStore) fsync() error {
	d.StoreFile.Seek(0, 0)
	err := json.NewEncoder(d.StoreFile).Encode(d.Pool)
	return err
}
