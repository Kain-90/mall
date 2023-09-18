package initialize

import "mall/global"

func Init() {
	SetupConfig()
	global.GVA_LOG = SetupLogger()
	global.GVA_DB = SetupDB()
}
