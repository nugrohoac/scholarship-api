// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// mutation.graphql
// query.graphql
// schema.graphql
// type/bank.graphql
// type/card_identity.graphql
// type/country.graphql
// type/image.graphql
// type/requirement.graphql
// type/scholarship.graphql
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

var _mutationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xc1\x6e\xe2\x40\x0c\xbd\xf3\x15\xc3\x0d\x7e\x21\x37\xc4\xee\x4a\x39\x2c\x5a\xc1\x72\xaa\x7a\x30\x19\x37\x58\x24\x9e\xa9\xed\x14\x45\x55\xff\xbd\x4a\x68\x49\x20\x23\xa4\xe6\x96\xf7\x6c\x3f\xcf\x7b\x33\xd6\x46\x74\x7f\x1b\x03\xa3\xc0\xee\x7d\xe6\x9c\x73\x5b\x2c\x49\x0d\x65\xaf\x28\x8b\x1e\xe9\xbe\xae\x32\x73\x8d\xa2\xfc\x6f\x23\xce\xaf\x38\xd6\x40\x55\xe6\x76\x26\xc4\xe5\x00\xc7\x63\x60\xdc\x84\x04\x01\xaa\xe7\x20\xfe\x96\x59\x5e\x46\xcf\xfa\x9f\x7d\xf4\x60\x78\x2b\x4f\x3e\x73\x39\xdb\x30\x87\xa1\xc6\xa4\xac\x85\xcc\x11\xc7\xc6\xf2\x1a\xca\xd1\xa2\x45\xa8\x23\x70\xbb\x19\xf5\x8d\xb8\x86\x4d\xda\xfc\x5e\x04\xbc\x17\x54\x4d\xe8\x04\x35\xa8\xd6\xc1\x27\x76\x28\x40\x7c\xee\x91\x8d\x8c\x50\x33\xf7\xd4\x6f\xb3\x1e\xd0\x76\xfe\x3c\x54\x1f\x80\x4f\x13\xdd\x0e\x5c\x15\xfd\x56\x29\x0f\xc7\xf4\xc4\x86\x1b\x2b\x57\x85\xd1\xdb\xc4\x4c\x0b\x27\xe4\x07\x5d\x5b\x54\xb4\x7f\x5f\x41\x2d\x7e\x12\xdd\x5a\x10\x0c\x77\xc5\x31\x54\x20\x7a\xa4\x38\x74\x6b\x0c\xac\x41\x26\x67\x4d\x07\x09\x75\x77\xba\xbb\x52\xea\x12\x4d\xa7\x0b\x67\x10\x8f\x78\xd7\xe0\x11\x7c\x45\x9c\x98\x8f\x15\x95\x74\xa0\x8a\xac\xfd\x85\x5a\x08\xc5\xee\x05\x4c\xeb\xb4\x39\x28\xf9\xc7\x35\x2f\x0d\x7b\xe2\x72\x67\x20\x36\x65\xff\x5c\xd8\xdf\xec\xa7\x9c\xe0\x6b\x43\x82\x35\xb2\x5d\x2f\xca\x76\xc0\xbe\xef\xc9\x32\x73\x3a\x58\x3a\xfb\xf8\x0c\x00\x00\xff\xff\x43\xce\x37\xf4\xb6\x03\x00\x00")

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

	info := bindataFileInfo{name: "mutation.graphql", size: 950, mode: os.FileMode(420), modTime: time.Unix(1643464321, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _queryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x90\xc1\x4a\xc4\x40\x10\x44\xef\xfb\x15\xed\xcd\xfc\xc2\x1c\x15\x17\x02\x1e\x74\x17\xbc\x88\x87\x71\xd2\x26\xcd\x26\xdd\xa1\xbb\x83\x04\xf1\xdf\x25\x89\x9b\x21\x20\xa2\xb0\x73\x9b\x1a\xaa\x5e\x4d\xf9\xd8\x23\x3c\x0e\xa8\x23\x7c\xec\x00\x00\xf6\xe8\xa9\xb9\x89\x7c\xba\x9e\xaf\xd3\x69\xa9\x23\x0f\x50\xb2\xaf\x52\x1a\xd4\x44\x03\x1c\x5d\x89\xeb\x55\xe6\xd8\xe1\x46\x2c\x02\xbc\x46\x3e\xed\x11\xab\x5d\x8e\xbf\x95\x81\x5d\xc7\x4b\x11\xd2\x12\x97\x21\xf7\x52\x13\xe7\x74\xec\x22\xb5\x67\xd3\xd5\x2a\xf7\xd1\xec\x5d\xb4\xda\xbe\x14\x01\xda\xc9\x7e\x40\xeb\x85\x0d\x97\xc4\x03\x1a\x72\x75\x37\x05\x3d\xa1\xd2\x1b\xa5\xe8\x24\xbf\x33\x8a\xb5\xe7\xf2\x73\xd1\x5a\xfc\xe1\x1b\xfa\x2f\xe7\xb4\xd9\x31\x35\xd2\x46\xb5\x86\xfa\xec\x9d\x2b\x8a\x96\xd5\x9f\xb6\xfb\x61\x65\xf3\xe8\x83\x05\x78\x2e\xd9\x5f\xce\x6c\xcb\xa8\x79\xd3\xcf\xaf\x00\x00\x00\xff\xff\xd6\x87\x76\x59\x25\x02\x00\x00")

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

	info := bindataFileInfo{name: "query.graphql", size: 549, mode: os.FileMode(420), modTime: time.Unix(1643463746, 0)}
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

var _typeRequirementGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xce\x51\x0a\xc3\x20\x0c\xc6\xf1\x77\x4f\xf1\x79\x8d\xde\xa0\xaf\xdb\x2e\x20\x6d\x58\x03\x35\x75\x36\x19\x8c\xe1\xdd\x87\xee\x61\x0e\x5f\x44\xfe\x24\xfc\xa2\xaf\x44\xc8\xf4\x30\xce\x14\x49\x14\x6f\x07\x00\xbc\x4e\x98\x45\xdb\xff\x5c\xb6\x63\x0f\xf9\xdc\x38\xcd\x5d\xae\x9b\x13\xae\x9a\x59\xee\x2d\x48\x88\xff\xe1\x19\x76\xfb\x95\xe2\x1c\x89\xc5\x1e\xbb\x55\xfc\x0b\xae\xc7\x62\x35\xd5\x31\x96\x64\x8a\xf6\x5e\x86\xcb\x7a\xd6\x0f\xae\x1f\x61\xef\xca\x27\x00\x00\xff\xff\x5f\x57\x3c\x68\xe5\x00\x00\x00")

func typeRequirementGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeRequirementGraphql,
		"type/requirement.graphql",
	)
}

func typeRequirementGraphql() (*asset, error) {
	bytes, err := typeRequirementGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/requirement.graphql", size: 229, mode: os.FileMode(420), modTime: time.Unix(1643410454, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeScholarshipGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x90\xc1\x6a\x03\x31\x0c\x44\xef\xfe\x0a\x7d\x87\x6f\x85\x36\x90\x73\x8e\x25\x07\xc5\x56\x37\x02\xaf\xec\x4a\x32\x25\x94\xfd\xf7\xd2\xdd\xc2\x7a\x5b\x7a\xb2\xe7\xcd\x48\x36\xe3\x8f\x46\x60\xe9\x5e\x0b\xaa\xdd\xb9\xc1\x67\x00\x00\xe0\x1c\xe1\x2c\xbe\xde\xad\x55\xb1\xaa\xe7\xbf\x28\x42\x37\xd2\x95\x08\xce\x14\xe1\xe2\xca\x32\xad\x00\xe7\xda\xc5\x87\x11\x47\xef\xb6\x6b\x9e\x71\xa2\xb8\x1d\xdb\xc0\x07\x6a\x26\xda\x13\xa9\xab\x92\xf8\x53\x6b\x85\x13\x8e\xbb\x32\x61\x2e\x2c\xc7\x07\xa9\xf0\xc4\x37\x2e\xec\x8f\x67\xb2\xa4\xdc\x9c\xab\x1c\x22\xd6\x6f\xc6\xf9\x5f\xfb\xad\x4b\x66\x99\x2e\x8e\xea\x07\xe3\xb4\x19\x2f\x92\x0f\x58\xe9\xbd\xb3\xd2\x4c\xe2\x16\xe1\x75\x90\xd7\xb0\x84\xf0\xbb\xda\x13\x51\xfe\xa9\x37\x75\x5d\xeb\x1b\xff\xb6\x07\xbf\x97\x0d\xf2\x1a\x96\xaf\x00\x00\x00\xff\xff\x53\x5b\x78\x5e\xa6\x01\x00\x00")

func typeScholarshipGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeScholarshipGraphql,
		"type/scholarship.graphql",
	)
}

func typeScholarshipGraphql() (*asset, error) {
	bytes, err := typeScholarshipGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/scholarship.graphql", size: 422, mode: os.FileMode(420), modTime: time.Unix(1643463746, 0)}
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
	"type/requirement.graphql":   typeRequirementGraphql,
	"type/scholarship.graphql":   typeScholarshipGraphql,
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
		"requirement.graphql":   &bintree{typeRequirementGraphql, map[string]*bintree{}},
		"scholarship.graphql":   &bintree{typeScholarshipGraphql, map[string]*bintree{}},
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
