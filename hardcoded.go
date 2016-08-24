// Code generated by go-bindata.
// sources:
// hardcoded/_doer.go
// hardcoded/_middleware.go
// DO NOT EDIT!

package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _hardcoded_doerGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xa4\x54\x51\x6f\xdb\x36\x10\x7e\xb6\x7e\xc5\x55\xc0\x50\x69\xf5\xe4\xbc\xf4\xa5\x5d\xb1\x07\xa7\x5d\xfa\xb0\x65\xe8\x8c\x6d\xaf\xb4\x48\xc9\x6c\x24\x92\x23\xa9\x24\x86\xe1\xff\xbe\x3b\x92\x92\x65\xa7\x0b\x0a\xa4\x28\xe0\x90\x77\xfa\xee\xfb\xee\x3e\x9e\x61\xf5\x1d\x6b\x05\xd4\x9d\x14\xca\x67\x99\xec\x8d\xb6\x1e\x8a\x6c\x91\x37\xbd\xcf\xf1\x47\x09\xbf\xda\x79\x6f\xf2\x0c\x0f\xad\xee\x98\x6a\x2b\x6d\xdb\xd5\xe3\x8a\x22\xb5\x56\x5e\x3c\x86\xc4\xff\x8b\xad\x6a\xff\x38\x02\x68\x83\x55\x2c\xab\xa5\x6a\x21\x6f\xa5\xdf\x0d\xdb\xaa\xd6\xfd\x6a\x76\x3f\xff\xfb\xa7\x56\xe7\x59\x99\x65\xab\x15\x70\x2d\x2c\x48\x07\x4c\x81\x44\x58\xdb\xb0\x5a\x40\xa3\x2d\xe4\x5c\x63\x66\x0e\x54\x02\xac\xf8\x77\x10\xce\x3b\x30\xda\x39\xb9\xed\xf6\xf0\x80\x45\xe0\xc1\x32\x63\x30\x2b\xf3\x7b\x23\x12\xd4\x04\x72\xc8\x16\xd7\xba\xa8\xe1\x47\x42\xa8\xd6\xa1\x11\x4b\xb0\xe9\xfc\x25\x22\x96\x50\x8c\x67\x67\xb4\x72\x62\x09\xc2\x5a\x6d\xcb\xec\x98\x45\x54\x6d\x7e\x67\xbd\x58\xfb\x47\x70\xde\x0e\xb5\x3f\x1c\x03\xef\x2d\x73\x22\x61\xdc\x30\xc5\x3b\x2c\x6d\xb0\xb0\xb6\xbd\x03\xbf\x13\x21\x7e\xc6\x3d\xa2\xd1\xf5\x35\xf1\x3c\x81\x35\x83\xaa\xa1\xe0\x53\xa8\x84\x17\xd0\x26\xd5\x56\xf8\xc1\x2a\x48\xe3\xa9\x10\xcd\x56\xeb\x38\xb3\xa2\x5c\x42\x8d\x68\x41\x1e\xaa\x48\xf3\x08\x8c\x18\xe7\x6e\xbc\x00\xaf\xcf\x1b\x1f\xd9\xcf\xd3\xa3\x00\xaa\xc7\x43\xe7\xb3\x99\x94\x59\xde\x4b\xd5\x64\x0b\xd4\x01\xef\x3e\xc0\x4c\x03\xf9\x8d\x86\x42\xd7\x18\xad\xfe\x62\xdd\x20\x8a\x69\x50\x87\x63\x59\x15\x48\x0f\x19\x60\xea\x3d\x43\xae\x06\x66\xee\xab\xfe\x34\x4c\x65\x0b\x94\xbf\xb9\xbd\xbe\x7d\x47\xc2\xc1\xb3\xd6\xa1\xd8\x8e\xf9\xa4\x5e\x2a\x33\x78\xe0\xcc\xb3\x5f\xb2\x85\x6c\xc0\x30\x8b\x00\xf4\x25\x55\xbd\x44\xfb\x64\x75\x3f\xd2\x43\x46\xe5\xfb\x79\xfe\xab\x0f\xa0\x64\x47\x9d\x5a\x20\x91\x8b\x8f\x3d\xb3\x21\x2b\xd1\x5f\x9e\x45\xd7\x3b\xd9\xf1\xdb\xa6\x38\x81\x9d\x9a\x50\xa2\xb6\x23\x88\xce\x89\xef\x43\xa6\xf4\xa0\x04\x5b\x4b\x12\x9c\xa9\x36\x98\x2a\x6c\x51\x56\x9f\xd5\x57\x51\xfb\x02\xaf\x4e\x3e\x41\xcc\x39\xdc\xcd\x66\xf3\xc7\x8d\x60\x5c\x58\xf7\x4c\x68\xcd\xac\x95\x08\x69\xab\x78\x51\x62\x27\xa8\xde\xac\x05\xc9\x9d\x78\x5c\x02\xae\xa2\xea\x23\x4d\xba\x29\xf2\x5a\x0f\x1d\x57\xaf\x3d\x76\x9e\xc8\x4c\x4e\xdc\x45\x64\x28\x7e\xb8\x2f\xf3\x60\x8c\xa8\x24\xe1\xf0\x8a\x93\xc3\xe7\xa6\xc6\x88\xdd\x8f\x8f\x92\x0e\x52\x38\x78\x7b\xf5\xcf\xb7\x1c\x1d\x72\x9f\xf8\xf9\xec\x5f\x30\xf7\x82\x8b\x86\x0d\x9d\xff\x92\xe0\x70\xcd\xa4\x6a\x7f\xe3\x22\xa2\xdb\x3d\x44\x46\xb8\xca\x40\x89\x07\x48\x5b\x12\x97\x01\xf3\xa0\xef\x91\xb7\xe4\x22\xee\x06\x35\xf4\x5b\xac\xa8\x9b\x89\x1d\x1a\x8e\xeb\xb0\xf6\x18\x39\xc7\xcb\x7a\xe8\x98\x8d\x62\x02\xdd\x2a\xbe\xad\xa9\x18\xb9\x6c\x2c\x31\x0e\x6d\x39\xc1\x21\xbb\xf2\x32\x3a\xdf\x0c\x29\x42\x68\xf1\xed\x20\x5a\xfc\x7a\x9f\xb2\x0f\xc7\x09\xed\xac\xad\x23\x98\x8c\x4a\xee\xc4\x3e\xfc\x86\x18\xe1\x36\xb2\x1d\x2c\x3e\x22\xad\x2a\xf8\x84\x72\xb8\xe8\xf1\x41\xfb\x78\x05\x66\xb0\xb8\xc0\x89\xa0\x7f\xed\xe0\xeb\xe0\x3c\xea\xc5\xc1\x63\x33\xb1\x00\xf6\xe3\xdb\xdd\xb1\xb2\xdd\x79\x50\xfa\xa1\x9a\x0d\x6d\x64\xf2\x74\x91\x4e\x33\x7d\xf9\xee\xb1\x78\xb9\x1c\x5f\xcc\x99\xd5\xc6\x97\xf4\xd4\xd9\xd3\x37\x64\xd3\xb0\x66\x3e\x37\xa9\x49\xa1\x22\xd4\x8c\x6c\xbe\x15\x49\x1f\xa7\xa0\x8a\xdd\x48\x10\xb2\xef\x05\x97\xcc\x8b\x6e\x1f\xdb\xe8\x77\xd8\x70\x63\x35\x36\x05\xff\x63\x9f\x6b\x61\xd0\x55\xaa\xdb\x87\x02\xbf\x7e\xdc\x38\x82\x9d\x81\x06\x82\xb6\xfa\x4d\xf8\x9d\xe6\xc4\x32\xc7\xa4\xfc\x19\x9e\xb4\x27\x67\xfe\x09\x66\xa1\x03\x6e\xa4\xbb\x8b\xfd\x9b\x16\xee\xb9\x5f\x70\xe7\x92\xed\x42\xdd\x57\xf8\x49\xaa\x14\xf0\x42\xef\xce\x1e\x50\xac\x49\x86\x97\x04\x7e\xf5\x1e\x7f\x7f\x1e\xeb\xe3\xe1\xcd\x9b\x00\x40\x1a\x90\x26\x2d\x33\x3f\xb8\xb5\xe6\x02\xb3\xde\x5e\x5d\x85\xe0\x62\x6b\x05\xbb\xc3\x3f\x8e\xa1\xd4\x38\xa9\xcb\x41\x9d\x76\xc5\x49\xf1\x31\xfb\x2f\x00\x00\xff\xff\x74\x68\xba\x23\x21\x09\x00\x00")

func hardcoded_doerGoBytes() ([]byte, error) {
	return bindataRead(
		_hardcoded_doerGo,
		"hardcoded/_doer.go",
	)
}

func hardcoded_doerGo() (*asset, error) {
	bytes, err := hardcoded_doerGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "hardcoded/_doer.go", size: 2337, mode: os.FileMode(420), modTime: time.Unix(1471457554, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _hardcoded_middlewareGo = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\x8c\x54\x41\x6f\xdb\x3a\x0c\x3e\xdb\xbf\x82\xcf\x27\xfb\x21\x4f\x39\xbc\x5b\x87\x1e\x8a\xac\x6d\x0a\x74\x6d\xd1\x64\xe8\x71\x50\x6c\xda\xd6\xe2\x48\x9e\xc4\x38\x29\x86\xfc\xf7\x51\xb2\x93\xba\x6b\x57\xec\x10\x20\xa2\x3e\x7e\xfc\x3e\x92\x56\x2b\xf3\xb5\xac\x10\x1c\xda\x0e\x6d\x1c\xab\x4d\x6b\x2c\x41\x1a\x47\x89\x46\x9a\xd6\x44\x6d\x12\xf3\xa1\x32\xed\xba\x12\x4a\x4f\x67\x0d\x32\x70\xba\x96\xcf\x1d\xe2\x7f\x95\x11\xdd\xff\xd3\xc6\x54\x15\xda\x24\x8e\xd6\xdd\x17\x55\x14\x0d\xee\xa4\x45\xf8\x38\x67\x73\x02\x7a\x7e\xd3\xa2\x26\x2b\x73\xa5\x2b\xce\x53\x54\x6f\x57\x22\x37\x9b\xe9\x28\x3e\xfe\xcf\x1c\x49\x9c\xc5\x71\xb9\xd5\x39\xec\x18\xfe\x52\x36\xb5\x66\x4b\x68\xc1\x2b\x17\x73\xa9\x39\x6c\xb3\x57\x27\xf8\xc9\x05\xeb\xe1\xff\xd9\x39\x8c\x45\x8b\x3b\xdc\x0d\x0c\x13\xe8\x6d\x85\x50\xb2\xbc\xff\x7c\x7f\x06\xb3\xf9\xc5\xdd\xf5\x25\x2c\xe7\x37\x8b\x24\xcb\x5e\x58\xce\x61\x10\x36\xd2\x31\xdc\x31\xca\x22\x6d\xad\x86\x21\x10\x1f\xe2\x78\x3a\x7d\x9b\x00\xb9\x45\x49\xe8\x40\x82\xc6\x1d\xb8\x56\x6a\xd0\x72\x83\x05\xc8\xd2\x1b\xa2\x1a\xe1\xeb\xe3\x2d\xb4\x92\x6a\x30\x65\x38\x5b\xfc\xb1\x45\x47\xc2\x13\xde\x10\xb4\x8d\xcc\x99\x80\x6a\xe5\xfa\x7c\xa5\xc7\x30\xc8\x8d\x26\xdc\xd3\x04\x4a\x63\x61\xeb\x10\x56\xcf\x60\x18\x60\x8f\xda\x1c\x74\x4a\xc2\xa8\xd1\x62\xc1\x34\x57\xd6\x6c\x66\x7d\x6a\x9a\x85\x52\x25\x8b\x0c\x05\x70\xaf\x1c\x39\x5f\xe7\x58\xa3\x46\x59\x30\xd1\x24\x14\x0e\x98\xde\x57\xe1\x8b\x05\x65\x2f\x93\xe7\xd1\x35\x0d\xac\x90\xd9\xf2\x5a\x35\x45\x6f\x4b\x52\xc8\x13\xfd\x74\xdf\x69\xec\xc7\xb3\x3d\xb5\x7b\x14\xbe\x62\xa6\xd4\xd3\xa5\xbb\x3e\xfe\x88\xae\x35\xda\xe1\x93\x55\x61\xd4\x16\xfe\x1d\xe2\xc1\x45\xe6\x79\x22\x76\x7a\x41\x84\x9b\x96\x80\x0c\x7c\x37\xec\x72\xb0\xcd\x56\x2a\x24\xf2\xdb\xea\xf5\x21\x37\xa0\x34\x50\x72\x9f\x82\xed\xa1\x07\x82\x39\x4c\x7b\xc7\x33\xf4\x6b\x66\x05\x4f\x4f\x3c\xf0\xf4\x38\xdc\x49\xcb\x4c\x6f\x3a\xcd\x37\xaa\x04\x97\x4f\x00\x6d\xd8\xcd\x31\xe0\xba\x31\x2b\xd9\x2c\x7d\x3d\x9b\x66\x9e\x3c\xba\xdc\xfb\x4b\x4a\xc7\xb0\xf9\x72\xf9\x30\x1f\x86\xe0\x31\x7f\xb8\x9b\x49\x6b\x15\x13\x59\xd1\x07\xb2\xec\x53\x28\xfa\xcf\x39\x68\xd5\x04\xff\x11\x2b\x7c\x2d\x61\x41\xd2\x92\x17\x9a\xf6\xbe\x78\xb9\xa3\x03\x60\xc3\xab\xf4\x77\x09\x93\x57\xb7\x33\x3f\xf4\xfb\x32\x75\x79\x16\x98\xf8\x57\x60\x89\xbe\x35\xe2\x4a\x69\xe5\xea\xd4\xc7\xf9\x74\x6b\xaa\xcb\x8e\x13\xd3\x64\xd8\xb3\x6f\x16\x73\x54\x1d\x16\x49\x76\xca\x0a\x03\xce\x8e\x4a\xde\xc9\x29\x03\xe7\x90\x73\x08\xdc\xfc\xa9\xcd\x68\xff\x7b\xab\x87\x6d\x7f\xe2\x97\x25\x88\xb7\xe2\xb4\xff\x13\x16\xc7\x2f\x4f\x14\xd5\x62\xe1\xdf\x4c\xdf\xd2\x74\xc7\x1b\x24\x3c\xfa\x08\xeb\x69\xbd\xab\x43\xc6\x1f\xfc\xaf\x00\x00\x00\xff\xff\xd6\x98\x0e\x95\x66\x05\x00\x00")

func hardcoded_middlewareGoBytes() ([]byte, error) {
	return bindataRead(
		_hardcoded_middlewareGo,
		"hardcoded/_middleware.go",
	)
}

func hardcoded_middlewareGo() (*asset, error) {
	bytes, err := hardcoded_middlewareGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "hardcoded/_middleware.go", size: 1382, mode: os.FileMode(420), modTime: time.Unix(1471474022, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"hardcoded/_doer.go":       hardcoded_doerGo,
	"hardcoded/_middleware.go": hardcoded_middlewareGo,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"hardcoded": &bintree{nil, map[string]*bintree{
		"_doer.go":       &bintree{hardcoded_doerGo, map[string]*bintree{}},
		"_middleware.go": &bintree{hardcoded_middlewareGo, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}