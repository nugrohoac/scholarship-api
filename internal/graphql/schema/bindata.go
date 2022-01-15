// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// mutation.graphql
// query.graphql
// schema.graphql
// type/bank.graphql
// type/country.graphql
// type/image.graphql
// type/user.graphql
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

var _mutationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\xf0\x2d\x2d\x49\x2c\xc9\xcc\xcf\x53\xa8\xe6\x52\x50\x50\x50\x08\x4a\x4d\xcf\x2c\x2e\x49\x2d\x0a\x2d\x4e\x2d\xd2\x00\x8b\x80\x00\x48\xa5\x95\x42\x69\x71\x6a\x51\x48\x65\x41\xaa\x22\x5c\x3c\x35\x37\x31\x33\xc7\x4a\x21\xb8\xa4\x28\x33\x2f\x1d\x21\x5c\x90\x91\x9f\x97\xea\x97\x8f\x45\x22\xb1\xb8\xb8\x3c\xbf\x28\x05\x55\x46\x13\x62\x34\x57\x2d\x20\x00\x00\xff\xff\xe5\xec\xd6\x1e\x91\x00\x00\x00")

func mutationGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_mutationGraphql,
		"mutation.graphql",
	)
}

func mutationGraphql() (*asset, error) {
	bytes, err := mutationGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "mutation.graphql", size: 145, mode: os.FileMode(420), modTime: time.Unix(1641989113, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _queryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x8d\x41\xca\xc2\x30\x10\x85\xf7\x39\xc5\xfc\xbb\xbf\x57\xc8\x52\xa1\x20\xb8\x51\x4f\x10\xd3\xa1\x86\x26\x93\x30\x99\x20\x41\xbc\xbb\x54\x31\xa1\x7b\x67\x37\xdf\xe3\x7d\x4f\x6a\x42\x38\x15\xe4\x0a\x0f\x05\x00\x30\xa2\xd8\xdb\xce\xd0\xf2\xff\x7e\xd7\xf3\x2e\x38\xd1\x70\x20\x69\xc8\x16\xce\x91\x35\x5c\x84\x1d\xcd\x0d\x93\x09\xb8\x81\x83\x86\xab\xa1\x65\x44\x9c\x54\xd7\xef\x63\x21\xe1\xfa\xab\x05\xfb\xd1\xf5\x91\x63\x9c\x1d\x75\x3b\x06\xe3\xfc\xb7\xf4\xd7\x70\x32\x39\xdf\x23\x4f\xdb\x64\xd0\xe0\xd7\xfa\x19\x73\x8a\x94\x51\x3d\x5f\x01\x00\x00\xff\xff\x8a\x14\x49\x7c\x23\x01\x00\x00")

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

	info := bindataFileInfo{name: "query.graphql", size: 291, mode: os.FileMode(420), modTime: time.Unix(1642269551, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _schemaGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\x4e\xce\x48\xcd\x4d\x54\xa8\xe6\x52\x50\x50\x50\x28\x2c\x4d\x2d\xaa\xb4\x52\x08\x04\x51\x60\x81\xdc\xd2\x92\xc4\x92\xcc\xfc\x3c\x2b\x05\x5f\x28\x8b\xab\x16\x10\x00\x00\xff\xff\x8e\x43\x79\x00\x32\x00\x00\x00")

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

	info := bindataFileInfo{name: "schema.graphql", size: 50, mode: os.FileMode(420), modTime: time.Unix(1641981243, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeBankGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x48\x4a\xcc\xcb\x56\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\xcc\x2b\x01\xb3\xf3\x12\x73\x53\xad\x14\x82\x4b\x8a\x32\xf3\xd2\xc1\x02\xc9\xf9\x29\x08\x81\x5a\x2e\x2e\xb8\x66\xb7\xd4\xd4\x14\xa8\x01\xc9\xa5\x45\xc5\xf9\x45\x28\xda\x40\x2a\x8a\xad\x14\xa2\x41\x74\x2c\x57\x2d\x20\x00\x00\xff\xff\x55\x04\xf7\xe5\x73\x00\x00\x00")

func typeBankGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeBankGraphql,
		"type/bank.graphql",
	)
}

func typeBankGraphql() (*asset, error) {
	bytes, err := typeBankGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/bank.graphql", size: 115, mode: os.FileMode(420), modTime: time.Unix(1641850924, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeCountryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x48\xce\x2f\xcd\x2b\x29\xaa\x54\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\xcc\x2b\x01\xb3\xf3\x12\x73\x53\xad\x14\x82\x4b\x8a\x32\xf3\xd2\xb9\x6a\xb9\xb8\x90\x95\xbb\xa5\xa6\xa6\x40\xb5\x24\x97\x16\x15\xe7\x17\xc1\x15\x82\x85\xc0\x8a\x32\x53\x8b\xad\x14\xa2\xa1\x1a\x62\xb9\x6a\x01\x01\x00\x00\xff\xff\xb5\xa0\xc5\x2d\x6f\x00\x00\x00")

func typeCountryGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeCountryGraphql,
		"type/country.graphql",
	)
}

func typeCountryGraphql() (*asset, error) {
	bytes, err := typeCountryGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/country.graphql", size: 111, mode: os.FileMode(420), modTime: time.Unix(1641984426, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeImageGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\xc8\xcc\x4d\x4c\x4f\x55\xa8\xe6\x52\x50\x50\x50\x28\x2d\xca\xb1\x52\x08\x2e\x29\xca\xcc\x4b\x07\xf3\xcb\x33\x53\x4a\x32\xac\x14\x3c\xf3\x4a\xc0\xdc\x8c\xd4\xcc\xf4\x8c\x12\x04\x3f\x37\x33\x37\x15\x45\x7d\x72\x62\x41\x49\x66\x7e\x1e\x5c\xac\x96\x8b\x2b\x33\xaf\xa0\xb4\x44\x01\x4c\x7a\x62\xb7\x49\x11\xcd\x2a\x45\x74\xbb\x14\x89\xb4\x0c\x10\x00\x00\xff\xff\x5b\xe8\x7b\x35\xcf\x00\x00\x00")

func typeImageGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeImageGraphql,
		"type/image.graphql",
	)
}

func typeImageGraphql() (*asset, error) {
	bytes, err := typeImageGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/image.graphql", size: 207, mode: os.FileMode(420), modTime: time.Unix(1641981243, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeUserGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x3d\x92\xc3\x20\x0c\x85\x7b\x4e\x41\xc6\xb7\x70\xb7\xb3\x55\x9a\x14\xfb\x73\x00\xd6\x68\x1d\x26\x46\x62\x40\x14\x9e\x1d\xdf\x7d\xc7\x40\x62\x83\x5d\x85\x3c\x3d\x21\xf9\x7b\xf0\xec\x40\xc6\x00\x5e\xfe\x09\x29\xa5\x34\xba\x97\x57\xe4\x74\x46\x65\xa1\x97\x9f\xec\x0d\x8e\x49\x58\xcd\x7d\x72\x7f\xcd\x0e\x92\x04\x56\x99\xa9\x32\xb9\x3b\x21\xdc\xa8\xd5\x98\x7a\x69\xac\x1a\x73\xdb\x40\xd6\x29\x9c\x6f\xed\x84\xc0\x8a\x63\xd8\x56\x18\x28\x22\xfb\xf9\xba\xdb\xca\x51\x60\x35\xbd\x93\xae\x3b\x95\xd6\x1e\x42\xa8\xb4\x11\x50\x83\xef\xcb\x6f\xde\x97\xef\x68\x86\xca\xf5\xa3\xf0\xb1\x1f\xb0\xfe\x7f\x1b\xd2\xe4\xe6\x33\xf6\x95\xfd\xe6\x8b\x10\x89\xe3\x44\xa3\xc1\x0f\x08\x8e\x30\x40\x01\xca\xf4\x00\xac\x6e\x59\xf9\x65\x8a\x6b\x63\x97\x3a\x0d\xba\xc8\xdf\x39\x86\xae\x65\x7f\xc9\x52\x4d\xbf\x88\x15\xff\xa2\x35\x09\x6c\xea\x2b\x83\x22\x9d\xa5\xd0\x1d\xb1\x3f\x2f\x38\x80\x2f\x85\x06\x7d\x51\x1b\xf8\xdd\x09\xfd\xae\xc1\x7f\xd9\x94\x43\x00\x27\xb5\x1a\xd1\x22\x04\x60\xb4\x2f\x40\x85\x7f\x0a\x83\x7c\x79\x5d\x51\x03\xb2\x78\x5a\xf3\x6a\xc5\x68\xd5\x94\xdf\xe6\x2f\xa4\xe3\xf2\x1f\x00\x00\xff\xff\x9a\x95\x68\x09\x1b\x03\x00\x00")

func typeUserGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeUserGraphql,
		"type/user.graphql",
	)
}

func typeUserGraphql() (*asset, error) {
	bytes, err := typeUserGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/user.graphql", size: 795, mode: os.FileMode(420), modTime: time.Unix(1642269551, 0)}
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
	"mutation.graphql":     mutationGraphql,
	"query.graphql":        queryGraphql,
	"schema.graphql":       schemaGraphql,
	"type/bank.graphql":    typeBankGraphql,
	"type/country.graphql": typeCountryGraphql,
	"type/image.graphql":   typeImageGraphql,
	"type/user.graphql":    typeUserGraphql,
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
	"mutation.graphql": &bintree{mutationGraphql, map[string]*bintree{}},
	"query.graphql":    &bintree{queryGraphql, map[string]*bintree{}},
	"schema.graphql":   &bintree{schemaGraphql, map[string]*bintree{}},
	"type": &bintree{nil, map[string]*bintree{
		"bank.graphql":    &bintree{typeBankGraphql, map[string]*bintree{}},
		"country.graphql": &bintree{typeCountryGraphql, map[string]*bintree{}},
		"image.graphql":   &bintree{typeImageGraphql, map[string]*bintree{}},
		"user.graphql":    &bintree{typeUserGraphql, map[string]*bintree{}},
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
