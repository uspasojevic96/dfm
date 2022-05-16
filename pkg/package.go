package pkg

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/uspasojevic96/dfm/util"
)

type Package struct {
	Name          string   `json:"name"`
	Version       string   `json:"version"`
	InstallScript string   `json:"installScript"`
	AffectedFiles []string `json:"affectedFiles"`
}

func (p *Package) Uninstall() error {
	return errors.New("not implemented")
}

func (p *Package) Install() error {
	cache := util.GetDFMInstallCachePath()
	path := cache + string(os.PathSeparator) + p.Name

	if !util.FileExists(path) {
		if err := os.Mkdir(path, 0644); err != nil {
			return err
		}
	}

	for _, file := range p.AffectedFiles {
		if util.FileExists(file) {
			if err := util.CopyFile(file, path+string(os.PathSeparator)+file); err != nil {
				return err
			}
		}
	}

	execPath := util.GetDFMRepoPath() + string(os.PathSeparator) + p.Name + string(os.PathSeparator) + p.InstallScript
	if err := exec.Command("sh", "-c", execPath).Run(); err != nil {
		return err
	}

	return nil
}

func (p *Package) IsInstalled() bool {
	return false
}

func (p *Package) Info() string {
	installed := "Not Installed"

	if p.IsInstalled() {
		installed = "Installed"
	}

	return p.Name + " [" + p.Version + "] " + " - " + installed
}

func LoadPackages(path string) ([]*Package, error) {
	files, err := os.ReadDir(path)

	var packages []*Package
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			pkgPath := path + string(os.PathSeparator) + file.Name() + string(os.PathSeparator) + "metadata.json"

			if util.FileExists(pkgPath) {
				var pkg Package
				content, err := ioutil.ReadFile(pkgPath)
				if err != nil {
					return nil, err
				}

				err = json.Unmarshal(content, &pkg)
				if err != nil {
					return nil, err
				}

				packages = append(packages, &pkg)
			}
		}
	}

	return packages, nil
}
