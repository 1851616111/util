package db

import (
	"errors"
	"fmt"
	"github.com/1851616111/tantan-test/router"
	"gopkg.in/pg.v4"
	"log"
	"path/filepath"
	"sync"
)

type DB struct {
	// value *pg.DB is only a tcp connections manager
	// it is not care of the database, username, password
	// but it cares the connection's other params like poolSize, poolTimeout, idleTimeout, idleCheckFrequency
	// so if user and relationship user the same address database, it needs different connection manager for that params
	TableKindToAddrToConsMappings map[string]map[string]*pg.DB

	//as the user and relationship id may start with different beginnings, and size, it needs to have different db router
	TableKindToRouterMappings map[string]router.Router

	// this is used for list query
	// thought this is a 3 layer map, is only contains limited kind->addr->database->table config metadata
	// so don'test worry about efficiency， its fast
	// it can change to a array list
	TableKindToAddrToDBToTableMappings map[string]map[string]map[string][]string

	sync.RWMutex
}

func NewDBFromJsonFile(path string) (*DB, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	config, err := loadConfig(absPath)
	if err != nil {
		return nil, err
	}

	if len(config.Databases) == 0 {
		return nil, fmt.Errorf("parse db config file %s, nil databases config", path)
	}

	var consMappings map[string]map[string]*pg.DB
	var routerMappings map[string]router.Router
	if consMappings, err = newDBConsMappings(config); err != nil {
		return nil, fmt.Errorf("new db connections err %v", err)
	}

	if routerMappings, err = newRouterMappings(config); err != nil {
		return nil, fmt.Errorf("new db router err %v", err)
	}

	return &DB{
		TableKindToAddrToConsMappings:      consMappings,
		TableKindToRouterMappings:          routerMappings,
		TableKindToAddrToDBToTableMappings: newtablesMappings(config),
	}, nil
}

func (db *DB) UsedDatabaseFunc(kind, addr, database string, fn func(db *pg.DB) error) error {

	if _, exists := db.TableKindToAddrToConsMappings[kind]; !exists {
		return ErrKindNotFound
	}
	if _, exists := db.TableKindToAddrToConsMappings[kind][addr]; !exists {
		return ErrDataBaseNotFound
	}

	d := db.TableKindToAddrToConsMappings[kind][addr]

	var err error = nil
	db.Lock()

	d.Options().Database = database
	err = fn(d)
	d.Options().Database = "" //clean database session

	db.Unlock()

	return err
}

func newDBConsMappings(config *config) (map[string]map[string]*pg.DB, error) {
	mm := map[string]map[string]*pg.DB{}

	for _, db := range config.Databases {
		if db.Host == "" || db.Port == 0 {
			log.Printf("database config addr(host=%s,port=%d) is invalid\n", db.Host, db.Port)
			continue
		}

		addr := fmt.Sprintf("%s:%d", db.Host, db.Port)
		opt := &pg.Options{
			Addr:     addr,
			User:     db.Credentials.User,
			Password: db.Credentials.Password,
			PoolSize: db.MaxPool,
		}

		if _, exists := mm[db.Label]; !exists {
			mm[db.Label] = map[string]*pg.DB{}
		}

		mm[db.Label][addr] = pg.Connect(opt)
	}

	return mm, nil
}

func newRouterMappings(config *config) (map[string]router.Router, error) {
	return map[string]router.Router{
		"user":                 newRouter("user", config.Databases),
		"relationship_like":    newRouter("relationship_like", config.Databases),
		"relationship_dislike": newRouter("relationship_dislike", config.Databases),
		"relationship_match":   newRouter("relationship_match", config.Databases),
	}, nil
}

func newtablesMappings(config *config) map[string]map[string]map[string][]string {
	m := map[string]map[string]map[string][]string{}

	for _, db := range config.Databases {
		if _, exists := m[db.Label]; !exists {
			m[db.Label] = map[string]map[string][]string{}
		}

		addr := fmt.Sprintf("%s:%d", db.Host, db.Port)
		if _, exists := m[db.Label][addr]; !exists {
			m[db.Label][addr] = map[string][]string{}
		}

		for _, instance := range db.Instances {
			if _, exists := m[db.Label][addr][instance.Name]; !exists {
				m[db.Label][addr][instance.Name] = instance.Tables
			}
		}
	}

	return m
}

var (
	ErrKindNotFound     = errors.New("kind not found")
	ErrDataBaseNotFound = errors.New("database not found")
)
