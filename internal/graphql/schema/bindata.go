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

var _mutationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xc1\x6e\xea\x40\x0c\xbc\xf3\x15\xcb\x0d\x7e\x21\x37\xc4\x7b\x87\x1c\x1e\x7a\x82\x72\xaa\x7a\x30\x59\x37\x58\x24\xde\xad\xed\x14\x45\x55\xff\xbd\x4a\x4a\x49\x20\x2b\xa4\xe6\x96\xf1\xd8\x63\xcf\x24\xd6\x46\x74\xff\x1a\x03\xa3\xc0\xee\x63\xe6\x9c\x73\x5b\x2c\x49\x0d\x65\xaf\x28\x8b\x1e\xe9\x9e\x8e\x99\xb9\x46\x51\x9e\xda\x88\xf3\x2b\x8e\x35\x50\x95\xb9\x9d\x09\x71\x39\xc0\xf1\x18\x18\x37\x21\x51\x00\xd5\x73\x10\x7f\x5b\x59\x7e\x8f\x9e\xf5\x2f\xfb\xe8\xc1\xf0\x56\x9e\x7c\xe6\x72\xb6\x61\x0e\x43\x8d\x49\x59\x0b\x99\x23\x8e\x8d\xe5\x35\x94\xa3\x45\x8b\x50\x47\xe0\x76\x33\xea\x1b\xd5\x1a\x36\x69\xf3\x7b\x11\xf0\x5e\x50\x35\xa1\x13\xd4\xa0\x5a\x07\x9f\xd8\xa1\x00\xf1\xb9\x47\x36\x32\x42\xcd\xdc\x73\xbf\xcd\x7a\x40\xdb\xf9\xcb\xc0\x3e\x00\x9f\x26\xba\x1d\xb8\x2a\xfa\xad\x52\x1e\x8e\xcb\x13\x1b\x6e\xac\x5c\x15\x46\xef\x13\x33\x2d\x9c\x90\x1f\x74\x6d\x51\xd1\xfe\x5f\x82\x5a\xfc\x26\xba\xb5\x20\x18\xee\x8a\x63\xa8\x40\xf4\x48\x71\xe8\xd6\x18\x58\x83\x4c\x6e\x4d\x07\x09\x75\x77\xdd\x1d\x95\xba\x44\xd3\xe9\xc2\x19\xc4\x23\xde\x35\x78\x04\x5f\x11\x27\xe6\x63\x45\x25\x1d\xa8\x22\x6b\xff\xa0\x16\x42\xb1\xfb\x03\xa6\x3c\x6d\x0e\x4a\xfe\x31\x47\xf0\xad\x21\xc1\x1a\xd9\x46\xbc\x2e\xf9\x0b\x73\x14\xf7\x6b\xc3\x9e\xb8\xdc\x19\x88\x4d\x27\x5d\xaa\x7f\xd9\x3f\x54\xb9\x7e\x54\xdb\x01\xfb\x11\x59\x66\x4e\x07\xfb\x67\x9f\x5f\x01\x00\x00\xff\xff\x61\xfb\x14\x42\xe2\x03\x00\x00")

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

	info := bindataFileInfo{name: "mutation.graphql", size: 994, mode: os.FileMode(420), modTime: time.Unix(1643592155, 0)}
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

var _typeRequirementGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8e\xc1\x0d\xc2\x30\x0c\x45\xef\x99\xe2\x67\x8d\x6e\xd0\x2b\xb0\x40\xd4\x5a\xd4\x52\xe3\x86\xd4\x46\x42\xa8\xbb\xa3\xa4\x07\x02\xb9\x44\xd1\x93\xed\xf7\xf4\x95\x08\x99\x1e\xc6\x99\x22\x89\xe2\xed\x00\x80\xe7\x01\xa3\x68\xfd\xef\xd3\xb2\xad\x21\xef\x0b\xa7\xb1\xc1\x65\x73\xc0\x55\x33\xcb\xbd\x02\x09\xf1\x17\x3c\xc3\x6a\x5f\x72\x38\x47\x62\xb1\x95\xdd\x8a\xfc\x14\xce\xdb\x64\x05\x95\x31\x96\x64\x8a\xfa\x5e\xba\xb2\x53\xfb\x77\xc4\x77\x01\xbe\x2f\xf0\xee\xf8\x04\x00\x00\xff\xff\x22\xdd\x74\x31\xee\x00\x00\x00")

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

	info := bindataFileInfo{name: "type/requirement.graphql", size: 238, mode: os.FileMode(420), modTime: time.Unix(1643464907, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeScholarshipGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\x41\x6b\xc3\x30\x0c\x85\xef\xfe\x15\xfa\x1d\xbe\x0d\xb6\x41\xcf\x3d\x8e\x1c\x5c\x5b\x4b\x05\x8e\xec\x49\x32\xa3\x8c\xfe\xf7\x51\xa7\x10\x67\x23\xa7\xe4\x3d\x7d\xef\x19\x5b\x76\xab\x08\x1a\xaf\x25\x07\xd1\x2b\x55\xf8\x71\x00\x00\x94\x3c\x9c\xd8\xfa\xbf\xd6\xc2\x5a\xe4\xf4\xdf\xf2\xd0\x14\xa5\x3b\x1c\x16\xf4\x70\x36\x21\x9e\xbb\x11\x96\xd2\xd8\x86\x88\x05\x6b\xba\x69\x5a\xc2\x8c\x7e\xfd\xac\x81\xef\x20\x09\x71\x23\x62\x13\x41\xb6\x97\x5a\x33\xc5\x30\x76\x25\x0c\x29\x13\xef\x0f\xc4\x4c\x33\x5d\x28\x93\xdd\x5e\x51\xa3\x50\x35\x2a\xbc\x43\xb4\x5d\x94\xd2\xe1\x58\xf0\xab\x91\xe0\x82\x6c\x03\xa2\x1e\x3e\x56\x68\xea\xd4\x67\xe3\x44\x3c\x9f\x2d\x88\xed\xe2\xcf\xc1\x1b\xa7\xa3\xd6\x47\xd5\x20\x27\x77\x77\xee\xef\x02\xde\x11\xd3\x73\x09\xb1\x49\x7f\xe4\xf1\x06\x1b\xf8\x28\x1b\xe4\xe4\xee\xbf\x01\x00\x00\xff\xff\xa2\xbe\x55\xe4\xcc\x01\x00\x00")

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

	info := bindataFileInfo{name: "type/scholarship.graphql", size: 460, mode: os.FileMode(420), modTime: time.Unix(1643592155, 0)}
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
