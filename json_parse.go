package vanConfig

import "github.com/bitly/go-simplejson"

func (self *Config) parseJson() error {

	json, err := simplejson.NewJson(self.Value)

	if err != nil {
		return err
	}

	self.ParseJson = json

	return nil
}


