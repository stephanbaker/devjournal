package main

import (
	"os"
	"testing"
	"time"
)

type statNotExists struct{}

func (fs statNotExists) Stat(name string) (os.FileInfo, error) {
	return nil, os.ErrNotExist
}

func (fs statNotExists) IsNotExist(error) bool {
	return true
}

func (fs statNotExists) Create(string) (*os.File, error) {
	return nil, nil
}

type statIsFile struct{}

func (fs statIsFile) Stat(name string) (os.FileInfo, error) {
	return fileInfoIsFile{}, nil
}

func (fs statIsFile) IsNotExist(error) bool {
	return false
}

func (fs statIsFile) Create(string) (*os.File, error) {
	return nil, nil
}

type statIsDir struct{}

func (fs statIsDir) Stat(name string) (os.FileInfo, error) {
	return fileInfoIsDir{}, nil
}

func (fs statIsDir) IsNotExist(error) bool {
	return false
}

func (fs statIsDir) Create(string) (*os.File, error) {
	return nil, nil
}

type fileInfoIsFile struct{}

func (f fileInfoIsFile) Name() string {
	return ""
}

func (f fileInfoIsFile) Size() int64 {
	return 0
}

func (f fileInfoIsFile) Mode() os.FileMode {
	return 0
}

func (f fileInfoIsFile) ModTime() time.Time {
	return time.Now()
}

func (f fileInfoIsFile) IsDir() bool {
	return false
}

func (f fileInfoIsFile) Sys() interface{} {
	return nil
}

type fileInfoIsDir struct{}

func (f fileInfoIsDir) Name() string {
	return ""
}

func (f fileInfoIsDir) Size() int64 {
	return 0
}

func (f fileInfoIsDir) Mode() os.FileMode {
	return 0
}

func (f fileInfoIsDir) ModTime() time.Time {
	return time.Now()
}

func (f fileInfoIsDir) IsDir() bool {
	return true
}

func (f fileInfoIsDir) Sys() interface{} {
	return nil
}

func TestFileNotExists(t *testing.T) {
	fs := statNotExists{}
	if fileExists("filename", fs) {
		t.Error("Test should have returned FileNotExists error")
	}
}

func TestStatFile(t *testing.T) {
	fs := statIsFile{}
	if !fileExists("filename", fs) {
		t.Error("Test should have indicated file exists")
	}

}

func TestStatDir(t *testing.T) {
	fs := statIsDir{}
	if fileExists("filename", fs) {
		t.Error("Test should have indicated file does not exist")
	}

}
