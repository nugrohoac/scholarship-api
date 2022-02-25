// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// mutation.graphql
// query.graphql
// schema.graphql
// type/bank.graphql
// type/bank_transfer.graphql
// type/card_identity.graphql
// type/country.graphql
// type/degree.graphql
// type/image.graphql
// type/major.graphql
// type/payment.graphql
// type/requirement.graphql
// type/scholarship.graphql
// type/school.graphql
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

var _mutationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x54\x3d\x6f\xdb\x40\x0c\xdd\xfd\x2b\x2e\x9b\xf3\x17\xb4\x05\x4e\x07\x0f\x0d\x02\x3b\x99\x8a\x0e\xb4\x8e\x91\x89\x48\xbc\x2b\x49\x35\x10\x8a\xfe\xf7\x42\x1f\xf6\xc9\xd6\xd5\x6d\xbc\x99\x7c\xe4\x23\xf9\x9e\xce\xba\x88\xee\x6b\x6b\x60\x14\xd8\xfd\x5a\x39\xe7\xdc\x0e\x2b\x52\x43\x79\x55\x94\xf5\x10\xe9\x7f\x3d\xb2\x70\xad\xa2\xbc\x74\x11\xef\xce\x71\x6c\x80\xea\xc2\xed\x4d\x88\xab\x14\x8e\xc7\xc0\xf8\x14\x32\x09\x50\xfd\x08\xe2\x2f\x33\xf7\x63\xeb\xd5\xf0\xe7\x35\x7a\x30\xbc\xa4\x27\x5f\xb8\x2d\x5b\xea\xc3\xd0\x60\x96\xd6\x42\xe1\x88\x63\x6b\xdb\x06\x2a\x3c\x67\xca\xd0\x44\xe0\xee\x69\x56\x36\xcb\xb5\x6c\xd2\x6d\xaf\x39\xc0\x7b\x41\xd5\x0c\x4d\x50\x83\x7a\x13\x7c\x66\x84\x12\xc4\x6f\x3d\xb2\x91\x11\x6a\xe1\xbe\x0d\xc3\x6c\x52\xb4\xbb\xfb\x9e\xd0\x07\xe0\xf7\x05\x6f\x1f\x7c\x28\x87\xa9\x72\x27\x9c\xa7\x17\x57\xb8\xb8\xe4\x43\x69\xf4\x73\x71\x4b\x0b\xef\xc8\x37\xaa\x76\xa8\x68\xcf\x93\x4e\xeb\xcf\x28\xb7\x11\x04\xc3\x7d\x79\x0c\x35\x88\x1e\x29\xa6\x6a\x8d\x81\x35\xc8\x62\xd7\xbc\x8e\xd0\xf4\xdb\x5d\x41\xa9\x17\x34\x2b\x2e\x7c\x80\x78\xc4\x6b\xf9\x62\xac\xa9\x1c\xac\xbd\x37\x10\xcb\xd0\x24\xc4\x17\xf6\x99\x3c\x73\x68\xb9\xc4\x06\xd9\x1e\xc1\x32\x83\x62\x4d\x15\x1d\xa8\x26\xeb\x1e\x51\x4b\xa1\xd8\x37\x5b\xe2\xb4\x3d\x28\xf9\xdb\x18\xc1\x1f\x2d\xc9\x48\x96\x70\xbd\x85\x26\xe4\xcc\x37\x6f\x2d\x7b\xe2\xea\x2f\x7b\x4d\xd9\xec\x4e\x33\x96\xb3\x3b\x77\x29\x76\x22\xb9\x2f\x9c\x26\x1d\x47\x79\xf7\xed\xa1\x21\x7b\x11\x60\x7d\x43\x79\x86\xae\x2f\x98\x49\x9c\xf0\x0b\x99\x6d\x2a\xca\x5f\xf1\xa6\xa3\xff\xe3\x8b\x58\x38\xe3\xb4\x43\x1c\x67\xbc\x76\x67\xa8\xd7\xff\x30\xe0\xf8\xdc\x2d\xfc\x90\x7b\x10\xa6\x4b\x85\x7a\xf5\xfb\x4f\x00\x00\x00\xff\xff\x01\xe7\x4e\x05\x4f\x05\x00\x00")

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

	info := bindataFileInfo{name: "mutation.graphql", size: 1359, mode: os.FileMode(420), modTime: time.Unix(1645787207, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _queryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\x4d\x6a\xf3\x30\x10\xdd\xe7\x14\x93\xdd\xe7\x2b\x68\x99\x2f\x4d\x31\xb4\xd0\x26\xd0\x4d\xc8\x42\xb5\xa6\xb6\x1a\x59\x63\x46\x63\x8a\x29\xbd\x7b\x91\xdd\x5a\x4e\x31\xa1\x05\x6b\x65\x3f\xf1\x7e\xe6\x31\x92\xae\x41\x78\x6c\x91\x3b\x78\x5f\x01\x00\xec\x50\x8a\x6a\xa3\xfd\xf9\x5f\xff\x1b\x8f\xb3\xb5\x15\x05\xb9\x97\x11\x2a\x5a\x0e\xc4\x0a\x0e\xc2\xd6\x97\x23\xec\x75\x8d\x17\x60\xa6\xe0\x59\xfb\xf3\x0e\xd1\xac\x92\xfc\x7f\x6a\xbd\x70\xb7\x94\x43\x31\xc8\x25\x93\x3b\x2a\xad\x4f\xea\x58\x6b\xeb\xbe\x49\xeb\x11\x6e\x74\x08\x6f\xc4\xe6\xf2\x26\x53\xe0\x22\x7d\x8f\xa1\x21\x1f\x70\x50\xdc\x63\x40\x6f\x6e\xa2\xd0\x13\xb2\x7d\xb1\x85\x16\x4b\xd7\x3d\xb2\x31\xe7\x30\x39\x71\x49\xf2\xf0\x65\xfa\x27\x66\xec\xec\x50\x54\xe4\x34\x87\xca\x36\x89\xdb\x47\x24\xce\xcd\xaf\xba\x9b\x69\x39\x88\x96\x36\x28\x38\xe6\x5e\x4e\x57\x4b\x0e\xc9\x3f\x15\x7d\x8b\x32\xc9\xb5\xe9\xf2\x6d\xca\x66\x87\x50\xeb\x19\x81\xc9\x58\x5b\x2c\x19\x51\xc1\xd1\xf4\x1f\xa7\xc9\xd5\xbd\x7e\x25\x5e\x6a\x47\xea\x28\xf6\x63\x0d\x63\x74\x72\x8b\x38\xc4\x13\x9f\xd2\x5c\x6b\xe4\x7a\xdf\x8f\xcf\x00\x00\x00\xff\xff\x12\xd1\xee\x0f\x6b\x03\x00\x00")

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

	info := bindataFileInfo{name: "query.graphql", size: 875, mode: os.FileMode(420), modTime: time.Unix(1645798043, 0)}
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

var _typeBank_transferGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x48\x4a\xcc\xcb\x0e\x29\x4a\xcc\x2b\x4e\x4b\x2d\x52\xa8\xe6\x52\x50\x50\x50\xc8\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x07\x0b\x24\x26\x27\xe7\x97\xe6\x95\xf8\xe1\x12\xcf\xb7\x52\xf0\xcc\x2b\x01\x0b\x65\xe6\x26\xa6\xa7\x5a\x41\x28\xae\x5a\x40\x00\x00\x00\xff\xff\x01\x6f\xbd\xff\x62\x00\x00\x00")

func typeBank_transferGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeBank_transferGraphql,
		"type/bank_transfer.graphql",
	)
}

func typeBank_transferGraphql() (*asset, error) {
	bytes, err := typeBank_transferGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/bank_transfer.graphql", size: 98, mode: os.FileMode(420), modTime: time.Unix(1645108896, 0)}
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

var _typeDegreeGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x48\x49\x4d\x2f\x4a\x4d\x55\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\xcc\x2b\x01\xb3\xf3\x12\x73\x53\xad\x14\x82\x4b\x8a\x32\xf3\xd2\xb9\x6a\x01\x01\x00\x00\xff\xff\x0e\xf6\x29\x1d\x2c\x00\x00\x00")

func typeDegreeGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeDegreeGraphql,
		"type/degree.graphql",
	)
}

func typeDegreeGraphql() (*asset, error) {
	bytes, err := typeDegreeGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/degree.graphql", size: 44, mode: os.FileMode(420), modTime: time.Unix(1645633111, 0)}
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

var _typeMajorGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\xc8\x4d\xcc\xca\x2f\x52\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\xcc\x2b\x01\xb3\xf3\x12\x73\x53\xad\x14\x82\x4b\x8a\x32\xf3\xd2\xb9\x6a\xb9\xb8\x10\x8a\xdd\x52\x53\x53\xa0\x1a\x92\x4b\x8b\x8a\xf3\x8b\xe0\xca\x40\x42\x60\x25\xc5\x56\x0a\xd1\x60\x46\x2c\x57\x2d\x20\x00\x00\xff\xff\xad\x54\xf4\xa3\x66\x00\x00\x00")

func typeMajorGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeMajorGraphql,
		"type/major.graphql",
	)
}

func typeMajorGraphql() (*asset, error) {
	bytes, err := typeMajorGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/major.graphql", size: 102, mode: os.FileMode(420), modTime: time.Unix(1645266495, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typePaymentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\xcb\x31\x0e\xc2\x30\x0c\x05\xd0\xbd\xa7\xf0\x39\xb2\x21\xb1\x74\x61\x81\x0b\x7c\x12\xd3\x5a\x34\x4e\x94\x98\x21\x42\xdc\x1d\x11\x18\xc8\xd0\xc9\xf6\x7f\xfe\xd6\x32\x53\x46\x8b\xac\x46\xcf\x89\x88\x48\x82\xa3\x59\xad\xef\xd5\xaf\x69\x43\xa9\xab\xe4\xf9\x2f\xbe\x42\xef\x97\x02\xad\x37\x2e\x6e\xb8\x3a\x07\x46\xd8\x44\xd9\xd1\xd9\x8a\xe8\xd2\x43\xfb\x7d\x1c\x61\x23\x7c\xea\x07\xef\xd3\x43\xed\x84\xb8\x6f\x69\x10\x89\x58\xd8\x7d\xc7\xf4\x7a\x07\x00\x00\xff\xff\x71\x28\x90\x61\xc7\x00\x00\x00")

func typePaymentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typePaymentGraphql,
		"type/payment.graphql",
	)
}

func typePaymentGraphql() (*asset, error) {
	bytes, err := typePaymentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/payment.graphql", size: 199, mode: os.FileMode(420), modTime: time.Unix(1645350005, 0)}
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

var _typeScholarshipGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x51\xcd\x6a\xc3\x30\x0c\xbe\xe7\x29\xfc\x1c\xbe\x15\xba\x41\xcf\x3d\x8e\x1e\x54\x5b\x4b\x05\x89\xec\x49\x32\x23\x8c\xbe\xfb\x48\xdc\x61\xa7\x65\xa7\xe8\xfb\xd1\x27\xc5\xb2\x25\xa3\xd3\x70\x4b\x13\x88\xde\x28\xbb\x9f\xc1\x39\xe7\x28\x7a\x77\x62\xdb\x6a\xcd\x89\x35\xc9\xe9\x95\xf2\xae\x28\xca\xc6\x30\xcc\xe8\xdd\xd9\x84\x78\xdc\x08\x98\x53\x61\xeb\x5a\x0c\xac\x68\xc3\x34\xc3\x88\xbe\x7e\x6a\xc3\x37\x48\x44\x6c\x8e\x50\x44\x90\xed\x90\xf3\x44\x01\xfa\x2c\xa8\x94\x51\xe2\xb3\x81\xd8\x7e\x70\x13\xdf\x38\xee\x25\xe6\x54\x38\xe0\x8c\x6c\x47\xb0\xfd\xc2\x38\xd1\x48\x57\x9a\xc8\x96\x23\x6a\x10\xca\x6b\xc4\xce\xa2\xe5\xaa\x14\xff\x95\x05\xbf\x0a\x49\x4d\x6f\x16\xf5\xee\xa3\x9a\x2e\x9b\xeb\xb3\x70\x24\x1e\x5f\x17\x7f\x08\xcf\x4b\x77\xa9\x6b\x54\x07\x6b\x5e\x86\x65\x05\xfe\xaf\xa8\x6f\x27\x08\x86\xf1\xd0\x26\xdc\x87\xe1\xf9\xd6\xef\x88\xf1\x71\xef\x50\x64\xbb\x67\xff\xb3\xcd\xb8\xce\xed\xe0\x65\xb8\xff\x06\x00\x00\xff\xff\x5d\x3f\x48\x05\x37\x02\x00\x00")

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

	info := bindataFileInfo{name: "type/scholarship.graphql", size: 567, mode: os.FileMode(420), modTime: time.Unix(1645108896, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeSchoolGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x28\x4e\xce\xc8\xcf\xcf\x51\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\xcc\x2b\x01\xb3\xf3\x12\x73\x53\xad\x14\x82\x4b\x8a\x32\xf3\xd2\xc1\x02\x20\xe5\x28\x02\x89\x29\x29\x45\xa9\xc5\xc5\x28\x62\xc5\x25\x89\x25\xa5\xc5\x10\x53\x6a\xb9\xb8\x90\xac\x70\x4b\x4d\x4d\x81\x5a\x93\x5c\x5a\x54\x9c\x5f\x84\xaa\x0f\xac\xa6\xd8\x4a\x21\x1a\xc2\x8a\xe5\xaa\x05\x04\x00\x00\xff\xff\x08\xdd\xa9\x2c\x9f\x00\x00\x00")

func typeSchoolGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeSchoolGraphql,
		"type/school.graphql",
	)
}

func typeSchoolGraphql() (*asset, error) {
	bytes, err := typeSchoolGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/school.graphql", size: 159, mode: os.FileMode(420), modTime: time.Unix(1645609664, 0)}
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
	"type/bank_transfer.graphql": typeBank_transferGraphql,
	"type/card_identity.graphql": typeCard_identityGraphql,
	"type/country.graphql":       typeCountryGraphql,
	"type/degree.graphql":        typeDegreeGraphql,
	"type/image.graphql":         typeImageGraphql,
	"type/major.graphql":         typeMajorGraphql,
	"type/payment.graphql":       typePaymentGraphql,
	"type/requirement.graphql":   typeRequirementGraphql,
	"type/scholarship.graphql":   typeScholarshipGraphql,
	"type/school.graphql":        typeSchoolGraphql,
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
		"bank_transfer.graphql": &bintree{typeBank_transferGraphql, map[string]*bintree{}},
		"card_identity.graphql": &bintree{typeCard_identityGraphql, map[string]*bintree{}},
		"country.graphql":       &bintree{typeCountryGraphql, map[string]*bintree{}},
		"degree.graphql":        &bintree{typeDegreeGraphql, map[string]*bintree{}},
		"image.graphql":         &bintree{typeImageGraphql, map[string]*bintree{}},
		"major.graphql":         &bintree{typeMajorGraphql, map[string]*bintree{}},
		"payment.graphql":       &bintree{typePaymentGraphql, map[string]*bintree{}},
		"requirement.graphql":   &bintree{typeRequirementGraphql, map[string]*bintree{}},
		"scholarship.graphql":   &bintree{typeScholarshipGraphql, map[string]*bintree{}},
		"school.graphql":        &bintree{typeSchoolGraphql, map[string]*bintree{}},
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
