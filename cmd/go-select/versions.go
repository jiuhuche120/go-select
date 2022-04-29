package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func versionsCMD() *cli.Command {
	return &cli.Command{
		Name:   "versions",
		Usage:  "prints out all installed golang versions",
		Action: versions,
	}
}

func versions(ctx *cli.Context) error {
	path, err := DefaultArtifacts()
	if err != nil {
		return err
	}
	if !Exist(path) {
		fmt.Println("No golang version set")
		err := Initialize()
		if err != nil {
			return err
		}
		err = os.Mkdir(path, 0777)
		if err != nil {
			return err
		}
	}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}
	var versions []string
	for _, file := range files {
		strs := strings.Split(file.Name(), "_")
		if len(strs) != 2 {
			continue
		}
		versions = append(versions, strs[1])
	}
	if len(versions) == 0 {
		fmt.Println("No golang version set")
		return nil
	}
	global, err := GlobalVersion()
	if err != nil {
		return err
	}
	for _, version := range versions {
		if global == version {
			fmt.Println(version + " (√)")
		} else {
			fmt.Println(version)
		}
	}
	return nil
}

func GlobalVersion() (string, error) {
	global, err := DefaultGlobalVersion()
	if err != nil {
		return "", err
	}
	if !Exist(global) {
		return "", errors.New("no global_version file")
	}
	data, err := ioutil.ReadFile(global)
	if err != nil {
		return "", err
	}
	version := strings.ReplaceAll(string(data), "\n", "")
	return version, nil
}
