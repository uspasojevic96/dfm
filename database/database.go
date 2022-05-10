package database

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/uspasojevic96/dfm/pkg"
)

var databasePath = string(os.PathSeparator) + ".dfm" + string(os.PathSeparator) + "database.json"

type Database struct {
	Packages   map[string]*pkg.Package `json:"packages"`
	LastUpdate int64                   `json:"lastUpdate"`
}

func init() {
	home, err := homedir.Dir()

	if err != nil {
		log.Fatal(err)
	}

	path := home + string(os.PathSeparator) + ".dfm"

	fileinfo, err := os.Stat(path)

	_ = fileinfo

	if err != nil {
		err := os.Mkdir(path, 0644)

		if err != nil {
			log.Fatal(err)
		}
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
	home, err := homedir.Dir()

	if err != nil {
		return err
	}

	input, err := ioutil.ReadFile(home + databasePath)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(input, d); err != nil {
		return err
	}

	return nil
}

func (d *Database) SaveDatabase() error {
	home, err := homedir.Dir()

	if err != nil {
		return err
	}

	output, err := json.Marshal(d)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(home+databasePath, output, 0644); err != nil {
		return err
	}

	return nil
}
