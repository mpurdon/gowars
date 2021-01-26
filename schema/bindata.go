// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// query.graphql
// schema.graphql
// types/film.graphql
// types/time.graphql
package schema

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

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _queryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xd0\xc1\x6a\xe3\x30\x18\x04\xe0\xbb\x9f\x62\x42\x2e\x59\x58\xf2\x00\xb9\xed\x65\xa1\x87\xd2\x94\x94\x5e\x4a\x0e\xaa\x33\x89\x04\xb2\x24\x7e\xfd\x4d\x31\xa5\xef\x5e\x2c\xdb\x90\x90\x9a\xf8\xaa\x7f\x3c\xf3\xe1\x25\x5e\x2c\xf1\xfc\x41\x69\xa1\x6d\x22\x84\x49\x98\x19\x34\xc3\x78\x8f\x78\x84\x5a\x82\x41\xa5\x45\x8a\xae\x7b\x77\x41\x63\x79\xfd\xb7\x7d\x58\x57\xe5\xab\xbe\xe0\xab\x02\x80\x25\x76\x34\x52\x5b\x1c\xa3\xc0\xe0\xe8\x7c\x83\xf7\x16\x4e\x33\xd4\xa9\xe7\x5f\x44\xc1\x89\x5a\xfa\xbb\x6b\xc6\xa7\x65\x40\x88\x48\x46\x4c\x43\xa5\x64\x18\x21\x92\xc4\xb3\x3b\xf0\xb0\x2e\xbd\x25\xba\x2a\x15\x1b\xec\x54\x5c\x38\xfd\xd9\xe0\xed\xbf\xf3\xcd\x62\x5f\x2d\x7f\x99\x4e\x94\x1c\x43\x37\xae\x96\x4e\x10\x4c\x73\xbd\x5e\x5b\x23\xa6\x2e\x7b\x77\x09\x65\x20\x31\x26\xcf\x55\x57\x74\x69\xd8\x96\xa1\x29\x85\x37\x81\x3a\xfe\x82\x1b\x43\x7f\x9e\x0d\xe8\xd3\xb7\x82\xf2\x3e\x21\xc8\x89\xb5\x63\x9e\x24\x8c\xf7\x79\x84\x21\x7d\x43\xd8\xf5\xef\x53\x06\x35\x92\xad\x4b\x97\x88\xce\xd0\xc4\x03\xfd\xb5\x66\x48\xce\xf6\x8c\xf9\x22\x7a\x92\xc7\xae\xf1\x0a\x36\x04\x26\x64\x67\x5a\x57\x7b\xde\x87\x0d\xc1\xb9\xae\x31\x3e\xc5\x7a\xed\xef\x8b\x7d\xf5\x5d\xfd\x04\x00\x00\xff\xff\x54\x57\x89\x6c\x87\x03\x00\x00")

func queryGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_queryGraphql,
		"query.graphql",
	)
}

func queryGraphql() (*asset, error) {
	bytes, err := queryGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "query.graphql", size: 903, mode: os.FileMode(420), modTime: time.Unix(1611554624, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x50\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x5c\xb5\x5c\x80\x00\x00\x00\xff\xff\x9e\xeb\xeb\x5e\x1c\x00\x00\x00")

func schemaGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_schemaGraphql,
		"schema.graphql",
	)
}

func schemaGraphql() (*asset, error) {
	bytes, err := schemaGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "schema.graphql", size: 28, mode: os.FileMode(420), modTime: time.Unix(1611552459, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typesFilmGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xa4\xd1\xb1\x6e\xdb\x30\x10\x06\xe0\x5d\x4f\xf1\x1b\x19\xba\x65\xf1\x54\x6d\x46\x82\x00\x59\x8a\xa2\x09\xda\xc1\xf0\x70\xa1\xce\xd2\x01\x12\xc9\x1e\x4f\x0d\x82\xa2\xef\x5e\x48\x24\x53\x37\x35\x92\x02\xd9\x8e\xe6\x7f\x9f\x8f\xba\xe6\x02\x3b\xdc\x19\x29\xbe\x91\x26\x1c\x65\x9c\x2e\x1b\x7b\x8a\x8c\x1b\x19\x27\xfc\x6c\x00\x60\xc9\xcc\x5e\xbe\xcf\x0c\xe9\xd8\x9b\x1c\x85\xf5\x72\xbd\x92\xae\xc5\xed\xf5\xa6\xc4\xee\x07\x06\x47\x49\xa1\x63\xf8\x79\x7a\x60\x45\x38\xc2\x06\xa9\xf2\x12\x2b\x81\x16\xb7\xde\x4e\x1b\x43\x64\x2f\xbe\x47\x24\xa5\x5e\x29\x0e\x09\x64\xb0\x81\xf1\xc0\xbd\xf8\xf5\xee\x5f\xad\x74\x5d\x29\x3d\x8e\x2d\xee\x4c\xc5\xf7\xa7\xaa\xa7\x89\xff\x6a\xfb\x90\xd0\x89\xb2\xb3\x50\x9e\x50\x4f\x9f\x68\xe2\x97\xc2\x0e\xa3\x24\x5b\xfa\x17\x27\x65\x88\xab\x13\x35\x74\xb3\x63\x4d\x19\xaa\xc7\x05\x4a\x2d\xf6\x85\x3a\x9c\x8e\xf3\xe5\xe6\x6a\xbb\xdd\x7e\x44\x47\xc6\x38\x06\x9d\xc8\x4e\x51\x28\x8f\x4c\x89\x21\x7e\xfd\x2d\x68\x2f\x9e\x46\x38\x65\xb2\xa0\x70\x61\xf6\xa6\x4f\xf9\xff\x4a\xf6\x9a\x8c\x5b\xdc\xcb\xc4\x9b\xe6\xe2\xe5\xd8\x29\xb2\x13\x4e\xb0\x81\x0c\xa4\x05\x7e\xfe\x82\x6b\xbe\x64\x96\x89\x73\xb5\x39\x9c\x81\x8c\x34\x0d\x12\xdf\xa0\x6a\x6a\x7d\x7e\xae\xcf\x69\x3f\x78\x10\x37\xbe\x31\x57\x0d\xb5\xd8\x7f\xcd\xe5\x39\xca\x0d\xa4\xe4\x8c\xf5\x75\xec\x4f\xac\xc5\xfe\x33\x6b\x0a\xfe\x9c\x16\x47\xf2\x6c\xaf\x53\x25\xb3\x38\x6b\xb5\x39\xfc\xdf\x82\x4d\x26\xce\xf0\x2a\x2a\xa7\x30\xab\x63\x3c\x52\xca\x0b\xe6\x2e\x2f\xb6\x1c\x76\x56\xd7\xfa\x6e\x9e\x3b\x79\xd6\x73\x5d\xf1\xe6\x57\xf3\x3b\x00\x00\xff\xff\xd7\x7e\x5c\x4d\x04\x04\x00\x00")

func typesFilmGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typesFilmGraphql,
		"types/film.graphql",
	)
}

func typesFilmGraphql() (*asset, error) {
	bytes, err := typesFilmGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "types/film.graphql", size: 1028, mode: os.FileMode(420), modTime: time.Unix(1611554624, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typesTimeGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x52\x56\x08\xc9\xcc\x4d\x55\xc8\x2c\x56\x48\xcc\x53\x08\x72\x73\x36\x36\x36\xb6\x54\x28\xc9\xcc\x4d\x2d\x2e\x49\xcc\x2d\xd0\xe3\x2a\x4e\x4e\xcc\x49\x2c\x02\x2b\xe2\x02\x04\x00\x00\xff\xff\x3c\x59\x30\xf9\x2c\x00\x00\x00")

func typesTimeGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typesTimeGraphql,
		"types/time.graphql",
	)
}

func typesTimeGraphql() (*asset, error) {
	bytes, err := typesTimeGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "types/time.graphql", size: 44, mode: os.FileMode(420), modTime: time.Unix(1611552604, 0)}
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
	"query.graphql":      queryGraphql,
	"schema.graphql":     schemaGraphql,
	"types/film.graphql": typesFilmGraphql,
	"types/time.graphql": typesTimeGraphql,
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
	"query.graphql":  &bintree{queryGraphql, map[string]*bintree{}},
	"schema.graphql": &bintree{schemaGraphql, map[string]*bintree{}},
	"types": &bintree{nil, map[string]*bintree{
		"film.graphql": &bintree{typesFilmGraphql, map[string]*bintree{}},
		"time.graphql": &bintree{typesTimeGraphql, map[string]*bintree{}},
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
