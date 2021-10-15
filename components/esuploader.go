package components

import (
	"errors"

	"github.com/daqnext/ESUploader/uploader"
	fj "github.com/daqnext/fastjson"
)

/*
elasticsearch_addr
elasticsearch_username
elasticsearch_password
*/
func InitESUploader(ConfigJson *fj.FastJson) (*uploader.Uploader, error) {

	elasticsearch_addr, elasticsearch_addr_err := ConfigJson.GetString("elasticsearch_addr")
	if elasticsearch_addr_err != nil {
		return nil, errors.New("elasticsearch_addr [string] in config.json not defined," + elasticsearch_addr_err.Error())
	}

	elasticsearch_username, elasticsearch_username_err := ConfigJson.GetString("elasticsearch_username")
	if elasticsearch_username_err != nil {
		return nil, errors.New("elasticsearch_username [string] in config.json not defined," + elasticsearch_username_err.Error())
	}

	elasticsearch_password, elasticsearch_password_err := ConfigJson.GetString("elasticsearch_password")
	if elasticsearch_password_err != nil {
		return nil, errors.New("elasticsearch_password [string] in config.json not defined," + elasticsearch_password_err.Error())
	}

	ESUploader, err := uploader.New(elasticsearch_addr, elasticsearch_username, elasticsearch_password)
	if err != nil {
		return nil, err
	}

	return ESUploader, nil
}
