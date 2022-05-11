package database

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/uspasojevic96/dfm/pkg"
	"github.com/uspasojevic96/dfm/util"
)

type Database struct {
	Packages   map[string]*pkg.Package `json:"packages"`
	LastUpdate int64                   `json:"lastUpdate"`
}

func New() *Database {
	path := util.GetDFMPath()

	_, err := os.Stat(path)

	if err != nil {
		err := os.Mkdir(path, 0644)

		if err != nil {
			log.Fatal(err)
		}
	}

	return &Database{
		Packages:   make(map[string]*pkg.Package),
		LastUpdate: 0,
	}
}

func (d *Database) AddPackage(p *pkg.Package) {
	d.Packages[p.Name] = p
}

func (d *Database) RemovePackage(name string) error {
	if d.Packages[name] != nil {
		err := d.Packages[name].Uninstall()

		if err != nil {
			return err
		}

		delete(d.Packages, name)
	}

	return errors.New("Package not found")
}

func (d *Database) LoadDatabase() error {
	path := util.GetDFMDatabasePath()
	input, err := ioutil.ReadFile(path)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(input, d); err != nil {
		return err
	}

	return nil
}

func (d *Database) SaveDatabase() error {
	path := util.GetDFMDatabasePath()

	output, err := json.Marshal(d)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(path, output, 0644); err != nil {
		return err
	}

	return nil
}
