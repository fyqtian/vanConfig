package vanConfig

import (
	"os"
	"io/ioutil"
	"github.com/bitly/go-simplejson"
	"sync"
	"errors"
)

type Config struct {
	//打开文件
	*os.File
	//文件类型json|xml
	FileType string
	//文件全路径
	FilePath string
	//文件内容
	Value []byte
	//
	ParseJson *simplejson.Json
	mu *sync.Mutex
}

func NewConfig(cf *Config)( error){
	return cf.Open()
}

func (self *Config) Open() error {

	_, err := os.Stat(self.FilePath)

	if err != nil {
		return err
	}

	self.File, err = os.Open(self.FilePath)

	if err != nil {
		return err
	}

	err = self.read()

	if err != nil{
		return err
	}
	defer self.File.Close()

	if self.FileType == "json" {
		err = self.parseJson()
	}else if self.FileType == "text"{
		err = self.parseText()
	}

	if err != nil{
		return err
	}

	return nil
}

func (self *Config) read() error {

	tmp, err := ioutil.ReadAll(self.File)

	if err != nil {
		return err
	}
	//检测长度
	if len(tmp) <= 0{
		return  errors.New("the file is empty")
	}

	self.Value = tmp

	return nil
}

func (self *Config)find(f func()){

	self.mu.Lock()
	defer self.mu.Unlock()

	if(self.FileType == "json"){
	}

}
