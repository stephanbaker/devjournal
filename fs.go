package main

import "os"

type fileSystem interface {
	Stat(string) (os.FileInfo, error)
	IsNotExist(error) bool
	Create(string) (*os.File, error)
}

type defaultFileSystem struct{}

func (d defaultFileSystem) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (d defaultFileSystem) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

func (d defaultFileSystem) Create(name string) (*os.File, error) {
	return os.Create(name)
}
