package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func ImgConvertToB64Str(imgPath string) (name, content string) {
	ext := filepath.Ext(imgPath)
	log.Println(ext)
	data, err := ioutil.ReadFile(imgPath)
	check(err)
	b64str := base64.StdEncoding.EncodeToString(data)
	return filepath.Base(imgPath), b64str
}
func B64StrSaveToImg(name, content string) {
	src, err := base64.StdEncoding.DecodeString(content)
	check(err)
	ioutil.WriteFile(name, src, 0777)
}
