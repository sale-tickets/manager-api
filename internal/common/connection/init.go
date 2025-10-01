package connection

func Init() {
	ConfigInfo.loadConfig()
	ConfigInfo.Database.connection()
}
