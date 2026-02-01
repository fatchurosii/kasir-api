package main

import (
	"fmt"
	"kasir-api/database"
	"kasir-api/router"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string `mapstructure:"APP_PORT"`
	DBConn string `mapstructure:"DB_CONN"`
}

func main() {

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig()
	}

	config := Config{
		Port:   viper.GetString("APP_PORT"),
		DBConn: viper.GetString("DB_CONN"),
	}

	db, err := database.InitDB(config.DBConn)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close()

	router.RegisterRoutes(db)

	addr := "0.0.0.0:" + config.Port

	fmt.Println("Listening on " + addr)

	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Println("Gagal memuat server:", err)
	}

}
