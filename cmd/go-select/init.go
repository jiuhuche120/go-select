package main

import (
	"os"

	"github.com/mitchellh/go-homedir"
)

const (
	defaultPath   = "~/.go-select"
	artifacts     = defaultPath + "/artifacts"
	globalVersion = defaultPath + "/global-version"
)

func Initialize() error {
	path, err := DefaultPath()
	if err != nil {
		return err
	}
	if !Exist(path) {
		err := os.Mkdir(path, 0777)
		if err != nil {
			return err
		}
	}
	return nil
}

func Exist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func DefaultPath() (string, error) {
	dir, err := homedir.Expand(defaultPath)
	if err != nil {
		return "", err
	}
	return dir, nil
}

func DefaultArtifacts() (string, error) {
	dir, err := homedir.Expand(artifacts)
	if err != nil {
		return "", err
	}
	return dir, nil
}

func DefaultGlobalVersion() (string, error) {
	dir, err := homedir.Expand(globalVersion)
	if err != nil {
		return "", err
	}
	return dir, nil
}
