package configs

import (
	"github.com/mynreden/microservices-go/common/utils"
)

type Config struct {
	Addr      string `json:"addr"`
	UsersAddr string `json:"usersAddr"`
	PostsAddr string `json:"postsAddr"`
	DB        struct {
		DSN string `json:"dsn"`
	} `json:"db"`
	StaticDir string `json:"static_dir"`
}

func GetConfig() (*Config, error) {
	addr := utils.EnvString("ADDR", ":8080")
	usersAddr := utils.EnvString("USER_ADDR", "localhost:50051")
	postsAddr := utils.EnvString("POST_ADDR", "localhost:50052")
	DSN := utils.EnvString("DSN", "pgsql:host=localhost;port=5432;dbname=finalAdvProg;user=postgres;password=sultan2004")
	staticDir := utils.EnvString("STATIC_DIR", "/static")
	c := &Config{addr, usersAddr, postsAddr, struct {
		DSN string `json:"dsn"`
	}{DSN}, staticDir}
	return c, nil
}
