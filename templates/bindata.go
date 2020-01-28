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

	info := bindataFileInfo{name: "bindata.go", size: 0, mode: os.FileMode(438), modTime: time.Unix(1580254231, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _errorTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xaa\xae\x2e\x49\xcd\x2d\xc8\x49\x2c\x49\x55\x50\xca\x48\x4d\x4c\x49\x2d\xd2\x2b\xc9\x2d\xc8\x51\x52\xd0\xab\xad\xe5\xe2\x0a\xc9\xc8\x2c\x56\xc8\x2c\x56\x48\xcc\x53\x48\x2d\x2a\xca\x2f\x52\x28\x48\x4c\x4f\x55\xe4\xb2\x49\x2a\xb2\xe3\xaa\xae\xd6\xf3\x2d\x4e\x07\xa9\x42\x36\x23\x2d\x3f\xbf\x04\xc5\x0c\x40\x00\x00\x00\xff\xff\x1b\xf6\x52\x38\x61\x00\x00\x00")

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

	info := bindataFileInfo{name: "error.tmpl", size: 97, mode: os.FileMode(438), modTime: time.Unix(1580235932, 0)}
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

var _gifTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\xd1\x6e\xab\x46\x10\x7d\xf7\x57\x8c\xb6\x51\xd5\x4a\x35\x54\x51\x9f\x30\xd0\x87\x46\xa9\xf2\x94\xaa\x49\x3e\x60\x60\x07\x58\x05\x76\xb9\xcb\x60\xc7\x42\xfc\xfb\x15\x18\x30\x8e\x1d\x27\x8a\xa5\x28\x91\x22\xec\xf5\xce\xd9\x33\x67\xe6\xcc\xd2\x34\x4c\x45\x99\x23\x13\x88\x8c\x50\x92\x75\xb8\x28\x73\x01\x4e\xdb\x2e\x16\xbe\x54\x6b\xa8\x78\x9b\x53\x20\x0a\xb4\xa9\xd2\x4b\x36\xa5\x07\x7f\xfd\x59\xbe\xac\x44\xb8\x00\x00\xf0\xb3\xeb\x71\x4b\x92\x1b\x64\x0f\x72\x4a\x78\x05\xbb\xfd\x1e\xec\xb6\x36\x8d\x73\x83\x8c\xce\xa3\xe2\x9c\xda\xd6\x77\xb3\xeb\x21\x7c\x76\xc4\x10\x6f\x55\x9a\xf1\x08\xdf\xfd\x35\x8d\x45\x9d\x12\x0c\x10\x98\x56\x6d\x3b\xfd\xd8\x83\x20\xc4\x39\x56\x55\x20\x18\x53\x01\x99\xa5\x24\x10\x6e\x45\x68\xe3\xec\xef\x1f\x01\x63\xea\x35\x8d\xd3\xb6\x22\xfc\xa5\x7f\xfa\x2e\xce\xd1\x49\xcb\x01\xd0\x77\xa5\x5a\x87\x8b\xf1\x31\xe3\xc6\xf4\xc2\x4b\xcc\x55\xaa\x3d\x88\x49\x33\xd9\x55\x9c\x13\x5a\x0f\x22\xc3\x59\xc7\xd6\x57\x45\x0a\x95\x8d\x03\xe1\x5a\xdc\xa4\x2a\xa9\xdc\x31\xeb\xbb\x9b\xb6\x75\x52\x95\x88\x09\x7a\xe1\x47\x36\xec\xfe\x17\x8b\xa6\x51\x09\x38\x4f\x15\xd9\x4e\xf2\x1d\xa3\x6e\xa5\x8f\xec\x96\xef\x74\x62\x9c\x5b\x5c\x1b\xab\x98\xe4\xb8\xa9\xa7\x5b\x86\x87\x3a\x24\xc6\x16\xa3\x14\x11\x56\x2a\xee\x16\x04\x60\xcc\xca\xe8\x40\xb8\x75\x45\xd6\xad\x4b\x89\x4c\x31\x32\xa5\xc6\x2a\xaa\x04\x14\xc4\x99\x91\x81\xf8\xef\xfe\xe1\x51\x1c\x42\xf6\xb0\x4a\x97\x35\x03\x6f\x4b\x0a\x44\xa6\xa4\x24\x2d\x40\x63\x41\x81\x48\x55\xa2\xa4\x80\x35\xe6\x35\x05\x62\x9e\xef\x09\x9c\xb1\x8e\x57\xcf\xb4\xfd\x03\xae\xfa\x20\xf0\x82\xd7\xa9\xfe\x33\x31\x7b\x55\xe6\x89\x4f\x8e\x11\xe5\xf3\x92\x1f\x9f\xb5\x3f\x53\x25\xc3\x51\x6f\xa0\x9d\xcc\x32\xce\x28\x7e\x8e\xcc\xcb\x98\x67\x8c\xbc\x6c\x9a\x8e\x78\xdb\x8a\x43\x4b\xec\x5a\x1c\x4a\x94\x52\xe9\x74\xf8\xb6\x51\x92\x33\x0f\xb0\x66\xb3\x9a\xe4\x31\x5a\x40\x0f\x4c\x32\x84\x11\xed\x0c\x73\xca\xab\x6f\x41\xfa\x63\x64\x27\x13\x1d\x71\x74\xfb\x7a\x9d\xea\x87\xd3\x41\xbd\xef\x86\xea\x5a\xb3\x11\xa1\x1f\xd5\xcc\x46\x0f\x69\x56\x75\x54\x28\x16\xe1\x53\xdf\xc7\xb0\x6f\x17\xf0\xd5\x18\x96\x20\x24\xb8\x24\xd9\xed\xf3\x5d\x15\xfa\xee\x0e\x22\x1c\xfc\x77\xc8\xaf\x73\xc9\x7e\xcd\x77\x67\xb6\xfa\xa4\xc5\x74\x32\xf8\xf5\xcb\xcc\xf5\x41\xd1\xfe\xa7\xc2\xac\x09\x6e\xad\x29\x60\x9c\x29\xc7\xc2\x65\x84\x96\x97\xe6\x93\xda\x0d\x43\x6c\x68\xdf\x0b\xa5\xfc\x72\x21\xe1\xf4\x84\xf9\xb5\x77\xd7\x0a\x46\x3e\xd5\x9b\x6d\x0d\xf3\x2b\xab\x9b\x6a\xef\x4f\xb4\x37\xce\xbc\xd4\xfa\xce\xe5\xbe\x77\xce\x31\x3e\xaf\xc0\x45\xe6\x7e\xc0\x35\xc1\xa3\x79\xaf\x47\x2f\xec\x50\x2d\xcf\xdd\xb9\xf7\x1b\xdd\x5f\xca\x53\xe8\xa7\x1a\x58\x52\x4e\xdf\x6f\x0e\xdc\xf4\xac\xe0\xdf\xbb\x5b\xf8\x6d\x6b\x6a\x40\x4b\xc0\x19\x81\xe9\x72\xfe\xfd\x48\xec\x08\xf5\x05\x83\x74\x52\x7a\xf6\x61\xff\xde\x99\x18\xc3\x07\xef\x9d\x3f\x03\x00\x00\xff\xff\x34\xd0\x46\x7e\x95\x0a\x00\x00")

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

	info := bindataFileInfo{name: "gif.tmpl", size: 2709, mode: os.FileMode(438), modTime: time.Unix(1580235937, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _headerTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xec\x58\x4b\x93\xdb\xb8\x11\xbe\xeb\x57\x20\x9c\x4a\x3c\x76\x49\xe4\x68\x3c\xf2\x66\x29\x4a\x55\x9b\xc9\xc6\x76\x55\x0e\x9b\xf5\x6e\x55\xaa\x52\x7b\x68\x02\x4d\x12\x1e\x10\x60\x00\x50\x8f\x4c\xcd\x7f\x4f\x81\xa4\x28\x3e\x44\xd9\xde\x4b\x2e\xd1\xc1\x96\xf0\xe8\xaf\xd1\xdf\xd7\xdd\xc0\x44\x99\xcd\xc5\x76\x16\x65\x08\x6c\x3b\x8b\x2c\xb7\x02\xb7\xef\x79\x91\x21\x94\x87\x28\xa8\x7f\xcf\x22\xc1\xe5\x13\xd1\x28\x36\x9e\xb1\x47\x81\x26\x43\xb4\x1e\xc9\x34\x26\x1b\x2f\xb3\xb6\x30\x61\x10\x50\x26\x3f\x1b\x9f\x0a\x55\xb2\x44\x80\x46\x9f\xaa\x3c\x80\xcf\x70\x08\x04\x8f\x4d\x90\x28\x69\x17\xb0\x47\xa3\x72\x0c\x1e\xfc\xef\xfc\xbb\x80\x9a\xfe\xb0\x9f\x73\xe9\x53\x63\xbc\xed\x2c\xaa\x70\xb6\x33\x42\x08\x89\x15\x3b\x92\xe7\xea\xab\xfb\xe4\xa0\x53\x2e\x43\x72\xb7\x6e\x87\x2a\x23\x09\xe4\x5c\x1c\x43\xf2\x83\xe6\x20\xe6\xe4\x03\x8a\x1d\x5a\x4e\x61\x4e\x0c\x48\xb3\x30\xa8\x79\x72\xde\x12\x03\x7d\x4a\xb5\x2a\x25\x5b\x50\x25\x94\x0e\x49\x2c\x80\x3e\x9d\x17\x34\xa3\xfb\x8c\x5b\x3c\x8f\x5a\x3c\xd8\x05\x08\x9e\xca\x90\x50\x94\x16\x75\x3d\xf7\x52\xfd\x7b\x43\x95\xb4\x28\xad\xe9\xb8\xcb\xb8\x29\x04\x1c\x43\xc2\xa5\xe0\x12\x17\xb1\x50\x5d\x98\x3d\x67\x36\x0b\xc9\xf2\xee\xfe\xa1\x38\x5c\xc6\x11\x98\xd8\xf3\x4c\xa1\x0c\xb7\x5c\xc9\xd0\xf1\x01\x96\xef\xb0\xe7\x81\x55\x85\x84\x5d\x07\xbf\x00\xc6\xb8\x4c\x1d\xc4\x09\xa0\xbf\x14\x3a\x8b\x2b\x58\x86\x54\x69\xa8\x31\xa4\x92\x78\x2d\x26\x03\x53\x7e\xa2\x74\x5e\x89\x65\xc8\xd7\xc2\xaa\xa2\xeb\x43\x67\xc6\x9d\x2f\x24\xbd\xe3\x27\x42\x81\x0d\x89\xe6\x69\xd6\x3d\xfa\xe9\x28\xef\xba\x6b\x2f\xd3\x37\xe6\x97\x81\x7e\x4a\x35\x1c\x2f\x39\x9e\x2d\x3b\x0e\x57\x6a\x32\xfc\x3f\x18\x92\x87\x55\x17\x69\x40\x65\xcf\x90\x50\xa9\xfa\x1a\x1b\x67\xbf\x42\xb2\xd8\x63\xfc\xc4\xed\xc2\x19\x03\xbd\x48\x35\x30\x8e\xd2\xde\x3e\xac\x18\xa6\x73\x42\x8f\x20\xe7\x24\xd5\x88\x72\x4e\x8e\x28\x84\xda\xcf\x49\x0e\x29\x4a\x0b\xaf\xcf\x16\x4f\x56\xba\x27\x16\xbc\x08\x2b\x32\xc7\xcb\x2a\x8a\x13\x2e\xc4\x29\x2e\x56\x83\x34\x05\x68\x94\xb6\x77\x22\x83\xa0\x69\xe6\x14\x0d\x5c\xa2\x9e\x62\x74\x35\x50\xd5\x70\x1b\x97\x45\x69\x3b\x9b\x1b\xc1\x7f\xbf\xea\x49\x21\x43\x47\x75\x48\x1e\x7a\xa3\x2d\xe1\xf7\xbd\xe1\x58\x69\x86\x7a\x28\xce\x4e\xc8\x97\xdf\x7d\xc1\xa9\xb8\xb4\x56\xc9\xb1\x57\x5f\xe3\x54\x97\xc1\x1b\xc6\xd8\x45\x0f\xee\xef\xbf\xec\x30\x2d\xb5\x71\x04\x14\x8a\x8f\x0a\xc9\x65\x87\xc3\x4c\xed\x7a\x4c\xf4\x7c\xa1\x94\x76\x6d\xf8\x4d\x31\x5a\xd4\xe8\xed\xae\x8e\x4e\xb8\x53\x53\x48\xae\xea\xaf\x28\x75\x21\x70\x4e\x76\x5c\x09\xb4\xad\xee\x4e\xb1\x19\x56\x15\x3f\x06\xc3\xa9\x2b\x03\xe4\xb9\xcd\x97\xba\xe6\x5d\xc8\xc9\x9b\xd5\x6a\xb5\xee\x94\x28\x27\xa6\x73\x81\x2f\x0e\xeb\x91\xd1\xec\xfe\x42\x2b\xa8\x9c\x98\xac\xcd\x43\xc9\x4e\x7b\xec\x6b\xb5\x27\x7e\xc2\x51\xb0\x09\x14\xb7\x99\xac\x8a\x03\xb9\x9f\x34\x52\x09\x7e\xde\x1d\x19\xa9\xad\x3d\xf0\x6a\x5c\x0e\x3b\xfe\x91\x6e\x7f\xb8\xfb\xe3\x50\x4f\x0b\xc7\x57\x69\x06\xa5\xf3\xa4\xb5\x9e\x95\x65\x71\x20\x46\x09\xce\xc8\xcd\x72\xb9\xbc\xe6\xf8\xbf\xec\xb1\xc0\x8d\x29\xe3\x9c\xdb\xdf\xe6\x13\xd3\xf5\x81\x7e\xbb\x7e\xc8\x4b\x6c\xc7\xef\xd6\xfd\x26\x52\xe7\xcc\xbe\x11\x53\xac\x04\x5b\xf7\x12\xd9\x5f\x62\xbe\x1e\x65\x4a\x8f\xce\xb7\x83\x5c\x17\x2a\xe5\xd2\xb9\x34\x27\x37\x1a\x53\x6e\x2c\xea\x5a\x8f\xe3\x42\x3e\xa8\x20\x17\xbb\xce\x29\xa0\xe7\x18\x0e\xee\x04\xe7\x8e\x0c\xb1\x51\xa2\xec\xdd\x17\xc6\x5d\xaf\x61\xf4\xdd\x5d\x6f\x54\xd7\x11\x18\x6a\xd3\x42\x7a\x3d\xa4\x6f\xff\xfc\x78\xa1\x5e\xae\xc6\x8a\x68\xd5\x72\x51\x73\x6f\x47\x97\x8f\x6f\xb9\x05\x5c\x0e\xd3\x0d\x00\x0c\x8b\xa3\xbb\xd5\x85\x84\x5b\x10\x9c\x4e\x36\xd6\xd1\x1d\xc9\x21\x2d\x4c\x01\x14\x9d\x2f\x7b\x0d\xc5\x28\x48\x53\x0c\x5f\xb6\x38\x4a\xb5\xa1\xa5\xc9\x94\x9d\xbe\xc0\x2c\x27\xba\xc4\x64\x9f\x25\x53\xd9\xda\x6c\x10\x4e\x12\xb1\x28\xbf\xae\x65\x0c\x7d\x1f\xf7\x8a\x2b\x50\x03\x2e\x27\x60\xa2\xa0\xb9\x95\xcf\x22\x43\x35\x2f\x6c\x7d\x3f\x4f\x4a\x49\x9d\x4e\x08\x55\xc5\xf1\x17\x3c\xd8\x5b\xce\x5e\x77\x80\x83\x37\xe4\x3d\x5a\x62\x33\xac\x94\x45\xea\xfa\xfa\x26\x68\x17\xec\x40\xb7\x7b\xc9\x86\x30\x45\xcb\x1c\xa5\xf5\x53\xb4\x3f\x0a\x74\x5f\xff\x72\xfc\xc8\x9c\xd5\xf5\x6c\xd6\xda\xfc\x84\x02\xe9\xb4\xd9\x93\x41\xdf\x54\xeb\x6e\xdb\xce\xd5\x99\xb0\xb5\x0d\xae\xe4\xcf\x20\x53\xbc\xbd\x9b\x93\xef\xdd\xe7\xf5\x9a\x04\x6f\xfe\xa6\x34\xc9\x55\xcc\x05\x12\x86\x3b\x4e\xd1\xbc\x09\x3a\xf0\x8f\xaa\x38\x9e\xc1\xb9\x34\x9c\xe1\xa4\x33\xed\x91\xf0\x80\xf4\x51\xe5\x39\x48\x76\xeb\x39\x4f\xbc\xd6\x2f\x8d\xb6\xd4\x92\x24\x20\x0c\x76\x22\xde\x04\x7a\x16\x05\xcd\x1b\xcd\x3d\x87\xb6\xb3\x88\xf1\x1d\xe1\x6c\xe3\x9d\x9e\x1c\x5e\x4d\x46\x3b\xde\x96\xc1\x66\xa2\x9a\xac\xf4\x41\x05\x18\xb3\xf1\xda\xd2\xed\x11\xa8\x62\xb0\xf1\x82\x6a\x93\x47\x72\xb4\x99\x62\x1b\xaf\x50\xc6\x76\xb6\x57\x26\xb2\xfb\xed\xdf\x55\x4a\x3e\xca\x28\xc8\xee\x07\x73\x0e\xbb\xb1\xae\xd5\xde\xdb\x76\x07\xaa\x88\x78\xdb\x5f\x0d\x6a\x09\x39\x46\x01\xe3\xbb\x6d\x54\xdf\x0e\xab\xa6\xe2\xb9\xc0\x79\xc4\x4d\x6e\xbc\xb2\x59\xe6\x6d\xeb\x85\xdf\x8c\xf3\x13\x18\xb3\x57\x9a\x7d\x01\xa7\x68\x96\xfd\x6e\x9c\x4f\x55\xa7\xbc\x80\x52\xb7\x50\x8f\xec\x40\x94\xb8\xf1\xea\x98\x91\x3f\x69\xd0\x7a\x7d\x42\x6f\xd6\x0c\xb1\xa3\xc0\xd1\xd2\xf0\x79\x9e\x6a\xa9\xed\x76\xb5\x6f\x61\xf7\xb4\xef\x8b\x04\x3f\x6a\x04\x8b\xc4\x71\xf5\x7f\x96\xff\x57\x2c\xd7\x4f\xd3\x2e\xbf\xd9\x72\x1b\x41\xf3\x57\x97\xc0\xdb\x46\xfc\x94\xe8\xca\x6b\xfd\x05\x92\xc0\x82\x42\x8e\x1a\x16\x1a\xad\x56\x0e\x97\x6f\xc9\xfb\x8f\x3f\x7d\xf8\xe1\xd7\x7f\x46\x01\x6c\xa3\x20\x5b\x9e\xad\x3e\x3f\xf3\x84\xf8\x8e\xb0\x97\x97\x7e\x5c\x5a\xa8\xf8\x88\x67\x80\xe6\x81\xef\xb9\x2a\xa0\x4a\xeb\x0c\x4e\x6c\x73\xec\x06\x17\x36\x7e\x50\x39\x86\xe4\xf9\xb9\x02\xf5\x4f\x52\x79\x79\xe9\x99\x7a\x7e\x46\x61\x70\xec\xd2\xd0\x1a\x51\xf2\x51\x70\xfa\xb4\xf1\x3e\xc3\x0e\xea\x6a\x19\x0e\x1a\xc7\xab\xb6\x16\xbe\x7a\xed\x57\x2d\xcc\x6f\xee\x07\x9b\x57\xd5\xc5\xe0\xd5\xda\x6b\x8b\xda\xf8\x38\xbf\x0b\xb2\x9b\xa3\x57\x50\x7f\x6e\x96\x0d\xcf\x2e\x59\x73\xf4\xa1\x32\x1a\x67\xfa\x6f\xbb\x9e\xb6\x5a\x01\x0d\xde\x90\xa3\x4a\xd1\xd6\x85\x7a\xe1\xb9\x2a\xbc\xff\xf1\x97\x61\x51\xa8\x35\x5f\xeb\xf9\xdf\x5e\x2f\xc5\x0a\x01\x14\x33\x25\x18\xea\x8d\xf7\xa9\x32\xe5\xee\x21\x24\xe5\x89\x21\x20\x19\x31\xb6\x4c\x92\x3f\xb4\x39\x52\x0b\xee\xaf\x60\xe1\xe5\xa5\xfe\xfe\x8f\x12\xf5\xd1\xfd\xa8\x46\xcf\x3f\xab\x28\x34\xff\x0d\x1d\x6a\x2e\x68\xbd\x2c\x74\x19\xd1\xcb\x83\xe6\x60\x55\x06\x44\x41\xbd\xe5\x6a\x02\xfe\x37\x00\x00\xff\xff\xec\xf7\xf6\xf0\x0f\x15\x00\x00")

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

	info := bindataFileInfo{name: "header.tmpl", size: 5391, mode: os.FileMode(438), modTime: time.Unix(1580235937, 0)}
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

var _searchTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x51\x4d\x6f\x9c\x30\x10\xbd\xf3\x2b\x46\xce\x19\xd8\x50\x71\x31\x2c\x97\x46\xaa\x7a\x6d\xfb\x07\x26\x78\xb0\xad\x1a\xdb\x32\x93\x84\xca\xf2\x7f\xaf\x4a\xc8\x76\xdb\xeb\x3c\xbf\xe7\xf7\x91\x33\xd3\x1a\x1d\x32\x81\x30\x84\x8a\x52\xc3\x6b\x74\x02\x9a\x52\xaa\x0a\x00\x60\x34\xdd\xf4\x9d\x30\xcd\x06\xbe\xd1\xf6\xe2\x78\x93\x90\x73\xf3\x84\x8c\xcd\x8f\xc0\xe8\xce\x6b\x29\xa0\xed\xb2\x8d\xad\xe9\xa6\x77\x66\xce\x09\xbd\x26\x78\x7f\xfb\xc5\x2e\x5b\x29\x07\x70\xc8\x2a\xfb\x0a\x1b\xff\x72\x74\x15\x6f\x56\xb1\x91\xf0\xe9\x72\x89\xfb\x00\xcf\x38\xff\xd4\x29\xbc\x78\x55\xcf\xc1\x85\x24\xe1\xa1\xeb\xba\x01\x98\x76\xae\xd1\x59\xed\x25\xcc\xe4\x99\xd2\x00\x11\x95\xb2\x5e\x4b\x78\x3c\xa8\x2b\x26\x6d\xbd\x84\xfe\xd0\x09\x49\x51\x92\xf0\x18\x77\xd8\x82\xb3\x0a\x1e\xfa\xbe\x1f\x60\x71\x01\x59\x82\xa3\x85\x07\x31\xdd\x0c\x1d\xa6\x10\x4c\xa2\xe5\x2a\xda\x3f\x51\xda\x9c\x9b\xaf\x4f\xa5\x88\x69\xb4\xab\x86\x2d\xcd\x57\xd1\x26\x7c\xbb\xc7\x1a\x6d\x17\xf1\x91\x63\xc5\xbd\x3e\xb3\x74\x97\xd3\xd0\x5e\x1b\xb2\xda\xf0\xc7\x49\x4c\x63\x8b\xff\xfd\x1a\xa7\x9c\x9b\xcf\x18\xd9\x06\x5f\xca\xd8\xc6\xbf\xf8\xd8\x2a\xfb\x3a\x9d\x75\x92\x57\xb7\x55\xee\xea\x9b\x1d\x61\x92\xf0\x1c\xd8\x1c\xf2\x07\xa3\xaa\xee\xa7\x5d\x42\xe0\x7f\xa6\xfd\x1d\x00\x00\xff\xff\x08\x4e\x54\x0d\xf8\x01\x00\x00")

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

	info := bindataFileInfo{name: "search.tmpl", size: 504, mode: os.FileMode(438), modTime: time.Unix(1580235935, 0)}
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

var _userTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xb4\x58\x6d\x6f\xe3\xb8\x11\xfe\x9e\x5f\x31\xd5\xde\xd5\x0e\x36\x96\x9d\x2c\xd2\xdd\xca\x92\x8b\xbd\x5c\xaf\x0d\x16\xb8\x5b\xe0\x72\x40\x8b\xa2\x1f\x28\x69\x24\x11\xa1\x48\x95\xa4\x62\xbb\x86\xff\x7b\x41\xea\x8d\xb2\x1d\xc7\xd8\xe2\xf2\xc5\x12\x39\x9c\xd7\xe7\x19\x71\xb2\xdb\x69\x2c\x2b\x46\x34\x82\x57\x20\x49\x51\xfa\xba\xac\x98\x07\xfe\x7e\x7f\x75\x15\x2a\xbd\x65\xb8\xba\x02\x00\xf0\x35\x89\x19\xe5\xcf\xb0\xb3\xaf\xe6\x2f\x26\xc9\x73\x2e\x45\xcd\xd3\x59\x22\x98\x90\x01\xbc\xbb\xbf\xbf\x5f\xf6\xfb\xed\xe2\xba\xa0\x1a\x87\xd5\x8c\x09\xa2\x03\x60\x98\xe9\x61\x31\x16\x32\x45\x19\x00\x17\xdc\x11\x15\xb5\x66\x94\xe3\xe1\x72\x52\x4b\x65\x14\x57\x82\x72\x8d\x72\xd8\xa8\x48\x9a\x52\x9e\x07\x70\xbb\xa8\x36\xc3\x72\x49\x64\x4e\x79\x00\x8b\x6a\x03\xf7\xd5\xc6\xfe\x8e\x04\x32\xc1\xf5\x4c\xd1\xff\x62\x00\xb7\x1f\xdd\x8d\x35\x4d\x75\x11\xc0\xed\xa7\xef\x9b\xb5\xfd\x28\x15\x41\x21\x5e\x50\x9e\x4f\xc8\xc7\x8f\x1f\x0f\x8f\x26\x82\x6b\xe4\xda\x39\x77\x3a\x51\x29\x55\x15\x23\xdb\xc3\xe8\xfb\x20\xef\x46\x31\x74\x19\xfc\x50\x6d\x40\x09\x46\xd3\x83\x62\x94\x94\xcf\x0a\xa4\x79\xa1\x03\xb8\x5f\x8c\x8e\x26\x0c\x89\x0c\x20\x16\xba\xe8\x7c\x0d\xe7\x6d\xe9\xaf\xc2\x94\xbe\x34\x08\x08\xe3\x5a\x6b\xc1\x81\xa6\x91\xa7\x49\xfc\x77\x51\xa2\x07\x09\x23\x4a\xd9\x77\x93\x11\x0f\x04\x4f\x18\x4d\x9e\x23\x4f\x54\xc8\xbf\x92\x1c\xa7\x13\x23\xf8\x44\xe2\xc9\x0d\xe8\x82\xaa\xeb\x25\x13\x09\xd1\x54\x70\xbf\x90\x98\x45\x93\x77\x66\x7f\xb2\xf4\x56\xe6\x37\x9c\x37\x46\x4e\x5a\x7c\x20\x1a\x73\x21\x29\xaa\x8b\xec\x0e\xe2\xe7\xac\x0f\x52\xc6\x87\xe1\xed\xac\x27\xbf\x55\x4c\x90\xf4\x22\x2f\x1a\xd1\x73\x1e\x34\x12\xc6\x7a\xf3\x04\x3f\xe3\xfa\xac\xf5\xcf\x5f\x1f\xbf\xe0\xf6\x22\xeb\x8d\xe8\x39\xeb\x8d\x84\xb1\xfe\xf9\xeb\x23\x7c\xc1\xed\x60\x3a\x9c\xdb\xda\x5b\x08\x58\xdb\x6d\x25\x5d\xcb\x2d\x98\xbd\xc6\xd1\xdd\x4e\x12\x9e\x23\x7c\xf7\x8c\xdb\x1b\xf8\xee\x85\xb0\x1a\x21\x88\xc0\xff\x91\x68\xe2\x0f\xc9\xdd\xef\x7b\xf0\x85\xc5\x87\xd5\x6e\x67\x0e\xec\xf7\xe1\xbc\xf8\xb0\xea\x77\x7a\x09\x63\xde\xc2\x31\xf2\x0c\xe1\x32\x26\xd6\x01\x14\x34\x4d\x91\x2f\x1b\xd2\xcc\x54\x45\x12\xdb\x27\xd6\x92\x54\x4b\xa8\x84\xa2\x26\xc8\x00\x24\x32\xa2\xe9\x0b\x2e\xbd\xf3\x8a\x87\x13\x24\x56\x82\xd5\x1a\x97\x20\x1b\xb6\x18\xae\x74\x9d\xe0\xde\xbe\x74\x3c\xba\x5d\x2c\xbe\x5f\x3a\xac\x0f\xc0\xb4\x2b\x22\x67\xb9\x24\x29\x45\xae\xa7\x7f\x5e\xa4\x98\xdf\x80\x96\x84\xab\x8a\x48\xe4\xfa\x06\xde\x2d\x16\x8b\xeb\xa5\xb7\x6a\xd3\x7b\xe8\x55\x9f\x43\x9b\x3d\x27\x53\x23\xd7\x09\xd8\xfa\x79\xf3\x9c\x66\x6a\xbe\xdb\xf9\x8f\x3f\xee\xf7\xde\x2a\xa4\x65\x0e\x4a\x26\x91\x37\x97\x64\xed\xee\xf9\x39\xcd\xbc\x2e\xda\x92\x6c\x66\x6d\x44\x77\xb6\x17\x80\x59\xe9\xc2\xea\x97\x4c\xdb\x9c\xb5\x59\xb8\xaf\x36\xd6\x67\x72\xca\x63\xe4\xe9\x81\xa3\x4e\x70\xdd\xf6\x11\x9c\x46\x04\x7d\x1d\x54\x61\x26\x64\xd9\xed\xc6\x44\xd1\xc4\x2c\x78\x40\x12\x53\xb0\xc8\x9b\xd7\x0a\xe5\x9c\xa4\x69\xd2\xe8\xdb\x7a\x50\xa2\x2e\x44\x6a\xab\xaa\x9d\xc2\x87\xc5\xdd\xea\x73\x9a\x02\x01\x8e\x6b\xe8\xc4\xff\x10\xce\x8b\x3b\x47\xc8\xb8\xd7\x5a\x93\x62\xed\xad\xdc\x85\x8c\x22\x4b\xfb\x36\xb1\x6d\x42\x0a\x29\xaf\x6a\x0d\x7a\x5b\x61\xe4\x69\xdc\x68\x0f\x38\x29\x31\xf2\x7a\x87\x0e\x6b\x7d\x6c\xc3\x55\xa1\xea\xb8\xa4\xda\x03\x0b\x80\xc8\xfb\xd5\xbe\xc2\x1f\xdf\x7d\xfa\xd3\xa7\xbb\x31\x6e\xc2\xb9\x49\x45\xfb\x5c\xdc\xad\xfe\x29\x6a\x09\x43\x5a\x83\x81\x50\x1d\xaa\xfc\xdf\x14\xca\x57\x98\x48\x87\x12\xe4\x5e\x87\x2f\x85\x44\x26\xc5\x5f\xfe\x13\x69\x92\x07\xbb\x9d\x6f\x50\x66\x7f\xc6\xe5\xb6\x35\x1a\x57\x44\x62\x29\x5e\xf0\xb8\x28\x5f\x7f\xf9\xf5\xc9\x41\x62\xff\x6d\x3e\x49\x51\x37\x2f\x0d\xe5\x8f\x92\xdb\xe5\xa9\x75\xee\x58\x47\xdb\x3d\x47\xc9\x5d\xfd\x63\xdc\x62\x07\xdc\x0e\x19\x6d\xde\xe9\x1b\x28\xee\x1b\xfc\xff\x8b\xe0\xba\xfd\xa8\x8c\xc0\x0b\xc8\x93\xc6\xf1\xb2\x66\x9a\x56\x44\x6a\xeb\xe1\x2c\x25\x9a\xb8\xd0\x7e\x3d\x51\x92\x68\xca\xf3\x3e\x4d\xf9\x01\x21\xda\x8f\x4e\xc3\x89\x9c\x66\xdf\x40\x87\x27\xaa\x19\xbe\xc1\x05\x6d\x64\x2e\x20\xc2\x09\xed\x24\x57\x6f\x29\x27\xb9\x7a\x5b\xf7\xb8\xd0\xc7\x86\x4c\x6b\xd0\x02\xdc\x3b\xc0\x51\x87\xbe\x80\x47\xbd\x05\x46\x62\x64\x2e\xa7\x8e\xb1\x79\x54\xba\xa4\xc0\xe4\x39\x16\x1b\x07\xe5\xb3\x06\xd8\xa7\x18\x33\x5c\x06\xdd\x8f\x14\xa9\xb5\x58\x7a\x2b\x38\xa6\x69\x83\x68\xeb\xd7\x61\x58\xe3\x16\xfe\x0d\x65\xfa\x89\x9e\xc4\x40\x46\x19\x76\xc1\x34\x08\xff\xc9\xae\x18\xea\x38\xef\xbf\x77\x87\x3c\xa2\x6d\x7f\x33\x3a\x43\xdb\x6a\xf5\x54\x50\x05\x54\xc1\xd6\x34\xd6\xf6\x82\xe4\x03\x3c\x6a\x48\x08\x87\x18\xa1\x56\x68\x41\x43\x92\x04\x95\x02\x5d\xa0\x11\xf3\xc3\x79\xd5\xeb\xf8\x82\x58\x01\xd5\xa0\x30\x91\xa8\x7d\x80\x7e\x81\x64\xe8\x48\x9e\x6e\x12\x97\xa6\xe4\x98\x6b\x4e\x5f\xb4\x60\x6d\x22\x36\x48\x32\xf1\x93\x8a\x3e\xe3\xd6\x8a\x78\x20\x91\xa4\x82\xb3\x6d\xe4\x75\x4f\x17\x14\xa4\x6d\xab\x82\x3f\x34\xb7\xce\x44\x54\xdb\x27\xdc\xe8\xe9\xc4\xd1\x3d\xb9\x5e\x4a\xd4\xb5\xe4\x90\x11\xa6\xcc\x35\xec\x41\x54\xdb\x2e\x95\xce\x37\x27\x23\x90\x91\x99\xd1\x61\x4c\xd3\x55\xdf\x9e\xcf\x15\x53\x25\x92\x56\xba\xd9\xcc\x6a\x6e\xdb\x29\xf4\xb7\xdf\x8a\xe4\xf8\x33\x29\xf1\x06\x90\x95\x5c\x5f\x3b\x13\xd7\x0b\x91\x40\x6f\x60\x28\xb9\x7d\x36\xd7\x68\x35\xcc\x45\xce\xac\x16\x41\x2a\x92\xba\x44\xae\xfd\x1c\xf5\x5f\x19\x9a\x47\xf5\xc3\xf6\xc1\x38\x6f\x6c\x4c\x5d\xf8\x5c\xbb\xa3\xa5\x84\x29\x85\x08\x16\x4b\xa0\x10\x3a\x3a\x7d\x86\x3c\xd7\xc5\x12\xe8\xfb\xf7\xae\x6b\x63\xcb\xff\xa2\xff\xf6\x2d\xf1\xfd\x76\x1c\x84\x08\x3c\x33\x10\x7a\x83\x8d\xbd\xeb\xb1\x8d\xe1\x42\x7f\xed\xd8\x70\xde\x59\xab\xee\x0d\x57\xad\xcc\xe0\xe8\x70\x1f\x7e\x30\x73\xad\x71\xf8\xa4\xb3\x27\x3c\xfc\x61\xfb\x98\xf6\x65\xbb\x3e\x0e\x3c\x66\x22\x79\x76\x94\xd9\xba\xbe\x6a\x76\x22\x6c\xaf\x9e\xb8\x03\xf8\x7c\x0e\x7f\x43\x6d\x79\x8a\x8d\x4d\x58\x53\x5d\x58\x4a\xa4\x98\x91\x9a\xe9\x5f\x2a\xf3\xed\x24\x3c\x05\x3b\x4e\x81\x19\xbc\xf4\x55\x87\x1a\x4d\x62\x88\x60\x4d\x79\x2a\xd6\xfe\x30\x4c\x11\x55\xf8\xaa\x8e\x95\x96\xd3\xdb\x6b\x2b\x4c\x33\x98\x6a\x12\xbb\xf9\x7a\x2d\x62\x53\x0b\xef\xbd\x11\xf6\xad\xc9\x69\xa3\x61\x0f\xc8\x14\x5e\x78\xde\x8e\xe3\xfd\xf9\x25\x0c\x83\x7c\xcb\x91\x2b\xf7\x7f\x3d\x99\x10\x7a\xf4\xbf\x9e\xff\x05\x00\x00\xff\xff\xd8\x7d\x55\x19\x09\x12\x00\x00")

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

	info := bindataFileInfo{name: "user.tmpl", size: 4617, mode: os.FileMode(438), modTime: time.Unix(1580254215, 0)}
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

