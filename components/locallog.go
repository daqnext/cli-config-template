package components

import (
	"errors"

	localLog "github.com/daqnext/LocalLog/log"
	fj "github.com/daqnext/fastjson"
)

func InitLocalLog(localLogger_ *localLog.LocalLog, ConfigJson *fj.FastJson) error {
	local_log_level, local_log_level_err := ConfigJson.GetString("local_log_level")
	if local_log_level_err != nil {
		return errors.New("local_log_level [string] in config.json not defined," + local_log_level_err.Error())
	}
	err := localLogger_.ResetLevel(local_log_level)
	if err != nil {
		return err
	}
	return nil
}
