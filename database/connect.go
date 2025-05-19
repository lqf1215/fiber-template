package database

import (
	"fiber-template/config"
	"fiber-template/model"
	"fiber-template/pkg"
	"fmt"
	"github.com/jinzhu/configor"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"strconv"
	"time"
)

func init() {
	projectPath := pkg.GetProjectPath()
	fmt.Println(projectPath)
	err := configor.Load(&config.Config, projectPath+"/config.yml")
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	fmt.Println(config.Config)

}

// ConnectDB connect to database
func ConnectDB() {
	allModels := []interface{}{
		&model.User{},
		&model.Manager{},
	}
	var err error
	dbConfig := config.Config.DB
	p := dbConfig.Port
	port, err := strconv.ParseUint(p, 10, 32)
	if err != nil {
		panic(err)
	}

	sqlLog := logger.New(log.New(os.Stdout, "[SQL] ", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Warn,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbConfig.Host, port,
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Name)
	if DB, err = gorm.Open(postgres.Open(dsn),
		&gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			PrepareStmt:                              true, // 开启自动更新UpdatedAt字段
			Logger:                                   sqlLog,
		}); err != nil {
		panic("failed to connect database")
	}

	//创表
	for _, m := range allModels {
		if !DB.Migrator().HasTable(m) {
			if err = DB.AutoMigrate(m); err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("Database Connected")
}
