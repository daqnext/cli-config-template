package components

import "github.com/daqnext/BGJOB_GO/bgjob"

func InitBGJobs() *bgjob.JobManager {
	//////// ini bGJob   //////////////////////
	return bgjob.New()
}
