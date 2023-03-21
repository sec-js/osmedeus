package database

// load driver

// // DB gorm connector
// var DB *gorm.DB

// // InitDB connect to db
// func InitDB(options libs.Options) (*gorm.DB, error) {
// 	newLogger := logger.New(
// 		log.New(ioutil.Discard, "\r\n", log.LstdFlags), // io writer
// 		logger.Config{
// 			SlowThreshold:             time.Second, // Slow SQL threshold
// 			LogLevel:                  logger.Info, // Log level
// 			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
// 			Colorful:                  false,       // Disable color
// 		},
// 	)

// 	config := gorm.Config{
// 		SkipDefaultTransaction:                   false,
// 		NamingStrategy:                           nil,
// 		FullSaveAssociations:                     false,
// 		Logger:                                   newLogger,
// 		NowFunc:                                  nil,
// 		DryRun:                                   false,
// 		PrepareStmt:                              false,
// 		DisableAutomaticPing:                     false,
// 		DisableForeignKeyConstraintWhenMigrating: false,
// 		DisableNestedTransaction:                 false,
// 		AllowGlobalUpdate:                        false,
// 		QueryFields:                              false,
// 		CreateBatchSize:                          0,
// 		ClauseBuilders:                           nil,
// 		ConnPool:                                 nil,
// 		Dialector:                                nil,
// 		Plugins:                                  nil,
// 	}

// 	var err error
// 	if options.Server.DBType == "mysql" {
// 		utils.InforF("Connect Database at %v:%v", options.Server.DBHost, options.Server.DBPort)
// 		utils.InforF("Use Database: %v", options.Server.DBName)
// 		DB, err = gorm.Open(mysql.Open(options.Server.DBConnection), &config)
// 	} else {
// 		DB, err = gorm.Open(sqlite.Open(options.Server.DBPath), &config)
// 	}

// 	if err != nil {
// 		fmt.Printf("Error Database connect at %v:%v -- %v\n", options.Server.DBHost, options.Server.DBPort, err)
// 		return nil, err
// 	}

// 	// scanning data
// 	// DB.AutoMigrate(&User{})
// 	//DB.AutoMigrate(&Org{})
// 	DB.AutoMigrate(&Target{})
// 	DB.AutoMigrate(&Scan{})
// 	DB.AutoMigrate(&Report{})
// 	DB.AutoMigrate(&CloudInstance{})
// 	// asset data
// 	return DB, nil
// }
