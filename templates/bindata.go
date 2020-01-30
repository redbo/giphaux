// Code generated by go-bindata.
// sources:
// .gitattributes
// bindata.go
// error.tmpl
// footer.tmpl
// gif.tmpl
// header.tmpl
// index.tmpl
// search.tmpl
// templates.go
// user.tmpl
// DO NOT EDIT!

package templates

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

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _Gitattributes = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x4a\xca\xcc\x4b\x49\x2c\x49\xd4\x4b\xcf\x57\xd0\x4d\xc9\x4c\x4b\xe3\x02\x04\x00\x00\xff\xff\xe5\xa1\x25\x74\x11\x00\x00\x00")

func GitattributesBytes() ([]byte, error) {
	return bindataRead(
		_Gitattributes,
		".gitattributes",
	)
}

func Gitattributes() (*asset, error) {
	bytes, err := GitattributesBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: ".gitattributes", size: 17, mode: os.FileMode(438), modTime: time.Unix(1580254215, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _bindataGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x01\x00\x00\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00")

func bindataGoBytes() ([]byte, error) {
	return bindataRead(
		_bindataGo,
		"bindata.go",
	)
}

func bindataGo() (*asset, error) {
	bytes, err := bindataGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(438), modTime: time.Unix(1580400919, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _errorTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\xca\x48\x4d\x4c\x49\x2d\xd2\x2b\xc9\x2d\xc8\x51\x52\xd0\xab\xad\xe5\xe2\xb2\x29\xb0\x0b\xc9\xc8\x2c\x56\xc8\x2c\x56\x48\xcc\x53\x48\x2d\x2a\xca\x2f\x52\x28\x48\x4c\x4f\x55\xb4\xd1\x2f\xb0\x03\xc9\x56\x57\xeb\xb9\x24\x96\x24\xea\xf9\x16\xa7\xd7\xd6\x82\x05\xb9\x90\x4d\x4c\xcb\xcf\x2f\x41\x31\x11\x10\x00\x00\xff\xff\x68\x1e\x54\x0a\x6f\x00\x00\x00")

func errorTmplBytes() ([]byte, error) {
	return bindataRead(
		_errorTmpl,
		"error.tmpl",
	)
}

func errorTmpl() (*asset, error) {
	bytes, err := errorTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "error.tmpl", size: 111, mode: os.FileMode(438), modTime: time.Unix(1580400913, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _footerTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb2\xd1\x4f\xc9\x2c\xb3\xe3\xb2\xd1\x4f\xca\x4f\xa9\x04\xd1\x19\x25\xb9\x39\x76\x5c\x80\x00\x00\x00\xff\xff\x55\x33\xd2\x97\x17\x00\x00\x00")

func footerTmplBytes() ([]byte, error) {
	return bindataRead(
		_footerTmpl,
		"footer.tmpl",
	)
}

func footerTmpl() (*asset, error) {
	bytes, err := footerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "footer.tmpl", size: 23, mode: os.FileMode(438), modTime: time.Unix(1579838624, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _gifTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\xd1\x6e\xab\x38\x10\x7d\xcf\x57\x8c\xbc\xd5\x6a\x57\xda\xc0\xaa\xda\x27\x02\xec\xc3\xad\x7a\xd5\xa7\x5e\xdd\xb6\x1f\x30\xe0\x01\xac\x82\xcd\x35\x43\xda\x08\xf1\xef\x57\x10\x20\xa4\x49\xd3\xaa\x91\xaa\x56\xaa\x48\x1c\xcf\xf1\x99\x33\x73\xc6\x34\x0d\x53\x51\xe6\xc8\x04\x22\x23\x94\x64\x1d\x2e\xca\x5c\x80\xd3\xb6\x8b\x85\x2f\xd5\x1a\x2a\xde\xe4\x14\x88\x02\x6d\xaa\xf4\x92\x4d\xe9\xc1\x7f\xff\x96\xcf\x2b\x11\x2e\x00\x00\xfc\xec\x72\xdc\x92\xe4\x06\xd9\x83\x9c\x12\x5e\xc1\x76\xbf\x07\xdb\xad\x4d\xe3\x5c\x21\xa3\x73\xaf\x38\xa7\xb6\xf5\xdd\xec\x72\x08\x9f\x1d\x31\xc4\x5b\x95\x66\x3c\xc2\x77\x7f\x4d\x63\x51\xa7\x04\x03\x04\xa6\x55\xdb\x4e\x3f\xf6\x20\x08\x71\x8e\x55\x15\x08\xc6\x54\x40\x66\x29\x09\x84\x5b\x11\xda\x38\xfb\xff\x57\xc0\x98\x7a\x4d\xe3\xb4\xad\x08\xff\xe8\x9f\xbe\x8b\x73\x74\xd2\x72\x00\xf4\x5d\xa9\xd6\xe1\x62\x7c\xcc\xb8\x31\x3d\xf3\x12\x73\x95\x6a\x0f\x62\xd2\x4c\x76\x15\xe7\x84\xd6\x83\xc8\x70\xd6\xb1\xf5\x55\x91\x42\x65\xe3\x40\xb8\xa9\x4a\xdc\x31\xe3\x9b\xab\xb6\x75\x52\x95\x88\x09\x76\xe1\x47\x36\xec\xfe\x17\x8b\xa6\x51\x09\x38\x0f\x15\xd9\x4e\xee\x2d\x9b\x6e\xa5\x8f\xec\x96\x6f\x74\x62\x9c\x6b\x5c\x1b\xab\x98\xe4\xb8\xa9\xa7\x5a\x86\xfb\x1a\x24\xc6\x16\xa3\x0c\x11\x56\x2a\xee\x16\x04\x60\xcc\xca\xe8\x40\xb8\x75\x45\xd6\xad\x4b\x89\x4c\x31\x32\xa5\xc6\x2a\xaa\x04\x14\xc4\x99\x91\x81\xf8\x71\x7b\x77\x2f\xf6\x21\x7b\x58\xa5\xcb\x9a\x81\x37\x25\x05\x22\x53\x52\x92\x16\xa0\xb1\xa0\x40\xa4\x2a\x51\x52\xc0\x1a\xf3\x9a\x02\x31\xcf\xf7\x08\xce\x58\xc3\x8b\x47\xda\xfc\x03\x17\x7d\x10\x78\xc1\xcb\x54\xbf\x4d\xcc\x5e\x94\x78\xe2\x93\x63\x44\xf9\xbc\xdc\x87\x67\xed\xce\x54\xc9\x70\xd4\x2b\x68\x47\xb3\x8c\x33\x8a\x1f\x23\xf3\x3c\xe6\x19\x23\x2f\x9b\xa6\x23\xde\xb6\x62\xdf\x0e\xdb\xf6\x86\x12\xa5\x54\x3a\x1d\xbe\x3d\x29\xc9\x99\x07\x58\xb3\x59\x4d\xf2\x18\x2d\xa0\x07\x26\x19\xc2\x88\x76\x82\x39\xe5\xd5\x97\x20\xfd\x3e\xb2\x93\x81\x0e\x38\xba\x7d\xbd\x8e\xf5\xc3\xf1\xa0\xde\x73\x43\x75\xad\x79\x12\xa1\x1f\xd5\xcc\x46\x0f\x69\x56\x75\x54\x28\x16\xe1\x43\xdf\xc7\xb0\x6b\x17\xf0\xd5\x18\x96\x20\x24\xb8\x24\xd9\xed\xf3\x5d\x15\xfa\xee\x16\x22\x1c\xfc\xb7\xcf\xaf\x73\xc9\x6e\xcd\x77\x67\xb6\xfa\xa0\xc5\x74\x32\xf8\xf5\xd3\xcc\xf5\x4e\xd1\x7e\x52\x61\xd6\x04\xd7\xd6\x14\x30\xce\x94\x43\xe1\x32\x42\xcb\x4b\xf3\x41\xed\x86\x21\x36\xb4\xef\x99\x52\x7e\xba\x90\x70\x7c\xc2\xfc\xd9\xbb\x6b\x05\x23\x9f\xea\xd5\xb6\x86\xf9\x75\xd5\x4d\xb5\xb7\x27\xda\x2b\x67\x9e\x6b\x7d\xe7\x7c\xdf\x3b\xa7\x18\x9f\x56\xe0\x2c\x73\xdf\xe1\x9a\xe0\xde\xbc\xd5\xa3\x67\x76\xa8\x96\xa7\xee\xdc\xdb\x27\xdd\x5f\xca\x53\xe8\x87\x1a\x58\x52\x4e\x5f\x6f\x0e\x5c\xf5\xac\xe0\xfb\xcd\x35\xfc\xb5\x31\x35\xa0\x25\xe0\x8c\xc0\x74\x39\xff\x7d\x20\x76\x84\xfa\x8c\x41\x3a\x29\x3d\xfb\xb0\x7b\xe7\x4c\x8c\xe1\xbd\x77\xce\xdf\x01\x00\x00\xff\xff\xae\xd4\x3a\x67\x91\x0a\x00\x00")

func gifTmplBytes() ([]byte, error) {
	return bindataRead(
		_gifTmpl,
		"gif.tmpl",
	)
}

func gifTmpl() (*asset, error) {
	bytes, err := gifTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "gif.tmpl", size: 2705, mode: os.FileMode(438), modTime: time.Unix(1580258612, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _headerTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x4b\x73\xe3\xb8\x11\xbe\xeb\x57\x20\x74\x25\xb6\xa7\x24\xd2\xf2\x58\xb3\x59\x8a\x52\xd5\xc6\xd9\xcc\x4c\x55\x0e\x9b\x9d\xdd\xaa\x54\xa5\xf6\xd0\x02\x9a\x24\xc6\x20\xc0\x00\xa0\x1e\x71\xf9\xbf\xa7\x40\x52\x14\x9f\x9e\x99\x3d\xe4\x14\x1d\x6c\x11\x8f\xee\x0f\xfd\x7d\xdd\x0d\x2a\x4a\x6d\x26\xb6\xb3\x28\x45\x60\xdb\x59\x64\xb9\x15\xb8\x7d\xcf\xf3\x14\xa1\x38\x46\x41\xf5\x3c\x8b\x04\x97\x4f\x44\xa3\xd8\x78\xc6\x9e\x04\x9a\x14\xd1\x7a\x24\xd5\x18\x6f\xbc\xd4\xda\xdc\x84\x41\x40\x99\xfc\x6c\x7c\x2a\x54\xc1\x62\x01\x1a\x7d\xaa\xb2\x00\x3e\xc3\x31\x10\x7c\x67\x82\x58\x49\xbb\x80\x03\x1a\x95\x61\xf0\xe0\x7f\xe7\xdf\x05\xd4\x74\x87\xfd\x8c\x4b\x9f\x1a\xe3\x6d\x67\x51\xe9\x67\x3b\x23\x84\x90\x9d\x62\x27\xf2\x5c\x7e\x75\x9f\x0c\x74\xc2\x65\x48\xee\xd6\xcd\x50\x69\x24\x86\x8c\x8b\x53\x48\x7e\xd0\x1c\xc4\x9c\x7c\x40\xb1\x47\xcb\x29\xcc\x89\x01\x69\x16\x06\x35\x8f\x2f\x5b\x76\x40\x9f\x12\xad\x0a\xc9\x16\x54\x09\xa5\x43\xb2\x13\x40\x9f\x2e\x0b\xea\xd1\x43\xca\x2d\x5e\x46\x2d\x1e\xed\x02\x04\x4f\x64\x48\x28\x4a\x8b\xba\x9a\x7b\x29\xff\x5e\x51\x25\x2d\x4a\x6b\x5a\x70\x19\x37\xb9\x80\x53\x48\xb8\x14\x5c\xe2\x62\x27\x54\xdb\xcd\x81\x33\x9b\x86\x64\x79\x77\xff\x90\x1f\xc7\xfd\x08\x8c\xed\x65\x26\x57\x86\x5b\xae\x64\xe8\xf8\x00\xcb\xf7\xd8\x41\x60\x55\x2e\x61\xdf\xf2\x9f\x03\x63\x5c\x26\xce\xc5\xd9\x41\x77\x29\xb4\x16\x97\x6e\x19\x52\xa5\xa1\xf2\x21\x95\xc4\xd7\x62\xd2\x33\xe5\xc7\x4a\x67\xa5\x58\xfa\x7c\x2d\xac\xca\xdb\x18\x5a\x33\xee\x7c\x21\xe9\x1c\x3f\x16\x0a\x6c\x48\x34\x4f\xd2\xf6\xd1\xcf\x47\x79\xd7\x5e\x3b\x4e\xdf\x90\x5f\x06\xfa\x29\xd1\x70\x1a\x03\x9e\x2e\x5b\x80\x4b\x35\x19\xfe\x1f\x0c\xc9\xc3\xaa\xed\xa9\x47\x65\xc7\x90\x50\x89\xfa\x1a\x1b\x17\x5c\x21\x59\x1c\x70\xf7\xc4\xed\xc2\x19\x03\xbd\x48\x34\x30\x8e\xd2\xde\x3c\xac\x18\x26\x73\x42\x4f\x20\xe7\x24\xd1\x88\x72\x4e\x4e\x28\x84\x3a\xcc\x49\x06\x09\x4a\x0b\xb7\x17\x8b\x67\x2b\xed\x13\x0b\x9e\x87\x25\x99\xc3\x65\x25\xc5\x31\x17\xe2\x1c\x17\xab\x41\x9a\x1c\x34\x4a\xdb\x39\x91\x41\xd0\x34\x75\x8a\x06\x2e\x51\x4f\x31\xba\xea\xa9\xaa\xbf\x8d\xcb\xbc\xb0\xad\xcd\xb5\xe0\xbf\x5f\x75\xa4\x90\xa2\xa3\x3a\x24\x0f\x9d\xd1\x86\xf0\xfb\xce\xf0\x4e\x69\x86\xba\x2f\xce\x56\xc8\x97\xdf\x7d\x01\xd4\xae\xb0\x56\xc9\x21\xaa\xaf\x01\xd5\x66\xf0\x8a\x31\x36\x8a\xe0\xfe\xfe\xcb\x80\x69\xa1\x8d\x23\x20\x57\x7c\x50\x48\xc6\x01\x87\xa9\xda\x77\x98\xe8\x60\xa1\x94\xb6\x6d\xf8\x75\x31\x5a\x54\xde\x9b\x5d\x2d\x9d\x70\xa7\xa6\x90\xbc\xaa\xbf\xbc\xd0\xb9\xc0\x39\xd9\x73\x25\xd0\x36\xba\x3b\xc7\xa6\x5f\x55\xfc\x1d\x18\x4e\x5d\x19\x20\xcf\x4d\xbe\x54\x35\x6f\x24\x27\xaf\x56\xab\xd5\xba\x55\xa2\x9c\x98\x2e\x05\x3e\x3f\xae\x07\x46\xd3\xfb\x91\x56\x50\x82\x98\xac\xcd\x7d\xc9\x4e\x23\xf6\xb5\x3a\x10\x3f\xe6\x28\xd8\x84\x17\xb7\x99\xac\xf2\x23\xb9\x9f\x34\x52\x0a\x7e\xde\x1e\x19\xa8\xad\x39\xf0\x6a\x58\x0e\x5b\xf8\x48\xbb\x3f\xdc\xfd\xb1\xaf\xa7\x85\xe3\xab\x30\xbd\xd2\x79\xd6\x5a\xc7\xca\x32\x3f\x12\xa3\x04\x67\xe4\x6a\xb9\x5c\xbe\x06\xfc\x5f\xf6\x94\xe3\xc6\x14\xbb\x8c\xdb\xdf\xe6\x13\xd3\xd5\x81\x7e\x7b\xfd\x90\x63\x6c\xef\xde\xad\xbb\x4d\xa4\xca\x99\x43\x2d\xa6\x9d\x12\x6c\xdd\x49\x64\x7f\x89\xd9\x7a\x90\x29\x1d\x3a\xdf\xf6\x72\x5d\xa8\x84\x4b\x07\x69\x4e\xae\x34\x26\xdc\x58\xd4\x95\x1e\x87\x85\xbc\x57\x41\x46\xbb\xce\x39\xa0\x97\x18\xf6\xee\x04\x97\x8e\x0c\x3b\xa3\x44\xd1\xb9\x2f\x0c\xbb\x5e\xcd\xe8\xbb\xbb\xce\xa8\xae\x22\xd0\xd7\xa6\x85\xe4\xf5\x90\xbe\xfd\xf3\xe3\x48\xbd\x5c\x0d\x15\xd1\xa8\x65\x54\x73\x6f\x07\x97\x8f\x6f\xb9\x05\x8c\x87\xe9\x0a\x00\xfa\xc5\xd1\xdd\xea\x42\xc2\x2d\x08\x4e\x27\x1b\xeb\xe0\x8e\xe4\x3c\x2d\x4c\x0e\x14\x1d\x96\x83\x86\x7c\x10\xa4\x29\x86\xc7\x2d\x0e\x52\xad\x6f\x69\x32\x65\xa7\x2f\x30\xcb\x89\x2e\x31\xd9\x67\xc9\x54\xb6\xd6\x1b\x84\x93\xc4\x4e\x14\x5f\xd7\x32\xfa\xd8\x87\xbd\xe2\x15\x57\x3d\x2e\x27\xdc\x44\x41\x7d\x2b\x9f\x45\x86\x6a\x9e\xdb\xea\x7e\x1e\x17\x92\x3a\x9d\x10\xaa\xf2\xd3\x2f\x78\xb4\x37\x9c\xdd\xb6\x1c\x07\x6f\xc8\x7b\xb4\xc4\xa6\x58\x2a\x8b\x54\xf5\xf5\x4d\xd0\x2c\xd8\x83\x6e\xf6\x92\x0d\x61\x8a\x16\x19\x4a\xeb\x27\x68\x7f\x14\xe8\xbe\xfe\xe5\xf4\x91\x39\xab\xeb\xd9\xac\xb1\xf9\x09\x05\xd2\x69\xb3\x67\x83\xbe\x29\xd7\xdd\x34\x9d\xab\x35\x61\x2b\x1b\x5c\xc9\x9f\x41\x26\x78\x73\x37\x27\xdf\xbb\xcf\xed\x9a\x04\x6f\xfe\xa6\x34\xc9\xd4\x8e\x0b\x24\x0c\xf7\x9c\xa2\x79\x13\xb4\xdc\x3f\xaa\xfc\x74\x71\xce\xa5\xe1\x0c\x27\xc1\x34\x47\xc2\x23\xd2\x47\x95\x65\x20\xd9\x8d\xe7\x90\x78\x0d\x2e\x8d\xb6\xd0\x92\xc4\x20\x0c\xb6\x22\x5e\x07\x7a\x16\x05\xf5\x3b\x9a\x7b\x1d\xda\xce\x22\xc6\xf7\x84\xb3\x8d\x77\x7e\xe5\xf0\x2a\x32\x9a\xf1\xa6\x0c\xd6\x13\xe5\x64\xa9\x0f\x2a\xc0\x98\x8d\xd7\x94\x6e\x8f\x40\x19\x83\x8d\x17\x94\x9b\x3c\x92\xa1\x4d\x15\xdb\x78\xb9\x32\xb6\xb5\xbd\x34\x91\xde\x6f\xff\xae\x12\xf2\x51\x46\x41\x7a\xdf\x9b\x73\xbe\x6b\xeb\x5a\x1d\xbc\x6d\x7b\xa0\x8c\x88\xb7\xfd\xd5\xa0\x96\x90\x61\x14\x30\xbe\xdf\x46\xd5\xed\xb0\x6c\x2a\x9e\x0b\x9c\x47\xdc\xe4\xc6\x2b\xea\x65\xde\xe5\x2c\xcd\xd0\xb6\xda\xfb\xcd\xae\x7f\x02\x63\x0e\x4a\xb3\x11\xd7\x79\x3d\x75\x76\xdf\x3c\xff\x5e\x5f\x9f\xca\x06\x3a\xe2\xa9\xea\xac\x1e\xd9\x83\x28\x70\xe3\x55\xa1\x24\x7f\xd2\xa0\xf5\xfa\xec\xbd\x5e\xd3\xf7\x1d\x05\x8e\xad\x9a\xe6\xcb\x54\xc3\x78\xbb\xd9\x7d\x0b\xe9\xe7\x7d\x5f\xe4\xfd\x51\x23\x58\x24\x8e\xc2\xff\x21\xf9\x67\x78\xff\xe7\xff\x0b\xfc\x57\xef\xb2\x6d\xe6\xd3\xe5\x36\x82\xfa\x67\x9a\xc0\xdb\x46\xfc\x9c\x4d\xca\x6b\xf0\x02\x89\x61\x41\x21\x43\x0d\x0b\x8d\x56\x2b\xe7\x97\x6f\xc9\xfb\x8f\x3f\x7d\xf8\xe1\xd7\x7f\x46\x01\x6c\xa3\x20\x5d\x5e\xac\x3e\x3f\xf3\x98\xf8\x8e\xca\x97\x97\x6e\x5c\x1a\x57\xbb\x13\x5e\x1c\xd4\xbf\x08\x78\xae\x6c\xa8\xc2\x3a\x83\x13\xdb\x1c\xc3\xc1\xc8\xc6\x0f\x2a\xc3\x90\x3c\x3f\x97\x4e\xfd\xb3\x88\x5e\x5e\x3a\xa6\x9e\x9f\x51\x18\x1c\x42\xea\x5b\x23\x4a\x3e\x0a\x4e\x9f\x36\xde\x67\xd8\x43\x55\x5e\xc3\x5e\xa7\xb9\x6e\x8a\xe7\xf5\xad\x5f\xf6\x3c\xbf\xbe\x50\x6c\xae\xcb\x9b\xc4\xf5\x7a\x74\xc7\x59\xa2\xd7\xb7\x7e\xac\x68\x61\x6e\x6e\xd7\x5e\x53\x2d\x87\xc7\xfe\x5d\xd0\xda\x59\x3e\x89\x6e\x00\xaf\x9f\x44\x1d\x84\x3f\xd7\x93\xfd\x78\x4a\x56\x87\xb3\xaf\xb6\x1a\x78\xf7\x05\xb3\xa3\xd7\x46\x94\xbd\x17\xd9\x41\x5d\x6a\xaa\x50\xb5\xf0\x52\x83\xde\xff\xf8\x4b\xbf\x04\x55\x79\x54\xe5\xc8\xbf\xbd\x4e\xe1\xc8\x05\x50\x4c\x95\x60\xa8\x37\xde\xa7\xd2\x94\xbb\x0c\x91\x84\xc7\x86\x80\x64\xc4\xd8\x22\x8e\xff\xd0\xe4\x5d\x25\xe2\xbf\x82\x85\x97\x97\xea\xfb\x3f\x0a\xd4\x27\xf7\x50\x8e\x5e\x1e\xcb\x28\xd4\xff\xfa\x80\xea\x5b\x62\x27\xb3\x5d\x96\x75\x72\xab\x3e\x58\x99\x55\x51\x50\x6d\x79\x35\xa9\xff\x1b\x00\x00\xff\xff\x4e\x00\x87\xd1\x94\x15\x00\x00")

func headerTmplBytes() ([]byte, error) {
	return bindataRead(
		_headerTmpl,
		"header.tmpl",
	)
}

func headerTmpl() (*asset, error) {
	bytes, err := headerTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "header.tmpl", size: 5524, mode: os.FileMode(438), modTime: time.Unix(1580397805, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _indexTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\xca\x48\x4d\x4c\x49\x2d\xd2\x2b\xc9\x2d\xc8\x51\x52\xd0\xab\xad\xe5\xe2\x0a\xce\xcf\x4d\x55\x28\x29\x4a\xcd\x4b\xc9\xcc\x4b\x57\xc8\xcc\x4d\x4c\x4f\x2d\x56\x54\x50\x08\x01\x89\x80\xd9\x5c\x5c\xc8\x06\xa4\xe5\xe7\x97\xa0\x18\x00\x08\x00\x00\xff\xff\x4c\x2e\x0a\x77\x5e\x00\x00\x00")

func indexTmplBytes() ([]byte, error) {
	return bindataRead(
		_indexTmpl,
		"index.tmpl",
	)
}

func indexTmpl() (*asset, error) {
	bytes, err := indexTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.tmpl", size: 94, mode: os.FileMode(438), modTime: time.Unix(1580235931, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _searchTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x51\xc1\x6e\x9d\x30\x10\xbc\xf3\x15\x2b\xe7\x0c\xbc\x50\x71\x31\x84\x4b\x23\x55\xbd\xb6\xfd\x81\x0d\x5e\x6c\xab\xc6\xb6\xcc\x26\xa5\xb2\xfc\xef\x55\x79\xbc\xf4\x35\xd7\x19\xcf\x78\x66\x27\x67\xa6\x35\x3a\x64\x02\x61\x08\x15\xa5\x86\xd7\xe8\x04\x34\xa5\x54\x15\x00\xc0\x68\xba\xe9\x3b\x61\x9a\x0d\x7c\xa3\xed\xd5\xf1\x26\x21\xe7\xe6\x19\x19\x9b\x1f\x81\xd1\x9d\x68\x29\xa0\xed\xb2\x8d\xad\xe9\xa6\xab\x32\xe7\x84\x5e\x13\x5c\xdf\x7e\xb1\xcb\x56\xca\x41\x1c\xb6\xca\xbe\xc1\xc6\xbf\x1d\x3d\x89\x5f\x56\xb1\x91\xf0\xe9\x72\x89\xfb\x00\x2f\x38\xff\xd4\x29\xbc\x7a\x55\xcf\xc1\x85\x24\xe1\xa1\xeb\xba\x01\x98\x76\xae\xd1\x59\xed\x25\xcc\xe4\x99\xd2\x00\x11\x95\xb2\x5e\x4b\x78\x3c\xa4\x2b\x26\x6d\xbd\x84\xfe\xf0\x09\x49\x51\x92\xf0\x18\x77\xd8\x82\xb3\x0a\x1e\xfa\xbe\x1f\x60\x71\x01\x59\x82\xa3\x85\x07\x31\xbd\x07\x3a\x42\x21\x98\x44\xcb\x93\x68\xff\x56\x69\x73\x6e\xbe\x3e\x97\x22\xa6\xd1\xae\x1a\xb6\x34\x5f\x89\x1b\xde\x68\xbb\x88\x5b\x87\x15\xf7\xfa\xec\xd1\x5d\xce\x30\x7b\x6d\xc8\x6a\xc3\x37\x48\x4c\x63\x8b\x1f\x7e\x8c\x53\xce\xcd\x67\x8c\x6c\x83\x2f\x65\x6c\xe3\x3f\x7e\x6c\x95\x7d\x9b\xce\x53\x92\x57\xef\x8b\xdc\x9d\x6e\x76\x84\x49\xc2\x4b\x60\x73\xd8\x1f\x8a\xaa\xba\x9f\x75\x09\x81\xff\x9b\xf5\x4f\x00\x00\x00\xff\xff\xd5\xb4\xca\x73\xf4\x01\x00\x00")

func searchTmplBytes() ([]byte, error) {
	return bindataRead(
		_searchTmpl,
		"search.tmpl",
	)
}

func searchTmpl() (*asset, error) {
	bytes, err := searchTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "search.tmpl", size: 500, mode: os.FileMode(438), modTime: time.Unix(1580311304, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _templatesGo = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x91\x31\x8f\x9b\x40\x10\x85\x6b\xf6\x57\x4c\xa8\x96\x88\xc3\xbd\x23\x17\x69\x52\x25\xa7\x14\xe9\x4e\xa7\x68\x0c\x03\x5e\x79\xd9\x45\x33\x8b\x22\xcb\xe2\xbf\x47\xb3\x80\x9d\x74\xd7\x50\xbc\x79\xfb\xcd\x7b\xc3\x84\xed\x15\x07\x82\x44\xe3\xe4\x31\x91\x18\xe3\xc6\x29\x72\x02\x6b\x8a\xf2\x92\x46\x7f\xd8\x47\xa5\xa9\x8c\x39\x1c\x86\x78\x1c\x28\x10\x63\x22\x18\xe2\xcb\xd9\x85\x0e\x13\xc2\xcb\x74\x1d\x9e\x14\x68\xd4\x0a\xdf\x23\x76\xbf\x1e\xda\x84\x2c\x24\x80\xde\xff\x63\xec\x39\x8e\xb0\x03\x3b\xd8\x71\x18\x3a\x60\x4a\x33\x07\x81\x74\xa1\x11\x50\x00\x41\x66\x97\xa8\x31\xfd\x1c\xda\xff\xd9\xb6\x02\xfb\x79\x87\x36\xbb\x5c\x03\x31\x47\xae\xe0\x6e\x0a\x7d\xf3\x03\x27\x38\x9e\x60\xc4\xe9\x4d\x12\xbb\x30\xbc\xbb\x90\x88\x7b\x6c\xe9\xbe\xdc\x4d\x51\x94\x6d\x0c\x09\x5d\x90\xf2\x08\xfa\xc0\x06\xa2\xce\x13\xac\xee\x1a\x2e\x78\x93\x84\xed\x15\xde\xde\x57\xa9\x82\x73\x8c\x5e\xf9\x45\xd1\x47\x86\xdf\x35\x88\xae\x60\x0c\x03\x3d\xed\x79\x5e\xb8\x1e\x36\xde\xe9\x04\xb2\x89\xc5\xda\x12\x12\xcf\x94\x85\xc5\xec\x9f\x6d\xd2\xa3\x17\x1d\x2d\xb5\x51\x9d\x63\x4c\xba\xe2\xd1\xf6\x95\xfe\xd8\xf2\x50\x56\xcd\xb7\x39\xb4\x62\xb7\xa2\x95\x79\x04\xe2\xf6\x19\xe9\xab\x08\xa5\x57\x1c\xf3\xc9\x34\x81\xeb\xe1\x1c\xbb\x5b\x3e\x95\xda\xb2\xc1\x0a\xb7\xd5\x97\x2c\x7d\x3a\x41\x70\x5b\xc3\x2d\x50\x70\x3e\xdb\x35\x13\x90\x17\x02\xd7\xeb\xa2\x8d\xa0\x01\x73\x28\x85\x34\x3f\xf5\xa7\xdb\xf5\x5a\x56\x37\x55\x1f\x03\xaf\x5d\x57\x5d\x89\xb5\x4e\xcd\x62\xfe\x06\x00\x00\xff\xff\x38\x6a\x1f\xdb\xb1\x02\x00\x00")

func templatesGoBytes() ([]byte, error) {
	return bindataRead(
		_templatesGo,
		"templates.go",
	)
}

func templatesGo() (*asset, error) {
	bytes, err := templatesGoBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "templates.go", size: 689, mode: os.FileMode(438), modTime: time.Unix(1580235937, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _userTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x6d\x6f\xe3\xb8\x11\xfe\x9e\x5f\x31\x55\xee\x6a\x07\x1b\xcb\x4e\x16\x69\xb6\xb2\xec\x62\x2f\xd7\x6b\x83\x05\xee\x02\x6c\x16\x68\x51\xf4\x03\x25\x8d\x24\x22\x14\xa9\x92\x54\x6c\xd7\xf0\x7f\x2f\x48\xbd\x51\xb6\xe3\x18\x5b\x5c\xbe\x58\x22\x47\xf3\xfa\x3c\x43\x4e\xb6\x5b\x8d\x45\xc9\x88\x46\xf0\x72\x24\x09\x4a\x5f\x17\x25\xf3\xc0\xdf\xed\x2e\x2e\x42\xa5\x37\x0c\x97\x17\x00\x00\xbe\x26\x11\xa3\xfc\x05\xb6\xf6\xd5\xfc\x45\x24\x7e\xc9\xa4\xa8\x78\x32\x89\x05\x13\x32\x80\xcb\xbb\xbb\xbb\x79\xb7\xdf\x2c\xae\x72\xaa\xb1\x5f\x4d\x99\x20\x3a\x00\x86\xa9\xee\x17\x23\x21\x13\x94\x01\x70\xc1\x1d\x51\x51\x69\x46\x39\xee\x2f\xc7\x95\x54\x46\x71\x29\x28\xd7\x28\xfb\x8d\x92\x24\x09\xe5\x59\x00\x37\xb3\x72\xdd\x2f\x17\x44\x66\x94\x07\x30\x2b\xd7\x70\x57\xae\xed\xef\x40\x20\x15\x5c\x4f\x14\xfd\x2f\x06\x70\x73\xef\x6e\xac\x68\xa2\xf3\x00\x6e\x3e\xfd\x58\xaf\xed\x06\xa9\x08\x72\xf1\x8a\xf2\x74\x42\xee\xef\xef\xf7\x3f\x8d\x05\xd7\xc8\xb5\xf3\xdd\xf1\x44\x25\x54\x95\x8c\x6c\xf6\xa3\xef\x82\xbc\x1d\xc4\xd0\x66\xf0\x63\xb9\x06\x25\x18\x4d\xf6\x8a\x51\x50\x3e\xc9\x91\x66\xb9\x0e\xe0\x6e\x36\xf8\x34\x66\x48\x64\x00\x91\xd0\x79\xeb\x6b\x38\x6d\x4a\x7f\x11\x26\xf4\xb5\x46\x40\x18\x55\x5a\x0b\x0e\x34\x59\x78\x9a\x44\x7f\x17\x05\x7a\x10\x33\xa2\x94\x7d\x37\x19\xf1\x40\xf0\x98\xd1\xf8\x65\xe1\x89\x12\xf9\x13\xc9\x70\x3c\x32\x82\xcf\x24\x1a\x5d\x83\xce\xa9\xba\x9a\x33\x11\x13\x4d\x05\xf7\x73\x89\xe9\x62\x74\x69\xf6\x47\x73\x6f\x69\x7e\xc3\x69\x6d\xe4\xa8\xc5\x07\xa2\x31\x13\x92\xa2\x3a\xcb\x6e\x2f\x7e\xca\x7a\x2f\x65\x7c\xe8\xdf\x4e\x7a\xf2\xad\x64\x82\x24\x67\x79\x51\x8b\x9e\xf2\xa0\x96\x30\xd6\xeb\x27\xf8\x15\x57\x27\xad\x7f\x7e\x7a\xfc\x82\x9b\xb3\xac\xd7\xa2\xa7\xac\xd7\x12\xc6\xfa\xe7\xa7\x47\xf8\x82\x9b\xde\x74\x38\xb5\xb5\xb7\x10\xb0\xb6\x9b\x4a\xba\x96\x1b\x30\x7b\xb5\xa3\xdb\xad\x24\x3c\x43\xf8\xe1\x05\x37\xd7\xf0\xc3\x2b\x61\x15\x42\xb0\x00\xff\x67\xa2\x89\xdf\x27\x77\xb7\xeb\xc0\x17\xe6\x1f\x97\xdb\xad\xf9\x60\xb7\x0b\xa7\xf9\xc7\x65\xb7\xd3\x49\x18\xf3\x16\x8e\x0b\xcf\x10\x2e\x65\x62\x15\x40\x4e\x93\x04\xf9\xbc\x26\xcd\x44\x95\x24\xb6\x7d\x62\x25\x49\x39\x87\x52\x28\x6a\x82\x0c\x40\x22\x23\x9a\xbe\xe2\xdc\x3b\xad\xb8\xff\x82\x44\x4a\xb0\x4a\xe3\x1c\x64\xcd\x16\xc3\x95\xb6\x13\xdc\xd9\x97\x96\x47\x37\xb3\xd9\x8f\x73\x87\xf5\x01\x98\x76\x45\xe4\x24\x93\x24\xa1\xc8\xf5\xf8\xcf\xb3\x04\xb3\x6b\xd0\x92\x70\x55\x12\x89\x5c\x5f\xc3\xe5\x6c\x36\xbb\x9a\x7b\xcb\x26\xbd\xfb\x5e\x75\x39\xb4\xd9\x73\x32\x35\x70\x9d\x80\xad\x9f\x37\xcd\x68\xaa\xa6\xdb\xad\xff\xf8\xf3\x6e\xe7\x2d\x43\x5a\x64\xa0\x64\x5c\x6f\xb4\xeb\x7e\x46\x53\xaf\x8d\xb4\x20\xeb\x49\x13\xcd\xad\xed\x03\x60\x56\xda\x90\xba\x25\xd3\x32\x27\x4d\x06\xee\xca\xb5\xf5\x97\x1c\xf3\x16\x79\xb2\xe7\xa4\x13\x58\xbb\x7d\x00\xa5\x01\x39\xdf\x06\x54\x98\x0a\x59\xb4\xbb\x11\x51\x34\x36\x0b\x1e\x90\xd8\x14\x6b\xe1\x4d\x2b\x85\x72\x4a\x92\x24\xae\xf5\x6d\x3c\x28\x50\xe7\x22\xb1\x15\xd5\x4e\xd1\xc3\xfc\x76\xf9\x39\x49\x80\x00\xc7\x15\xb4\xe2\x7f\x08\xa7\xf9\xad\x23\x64\xdc\x6b\xac\x49\xb1\xf2\x96\xee\x42\x4a\x91\x25\x5d\x8b\xd8\xd4\x21\x85\x94\x97\x95\x06\xbd\x29\x71\xe1\x69\x5c\x6b\x0f\x38\x29\x70\xe1\x75\x0e\xed\xd7\xf9\xd0\x86\xab\x42\x55\x51\x41\xb5\x07\xb6\xf8\x0b\xef\xab\x7d\x85\x3f\x5e\x7e\xfa\xd3\xa7\xdb\x21\x66\xc2\xa9\x49\x45\xf3\x9c\xdf\x2e\xff\x29\x2a\x09\x7d\x5a\x83\x9e\x4c\x2d\xa2\xfc\x6f\x0a\xe5\x1b\x2c\xa4\x7d\x09\x32\xaf\xc5\x96\x42\x22\xe3\xfc\x2f\xff\x59\x68\x92\x05\xdb\xad\x6f\x10\x66\x7f\x86\xe5\xb6\x35\x1a\x56\x44\x62\x21\x5e\xf1\xb0\x28\x4f\xbf\x7d\x7d\x76\x90\xd8\x9d\xcb\x47\xe9\xe9\xe6\xa5\xa6\xfb\x41\x72\xdb\x3c\x35\xce\x1d\xea\x68\x3a\xe7\x20\xb9\xcb\x7f\x0c\xdb\x6b\x8f\xdb\x3e\xa3\xf5\x3b\x7d\x07\xc5\x5d\x73\xff\x7f\x11\x5c\x35\x07\xca\x00\xbc\x80\x3c\xae\x1d\x2f\x2a\xa6\x69\x49\xa4\xb6\x1e\x4e\x12\xa2\x89\x0b\xed\xb7\x13\x25\x89\xa6\x3c\xeb\xd2\x94\xed\x11\xa2\x39\x70\x6a\x4e\x64\x34\xfd\x0e\x3a\x3c\x53\xcd\xf0\x1d\x2e\x68\x23\x73\x06\x11\x8e\x68\x27\x99\x7a\x4f\x39\xc9\xd4\xfb\xba\x87\x85\x3e\x34\x64\x5a\x83\x16\xe0\x9e\xff\x07\xdd\xf9\x0c\x1e\x75\x16\x18\x89\x90\xb9\x9c\x3a\xc4\xe6\x41\xe9\xe2\x1c\xe3\x97\x48\xac\x1d\x94\x4f\x6a\x60\x1f\x63\x4c\x7f\x11\x74\x0f\x28\x52\x69\x31\xf7\x96\x70\x48\xd3\x1a\xd1\xd6\xaf\xfd\xb0\x86\x2d\xfc\x3b\xca\xf4\x0b\x3d\x8a\x81\x94\x32\x6c\x83\xa9\x11\xfe\x8b\x5d\x31\xd4\x71\xde\x7f\xef\x0e\x79\x40\xdb\xee\x56\x74\x82\xb6\xe5\xf2\x39\xa7\x0a\xa8\x82\x8d\x69\xac\xcd\xe5\xc8\x07\x78\xd4\x10\x13\x0e\x11\x42\xa5\xd0\x82\x86\xc4\x31\x2a\x05\x3a\x47\x23\xe6\x87\xd3\xb2\xd3\xf1\x05\xb1\x04\xaa\x41\x61\x2c\x51\xfb\x00\xdd\x02\x49\xd1\x91\x3c\xde\x24\xce\x4d\xc9\x21\xd7\x9c\xbe\x68\xc1\x5a\x47\x6c\x90\x64\xe2\x27\x25\x7d\xc1\x8d\x15\xf1\x40\x22\x49\x04\x67\x9b\x85\xd7\x3e\x9d\x51\x90\xa6\xad\x0a\xfe\x50\xdf\x38\x63\x51\x6e\x9e\x71\xad\xc7\x23\x47\xf7\xe8\x6a\x2e\x51\x57\x92\x43\x4a\x98\x32\x57\xb0\x07\x51\x6e\xda\x54\x3a\x67\x4e\x4a\x20\x25\x13\xa3\xc3\x98\xa6\xcb\xae\x3d\x9f\x2a\xa6\x8a\x25\x2d\x75\xbd\x99\x56\xdc\xb6\x53\xe8\x6e\xbe\x25\xc9\xf0\x57\x52\xe0\x35\x20\x2b\xb8\xbe\x72\xa6\xad\x57\x22\x81\x5e\x43\x5f\x72\xfb\x6c\xae\xd0\xaa\x9f\x89\x9c\x39\x6d\x01\x89\x88\xab\x02\xb9\xf6\x33\xd4\x7f\x65\x68\x1e\xd5\x4f\x9b\x07\xe3\xbc\xb1\x31\x76\xe1\x73\xe5\x8e\x95\x12\xc6\x14\x16\x30\x9b\x03\x85\xd0\xd1\xe9\x33\xe4\x99\xce\xe7\x40\x3f\x7c\x70\x5d\x1b\x5a\xfe\x17\xfd\xb7\x6f\x89\xef\x37\xa3\x20\x2c\xc0\x33\xc3\xa0\xd7\xdb\xd8\xb9\x1e\xdb\x18\xce\xf4\xd7\x8e\x0c\xa7\x9d\xb5\xea\xde\x71\xd5\xca\xf4\x8e\xf6\x77\xe1\x07\x33\xd3\x1a\x87\x8f\x3a\x7b\xc4\xc3\x9f\x36\x8f\x49\x57\xb6\xab\xc3\xc0\x23\x26\xe2\x17\x47\x99\xad\xeb\x9b\x66\x47\xc2\xf6\xea\x91\x3b\x7c\x4f\xa7\xf0\x37\xd4\x96\xa7\x58\xdb\x84\x15\xd5\xb9\xa5\x44\x82\x29\xa9\x98\xfe\xad\x34\x67\x27\xe1\x09\xd8\x51\x0a\xcc\xd0\xa5\x2f\x5a\xd4\x68\x12\xc1\x02\x56\x94\x27\x62\xe5\xf7\x83\x14\x51\xb9\xaf\xaa\x48\x69\x39\xbe\xb9\xb2\xc2\x34\x85\xb1\x26\x91\x9b\xaf\xb7\x22\x36\xb5\xf0\x3e\x18\x61\xdf\x9a\x1c\xd7\x1a\x76\x80\x4c\xe1\x99\xdf\xdb\x51\xbc\xfb\x7e\x0e\xfd\x10\xdf\x70\xe4\xc2\xfd\x3f\x4f\x2a\x84\x1e\xfc\x9f\xe7\x7f\x01\x00\x00\xff\xff\x25\x94\xa2\x19\x05\x12\x00\x00")

func userTmplBytes() ([]byte, error) {
	return bindataRead(
		_userTmpl,
		"user.tmpl",
	)
}

func userTmpl() (*asset, error) {
	bytes, err := userTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "user.tmpl", size: 4613, mode: os.FileMode(438), modTime: time.Unix(1580311294, 0)}
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
	".gitattributes": Gitattributes,
	"bindata.go": bindataGo,
	"error.tmpl": errorTmpl,
	"footer.tmpl": footerTmpl,
	"gif.tmpl": gifTmpl,
	"header.tmpl": headerTmpl,
	"index.tmpl": indexTmpl,
	"search.tmpl": searchTmpl,
	"templates.go": templatesGo,
	"user.tmpl": userTmpl,
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
	".gitattributes": &bintree{Gitattributes, map[string]*bintree{}},
	"bindata.go": &bintree{bindataGo, map[string]*bintree{}},
	"error.tmpl": &bintree{errorTmpl, map[string]*bintree{}},
	"footer.tmpl": &bintree{footerTmpl, map[string]*bintree{}},
	"gif.tmpl": &bintree{gifTmpl, map[string]*bintree{}},
	"header.tmpl": &bintree{headerTmpl, map[string]*bintree{}},
	"index.tmpl": &bintree{indexTmpl, map[string]*bintree{}},
	"search.tmpl": &bintree{searchTmpl, map[string]*bintree{}},
	"templates.go": &bintree{templatesGo, map[string]*bintree{}},
	"user.tmpl": &bintree{userTmpl, map[string]*bintree{}},
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

