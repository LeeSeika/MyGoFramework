package orm

import (
	"context"
	"github.com/godemo/coredemo/framework"
	"github.com/godemo/coredemo/framework/contract"
	"gorm.io/driver/clickhouse"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"sync"
	"time"
)

func NewHadeOrmService(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	ormService := &HadeOrmService{
		container: container,
		dbs:       map[string]*gorm.DB{},
		lock:      &sync.RWMutex{},
	}
	return ormService, nil
}

type HadeOrmService struct {
	container  framework.Container
	dbs        map[string]*gorm.DB
	lock       *sync.RWMutex
	configPath string
	gormConfig *gorm.Config
}

func (hos *HadeOrmService) GetBasicConfig(container framework.Container) *contract.DBConfig {
	configService := container.MustMake(contract.ConfigKey).(contract.Config)
	logService := container.MustMake(contract.LogKey).(contract.Log)
	config := &contract.DBConfig{}
	// 直接使用配置服务的load方法读取,yaml文件
	err := configService.Load("database", config)
	if err != nil {
		// 直接使用logService来打印错误信息
		logService.Error(context.Background(), "parse database config error", nil)
		return nil
	}
	return config
}

func (hos *HadeOrmService) GetDB(options ...contract.DBOption) (*gorm.DB, error) {

	for _, opt := range options {
		if err := opt(hos); err != nil {
			return nil, err
		}
	}

	if hos.configPath == "" {
		WithConfigPath("database.default")
	}

	logService := hos.container.MustMake(contract.LogKey).(contract.Log)
	configService := hos.container.MustMake(contract.LogKey).(contract.Config)

	config := hos.GetBasicConfig(hos.container)
	if err := configService.Load(hos.configPath, config); err != nil {
		return nil, err
	}
	if config.Dsn == "" {
		dsn, err := config.FormatDsn()
		if err != nil {
			return nil, err
		}
		config.Dsn = dsn
	}
	if db, ok := hos.dbs[config.Dsn]; ok {
		return db, nil
	}

	ormLogger := NewOrmLogger(logService)
	gormConfig := hos.gormConfig
	if gormConfig == nil {
		gormConfig = &gorm.Config{}
	}
	if gormConfig.Logger == nil {
		gormConfig.Logger = ormLogger
	}

	var db *gorm.DB
	var err error
	switch config.Dsn {
	case "mysql":
		db, err = gorm.Open(mysql.Open(config.Dsn))
	case "postgres":
		db, err = gorm.Open(postgres.Open(config.Dsn), gormConfig)
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(config.Dsn), gormConfig)
	case "sqlserver":
		db, err = gorm.Open(sqlserver.Open(config.Dsn), gormConfig)
	case "clickhouse":
		db, err = gorm.Open(clickhouse.Open(config.Dsn), gormConfig)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if config.ConnMaxIdle > 0 {
		sqlDB.SetMaxIdleConns(config.ConnMaxIdle)
	}
	if config.ConnMaxOpen > 0 {
		sqlDB.SetMaxOpenConns(config.ConnMaxOpen)
	}
	if config.ConnMaxLifetime != "" {
		duration, err := time.ParseDuration(config.ConnMaxLifetime)
		if err != nil {
			logService.Error(context.Background(), "conn max life time parse error", map[string]interface{}{"err": err})
		} else {
			sqlDB.SetConnMaxLifetime(duration)
		}
	}

	if err == nil {
		hos.dbs[config.Dsn] = db
	}

	return db, err
}
