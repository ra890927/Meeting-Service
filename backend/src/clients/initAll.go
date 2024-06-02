package clients

import (
	"meeting-center/src/utils"
)

var initialized = false

// avoid using init function for testing purpose
func Init() {
	if !initialized {
		utils.InitConfig()
		initRedis()
		initDB()
		initialized = true
	}
}
