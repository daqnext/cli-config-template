package components

import (
	"github.com/daqnext/BGJOB_GO/bgjob"
	"github.com/daqnext/LocalLog/log"
)

func InitBGJobs(localLogger *log.LocalLog) *bgjob.JobManager {
	//////// ini bGJob   //////////////////////
	return bgjob.New(localLogger)
}
