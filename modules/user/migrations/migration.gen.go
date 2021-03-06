// Code generated by go-bindata. DO NOT EDIT.
// sources:
// postgres/20170727210828_users.sql
// postgres/20181228200758_display_name.sql
// postgres/20190817143814_force_change_pass.sql

package migrations


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

func bindataRead(data, name string) ([]byte, error) {
	gz, err := gzip.NewReader(strings.NewReader(data))
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
	info  fileInfoEx
}

type fileInfoEx interface {
	os.FileInfo
	MD5Checksum() string
}

type bindataFileInfo struct {
	name        string
	size        int64
	mode        os.FileMode
	modTime     time.Time
	md5checksum string
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
func (fi bindataFileInfo) MD5Checksum() string {
	return fi.md5checksum
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _bindataPostgres20170727210828userssql = "" +
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x92\x4b\x6f\xdb\x30\x10\x84\xcf\xe6\xaf\xd8\x43\x00\xd9\x28\xe3\x34" +
	"\xe9\xe3\xe2\x4b\x59\x49\xad\x8d\x2a\xb2\xa3\x47\x00\x9f\x04\x5a\xdc\x5a\x1b\xeb\x15\x92\xaa\xea\xfc\xfa\x42\x4e" +
	"\x63\xbb\x40\x03\xb4\x27\x62\x67\x67\x09\xcc\x87\x61\x97\x97\xf0\xa6\xa2\xad\x96\x16\x21\x6d\x87\x31\xbe\x0b\x80" +
	"\x6a\x30\x98\x5b\x6a\x6a\x70\xd2\xd6\x01\x32\x80\x3f\x31\xef\x2c\x2a\xe8\x0b\xac\xc1\x16\x64\xe0\xf9\x6e\x30\x91" +
	"\x01\xd9\xb6\x25\xa1\x62\xcc\x8d\x7c\x91\xf8\x10\xbb\x73\xff\x56\x80\x94\x72\x76\xd4\x12\xf1\x39\xf0\x07\x69\xda" +
	"\x19\xd4\x66\xcc\x46\xa4\x60\x43\x5b\x83\x9a\x64\x09\xe1\x32\x81\x30\x0d\x02\xce\x46\x58\x49\x2a\xe1\x87\xd4\x79" +
	"\x21\xf5\xf9\xa2\x95\xc6\xf4\x8d\x56\x7f\xdb\x19\x2b\x6d\x67\x80\x6a\x7b\x54\xc1\xf3\xbf\x88\x34\x48\xe0\x9a\xb3" +
	"\x51\xae\x51\x5a\x54\x99\xb4\x60\xa9\x42\x63\x65\xd5\x42\x4f\xb6\x38\x8c\xf0\xd4\xd4\x78\xfe\x5d\xd7\xaa\xff\xb1" +
	"\x97\xd2\xd8\xac\x6c\xb6\x54\xbf\x6a\xe7\x6c\xe4\x2e\xc3\x38\x89\xc4\x22\x4c\x06\x0e\xd9\x81\x43\x46\x2a\x6b\x35" +
	"\x55\x52\xef\x61\x15\x2d\x6e\x45\xb4\x86\x6f\xfe\x1a\xc6\xa4\x26\xaf\x9d\x1c\x00\x65\x5d\x4d\x8f\x1d\x42\x1a\x2e" +
	"\xee\x52\x1f\xc6\x07\x71\xc2\xd8\x64\xc6\xd8\x22\x8c\xfd\x28\x81\x45\x98\x2c\x4f\xc4\x7f\x5b\x38\xbc\x60\xe4\xf0" +
	"\x0c\x8d\xc3\x09\x0e\x87\x53\x72\x0e\xa7\x58\x13\xb8\x17\x41\xea\xc7\x6c\x34\x76\x2a\x69\x2c\xea\x4f\x39\xea\xae" +
	"\x44\x59\x4f\x49\x3b\x1c\x9c\x8b\x1b\x79\xf1\xf6\xe3\xc5\x83\x27\xf6\x1f\xae\xdf\xc7\x58\x7d\xed\xdd\x79\xe1\x4d" +
	"\xa7\xbb\xef\x0a\x55\x7f\x45\x1b\xf7\xdd\xd3\xfe\x61\xba\x7b\x5c\x59\x77\xbe\x6c\x44\x7f\xbf\x9e\xbb\x57\x51\xe9" +
	"\x05\xd2\xe1\x70\xc3\xc1\xa9\x9b\xde\xf9\xf3\x19\xc2\x9c\xb7\xd4\x6b\xfa\xfa\xa5\xa7\xc7\x92\x0e\xe2\x3f\xd5\x54" +
	"\x37\x65\x89\x0a\x36\x32\xdf\x31\xe6\x45\xcb\xd5\x59\x51\xc1\x15\xb1\x2b\x3c\x1f\x66\xec\x57\x00\x00\x00\xff\xff" +
	"\x21\x5f\x15\xb1\x18\x03\x00\x00"

func bindataPostgres20170727210828userssqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataPostgres20170727210828userssql,
		"postgres/20170727210828_users.sql",
	)
}



func bindataPostgres20170727210828userssql() (*asset, error) {
	bytes, err := bindataPostgres20170727210828userssqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "postgres/20170727210828_users.sql",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(0),
		modTime: time.Unix(0, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataPostgres20181228200758displaynamesql = "" +
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\x8f\xc1\x4e\xc3\x30\x10\x44\xef\xfe\x8a\xb9\xf5\x80\xc2\x0f\x44\x1c" +
	"\x4c\x1d\x89\x83\x49\x4a\x6a\x73\x45\x4b\xb2\xa2\x2b\x1c\xc7\x8a\x53\x15\xfe\x1e\x05\x04\x2a\x52\x23\xf5\x3a\x9a" +
	"\xd9\xb7\xaf\x28\x70\x33\xc8\xdb\x44\x33\xc3\x27\x55\x14\xd8\x3f\x59\x48\x44\xe6\x6e\x96\x31\x62\xe3\xd3\x06\x92" +
	"\xc1\x1f\xdc\x1d\x67\xee\x71\x3a\x70\xc4\x7c\x90\x8c\x9f\xdd\x52\x92\x0c\x4a\x29\x08\xf7\x4a\x69\xeb\xaa\x16\x4e" +
	"\xdf\xdb\x0a\x44\x74\x7b\xcc\x3c\x65\x68\x63\xb0\x6d\xac\x7f\xac\xd1\x4b\x4e\x81\x3e\x5f\x22\x0d\x8c\x67\xdd\x6e" +
	"\x1f\x74\x5b\x2a\xbf\x33\xda\x9d\x2f\xf6\x95\xfb\x5f\xbd\x03\x0f\x24\xa1\x5c\x23\x7c\xa7\x97\x18\xcb\xa5\xba\x71" +
	"\xa8\xbd\xb5\xa5\x52\xe7\xca\x66\x3c\xc5\x5f\xe9\x3f\xe3\x25\xbc\xca\x79\x1a\x43\xe0\x1e\xaf\xd4\xbd\xaf\x79\x9b" +
	"\xb6\xd9\x5d\x7a\xaa\xfc\x0a\x00\x00\xff\xff\x94\x93\x3f\x15\x79\x01\x00\x00"

func bindataPostgres20181228200758displaynamesqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataPostgres20181228200758displaynamesql,
		"postgres/20181228200758_display_name.sql",
	)
}



func bindataPostgres20181228200758displaynamesql() (*asset, error) {
	bytes, err := bindataPostgres20181228200758displaynamesqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "postgres/20181228200758_display_name.sql",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(0),
		modTime: time.Unix(0, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}

var _bindataPostgres20190817143814forcechangepasssql = "" +
	"\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x8c\xce\xb1\x4e\xc3\x40\x10\x04\xd0\x7e\xbf\x62\xba\x14\xc8\xfc\x40\x2a" +
	"\x83\xd3\x1d\x04\x42\x5c\x47\xcb\x79\x95\x5b\x61\xdf\x9d\xbc\x1b\x19\xf1\xf5\xc8\x20\x10\x05\x42\x94\x33\x9a\x91" +
	"\x5e\xd3\xe0\x6a\xd2\xf3\xcc\x2e\xe8\x2b\x35\x0d\x9e\x1e\x03\x34\xc3\x24\xba\x96\x8c\x4d\x5f\x37\x50\x83\xbc\x4a" +
	"\xbc\xb8\x0c\x58\x92\x64\x78\x52\xc3\xe7\x6f\x1d\xa9\x81\x6b\x1d\x55\x06\x22\x6a\xc3\x71\x77\xc0\xb1\xbd\x09\x3b" +
	"\x30\xf3\xf5\xc5\x64\x36\x02\x80\xb6\xeb\x70\xbb\x0f\xfd\xdd\x3d\x62\xe2\x7c\x96\x53\x65\xb3\x13\x3b\x5c\x27\x31" +
	"\xe7\xa9\x62\x51\x4f\x1f\x11\x6f\x25\xcb\x96\xe8\xa7\xb0\x2b\x4b\xfe\x32\x7e\x03\xd7\xf2\x5f\xc4\xb9\x8c\xa3\x0c" +
	"\x78\xe6\xf8\xf2\x97\xb2\x3b\xec\x1f\x7e\x67\x6e\xe9\x3d\x00\x00\xff\xff\xf4\xff\x40\xfd\x2f\x01\x00\x00"

func bindataPostgres20190817143814forcechangepasssqlBytes() ([]byte, error) {
	return bindataRead(
		_bindataPostgres20190817143814forcechangepasssql,
		"postgres/20190817143814_force_change_pass.sql",
	)
}



func bindataPostgres20190817143814forcechangepasssql() (*asset, error) {
	bytes, err := bindataPostgres20190817143814forcechangepasssqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{
		name: "postgres/20190817143814_force_change_pass.sql",
		size: 0,
		md5checksum: "",
		mode: os.FileMode(0),
		modTime: time.Unix(0, 0),
	}

	a := &asset{bytes: bytes, info: info}

	return a, nil
}


//
// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
//
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
// nolint: deadcode
//
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

//
// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or could not be loaded.
//
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, &os.PathError{Op: "open", Path: name, Err: os.ErrNotExist}
}

//
// AssetNames returns the names of the assets.
// nolint: deadcode
//
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

//
// _bindata is a table, holding each asset generator, mapped to its name.
//
var _bindata = map[string]func() (*asset, error){
	"postgres/20170727210828_users.sql":             bindataPostgres20170727210828userssql,
	"postgres/20181228200758_display_name.sql":      bindataPostgres20181228200758displaynamesql,
	"postgres/20190817143814_force_change_pass.sql": bindataPostgres20190817143814forcechangepasssql,
}

//
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
//
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, &os.PathError{
					Op: "open",
					Path: name,
					Err: os.ErrNotExist,
				}
			}
		}
	}
	if node.Func != nil {
		return nil, &os.PathError{
			Op: "open",
			Path: name,
			Err: os.ErrNotExist,
		}
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

var _bintree = &bintree{Func: nil, Children: map[string]*bintree{
	"postgres": {Func: nil, Children: map[string]*bintree{
		"20170727210828_users.sql": {Func: bindataPostgres20170727210828userssql, Children: map[string]*bintree{}},
		"20181228200758_display_name.sql": {Func: bindataPostgres20181228200758displaynamesql, Children: map[string]*bintree{}},
		"20190817143814_force_change_pass.sql": {Func: bindataPostgres20190817143814forcechangepasssql, Children: map[string]*bintree{}},
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
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
