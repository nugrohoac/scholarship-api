// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// mutation.graphql
// query.graphql
// schema.graphql
// type/bank.graphql
// type/card_identity.graphql
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

var _mutationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x90\x31\x4f\xc4\x30\x0c\x85\xf7\xfb\x15\xbe\x0d\xfe\x42\xb6\xd3\x4d\x1d\xb8\x01\xb8\x09\x31\x98\xc6\x2a\xd6\xb5\x76\x14\xbb\xa0\x08\xf1\xdf\x51\x8b\x44\x5b\x9a\x6c\x79\xcf\xf6\xf7\xf4\xbc\x24\x82\x87\xd1\xd1\x59\x05\xbe\x0e\x00\x00\x8f\xd4\xb1\x39\xe5\xab\x51\xbe\x9b\x95\xe9\x4d\x93\x01\x46\xa3\xfc\x5c\x12\x1d\xff\x74\x1a\x90\xfb\x00\x4f\x9e\x59\xba\x45\x4e\xef\x2a\x74\xd1\x8a\x81\x66\x9f\x9a\xe3\xd6\xb9\xff\x3d\x7d\x98\x3f\xd7\x14\xd1\x69\x8b\xe7\x18\xa0\x11\x5f\xee\x08\x0e\x54\xc5\xba\x06\x60\x49\xa3\x37\x03\x76\xab\xa0\xad\x0e\x09\xa5\x5c\x56\x7b\x2b\x6f\x14\xcf\xa5\xf9\x0f\xc1\x18\x33\x99\x55\x38\x6a\x8e\xfd\x59\x63\x25\x43\x8b\x39\x36\x91\xc4\xd9\x99\x2c\xc0\xcb\x9c\xe6\xbc\xa8\xe5\xf8\xba\x4c\xbf\xa1\xdc\x76\xdc\x49\x3c\xb5\x73\xaa\x5a\x87\x6b\x7b\x57\xc3\xa6\xca\x53\xeb\xfc\xb1\x2b\xd3\xf5\x46\x52\xdf\xfa\xfe\x09\x00\x00\xff\xff\x32\x12\xe6\x94\x12\x02\x00\x00")

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

	info := bindataFileInfo{name: "mutation.graphql", size: 530, mode: os.FileMode(420), modTime: time.Unix(1642828500, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _queryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x8f\xc1\x8a\xc2\x40\x0c\x86\xef\x7d\x8a\xec\x6d\xfb\x0a\x73\xdc\x65\x0b\x0b\x7b\x58\x2b\x78\x1f\xa7\xb1\x86\xb6\x49\xc9\xa4\x48\x11\xdf\x5d\x6a\x75\x4a\x2f\xe2\xc1\xb9\xcd\x1f\xfe\xef\x4b\x6c\xec\x11\x36\x03\xea\x08\xe7\x0c\x00\xa0\x40\x0b\xc7\x2f\xcf\xcd\xe7\xed\x3b\xbd\x96\x3a\x32\x07\xbf\x6c\x29\x0a\x83\x46\x51\x07\x5b\x53\xe2\x3a\xc5\xec\x3b\x5c\x85\xb9\x83\xbd\xe7\xa6\x40\xac\xb2\x05\xff\x2d\x03\x9b\x8e\xef\x32\x84\x19\xb7\x48\xfe\xa4\x26\x5e\xe8\xd8\x79\x6a\x1f\xa5\x8f\x14\xf7\x3e\xc6\x93\x68\xb5\x9e\xe4\x0e\xda\xa9\x5e\x62\xec\x85\x23\xce\xc4\x12\x23\x72\xf5\x33\x81\x76\xa8\x74\xa0\xe0\x8d\xe4\xb9\x23\x4f\x7b\xce\x97\x8b\xd6\x62\xff\x77\xe9\x8b\xcd\xcb\x35\x00\x00\xff\xff\x05\x95\xbd\xc5\xa0\x01\x00\x00")

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

	info := bindataFileInfo{name: "query.graphql", size: 416, mode: os.FileMode(420), modTime: time.Unix(1642877868, 0)}
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

var _typeCard_identityGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xce\x41\x0a\x03\x21\x0c\x05\xd0\xbd\xa7\xf8\x5e\xc3\x6d\x57\xae\xdb\x0b\x0c\x35\x0c\xa1\x4c\x1a\x6c\xa4\x48\x99\xbb\x17\x75\x16\x9d\xba\xd1\xf8\x91\xf7\x63\x55\x09\xf7\x25\xa7\x98\x48\x8c\xad\xe2\xe3\x00\x80\x53\x40\x14\xeb\x73\xfb\x13\x70\xb5\xcc\xb2\xf6\x40\x9e\xa7\x27\x6f\xcb\x4a\x61\x5c\x3d\x28\x2f\xca\xf1\x00\x76\xe7\x58\xb4\x18\xfa\x79\x99\x9b\x86\xfe\xbb\xc2\xad\x2a\xf9\xbf\x22\x7f\x6a\x6a\x54\x6c\xb3\x6f\x3e\x49\xd9\x26\xe0\xd0\x1f\xa6\x43\xd2\xb7\xba\xfd\x1b\x00\x00\xff\xff\xef\x69\xfb\x48\xee\x00\x00\x00")

func typeCard_identityGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeCard_identityGraphql,
		"type/card_identity.graphql",
	)
}

func typeCard_identityGraphql() (*asset, error) {
	bytes, err := typeCard_identityGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/card_identity.graphql", size: 238, mode: os.FileMode(420), modTime: time.Unix(1642350756, 0)}
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

var _typeUserGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x90\xbd\x6e\xc3\x30\x0c\x84\xf7\x3c\x05\x9f\x43\x5b\xd1\xc9\x4b\x86\xb6\x5b\xd1\x81\xb5\x58\x47\x88\x45\x0a\x12\x3d\x08\x45\xde\xbd\xd0\x0f\x52\xc9\x93\xe9\xf3\x1d\x7d\xfc\x34\x07\x82\x23\x51\x84\xdf\x0b\x00\x80\xb3\x06\x16\xd6\x3a\x33\x7a\x32\xf0\xae\xd1\xf1\x56\x85\x62\x36\xd5\xfd\x91\x03\x55\x89\x3c\xba\x7d\x32\x85\x9b\x30\x5d\xe5\xac\xa9\x18\x70\x1e\xb7\x16\x5b\xc5\x07\xe4\x7c\x3d\xff\x21\x29\xea\x91\xfe\x2b\xac\x72\xb0\xc6\xbc\x0c\xad\x82\x24\xc5\xfd\x55\xec\x9c\x44\x6b\x23\xa5\x34\x69\x1b\xb1\xa5\x68\xfa\xb3\xf5\xd5\x1b\xbb\x75\x72\xad\x18\xed\x62\x89\xd5\xa9\xa3\x64\xe0\x73\x10\xf2\x57\xb5\x7c\x23\xdf\xc7\x0e\xe5\xfd\x65\xad\xe5\x4e\x97\x8e\x5f\xc6\xe3\x1e\x97\x4b\x45\xbd\xcb\xe6\xf8\x8d\x52\x10\x4e\xd4\x99\xab\xdc\x89\xa7\x2d\x05\x71\x03\x5d\x82\xc4\x87\x7f\x52\xef\x99\xba\x40\x62\x87\x76\x94\xb6\x4f\x6b\x3b\xb7\x1b\x3d\xee\x0d\xf9\x0f\xd5\xf1\xf1\x17\x00\x00\xff\xff\xf2\x19\xc5\xc0\xf2\x01\x00\x00")

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

	info := bindataFileInfo{name: "type/user.graphql", size: 498, mode: os.FileMode(420), modTime: time.Unix(1642350756, 0)}
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
	"mutation.graphql":           mutationGraphql,
	"query.graphql":              queryGraphql,
	"schema.graphql":             schemaGraphql,
	"type/bank.graphql":          typeBankGraphql,
	"type/card_identity.graphql": typeCard_identityGraphql,
	"type/country.graphql":       typeCountryGraphql,
	"type/image.graphql":         typeImageGraphql,
	"type/user.graphql":          typeUserGraphql,
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
		"bank.graphql":          &bintree{typeBankGraphql, map[string]*bintree{}},
		"card_identity.graphql": &bintree{typeCard_identityGraphql, map[string]*bintree{}},
		"country.graphql":       &bintree{typeCountryGraphql, map[string]*bintree{}},
		"image.graphql":         &bintree{typeImageGraphql, map[string]*bintree{}},
		"user.graphql":          &bintree{typeUserGraphql, map[string]*bintree{}},
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
