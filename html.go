// Code generated by go-bindata.
// sources:
// html/default.css
// html/doc.html
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

var _htmlDefaultCss = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xac\xcd\x41\x0a\xc2\x30\x10\x85\xe1\xfd\x9c\x62\xc0\x75\xa1\x55\x11\x69\x2e\xe0\x35\x26\xc9\xa4\x06\xc7\x4c\x89\xad\x58\xc5\xbb\x0b\x36\xb8\x52\x74\xe1\xfa\xf1\xfe\x0f\xac\xfa\x09\x6f\x80\x88\x68\xc9\x1d\xba\xac\x63\xf2\x95\x53\xd1\xdc\xe2\xc2\xd7\xbc\x0e\x6c\x9e\x73\xd0\x34\x54\x81\x8e\x51\xa6\x16\x77\x2c\x67\x1e\xa2\x23\x03\x77\x80\x7d\x53\x12\xe5\xa7\x99\x52\xc7\x99\xbd\xf9\x10\x26\x89\x8e\xad\x8c\x5f\xd2\xaf\xed\x14\xaf\xdc\xe2\x6a\xdb\x5f\x66\x6f\xf9\xce\xfb\x33\x56\x17\xac\x2f\xd6\x8f\xbf\x66\x33\xff\x1e\x01\x00\x00\xff\xff\x9f\xb0\x6e\x3e\x5a\x01\x00\x00")

func htmlDefaultCssBytes() ([]byte, error) {
	return bindataRead(
		_htmlDefaultCss,
		"html/default.css",
	)
}

func htmlDefaultCss() (*asset, error) {
	bytes, err := htmlDefaultCssBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/default.css", size: 346, mode: os.FileMode(420), modTime: time.Unix(1477443166, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _htmlDocHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x58\x4f\x6f\xdc\xb6\x12\xbf\xef\xa7\x98\xb7\x87\x1c\xf2\xf2\xa4\xc4\xc8\xe9\x45\x56\xe1\x3a\x41\x8b\x22\x7f\x8c\xd8\x45\xd1\x5e\x02\x8a\x1a\x49\x8c\x29\x52\x21\x87\xbb\xde\x34\xfd\xee\xc5\x90\x5a\x7b\xff\xd8\x40\x94\x28\x17\xed\x92\x1c\xf2\x37\x33\x1c\xce\xfc\xc8\xe2\x3f\x2f\xdf\x9d\x5f\xfd\x79\xf1\x0a\x3a\xea\x75\xb9\x28\xb6\x3f\x28\xea\x72\x51\x90\x22\x8d\xe5\xe5\xa7\x20\x1c\x5e\x68\xb1\x01\x8f\x6e\x85\x0e\x6a\x2b\x43\x8f\x86\x04\x29\x6b\x8a\x3c\x89\x2d\x0a\xad\xcc\x35\xd0\x66\xc0\xd3\x25\xe1\x0d\xe5\xd2\xfb\x25\xf4\x58\x2b\x71\xba\xf4\xd2\x21\x9a\x25\x38\xd4\xa7\x4b\x4f\x1b\x8d\xbe\x43\xa4\x25\x74\x0e\x9b\xd3\x65\xce\xc0\x79\x8d\x8d\x08\x9a\x32\x9e\x58\x2e\x8a\x7c\x54\xa3\xb2\xf5\x86\x95\x7a\x36\xaa\x32\x3c\xac\x4a\xf7\x8c\x05\x4f\xca\xdf\xdf\xbf\xbe\x84\xe0\xb1\x86\x6a\x03\xd4\x21\xdc\x19\x81\x6e\x9c\x5b\xe4\xdd\x09\x1b\x29\xaa\xa8\x3d\x39\xfe\xd4\x65\x31\x38\x2c\x3b\xa2\xe1\xff\x79\xfe\x48\xd3\x8b\x24\xfc\xa8\xa5\x17\x79\x91\xf3\x58\x91\x53\x9d\x44\x2f\x3b\xbb\x06\x19\x9c\x43\x43\x30\xc4\xb5\x3d\x08\x53\x83\xb7\x3d\x82\x32\x8d\x05\x51\xd9\x40\xac\x41\x3f\x4e\xcb\xc9\x95\x5f\x83\x95\x1c\x62\x65\xc6\x7f\x0e\x81\xaf\x3a\xe5\x8f\x76\x61\xca\xea\xc6\x92\x6a\x94\x8c\x53\x7d\xf6\xd1\xf3\x02\xfb\x10\xbf\x20\xc1\x9e\x14\x34\xce\xf6\xd1\x97\x69\x9d\x27\xe0\x83\xec\x40\x78\x98\x88\x2d\xad\x21\x67\x75\xd2\x61\x73\x88\xfb\x8d\x8b\x79\x12\x8e\x8e\xf6\x87\x3b\x41\xc0\x41\x04\x2b\x89\x19\x44\x17\x56\x28\x6d\x8f\x1e\x04\x18\x5c\x6f\xc7\xa0\xb1\x0e\xce\x94\x8b\xf2\x35\x72\x97\x07\xb2\x20\xad\x31\x28\x09\xc8\x7e\xa3\x8e\xda\xb6\x2d\xba\xdc\x09\x3b\xd4\x99\x08\xb5\xb2\x3f\xb1\xa4\x6d\x9a\x2f\x1c\x2b\x5f\x6a\xac\x42\xcb\x33\x8e\xec\x40\x82\x28\x0f\x69\x89\x39\xf0\x6b\x21\x87\x09\xf0\x2f\xcf\xce\x2f\x46\xf4\xe8\x1f\x87\xbd\x25\x84\x71\x69\xb0\x0d\xa8\x97\xd1\x55\x73\xa8\xd6\x8b\x49\xaa\xbd\x39\xdb\xaa\xf6\x04\x66\xd9\x19\xea\xbe\x12\xfe\xb5\x6d\x81\xc5\xad\x53\x9f\xd3\x29\x99\x63\x6b\x0c\xd2\x04\x7c\x6d\xd7\xa0\x71\x85\x1a\x0c\xd2\xda\xba\x6b\x10\x92\xd4\x4a\xd1\x66\x0e\x5d\xf8\x3b\x41\x99\xf7\x67\xef\x2e\x40\xda\xbe\x17\xa6\x9e\xc5\x17\x8e\x26\xc1\x5f\x5d\xcc\x6b\x3d\xf9\x69\xf0\x97\x77\xf8\x33\xc0\x7b\xfc\x14\xd0\x48\x74\x13\x74\xb8\x9d\x33\xab\x23\x56\x56\x87\x1e\x27\xa8\x91\x26\x40\x27\x4c\xad\x95\x69\xe7\xd0\xe1\x33\x3a\x2b\xad\x69\x26\x68\xf1\x17\x3a\x7b\x6e\x4d\x93\xff\x6c\xcd\x47\x1b\xe6\xf2\xc9\xe3\x09\xd9\x29\x1d\x4d\x4e\x99\x42\xeb\x31\x4d\x7d\xe7\xc9\xd8\x6e\x31\x39\x21\xd3\x9e\x34\x42\x7b\xfc\x42\x2e\xe0\x7d\x6a\xbc\x32\x4c\x71\x60\x6f\x1a\x6b\xf2\x04\x88\x8b\xe0\x5a\x69\x0d\xd2\xa1\x20\x04\x01\x1e\x07\xe1\xf8\xaf\xb6\x6d\xa3\x78\xda\x80\x32\xb2\x00\xad\x37\xd1\x8e\xbb\x08\x53\x06\x72\xea\x87\xfc\x71\xb6\xb7\x76\xa6\xed\x77\x6e\x78\x8a\x9e\x89\xf6\xed\x4c\xfa\x0e\xeb\x76\x03\x17\x77\x2c\xdc\x59\xfd\x1b\xec\xe3\xa6\xd4\x0a\x0d\xc5\x66\x8f\x24\x6a\x41\xe2\x41\xe6\xc5\x34\x6b\x2b\xc4\xc5\x95\xdb\x5b\xa2\xe9\xed\xe4\xf3\x74\x00\x2f\xed\x0a\x5d\xf6\x71\x68\x8f\x79\x25\x97\xf4\x15\x3a\xe6\x4d\xb3\xc3\x46\x0a\x93\x0d\xb2\x7f\xc8\xe2\xc4\x71\x84\x87\x8b\xf3\x37\x73\x40\xad\xc5\xea\xc7\x43\x6d\x43\x37\x05\xc8\x69\x91\x7e\xcb\xfb\x72\x01\x03\x8f\xf1\x35\x3a\x77\x24\x4d\x23\x1d\x8d\xb1\x6a\x2c\x71\x9a\x0a\x31\x22\x65\x27\x4c\x8b\x7b\x13\x4d\x6c\x25\x5e\x0a\x55\x20\x50\x3e\xdd\x74\xc8\x82\x32\x35\xd3\x75\xbc\x07\xe9\xf2\x53\x40\xfc\x8c\x95\xbd\x99\xc7\xdc\x0a\x5b\x65\x9a\xe6\xd0\xcc\xd8\x0d\x8d\xf0\xc4\xa7\x69\x2d\x5c\x3d\x23\x9c\xc3\xf5\xfd\x78\x0e\xd7\xca\xcc\x84\xd4\x07\x42\xb2\x6d\xab\xf1\x10\x2b\xf5\x02\x0b\x80\x27\x41\x61\x6a\x1e\x7f\x00\xd1\xe0\x0d\x29\xc2\xa3\x63\x11\xef\xb9\x3c\x08\x3c\xca\xb9\x88\x7b\xb4\xf2\x34\x0f\xee\xe0\x70\xf5\x20\x6e\x1c\xb4\xc1\xff\x28\x6c\x11\xfc\x91\x83\x63\x67\x04\xaa\x84\xbc\x9e\x09\x48\x8b\xcd\xbd\x60\xe3\x6e\x56\x48\x6b\xc4\x64\x5e\xbc\xbb\x8f\xc2\x73\x61\x1f\xc2\xc6\x8b\xea\xcc\x36\x7a\xb2\xc3\x31\x8e\x1d\x7e\x80\x2b\x1d\xfa\xd0\x1f\x6f\x5c\x74\x5e\x43\x7c\x39\xdc\x39\xfa\x10\x2f\x8a\xf3\x1d\x4d\xdf\x85\xa6\xd1\xf8\x81\x4b\x91\x3f\x32\x38\x0d\xce\x1c\xa7\x29\x89\xd6\x76\x7d\x54\xa6\x29\x38\x33\x96\x91\x31\xd3\x26\xa9\xf9\x50\xc3\xd1\xa6\x1e\x63\xb2\xcc\x2d\x22\x7f\xc7\xd7\xac\xee\xa4\x7c\xbb\xfb\x7a\x93\x1e\xbb\xf6\xba\x40\x38\x26\x85\x86\xb8\xa6\x08\x18\xd0\xf9\xe0\x89\xdb\xb2\x0b\xe6\x1a\x6b\xf8\xf5\xea\xea\x02\x1c\xfa\xc1\x1a\x8f\x49\x6a\x34\xa3\x48\x36\x94\x47\x0f\x49\x0b\xc7\x4c\xd0\x13\x57\x34\xdc\x8c\x10\x89\x1b\xc2\x5a\x51\x17\x39\xd8\xd8\x36\xa1\xaf\x12\x8b\xe4\x02\xb5\x8c\xa8\x1f\xf0\x86\xd0\x78\x65\xcd\x12\x3a\x14\x75\x1a\x47\x21\xbb\xa4\x55\xb6\x28\xba\xe7\xcc\x53\x1c\xc6\xc5\xa9\x73\x88\x7b\x0f\x55\xd0\xa3\xf7\xa2\x45\x9f\x15\x79\xf7\xbc\x5c\x14\x41\xc7\xb7\xc9\xb2\xee\xc5\x90\x71\x64\x28\xd3\xa6\xb4\xa7\x55\x09\x67\xc0\x1c\x0c\x3a\xe1\x3b\xd0\x8a\xd0\x09\x1d\x5f\x35\x84\x32\xaa\xbd\xe3\x60\xb7\xef\x78\xb7\x7c\x88\xa9\xe0\x75\x5c\x38\x6d\xc5\x76\xb9\x71\x63\xc6\xca\x7d\xfb\x68\xb6\x53\xe9\x11\x56\x42\x07\x84\x5e\x6c\xa0\x42\x10\x06\x44\xe5\xad\xe6\xb2\xb2\xcd\x46\x4f\xb3\xec\xd9\xd3\xa7\x7c\x7e\xfe\x7b\x12\xb3\xd2\xff\x4e\x16\xf1\x0e\x01\x0e\xb5\x20\xb5\xda\x96\xf7\x2c\x6a\x30\x38\xdb\x3a\xf4\x3e\xea\x70\x06\xbf\x5d\xbe\x7b\x7b\xbf\x45\xa6\x85\xc5\xdf\x0b\x00\xd8\x9a\x71\x5a\x90\xea\xb1\x7c\x11\xfb\x34\x9a\x96\xba\xdb\xae\x7f\x16\x7f\x44\x3f\x73\x93\x59\x86\x32\xd0\x2b\xad\x95\x47\x69\x4d\xed\xb3\xc5\xa2\xc8\xc7\xb7\xd9\x3c\x3d\x1c\xff\x1b\x00\x00\xff\xff\x5c\x61\xc0\x3a\x50\x16\x00\x00")

func htmlDocHtmlBytes() ([]byte, error) {
	return bindataRead(
		_htmlDocHtml,
		"html/doc.html",
	)
}

func htmlDocHtml() (*asset, error) {
	bytes, err := htmlDocHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "html/doc.html", size: 5712, mode: os.FileMode(420), modTime: time.Unix(1477443759, 0)}
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
	"html/default.css": htmlDefaultCss,
	"html/doc.html": htmlDocHtml,
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
	"html": &bintree{nil, map[string]*bintree{
		"default.css": &bintree{htmlDefaultCss, map[string]*bintree{}},
		"doc.html": &bintree{htmlDocHtml, map[string]*bintree{}},
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
