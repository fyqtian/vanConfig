package vanConfig

import (
	"strings"
	"bytes"
	"encoding/json"
	"github.com/bitly/go-simplejson"
)

const LINE_A string = "\n"
const BRACKETS_LEFT byte = '['
const BRACKETS_RIGHT byte = ']'
const EQUAL byte = '='
const POUND byte = '#'

func (self *Config) parseText() error {

	m := make(map[string]map[string]string)

	tmpStringSlice := strings.Split(string(self.Value), LINE_A)

	var tmpKey string
	for _, v := range tmpStringSlice {
		v = strings.TrimSpace(v)
		//if line is empty
		if len(v) == 0 {
			continue
		}
		bTmp := []byte(v)
		tLen := len(bTmp)

		if bTmp[0] == BRACKETS_LEFT && bTmp[tLen-1] == BRACKETS_RIGHT {
			value := string(bTmp[1:tLen-1])
			m[value] = make(map[string]string)
			tmpKey = value
		}

		if bytes.Contains(bTmp, []byte{EQUAL}) && bTmp[0] != EQUAL && bTmp[0] != POUND {
			equalIndex := bytes.Index(bTmp, []byte{EQUAL})
			key := string(bTmp[0:equalIndex])
			value := string((bTmp[equalIndex+1:]))

			if _,ok := m[tmpKey];ok {
				m[tmpKey][key] = value
			}
		}
	}

	json ,err := json.Marshal(&m)

	if err != nil{
		return err
	}

	self.ParseJson,err = simplejson.NewJson(json)

	if err != nil{
		return err
	}
	return nil
}
