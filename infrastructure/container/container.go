package container

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type (
	Container struct {
		Apps  *Apps
		Pgsql *Pgsql
		Redis *Redis
	}

	Apps struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Host    string `json:"host"`
		Port    int    `json:"port"`
	}

	Pgsql struct {
		Username           string `json:"username"`
		Password           string `json:"password"`
		DbName             string `json:"Dbname"`
		Host               string `json:"host"`
		Port               int    `json:"port"`
		MinIdleConnections int    `json:"minIdleConnections"`
		MaxOpenConnections int    `json:"maxOpenConnections"`
		MaxLifetime        int    `json:"maxLifetime"`
	}

	Redis struct {
		RedisAddr      string
		RedisPassword  string
		RedisDB        int
		RedisDefaultdb int
		MinIdleConns   int
		PoolSize       int
		PoolTimeout    int
	}
)

func (c *Container) Validate() *Container {
	if c.Apps == nil {
		panic("Apps config is nill")
	}
	if c.Pgsql == nil {
		panic("Pgsql config is nill")
	}
	if c.Redis == nil {
		panic("Redis config is nill")
	}

	return c
}

func New(envpath string) *Container {
	v := viper.New()
	v.SetConfigFile(envpath)
	pathDir, err := os.Executable()
	if err != nil {
		panic(err)
	}
	dir := filepath.Dir(pathDir)
	v.AddConfigPath(dir)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	appHost := v.GetString("app.host")
	appPort := v.GetInt("app.port")
	appName := v.GetString("app.name")
	appVersion := v.GetString("app.version")

	psqlDbName := v.GetString("psql.dbname")
	psqlUser := v.GetString("psql.username")
	psqlPass := v.GetString("psql.password")
	psqlHost := v.GetString("psql.host")
	psqlPort := v.GetInt("psql.port")
	psqlMaxOpenConnections := v.GetInt("psql.MaxOpenConnections")
	psqlMaxLifetime := v.GetInt("psql.MaxLifetime")
	psqlMinIdleConn := v.GetInt("psql.MinIdleConnection")

	redisAddr := v.GetString("redis.Addr")
	redisPassword := v.GetString("redis.Password")
	redisDB := v.GetInt("redis.Db")
	redisDefaultDB := v.GetInt("redis.Defaultdb")
	redisMinIdleConns := v.GetInt("redis.MinIdleConns")
	redisPoolSize := v.GetInt("redis.PoolSize")
	redisPoolTimeout := v.GetInt("redis.PoolTimeout")

	appConf := &Apps{
		Name:    appName,
		Version: appVersion,
		Host:    appHost,
		Port:    appPort,
	}

	psqlConf := &Pgsql{
		Username:           psqlUser,
		Password:           psqlPass,
		DbName:             psqlDbName,
		Host:               psqlHost,
		Port:               psqlPort,
		MinIdleConnections: psqlMinIdleConn,
		MaxOpenConnections: psqlMaxOpenConnections,
		MaxLifetime:        psqlMaxLifetime,
	}

	redisConf := &Redis{
		RedisAddr:      redisAddr,
		RedisPassword:  redisPassword,
		RedisDB:        redisDB,
		RedisDefaultdb: redisDefaultDB,
		MinIdleConns:   redisMinIdleConns,
		PoolSize:       redisPoolSize,
		PoolTimeout:    redisPoolTimeout,
	}

	cont := &Container{
		Apps:  appConf,
		Pgsql: psqlConf,
		Redis: redisConf,
	}

	cont.Validate()

	return cont

}
