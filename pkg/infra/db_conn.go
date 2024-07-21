package infra

import (
	"fmt"

	"golang.org/x/xerrors"

	"github.com/chaki8923/wedding-backend/pkg/lib/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func NewDBConnector(cfg *config.Config) (*gorm.DB, error) {
	// ロガーの設定
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // ログの出力先
		logger.Config{
				SlowThreshold: time.Second,   // 遅いクエリの閾値
				LogLevel:      logger.Info,   // ログレベル
				Colorful:      false,         // カラー出力を無効化
		},
)

dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, // デバッグロガーを設定
})

if err != nil {
		return nil, xerrors.Errorf("db connection failed：%w", err)
}

return db, nil
}
