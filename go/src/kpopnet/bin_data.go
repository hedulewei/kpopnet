// Code generated by go-bindata. DO NOT EDIT.
// sources:
// sql/get_bands.sql
// sql/get_idols.sql
// sql/init_db.sql
// sql/upsert_band.sql
// sql/upsert_idol.sql

package kpopnet

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _get_bandsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x0a\x76\xf5\x71\x75\x0e\x51\xc8\x4c\xd1\x51\x48\x49\x2c\x49\x54\x70\x0b\xf2\xf7\x55\x48\x4a\xcc\x4b\x29\xe6\x02\x04\x00\x00\xff\xff\x14\x21\x0f\x5f\x1b\x00\x00\x00")

func get_bandsSqlBytes() ([]byte, error) {
	return bindataRead(
		_get_bandsSql,
		"get_bands.sql",
	)
}

func get_bandsSql() (*asset, error) {
	bytes, err := get_bandsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "get_bands.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xce, 0xf3, 0x28, 0xfe, 0x9b, 0x8b, 0x77, 0x51, 0x29, 0x48, 0xac, 0xb1, 0xe6, 0x2e, 0x6a, 0xe4, 0xa7, 0x94, 0x74, 0xeb, 0xc4, 0x3b, 0x95, 0xf, 0x22, 0xa7, 0x3c, 0xd8, 0x26, 0xf4, 0xe8, 0xa2}}
	return a, nil
}

var _get_idolsSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x0a\x76\xf5\x71\x75\x0e\x51\xc8\x4c\xd1\x51\x48\x4a\xcc\x4b\x89\x07\x31\x52\x12\x4b\x12\x15\xdc\x82\xfc\x7d\x15\x32\x53\xf2\x73\x8a\xb9\x00\x01\x00\x00\xff\xff\xd9\x11\x11\x34\x24\x00\x00\x00")

func get_idolsSqlBytes() ([]byte, error) {
	return bindataRead(
		_get_idolsSql,
		"get_idols.sql",
	)
}

func get_idolsSql() (*asset, error) {
	bytes, err := get_idolsSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "get_idols.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x38, 0x44, 0xfe, 0xcc, 0x62, 0x3d, 0x28, 0x78, 0xf2, 0xd7, 0x53, 0xfa, 0x71, 0xf0, 0xe5, 0xcb, 0x5c, 0xa8, 0xd2, 0xf, 0x89, 0xcb, 0x9, 0xa4, 0x8f, 0xcb, 0x78, 0xdd, 0x79, 0xfd, 0x28, 0xa9}}
	return a, nil
}

var _init_dbSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\xcf\x31\x6f\x83\x30\x10\x05\xe0\xdd\xbf\xe2\x6d\x80\xd4\x7f\xd0\xa1\x72\xcd\x43\x45\xb8\xa6\x32\xae\x14\xa6\xc8\xc8\x19\x1c\x25\x30\x10\xfe\x7f\x04\x8a\xa5\x2c\x61\xc8\xfa\xee\x4e\xf7\x3d\x65\x29\x1d\xe1\xe4\xb7\x26\xea\x0a\xa6\x75\xe0\xa1\xee\x5c\x87\xc1\x8f\x61\x46\x2e\x80\x18\xb0\x2c\x31\xe0\xcf\xd6\xbf\xd2\xf6\x68\xd8\x7f\x08\x20\xf8\x9b\xc7\x79\x9e\xc6\x61\x3b\x33\xff\x5a\xaf\xb1\xfa\xa1\x6a\x90\x9b\xd6\xe5\xdb\xc6\x17\xb2\x18\xb2\x02\xd2\x94\x48\xc1\xe8\xaf\xa7\xac\x10\xc5\xa7\x10\x3b\x82\x18\xa6\xcb\xae\x60\x25\x1e\xd3\x2c\x19\x60\x59\xd1\xd2\x28\xa6\x0e\xad\x41\x49\x4d\x47\x28\xd9\x29\x59\xf2\x4d\xfd\x73\xf8\x78\xfd\xb2\xd7\x3d\x00\x00\xff\xff\x89\xaa\xc4\xea\x59\x01\x00\x00")

func init_dbSqlBytes() ([]byte, error) {
	return bindataRead(
		_init_dbSql,
		"init_db.sql",
	)
}

func init_dbSql() (*asset, error) {
	bytes, err := init_dbSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "init_db.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x9f, 0x70, 0x22, 0x1b, 0xc, 0xe3, 0x5, 0xb5, 0x31, 0xff, 0x55, 0x56, 0x71, 0x73, 0x78, 0x2f, 0x6f, 0xaf, 0x3d, 0xd8, 0x1e, 0xde, 0x13, 0xf2, 0x3, 0x2, 0xbc, 0xaf, 0x13, 0xb8, 0xc9, 0x74}}
	return a, nil
}

var _upsert_bandSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xf2\xf4\x0b\x76\x0d\x0a\x51\xf0\xf4\x0b\xf1\x57\x48\x4a\xcc\x4b\x29\x56\xd0\xc8\x4c\xd1\x51\x48\x49\x2c\x49\xd4\x54\x08\x73\xf4\x09\x75\x0d\x56\xd0\x50\x31\xd4\x51\x50\x31\xd2\xe4\xf2\xf7\x53\x70\xf6\xf7\x73\xf3\xf1\x74\x0e\x01\x29\xd3\x54\x70\xf1\xe7\x52\x50\x08\x0d\x70\x71\x0c\x71\x55\x08\x76\x0d\x01\x6b\x53\xb0\x55\x50\x31\xe2\x02\x04\x00\x00\xff\xff\x7c\xf2\x04\xd0\x58\x00\x00\x00")

func upsert_bandSqlBytes() ([]byte, error) {
	return bindataRead(
		_upsert_bandSql,
		"upsert_band.sql",
	)
}

func upsert_bandSql() (*asset, error) {
	bytes, err := upsert_bandSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "upsert_band.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0x73, 0x18, 0x63, 0xe3, 0x95, 0x2, 0x52, 0x53, 0xb4, 0xa9, 0x8, 0xf5, 0x13, 0xc7, 0xfc, 0x9, 0xb0, 0xe0, 0x12, 0xb1, 0x99, 0x46, 0x3f, 0x8d, 0xcd, 0xfa, 0x19, 0x80, 0xc2, 0xd4, 0x0, 0xe}}
	return a, nil
}

var _upsert_idolSql = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xf2\xf4\x0b\x76\x0d\x0a\x51\xf0\xf4\x0b\xf1\x57\xc8\x4c\xc9\xcf\x29\x56\xd0\xc8\x4c\xd1\x51\x48\x4a\xcc\x4b\x89\x07\x31\x52\x12\x4b\x12\x35\x15\xc2\x1c\x7d\x42\x5d\x83\x15\x34\x54\x0c\x75\x14\x54\x8c\x74\x14\x54\x8c\x35\xb9\xfc\xfd\x14\x9c\xfd\xfd\xdc\x7c\x3c\x9d\x43\x40\x7a\x34\x15\x5c\xfc\xb9\x14\x14\x42\x03\x5c\x1c\x43\x5c\x15\x82\x5d\x43\x60\x66\x28\xd8\x82\xb5\x80\x0c\x02\x31\x8d\xb9\x00\x01\x00\x00\xff\xff\xf9\xf9\xc7\x4a\x73\x00\x00\x00")

func upsert_idolSqlBytes() ([]byte, error) {
	return bindataRead(
		_upsert_idolSql,
		"upsert_idol.sql",
	)
}

func upsert_idolSql() (*asset, error) {
	bytes, err := upsert_idolSqlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "upsert_idol.sql", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xf2, 0xcf, 0x2a, 0xde, 0x16, 0xb6, 0x2b, 0x87, 0xe2, 0x4b, 0x8, 0x74, 0xc3, 0x66, 0xa9, 0x82, 0x72, 0x3b, 0xdb, 0xf5, 0x96, 0x6c, 0xab, 0x8e, 0x53, 0x72, 0xe4, 0xc3, 0x73, 0xa4, 0xa, 0xb}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"get_bands.sql": get_bandsSql,

	"get_idols.sql": get_idolsSql,

	"init_db.sql": init_dbSql,

	"upsert_band.sql": upsert_bandSql,

	"upsert_idol.sql": upsert_idolSql,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"get_bands.sql":   &bintree{get_bandsSql, map[string]*bintree{}},
	"get_idols.sql":   &bintree{get_idolsSql, map[string]*bintree{}},
	"init_db.sql":     &bintree{init_dbSql, map[string]*bintree{}},
	"upsert_band.sql": &bintree{upsert_bandSql, map[string]*bintree{}},
	"upsert_idol.sql": &bintree{upsert_idolSql, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}
