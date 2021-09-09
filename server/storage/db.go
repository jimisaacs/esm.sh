package storage

import (
	"errors"
	"fmt"
	"sync"
	"time"

	logx "github.com/ije/gox/log"
	"github.com/ije/gox/utils"
)

var ErrorNotFound = errors.New("record not found")

type Store map[string]string

type DB interface {
	Open(config string, log *logx.Logger, isDev bool) (conn DBConn, err error)
}

type DBConn interface {
	Get(id string) (store Store, modtime time.Time, err error)
	Put(id string, store Store) error
	Delete(id string) error
	Close() error
}

var dbs = sync.Map{}

func OpenDB(dbUrl string, log *logx.Logger, isDev bool) (DBConn, error) {
	name, config := utils.SplitByFirstByte(dbUrl, ':')
	db, ok := dbs.Load(name)
	if ok {
		return db.(DB).Open(config, log, isDev)
	}
	return nil, fmt.Errorf("unregistered db '%s'", name)
}

func RegisterDB(name string, db DB) error {
	_, ok := dbs.Load(name)
	if ok {
		return fmt.Errorf("db '%s' has been registered", name)
	}

	dbs.Store(name, db)
	return nil
}
