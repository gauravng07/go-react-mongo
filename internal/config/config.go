package config

import "github.com/spf13/viper"

const (
	Port			=	"PORT"
	Env				= 	"ENV"
	LogLevel		=	"LOGLEVEL"
	BuildDir		= 	"BUILD_DIR"
	UserName		=	"USER_NAME"
	Password		=	"PASSWORD"
	ClusterAddress	= 	"CLUSTER_ADDRESS"
	DBName			= 	"DB_NAME"
	Collection 		= 	"COLLECTION"
)

func init()  {
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetDefault(Port, "8080")
	viper.SetDefault(Env, "dev")
	viper.SetDefault(LogLevel, "debug")
	viper.SetDefault(BuildDir, "frontend/shopper-stop/build")
}

func ReadConfig(env string) error {
	viper.SetConfigFile("app-" + env + ".yaml")
	return viper.ReadInConfig()
}