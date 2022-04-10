// Code generated for package schema by go-bindata DO NOT EDIT. (@generated)
// sources:
// mutation.graphql
// query.graphql
// schema.graphql
// type/applicant.graphql
// type/assessment.graphql
// type/bank.graphql
// type/bank_transfer.graphql
// type/card_identity.graphql
// type/country.graphql
// type/degree.graphql
// type/ethnic.graphql
// type/image.graphql
// type/major.graphql
// type/payment.graphql
// type/requirement.graphql
// type/scholarship.graphql
// type/school.graphql
// type/sponsor.graphql
// type/student.graphql
// type/user.graphql
// type/user_document.graphql
// type/user_school.graphql
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

var _mutationGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x55\xc1\x6e\xdb\x3c\x0c\xbe\xf7\x29\xdc\x5b\xfb\x0a\xbe\x05\x49\xf0\x23\xc0\xbf\xa2\x48\xda\xc3\x30\xf4\xc0\x48\xac\x23\xd4\xa6\x34\x91\x5a\x61\x0c\x7b\xf7\x41\xb6\x6b\x39\x91\x9a\x62\xb9\x04\x26\x3f\x8a\xd4\xc7\x8f\x94\xf4\x0e\xab\x6f\x41\x40\x8c\xa5\xea\xf7\x4d\x55\x55\xd5\x1e\x1b\xc3\x82\xfe\x99\xd1\xdf\x0d\x96\xf8\x8b\xc8\xba\x0a\x8c\xfe\xa9\x77\x78\x3b\xdb\xb1\x03\xd3\xd6\xd5\x41\xbc\xa1\x26\x99\xdd\xc9\x12\x3e\xd8\x82\x03\x98\xdf\xad\xd7\xe7\x9e\xfb\xf1\xe8\x9b\xe1\xe3\xd9\x69\x10\x3c\x4f\x6f\x74\x5d\xed\x48\xd2\x39\x04\x1d\x16\xd3\x8a\xad\x2b\x43\x2e\xc8\xae\x83\x06\x67\x8f\xb2\x9d\x03\xea\x1f\x16\x61\x0b\x5f\x20\xf1\xfd\xee\x32\x07\x68\xed\x91\xb9\x90\xc6\xb2\x40\xbb\xb6\xba\x50\x82\x02\xaf\x77\x1a\x49\x8c\x18\xe4\xba\xfa\x31\x14\xb3\x4e\xd6\xfe\xf6\x25\xa1\x8f\x40\x6f\x59\xde\x68\x5c\xa9\xa1\xaa\x12\x85\x4b\x77\x91\x05\x94\x13\x19\x35\xd1\xb0\x1d\x3e\x66\x5f\x83\xa4\xd1\xd7\xd3\x7f\xec\x65\x3a\xd7\x78\x39\x6d\x40\x0a\x27\x0e\xae\xc7\x16\x14\x5e\xe9\xdb\x4a\x89\xf9\x95\x75\x4e\xec\x1b\xd2\x95\xa8\x3d\x32\xca\xe3\xa4\x8a\xbb\x7f\xd1\xc9\xda\x23\x08\x1e\xd4\xc9\xb6\xe0\xf9\x64\x5c\x8a\x66\x67\x89\xad\xcf\x98\x2d\xab\x06\xba\xc8\xe5\x05\xd4\x44\xf9\x14\xa5\x04\xef\xe0\x35\xe2\xa5\x58\x9c\x6b\x8d\x1a\x06\xe9\x20\xe0\xa5\x90\x26\x21\xb6\xa4\x0b\x7e\x22\x1b\x48\x61\x87\x24\xe5\x36\x60\x6b\x1a\x73\x34\xad\x91\x7e\x83\xac\xbc\x71\xf1\xb0\x1c\xc7\xe1\xc8\x46\x5f\xc7\x78\xfc\x19\x8c\x1f\x93\x25\x5c\x14\xec\x84\x5c\xa8\xf4\x35\x90\x36\xd4\x7c\x72\xaf\xc9\x5b\xbc\xd3\x22\xcb\x3c\x0b\xfb\x64\xfb\x48\x72\x5f\x57\x9c\xfa\x38\xb6\xf7\x10\x8e\x9d\x91\x27\x0f\xc4\xaf\xe8\x1f\xa1\x8f\x01\x8b\x16\x27\x7c\xd6\x66\x99\x82\x3e\x11\xf3\x57\xf3\xf3\xc5\xfc\x65\xca\xf8\xb8\x83\x1b\x6b\xbc\x54\xa7\x6d\xef\xbe\x10\xe0\xb8\x5c\x33\x3d\x94\xd6\xcf\xc4\x94\x6d\x27\x92\x50\x82\xdb\xea\x30\xca\x2a\xe5\x89\x53\x92\xd1\xd2\x80\xfb\x8e\xe0\xf7\x08\x9c\x14\xb1\x5c\x5c\x88\xfe\x3f\x0b\x85\x75\xce\x12\x74\xbf\x1e\xd7\xe4\x88\x18\x6e\x1f\x27\x7d\x61\xbd\xc0\x6f\x90\xc5\x10\x94\xe5\x17\x0b\x1c\xc9\x99\x75\xf1\x3c\x9b\x96\xda\x8b\xc0\x8d\x55\xe1\x4c\x42\x23\xed\x2f\xa5\x25\xe4\x5c\xdb\x17\x77\x42\x91\x92\x6b\x2a\x42\x66\xe8\x33\x9a\x3c\x2a\xdb\x75\x48\x7a\xb8\xd8\xff\x28\x12\xd7\xe9\xc5\x9a\xb8\x9f\xc3\x16\x52\x5e\x31\x23\xf3\xb9\x8a\xa7\x9d\x40\x92\x3f\x3e\x1f\x9e\xed\x3c\xf4\x8b\xf7\x64\x95\x7b\xfb\x97\x3c\xf6\xa0\xac\x2f\x04\x0d\xe6\x05\x7d\x53\xad\x7f\xfe\x06\x00\x00\xff\xff\x6c\x55\x67\x85\x0f\x08\x00\x00")

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

	info := bindataFileInfo{name: "mutation.graphql", size: 2063, mode: os.FileMode(420), modTime: time.Unix(1649558190, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _queryGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xac\x53\x41\x6e\xc2\x30\x10\xbc\xf3\x0a\x73\x23\x5f\xf0\xad\x14\xa8\x22\xb5\x52\x0b\x55\x2f\x88\x83\xeb\x2c\x89\x4b\xe2\x8d\xec\x8d\xda\xa8\xea\xdf\x2b\xe3\x12\x87\x10\x50\x41\xc9\xc9\x59\x6b\x67\x66\xc7\xb3\x54\x97\xc0\x5e\x2a\x30\x35\xfb\x1e\x31\xc6\xd8\x02\x48\x66\x53\xa1\x77\x93\xfd\xaf\xfb\x72\x55\x28\xe2\x2c\xd6\xd4\x94\x64\x65\x2c\x1a\xce\x56\x64\x94\x4e\x9b\xb2\x16\x05\x1c\x15\x23\xce\xde\x85\xde\x2d\x00\x92\x51\x80\xbf\xc7\x4a\x93\xa9\x87\x62\x90\x1e\x2e\x90\x3c\x62\xaa\x74\x40\x87\x42\xa8\xfc\xd0\x34\x6e\xca\xa5\xb0\xf6\x13\x4d\x72\x7c\x13\x71\x96\xbb\xf6\x25\xd8\x12\xb5\x05\x8f\xb8\x04\x0b\x3a\x99\x3b\xa0\x37\x30\x6a\xab\xa4\x20\x85\x97\x39\xa2\x46\xa7\x9f\x1c\x4d\x8a\xf4\xfc\x47\x7a\x55\xa7\xf3\x6c\x25\x33\xcc\x85\xb1\x99\x2a\x43\xef\x5e\x22\x9a\x38\xf9\x97\x77\x3d\x2e\x5b\x12\x54\x59\xce\xd6\xb1\xa6\xcd\x45\x93\x6d\xe0\x0f\x46\x3f\x00\xb5\x74\x4d\xeb\x78\x16\xb4\x29\x2f\x6a\xdc\x03\xd0\x1a\x6b\x06\xa9\x01\xe0\x6c\x9d\xec\x0f\x9b\xd6\xd5\x93\xf8\x40\x33\x54\x46\x0a\x07\xd6\x89\xa1\x93\x8e\xf9\x20\x0c\xee\x73\xab\xd4\xe7\x1a\xe6\x5d\x5e\xff\x6a\xb7\x10\x5b\x10\x46\x66\xaf\xf0\x45\x27\x4c\x1e\xb4\x43\x35\xa7\x4c\x2b\xc9\xd9\xda\x1f\xda\xf6\xde\x95\x65\xae\xa4\xd0\x74\x2e\x4e\x61\x53\x5a\x4f\x77\x72\x77\x85\xf6\x93\xac\x45\x9c\x89\x83\x8a\xa3\x50\x35\xda\x2e\x47\xaa\x69\x6e\x9b\x4b\x55\x02\xed\xa1\x06\x31\xd7\x83\xf6\xe4\xe7\x10\x7d\x21\x77\xb8\xdd\x2a\x09\x93\x5b\xb6\xf0\x3c\x71\x67\xeb\x7e\x7e\x03\x00\x00\xff\xff\x6e\x4d\xe1\x5a\xb0\x05\x00\x00")

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

	info := bindataFileInfo{name: "query.graphql", size: 1456, mode: os.FileMode(420), modTime: time.Unix(1649565722, 0)}
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

	info := bindataFileInfo{name: "schema.graphql", size: 50, mode: os.FileMode(420), modTime: time.Unix(1649397699, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeApplicantGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x8f\xbd\x0a\xc3\x30\x0c\x84\x77\x3f\x85\x9e\xc3\x73\x29\x04\xba\x75\x2c\x1d\x84\x23\x12\x43\xfc\x83\x24\x17\x42\xc9\xbb\x97\xb8\x26\xa9\xbb\x18\x9d\x38\x7f\xa7\xd3\x35\x13\x60\xce\x8b\x77\x18\x15\xde\x06\x00\xc0\x8f\x16\x86\xa8\x75\x2e\x42\x3c\xfc\x69\x5b\xdf\x2a\xc5\xcd\x69\x41\x96\xd9\xe7\x5f\x97\x28\x6a\x91\x53\xef\x09\xeb\x05\x95\x2c\xdc\x95\x7d\x9c\xea\x96\x44\x70\xed\x36\x4c\x2e\x85\x40\x71\x44\xf5\x29\xde\x48\x75\x4f\xf3\x01\x27\x6a\x71\x89\x49\x2c\x3c\xea\xf0\x34\x9b\x31\x7d\x83\x2b\xd1\xd8\x5a\xb8\xc2\x92\xb8\xc3\x1f\xb6\x1d\x71\x88\x13\x53\xa9\xed\x7b\xc4\xd0\x5f\xfb\xc2\xa5\xd0\xb7\xd2\xf6\x09\x00\x00\xff\xff\x18\x9e\x62\x4a\x37\x01\x00\x00")

func typeApplicantGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeApplicantGraphql,
		"type/applicant.graphql",
	)
}

func typeApplicantGraphql() (*asset, error) {
	bytes, err := typeApplicantGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/applicant.graphql", size: 311, mode: os.FileMode(420), modTime: time.Unix(1649558190, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeAssessmentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8e\xc1\x0a\xc2\x30\x0c\x86\xef\x79\x8a\xec\x35\x76\x53\xf0\xb0\x8b\x17\x1f\x40\x62\x17\x46\xa0\x4d\x6b\x9b\x0a\x43\xf6\xee\xb2\x30\x10\xc1\x4b\x48\x3e\xf2\x7f\xfc\xa2\xa5\x1b\xfa\x3c\x95\x12\x25\x90\xda\x25\xca\x22\x0f\x89\x62\x2b\xbe\x01\x11\xb1\xf2\xb3\x4b\xe5\xc4\x6a\xd3\x3c\xe2\xa4\x36\x38\x7f\x51\xec\x3c\xe2\x39\xe7\xc8\xa4\x03\x6c\x00\xff\x84\xb7\x90\x2b\x1f\x2a\xa5\xc4\x23\xb6\x9d\x5c\x29\xf1\x8f\xc7\xbd\x1b\x00\x6b\x4f\xdf\x97\x23\x48\x21\x70\x31\xd2\xc0\x7e\x72\x6b\xb4\xfa\xd6\x7a\x29\xb9\x9a\xe8\x72\x9f\x73\xe8\x7b\x49\xe7\x96\x8d\x22\x6c\x9f\x00\x00\x00\xff\xff\x72\xa6\x57\x60\xe2\x00\x00\x00")

func typeAssessmentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeAssessmentGraphql,
		"type/assessment.graphql",
	)
}

func typeAssessmentGraphql() (*asset, error) {
	bytes, err := typeAssessmentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/assessment.graphql", size: 226, mode: os.FileMode(420), modTime: time.Unix(1649558190, 0)}
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

	info := bindataFileInfo{name: "type/bank.graphql", size: 115, mode: os.FileMode(420), modTime: time.Unix(1649397699, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeBank_transferGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x48\x4a\xcc\xcb\x0e\x29\x4a\xcc\x2b\x4e\x4b\x2d\x52\xa8\xe6\x52\x50\x50\x50\xc8\x4b\xcc\x4d\xb5\x52\x08\x2e\x29\xca\xcc\x4b\x07\x0b\x24\x26\x27\xe7\x97\xe6\x95\xf8\xe1\x12\xcf\x47\x11\xcd\xcc\x4d\x4c\x4f\xb5\x82\x50\x5c\xb5\x80\x00\x00\x00\xff\xff\x98\x39\xe2\x52\x65\x00\x00\x00")

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

	info := bindataFileInfo{name: "type/bank_transfer.graphql", size: 101, mode: os.FileMode(420), modTime: time.Unix(1649397699, 0)}
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

	info := bindataFileInfo{name: "type/card_identity.graphql", size: 238, mode: os.FileMode(420), modTime: time.Unix(1649397699, 0)}
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

	info := bindataFileInfo{name: "type/country.graphql", size: 111, mode: os.FileMode(420), modTime: time.Unix(1649397699, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeDegreeGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x48\x49\x4d\x2f\x4a\x4d\x55\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\xcc\x2b\x01\xb3\xf3\x12\x73\x53\xad\x14\x82\x4b\x8a\x32\xf3\xd2\xc1\x02\x45\x89\x79\xd9\x10\xe9\x5a\x40\x00\x00\x00\xff\xff\x45\xbe\x50\xd7\x3a\x00\x00\x00")

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

	info := bindataFileInfo{name: "type/degree.graphql", size: 58, mode: os.FileMode(420), modTime: time.Unix(1649397699, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeEthnicGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x70\x2d\xc9\xc8\xcb\x4c\x56\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\xcc\x2b\x01\xb3\xf3\x12\x73\x53\xad\x14\x82\x4b\x8a\x32\xf3\xd2\xb9\x6a\x01\x01\x00\x00\xff\xff\x53\xf6\x00\xd4\x2c\x00\x00\x00")

func typeEthnicGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeEthnicGraphql,
		"type/ethnic.graphql",
	)
}

func typeEthnicGraphql() (*asset, error) {
	bytes, err := typeEthnicGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/ethnic.graphql", size: 44, mode: os.FileMode(420), modTime: time.Unix(1649397699, 0)}
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

	info := bindataFileInfo{name: "type/image.graphql", size: 207, mode: os.FileMode(420), modTime: time.Unix(1649397699, 0)}
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

	info := bindataFileInfo{name: "type/major.graphql", size: 102, mode: os.FileMode(420), modTime: time.Unix(1649397699, 0)}
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

	info := bindataFileInfo{name: "type/payment.graphql", size: 199, mode: os.FileMode(420), modTime: time.Unix(1649397699, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeRequirementGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\x8f\x51\x0e\xc2\x20\x0c\x86\xdf\x39\x45\x77\x8d\xdd\x60\xaf\xea\xbb\x21\x5b\xdd\x9a\xac\x65\xb2\xd6\xcc\x98\xdd\xdd\x00\x89\x4e\x79\x81\x9f\x9f\xc0\xf7\x55\x9f\x0b\x42\xc4\xbb\x51\x44\x46\x51\x78\x39\x00\x00\x1a\x5a\xe8\x44\x73\x5e\xfb\x29\xcc\x3e\xae\x13\x2d\xdd\xa1\x4e\x2f\x5b\x38\x6b\x24\x19\x73\x21\x9e\x7f\x8b\x87\x9f\xed\xdb\xec\xce\xa1\x18\x1f\x61\x97\x04\x2f\x40\xc5\xad\x7c\x3b\x84\xde\xd2\x5d\x3e\x04\xc1\xbc\x47\x2f\x63\x49\xec\x37\x62\xe3\x92\x49\x3e\x39\x08\x5e\xc3\x2d\x41\x48\x16\x53\xc8\xeb\xa9\x9a\xab\x48\xff\x29\x34\x95\x7e\x53\xfb\x37\x6e\x7f\x07\x00\x00\xff\xff\xdd\x6d\x74\xdb\x2c\x01\x00\x00")

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

	info := bindataFileInfo{name: "type/requirement.graphql", size: 300, mode: os.FileMode(420), modTime: time.Unix(1649397699, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeScholarshipGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x74\x52\xc1\x6e\x83\x30\x0c\xbd\xf3\x15\xf9\x8e\xdc\x2a\x75\x93\x7a\xe6\x38\xf5\xe0\x26\x1e\xb5\x04\x4e\x66\x3b\xda\xaa\xa9\xff\x3e\x41\x98\x08\xad\x7a\xc2\x7e\xef\xf9\xf1\x30\xb6\x5b\x46\xa7\xe1\x9a\x46\x10\xbd\x52\x76\xbf\x9d\x73\xce\x51\xf4\xee\xc4\xb6\xd4\x9a\x13\x6b\x92\xd3\x33\xe4\x5d\x51\x94\x05\x61\x98\xd0\xbb\xde\x84\x78\x58\x00\x98\x52\x61\x6b\x46\x0c\xac\xe8\xd6\x1b\xfe\x58\xbf\x62\xcd\x18\x4d\x30\xa0\xaf\x8f\xea\xf3\x0d\x12\x11\xb7\xc1\x50\x44\x90\xed\x90\xf3\x48\x01\xda\x57\x40\x85\x8c\x12\xf7\x06\x62\xfb\x3c\x1b\xf9\xc6\x71\x4f\x31\xa7\xc2\x01\x27\x64\x3b\x82\xed\xbf\x03\x47\x1a\xe8\x42\x23\xd9\xed\x88\x1a\x84\xf2\x6c\xb1\x93\x68\xb9\x28\xc5\x97\xb4\xe0\x57\x21\xa9\xee\x9b\x44\xbd\xfb\xa8\xa2\xf3\xa2\xfa\x2c\x1c\x89\x87\xe7\xe0\x2b\xf1\x18\xba\x71\x9d\xad\x9a\xb6\xfa\x65\xb8\xcd\x8d\xff\x2f\xea\xee\x04\xc1\x30\x1e\x5e\xae\xa6\x6f\xfe\xd2\xbd\xeb\x1e\xaf\xe3\x1d\x31\xae\x17\x12\x8a\x2c\x17\xd0\xee\x61\x13\xce\x91\x9a\xf6\xdc\xdd\xff\x02\x00\x00\xff\xff\x57\x6f\xe0\xd8\x69\x02\x00\x00")

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

	info := bindataFileInfo{name: "type/scholarship.graphql", size: 617, mode: os.FileMode(420), modTime: time.Unix(1649583647, 0)}
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

	info := bindataFileInfo{name: "type/school.graphql", size: 159, mode: os.FileMode(420), modTime: time.Unix(1649397699, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeSponsorGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x28\x2e\xc8\xcf\x2b\xce\x2f\x72\x4b\x4d\x4d\x51\xa8\xe6\x52\x50\x50\x50\x48\x2e\x2d\x2a\xce\x2f\xb2\x52\x08\x2e\x29\xca\xcc\x4b\x07\x0b\x41\x15\x15\x5b\x29\x44\x97\x16\xa7\x16\xc5\x72\xd5\x72\x01\x02\x00\x00\xff\xff\x41\x88\xd4\x71\x3d\x00\x00\x00")

func typeSponsorGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeSponsorGraphql,
		"type/sponsor.graphql",
	)
}

func typeSponsorGraphql() (*asset, error) {
	bytes, err := typeSponsorGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/sponsor.graphql", size: 61, mode: os.FileMode(420), modTime: time.Unix(1649397699, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeStudentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x64\xce\xb1\xca\x02\x41\x0c\x04\xe0\x7e\x9f\x22\xcf\x71\xdd\xcf\x2f\x8a\x9d\x60\x29\x16\xf1\x32\xe8\xc2\xde\xee\x91\xe4\x0a\x91\x7b\x77\xe1\x58\xd0\x60\x37\x7c\x4c\x86\xf8\x73\x06\x99\x2f\x82\xea\x7b\x40\xe8\x95\x88\x88\xc6\x45\xad\xe9\x40\x67\xd7\x5c\xef\x1b\xf5\x92\x0d\x74\xe9\xf1\x9a\xd6\x94\xbe\x07\xfa\x71\x96\x81\x8e\xd5\xb7\x5c\x79\x42\x98\xc1\xc4\xb9\x04\x99\x9b\x39\x97\xff\x26\xb1\xc8\x22\x0a\xb3\x60\xb7\xac\xfe\xd8\xb1\xe3\x57\x4f\x85\xc7\xc8\xa3\x82\x1d\xf2\xe7\x51\x59\x01\x3d\x34\xfe\x3c\xb1\xa6\x77\x00\x00\x00\xff\xff\x90\xb8\xcb\xba\x06\x01\x00\x00")

func typeStudentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeStudentGraphql,
		"type/student.graphql",
	)
}

func typeStudentGraphql() (*asset, error) {
	bytes, err := typeStudentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/student.graphql", size: 262, mode: os.FileMode(420), modTime: time.Unix(1649563595, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeUserGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x54\xb1\x8e\xdb\x30\x0c\xdd\xfd\x15\xba\xdf\xd0\x76\xb8\xb4\x45\x86\x1e\x8a\x5e\x3b\x15\x19\x58\x8b\x75\xd4\x93\x49\x43\xa2\x07\xa3\xc8\xbf\x17\xa2\x94\xd8\x56\x0e\x59\x2c\xf1\xf1\x51\x7a\x14\x5f\x22\xcb\x84\x66\x4e\x18\xcd\xbf\xce\x18\x63\xbc\xb3\xe6\x48\xa2\x7b\x82\x11\xad\x79\x93\xe8\x69\x50\x20\x93\xad\xb2\x7f\x2c\x13\x2a\x84\x23\xf8\xb0\x23\x4d\x67\x26\x7c\xe5\x16\x13\xb6\xc6\x8f\x30\x94\xb2\x9e\xc7\x09\x68\x79\x6d\x6f\x48\x02\x32\xa7\x55\x42\xcf\x33\x49\x5c\x8e\xee\x0e\xb2\xd7\x4d\xb9\x80\x93\x40\x78\x61\xb7\x3f\x0e\x9c\x8b\x98\xd2\x0e\x1b\x90\x1c\x46\x5b\x57\x85\x7e\xfb\x28\xe7\x03\xc8\xbe\x58\xd1\x6f\x01\xfa\x3d\x8c\x72\x26\xdf\x6f\x15\x15\xc4\x9a\x4f\xba\x16\x91\x10\xdd\xd1\x21\x89\x17\x8f\xc9\x9a\x5f\x1b\x60\x39\x95\xe3\x81\xde\xb7\xa7\xe4\xd8\xea\xf7\x16\x3e\xf7\xda\x63\xf3\x98\xdb\x4c\xfb\x7e\x7d\x44\x10\x74\xcf\xb2\x47\x21\x22\xc6\x2f\x0c\xfb\x49\xe5\x41\xbe\xf5\x67\xe6\x90\x15\xae\xd1\xe9\x96\x3d\x70\x3f\x8f\x48\x72\xcd\x5f\xe3\x53\x77\xe9\x3a\xb5\x4e\xe0\xc1\xd3\x77\x4c\x13\x53\xc2\xea\x21\xe1\x77\xa4\xbb\x9b\x8a\x71\x72\x21\xd2\x3c\xde\x5c\x54\x6b\xf4\x00\x8e\xd5\x04\x73\x7e\xa9\x3a\xc1\xd1\xd3\xad\xa8\xcc\xac\x96\x8c\x10\x8a\x99\xfe\xa0\x6e\x2f\x5d\xe7\x69\x9a\xc5\xe8\xf7\x67\xc2\xf8\x52\x1c\x92\x1b\xdf\xdb\xfb\xe9\x03\x72\xe9\xfd\x2a\x47\x03\x5b\xb2\x25\xb3\x72\x9e\x94\xe2\x70\x88\x88\x95\x72\xd0\x60\xa5\x54\x81\x7f\x39\x56\xc2\xd7\xbc\x6f\xf2\x48\x91\x43\xc8\xef\x79\xe7\xbd\x21\x82\x9b\x41\x3c\xd3\x36\x55\x2e\x1e\x26\xb0\xe6\x73\x60\x90\xa6\x8b\x56\xe7\xe3\x9e\x1b\x45\x8f\xc9\x6d\x7f\x1f\xb0\x37\x03\xda\xcc\xf5\xf1\x90\xca\x2f\xa6\xf9\xeb\xb9\xfc\x0f\x00\x00\xff\xff\x2d\x2d\x1b\x68\x94\x04\x00\x00")

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

	info := bindataFileInfo{name: "type/user.graphql", size: 1172, mode: os.FileMode(420), modTime: time.Unix(1649559835, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeUser_documentGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x2a\xa9\x2c\x48\x55\x28\x2d\x4e\x2d\x72\xc9\x4f\x2e\xcd\x4d\xcd\x2b\x51\xa8\xe6\x52\x50\x50\x50\xc8\x4c\xb1\x52\xf0\xcc\x2b\x01\xb3\x41\xd2\x9e\x48\xfc\x14\xa8\x52\x2b\x85\xcc\xdc\xc4\xf4\x54\xae\x5a\x40\x00\x00\x00\xff\xff\x4c\x8a\xd9\xa7\x45\x00\x00\x00")

func typeUser_documentGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeUser_documentGraphql,
		"type/user_document.graphql",
	)
}

func typeUser_documentGraphql() (*asset, error) {
	bytes, err := typeUser_documentGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/user_document.graphql", size: 69, mode: os.FileMode(420), modTime: time.Unix(1649397699, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _typeUser_schoolGraphql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x8a\x31\x0e\xc2\x40\x0c\x04\xfb\xbc\x62\x1f\xc0\x0b\x5c\x23\xa4\xd4\x79\x81\xc5\x59\xc7\xa1\x8b\x1d\x39\x4e\x81\x10\x7f\x47\xc4\x29\x28\x52\x8d\x76\x66\xe3\xb5\x08\xb6\x55\x7c\xba\x3f\xcc\x3a\xde\x03\x00\xb4\x42\x18\x35\x2e\xfb\xf8\xd5\xf1\x5f\xac\xfb\x95\x0e\xa6\x2b\x52\x5d\x84\x0e\xa6\x9b\xf9\x69\x4e\x89\x34\xa2\x6e\xbd\xcf\xa2\x71\xe5\x10\xc2\x14\xde\xb4\x66\xab\xce\x65\xe3\x68\xa6\x27\x6d\x61\xc2\xad\x1b\xc7\xf0\xf9\x06\x00\x00\xff\xff\x46\x48\x4d\x37\xb2\x00\x00\x00")

func typeUser_schoolGraphqlBytes() ([]byte, error) {
	return bindataRead(
		_typeUser_schoolGraphql,
		"type/user_school.graphql",
	)
}

func typeUser_schoolGraphql() (*asset, error) {
	bytes, err := typeUser_schoolGraphqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "type/user_school.graphql", size: 178, mode: os.FileMode(420), modTime: time.Unix(1649397699, 0)}
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
	"type/applicant.graphql":     typeApplicantGraphql,
	"type/assessment.graphql":    typeAssessmentGraphql,
	"type/bank.graphql":          typeBankGraphql,
	"type/bank_transfer.graphql": typeBank_transferGraphql,
	"type/card_identity.graphql": typeCard_identityGraphql,
	"type/country.graphql":       typeCountryGraphql,
	"type/degree.graphql":        typeDegreeGraphql,
	"type/ethnic.graphql":        typeEthnicGraphql,
	"type/image.graphql":         typeImageGraphql,
	"type/major.graphql":         typeMajorGraphql,
	"type/payment.graphql":       typePaymentGraphql,
	"type/requirement.graphql":   typeRequirementGraphql,
	"type/scholarship.graphql":   typeScholarshipGraphql,
	"type/school.graphql":        typeSchoolGraphql,
	"type/sponsor.graphql":       typeSponsorGraphql,
	"type/student.graphql":       typeStudentGraphql,
	"type/user.graphql":          typeUserGraphql,
	"type/user_document.graphql": typeUser_documentGraphql,
	"type/user_school.graphql":   typeUser_schoolGraphql,
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
		"applicant.graphql":     &bintree{typeApplicantGraphql, map[string]*bintree{}},
		"assessment.graphql":    &bintree{typeAssessmentGraphql, map[string]*bintree{}},
		"bank.graphql":          &bintree{typeBankGraphql, map[string]*bintree{}},
		"bank_transfer.graphql": &bintree{typeBank_transferGraphql, map[string]*bintree{}},
		"card_identity.graphql": &bintree{typeCard_identityGraphql, map[string]*bintree{}},
		"country.graphql":       &bintree{typeCountryGraphql, map[string]*bintree{}},
		"degree.graphql":        &bintree{typeDegreeGraphql, map[string]*bintree{}},
		"ethnic.graphql":        &bintree{typeEthnicGraphql, map[string]*bintree{}},
		"image.graphql":         &bintree{typeImageGraphql, map[string]*bintree{}},
		"major.graphql":         &bintree{typeMajorGraphql, map[string]*bintree{}},
		"payment.graphql":       &bintree{typePaymentGraphql, map[string]*bintree{}},
		"requirement.graphql":   &bintree{typeRequirementGraphql, map[string]*bintree{}},
		"scholarship.graphql":   &bintree{typeScholarshipGraphql, map[string]*bintree{}},
		"school.graphql":        &bintree{typeSchoolGraphql, map[string]*bintree{}},
		"sponsor.graphql":       &bintree{typeSponsorGraphql, map[string]*bintree{}},
		"student.graphql":       &bintree{typeStudentGraphql, map[string]*bintree{}},
		"user.graphql":          &bintree{typeUserGraphql, map[string]*bintree{}},
		"user_document.graphql": &bintree{typeUser_documentGraphql, map[string]*bintree{}},
		"user_school.graphql":   &bintree{typeUser_schoolGraphql, map[string]*bintree{}},
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
