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

var _mutationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x52\xc1\x6e\xdb\x30\x0c\xbd\xe7\x2b\xd4\x5b\xfb\x0b\xbe\x15\xe9\x0e\x39\xac\x28\x92\xf6\x34\xec\xc0\x58\xac\x43\xd4\xa6\x34\x92\x5a\x61\x0c\xfb\xf7\xc1\x4e\x5a\x29\xb1\x16\xa0\xbe\x99\xef\x51\xef\x91\x8f\x36\x46\x74\xdf\x93\x81\x51\x60\xf7\x67\xe5\x9c\x73\x5b\xec\x48\x0d\xe5\x45\x51\x6e\xe7\xca\xf4\x4d\xcc\xc6\x25\x45\x79\x1e\x23\xde\x7c\xd6\x71\x00\xea\x1b\xb7\x33\x21\xee\x72\x39\x1e\x02\xe3\x63\xa8\x00\xa0\xfa\x1e\xc4\x9f\x23\x77\xc7\xa7\x57\xf3\xcf\x4b\xf4\x60\x78\x2e\x4f\xbe\x71\x1b\xb6\xfc\x0e\xc3\x80\x55\x59\x0b\x8d\x23\x8e\xc9\x36\x03\x74\x85\xd1\x36\x0c\x11\x78\x7c\x2c\xfa\x0a\x2c\xb1\xc9\xb8\xb9\x14\x01\xef\x05\x55\x2b\x3a\x41\x0d\xfa\x75\xf0\x15\x0f\x2d\x88\xdf\x78\x64\x23\x23\xd4\xc6\xfd\x98\xdd\xac\x73\x75\xbc\xf9\x99\xd9\x7b\xe0\xb7\x85\xee\x54\xbc\x6f\x67\x57\xb5\x1d\x96\xf0\x62\x0d\x67\xab\xbc\x6f\x8d\x7e\x2f\x96\x69\xe1\x0d\xf9\x4a\xd7\x16\x15\xed\xe9\x14\xd4\xed\x57\xa2\x5b\x0b\x82\xe1\xae\x3d\x84\x1e\x44\x0f\x14\x73\xb7\xc6\xc0\x1a\x64\x31\x6b\x3d\x48\x18\xa6\xe9\x2e\xa8\x34\x25\x5a\xa6\x9b\xe9\xef\x20\x1e\xf1\x32\xbe\x18\x7b\x6a\xe7\xdb\xde\x19\x88\x55\x64\x32\xe3\x1b\xfb\x0a\xce\x1c\x12\xb7\x38\x20\xdb\x03\x58\xc5\x28\xf6\xd4\xd1\x9e\x7a\xb2\xf1\x01\xb5\x15\x8a\xd3\x63\x4b\x9e\xa6\xbd\x92\xbf\xce\x11\xfc\x95\x48\x8e\x62\x99\x37\x9d\xd0\x89\x59\xdc\xcd\x6b\x62\x4f\xdc\xfd\x67\xae\x13\x5a\x9d\xa9\x50\xf9\xbc\xce\x6d\xae\x7d\x88\xdc\x35\x4e\x73\x8e\xc7\x78\x77\x69\x3f\x90\x3d\x0b\xb0\xbe\xa2\x3c\xc1\x38\x35\x14\x11\x67\xfe\x22\x66\x3b\x35\xd5\xb7\x78\xf5\xa2\xab\xd1\x7f\x98\x8c\x47\x13\xab\xbf\xff\x02\x00\x00\xff\xff\xbf\x65\x3b\x97\xcb\x04\x00\x00")

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

	info := bindataFileInfo{name: "mutation.graphql", size: 1227, mode: os.FileMode(420), modTime: time.Unix(1645176865, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _queryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x90\x41\x4b\xc3\x40\x10\x85\xef\xfd\x15\xd3\x9b\xf9\x0b\x7b\xac\xb5\x12\xf0\xa0\x2d\x78\x29\x3d\xac\xc9\x98\x0c\x4d\x66\xc2\xcc\x04\x09\xe2\x7f\x97\x34\x9a\xb5\x50\x8a\x82\x7b\xdb\xb7\xbc\xf7\xbd\x7d\x3e\x74\x08\x4f\x3d\xea\x00\xef\x0b\x00\x80\x0d\x7a\x51\xaf\x22\x1f\x6f\x4e\xd7\xf1\x34\xd4\x92\x07\xc8\xd9\x67\xa9\xe8\xd5\x44\x03\xec\x5c\x89\xab\x59\xe6\xd8\xe2\x99\x98\x05\x78\x89\x7c\xdc\x20\x96\x8b\x14\x7f\x2b\x3d\xbb\x0e\xff\x45\x28\xa6\xb8\x04\x79\x90\x8a\x38\xa5\x63\x1b\xa9\xf9\x36\x2d\x67\xb9\x8b\x66\x6f\xa2\xe5\xf9\x4b\x16\xa0\x19\xed\x5b\xb4\x4e\xd8\x70\x4a\xdc\xa2\x21\x97\x77\x63\xd0\x33\x2a\xbd\x52\x11\x9d\xe4\x3a\x23\x9b\x7b\x4e\x3f\x17\xad\xc4\x1f\xbf\xa0\x7f\x72\x8e\x9b\xed\x8a\x5a\x9a\xa8\x56\x53\x97\xbc\xa7\x8a\xa2\x79\xf9\xab\xed\x2e\xac\x6c\x1e\xbd\xb7\x00\xfb\x9c\xfd\x70\x75\x64\x4b\xfc\x34\xf4\x3d\xfa\x8f\x5e\xab\x21\x5f\xa7\x6e\x34\x95\x5a\x5e\x08\x98\xcd\x6b\xac\x14\x31\xc0\x7e\x62\x1d\x16\x1f\x9f\x01\x00\x00\xff\xff\xeb\x87\x1f\xf2\x90\x02\x00\x00")

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

	info := bindataFileInfo{name: "query.graphql", size: 656, mode: os.FileMode(420), modTime: time.Unix(1645176669, 0)}
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

var _typePaymentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x54\x8b\x31\x0e\xc2\x30\x0c\x45\xf7\x9e\xc2\xe7\xc8\x86\xc4\xd2\x85\x05\x2e\xf0\x49\x4c\x6b\xd1\xb8\x91\x63\x86\x08\x71\x77\x44\x60\x20\x93\xfd\xdf\xfb\xdf\x5b\x61\x2a\x68\x99\xd5\xe9\x39\x11\x11\x49\x0a\x34\xab\xf7\xbf\xc6\x75\xdf\x60\x75\x95\x32\xff\xe1\x2b\xf4\x7e\x31\x68\xbd\xb1\x85\x21\x75\x9d\x18\x69\x13\xe5\x40\x67\x37\xd1\xa5\x43\xff\x35\x8e\xf0\x51\x7c\xe6\x87\x18\xf7\x87\xfa\x09\x79\x74\x92\xb1\x70\xf8\x9e\xe9\xf5\x0e\x00\x00\xff\xff\xee\xaf\xbf\x56\xad\x00\x00\x00")

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

	info := bindataFileInfo{name: "type/payment.graphql", size: 173, mode: os.FileMode(420), modTime: time.Unix(1645108896, 0)}
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
