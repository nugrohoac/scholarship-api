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

var _mutationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\xc1\x6e\xdb\x30\x0c\xbd\xe7\x2b\xd4\x5b\xfb\x0b\xbe\x15\xe9\x0e\x3e\xac\x18\x92\xf5\x34\xec\xc0\x48\x9c\x43\xd4\xa6\x34\x92\x5a\x61\x0c\xfb\xf7\xc1\x6e\x16\x39\xb1\x16\xa0\xbe\xf9\xbd\x27\x3e\x92\x8f\x36\x26\x74\x9f\xb3\x81\x51\x64\xf7\x7b\xe3\x9c\x73\x3b\xec\x48\x0d\xe5\x45\x51\xee\x67\x64\xfa\x26\x65\xe3\xb2\xa2\x7c\x1d\x13\xde\x9d\x71\x1c\x80\xfa\xc6\xed\x4d\x88\xbb\x02\xa7\x63\x64\x7c\x8e\x15\x02\x54\xdf\xa2\x84\x4b\xe6\xe1\xbd\xf4\x66\xfe\x79\x49\x01\x0c\x2f\xed\x29\x34\xae\x65\x2b\x75\x18\x06\xac\xda\x5a\x6c\x1c\x71\xca\xd6\x0e\xd0\x2d\x1a\xf5\x71\x48\xc0\xe3\xf3\xe2\xdd\x82\xcb\x6c\x32\xb6\xd7\x26\x10\x82\xa0\x6a\xc5\x27\xaa\x41\xbf\x8d\xa1\xd2\x83\x07\x09\x6d\x40\x36\x32\x42\x6d\xdc\xb7\xb9\x9b\x6d\x41\xc7\xbb\xef\x45\x7d\x00\x7e\x5d\xf9\x4e\xe0\xa3\x9f\xbb\xaa\xed\x70\x49\xaf\xd6\x70\xb1\xca\x47\x6f\xf4\x6b\xb5\x4c\x8b\xaf\xc8\x37\x5e\xed\x50\xd1\xbe\x9c\x82\xba\xff\x48\x74\x5b\x41\x30\xdc\xfb\x63\xec\x41\xf4\x48\xa9\xbc\xd6\x14\x59\xa3\xac\x66\xad\x07\x09\xc3\x34\xdd\x95\x94\xa6\x44\xeb\xe9\xc2\x1b\x48\x40\xbc\xce\x2f\xa5\x9e\xfc\x7c\xdc\x7b\x03\xb1\x8a\x4f\x51\x7c\xe2\x50\xe1\x99\x63\x66\x8f\x03\xb2\x3d\x81\x55\x3a\xc5\x9e\x3a\x3a\x50\x4f\x36\x3e\xa1\x7a\xa1\x34\x15\x5b\xeb\x34\x1f\x94\xc2\x6d\x8d\xe0\xcf\x4c\xf2\x6e\x56\x74\xd3\x0d\x9d\x94\x8b\xc3\xf9\x91\x39\x10\x77\xff\x99\xeb\xc4\x56\x67\x5a\xb8\x9c\xcf\x73\x57\xb0\x7f\x26\x0f\x8d\xd3\x12\xe4\xe6\xcf\xdf\x00\x00\x00\xff\xff\x34\x81\xc8\xb8\x2c\x04\x00\x00")

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

	info := bindataFileInfo{name: "mutation.graphql", size: 1068, mode: os.FileMode(420), modTime: time.Unix(1644665343, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _queryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x90\x41\x4b\xc3\x40\x10\x85\xef\xf9\x15\xd3\x9b\xf9\x0b\x39\x56\xad\x04\x3c\x68\x0b\x5e\xc4\xc3\xba\x19\x93\xa1\xc9\x4c\x98\x99\x20\x41\xfc\xef\x92\x46\xb3\x16\x4a\x51\xe8\xde\xf6\x2d\xef\x7b\x6f\x9f\x8f\x3d\xc2\xe3\x80\x3a\xc2\x47\x06\x00\xb0\x41\x8f\xcd\x3a\xf0\xfe\xea\x70\x9d\x4e\x4b\x1d\x79\x01\x25\xfb\x22\xc5\x41\x4d\xb4\x80\x9d\x2b\x71\xbd\xc8\x1c\x3a\x3c\x12\xf3\x02\x5e\x03\xef\x37\x88\x55\x96\xf0\xd7\x32\xb0\xeb\x78\xa9\x84\x38\xe3\x52\xc8\xbd\xd4\xc4\x89\x8e\x5d\xa0\xf6\xc7\xb4\x5a\xe4\x3e\x98\xbd\x8b\x56\xc7\x2f\x79\x01\xed\x64\xdf\xa2\xf5\xc2\x86\x33\x71\x8b\x86\x5c\xdd\x4e\xa0\x27\x54\x7a\xa3\x18\x9c\xe4\x7c\x46\xbe\xf4\x9c\x7f\x2e\x5a\x8b\x3f\x7c\x87\xfe\xcb\x39\x6d\xb6\x8b\x8d\xb4\x41\xad\xa1\x3e\x79\x0f\x15\x45\xcb\xea\x4f\xdb\x9d\x58\xd9\x3c\xf8\x60\x05\x3c\x97\xec\x2f\x67\x47\xb6\x94\x9f\x86\xbe\x43\xff\xd5\x6b\x3d\x96\x37\xa9\x1b\xcd\xa5\x56\x27\x00\xd9\xe7\x57\x00\x00\x00\xff\xff\x7d\xec\xf3\x5b\x77\x02\x00\x00")

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

	info := bindataFileInfo{name: "query.graphql", size: 631, mode: os.FileMode(420), modTime: time.Unix(1644261771, 0)}
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

var _typeScholarshipGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x91\x41\x8a\xf3\x30\x0c\x85\xf7\x39\x85\xce\x91\x5d\xa1\xff\x0f\x5d\x77\x39\x74\xa1\xda\x9a\x54\x90\xc8\x1e\x49\x66\x28\x43\xef\x3e\x24\x2e\xd8\x69\x99\x55\xa2\xf7\x9e\x3e\xc9\xc8\xef\x99\xc0\xc2\x2d\xcd\xa8\x76\xe3\x0c\x3f\x03\x00\x00\xc7\x11\x4e\xe2\xdb\xbf\xe5\x24\x96\xf4\xf4\x2e\x8d\x50\x8c\x74\x53\x04\x17\x1a\xe1\xec\xca\x32\x6d\x02\x2e\xa9\x88\x77\x2d\x8e\x5e\xac\xd5\xbc\xe0\x44\x63\xfd\xd4\x86\x6f\xd4\x48\xd4\x12\xa1\xa8\x92\xf8\x21\xe7\x99\x03\xf6\x2c\xac\x92\x73\x92\xb3\xa3\xfa\x7e\x70\x33\xff\x49\xdc\x5b\x22\xa9\x48\xa0\x85\xc4\x8f\xe8\xfb\x85\x69\xe6\x89\xaf\x3c\xb3\xdf\x8f\x64\x41\x39\xaf\x88\x5d\xc4\xca\xd5\x38\xfe\x69\x2b\x7d\x15\xd6\x4a\x6f\x11\x1b\xe1\xa3\x86\x2e\x5b\xea\xb3\x48\x64\x99\xde\x17\x7f\x1a\xaf\x4b\x77\xd4\x15\xd5\x95\x95\x17\x94\xd0\x29\x1e\x1a\xec\x31\x0c\xaf\x67\xfd\x4f\x14\x9f\xa7\x0d\x45\xb7\xd3\xf5\xef\x6a\xc1\x75\x44\x57\x5e\x86\xc7\x6f\x00\x00\x00\xff\xff\x3f\x12\xc4\x4e\x22\x02\x00\x00")

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

	info := bindataFileInfo{name: "type/scholarship.graphql", size: 546, mode: os.FileMode(420), modTime: time.Unix(1644665722, 0)}
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
