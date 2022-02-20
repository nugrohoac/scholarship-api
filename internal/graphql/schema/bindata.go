// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// mutation.graphql
// query.graphql
// schema.graphql
// type/bank.graphql
// type/bank_transfer.graphql
// type/card_identity.graphql
// type/country.graphql
// type/image.graphql
// type/major.graphql
// type/payment.graphql
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

var _mutationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x53\xc1\x6e\xdb\x30\x0c\xbd\xe7\x2b\xd4\x5b\xfb\x0b\xbe\x15\xe9\x0e\x39\xac\x28\x92\xf6\x34\xec\xc0\x58\xac\x43\xd4\xa6\x34\x92\x5a\x61\x0c\xfb\xf7\xc1\x4e\x5a\x29\xb1\x16\x6c\xb9\x85\x7c\xd4\x7b\xe4\x7b\xb6\x31\xa2\xfb\x9a\x0c\x8c\x02\xbb\x5f\x2b\xe7\x9c\xdb\x62\x47\x6a\x28\x2f\x8a\x72\x3b\x57\xa6\xdf\x84\x6c\x5c\x52\x94\xe7\x31\xe2\xcd\x67\x1d\x07\xa0\xbe\x71\x3b\x13\xe2\x2e\x97\xe3\x21\x30\x3e\x86\x4a\x03\x54\xdf\x83\xf8\xf3\xce\xdd\xf1\xe9\xd5\xfc\xe7\x25\x7a\x30\x3c\xa7\x27\xdf\xb8\x0d\x5b\x7e\x87\x61\xc0\x2a\xad\x85\xc6\x11\xc7\x64\x9b\x01\xba\x42\x68\x1b\x86\x08\x3c\x3e\x16\x73\x45\x2f\xb1\xc9\xb8\xb9\x24\x01\xef\x05\x55\x2b\x3c\x41\x0d\xfa\x75\xf0\x15\x0d\x2d\x88\xdf\x78\x64\x23\x23\xd4\xc6\x7d\x9b\xd5\xac\x73\x75\xbc\xf9\x9e\xd1\x7b\xe0\xb7\x05\xef\x54\xbc\x6f\x67\x55\xb5\x1b\x96\xed\xc5\x19\xce\x4e\x79\xdf\x1a\xfd\x5c\x1c\xd3\xc2\x1b\xf2\x95\xa9\x2d\x2a\xda\xd3\xc9\xa8\xdb\xff\xb1\x6e\x2d\x08\x86\xbb\xf6\x10\x7a\x10\x3d\x50\xcc\xd3\x1a\x03\x6b\x90\xc5\xae\x75\x23\x61\x98\xb6\xbb\x80\xd2\xe4\x68\xe9\x6e\x86\xbf\x83\x78\xc4\x4b\xfb\x62\xec\xa9\x9d\xb3\xbd\x33\x10\xab\xd0\x64\xc4\x17\xf6\x95\x3e\x73\x48\xdc\xe2\x80\x6c\x0f\x60\x15\xa1\xd8\x53\x47\x7b\xea\xc9\xc6\x07\xd4\x56\x28\x4e\x8f\x2d\x71\x9a\xf6\x4a\xfe\x3a\x46\xf0\x47\x22\x39\x92\x65\xdc\x14\xa1\x13\xb2\xc8\xcd\x6b\x62\x4f\xdc\xfd\x65\xaf\x53\xb7\xba\x53\xc1\xf2\x99\xce\x6d\xae\x7d\x90\xdc\x35\x4e\xb3\x8f\x47\x7b\x77\x69\x3f\x90\x3d\x0b\xb0\xbe\xa2\x3c\xc1\x38\x0d\x14\x16\x67\xfc\xc2\x66\x3b\x0d\xd5\xaf\x78\x35\xd1\xff\xf0\x45\x2c\x92\xf1\xb1\x43\x3c\x6a\x5c\xfd\xfe\x13\x00\x00\xff\xff\x2c\x3a\x3f\x65\xea\x04\x00\x00")

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

	info := bindataFileInfo{name: "mutation.graphql", size: 1258, mode: os.FileMode(420), modTime: time.Unix(1645179535, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _queryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x92\xc1\x4a\xf3\x50\x10\x85\xf7\x7d\x8a\xe9\xee\xcf\x2b\xdc\x65\xff\x5a\x09\x28\x68\x0b\x6e\x4a\x17\xd7\x64\x4c\xae\x4d\x66\xc2\xcc\x04\x09\xe2\xbb\x4b\x12\xbd\xd7\x42\x29\x0a\xcd\x2e\x67\x38\xdf\x39\x39\xc4\x86\x0e\xe1\xb1\x47\x19\xe0\x7d\x01\x00\xb0\x41\x2b\xea\x95\xa7\xe3\xbf\xe9\x75\x7c\x9a\xd0\x06\x73\x90\x93\x45\xa9\xe8\x45\x59\x1c\xec\x4c\x02\x55\x51\x26\xdf\xe2\x89\x98\x39\x78\xf6\x74\xdc\x20\x96\x8b\x84\xff\xcf\x3d\x99\x0c\xd7\x4a\x28\x66\x5c\x0a\xb9\xe3\x2a\x50\xa2\x63\xeb\x43\xf3\x6d\x5a\x46\xb9\xf3\xaa\x6f\x2c\xe5\xe9\x25\x73\xd0\x8c\xf6\x2d\x6a\xc7\xa4\x38\x13\xb7\xa8\x48\xe5\xcd\x08\x7a\x42\x09\x2f\xa1\xf0\x16\xf8\x72\x46\x16\x7b\xce\x5f\xce\x52\xb1\x3d\x7c\x85\xfe\xc9\x39\x6e\xb6\x2b\x6a\x6e\xbc\x68\x1d\xba\xe4\x9d\x2a\xb2\xe4\xe5\xaf\xb6\x3b\xb3\xb2\x9a\xb7\x5e\x1d\xec\x73\xb2\xc3\xc5\x91\x35\xe5\xa7\xa1\x6f\xd1\x7e\xf4\x5a\x0d\xf9\x3a\x75\x0b\x73\xa9\xe5\x19\x40\x34\xaf\xb1\x12\x44\x07\xfb\x39\xeb\x10\x0f\xf7\xfe\x95\xe5\x5a\xff\x47\x3b\xc2\xa6\xd2\x1f\x9f\x01\x00\x00\xff\xff\x36\xa1\xcc\x40\xef\x02\x00\x00")

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

	info := bindataFileInfo{name: "query.graphql", size: 751, mode: os.FileMode(420), modTime: time.Unix(1645266495, 0)}
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

	info := bindataFileInfo{name: "type/payment.graphql", size: 199, mode: os.FileMode(420), modTime: time.Unix(1645349145, 0)}
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
	"type/image.graphql":         typeImageGraphql,
	"type/major.graphql":         typeMajorGraphql,
	"type/payment.graphql":       typePaymentGraphql,
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
		"bank_transfer.graphql": &bintree{typeBank_transferGraphql, map[string]*bintree{}},
		"card_identity.graphql": &bintree{typeCard_identityGraphql, map[string]*bintree{}},
		"country.graphql":       &bintree{typeCountryGraphql, map[string]*bintree{}},
		"image.graphql":         &bintree{typeImageGraphql, map[string]*bintree{}},
		"major.graphql":         &bintree{typeMajorGraphql, map[string]*bintree{}},
		"payment.graphql":       &bintree{typePaymentGraphql, map[string]*bintree{}},
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
