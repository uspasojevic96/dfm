package pkg

import (
	"errors"
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
	return nil, errors.New("not implemented")
}
