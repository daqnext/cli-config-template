package components

import (
	"errors"

	"github.com/daqnext/cli-config-template/cli"
)

func InitLocalLog() error {
	local_log_level, local_log_level_err := cli.AppToDO.ConfigJson.GetString("local_log_level")
	if local_log_level_err != nil {
		return errors.New("local_log_level [string] in config.json not defined," + local_log_level_err.Error())
	}
	cli.LocalLogger.ResetLevel(local_log_level)
	return nil
}
