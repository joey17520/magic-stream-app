package database

import (
	"context"
	"time"

	"github.com/joey17520/magic-stream-app/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.uber.org/zap"
)

var (
	Client *mongo.Client
	logger *zap.Logger
	cfg    *config.Config
)

// initLogger 初始化数据库日志记录器
func initLogger() {
	if logger == nil {
		// 创建简单的控制台日志记录器
		zapConfig := zap.NewProductionConfig()
		zapConfig.OutputPaths = []string{"stdout"}
		zapConfig.ErrorOutputPaths = []string{"stderr"}

		var err error
		logger, err = zapConfig.Build()
		if err != nil {
			// 如果无法创建zap记录器，使用默认的
			logger = zap.NewNop()
		}
	}
}

// getLogger 获取日志记录器
func getLogger() *zap.Logger {
	if logger == nil {
		initLogger()
	}
	return logger
}

// InitDB 初始化数据库连接
func InitDB(config *config.Config) error {
	logger := getLogger()
	cfg = config

	startTime := time.Now()

	MongoDB := cfg.MongoDBURI
	if MongoDB == "" {
		logger.Fatal("MONGODB_URI environment variable is not set")
	}

	databaseName := cfg.DatabaseName
	if databaseName == "" {
		databaseName = "magicstream"
	}

	logger.Info("Connecting to MongoDB",
		zap.String("database", databaseName),
		zap.String("uri_length", string(rune(len(MongoDB)))),
	)

	clientOptions := options.Client().ApplyURI(MongoDB)

	// 创建客户端
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		logger.Error("Failed to create MongoDB client",
			zap.Error(err),
			zap.String("database", databaseName),
		)
		return err
	}

	// 设置连接超时
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 测试连接
	err = client.Ping(ctx, nil)
	if err != nil {
		logger.Error("Failed to ping MongoDB",
			zap.Error(err),
			zap.String("database", databaseName),
		)
		return err
	}

	Client = client

	duration := time.Since(startTime)
	logger.Info("MongoDB connection established successfully",
		zap.String("database", databaseName),
		zap.Duration("connection_time", duration),
	)

	return nil
}

// GetDBInstance 获取数据库实例（单例模式）
func GetDBInstance() *mongo.Client {
	if Client == nil {
		if cfg == nil {
			getLogger().Fatal("Database configuration not initialized. Call InitDB first.")
		}
		if err := InitDB(cfg); err != nil {
			getLogger().Fatal("Failed to initialize database", zap.Error(err))
		}
	}
	return Client
}

// OpenCollection 打开指定集合
func OpenCollection(collectionName string) *mongo.Collection {
	logger := getLogger()

	if Client == nil {
		GetDBInstance()
	}

	databaseName := cfg.DatabaseName
	if databaseName == "" {
		databaseName = "magicstream"
	}

	collection := Client.Database(databaseName).Collection(collectionName)

	if collection == nil {
		logger.Error("Failed to open collection",
			zap.String("collection", collectionName),
			zap.String("database", databaseName),
		)
		return nil
	}

	logger.Debug("Collection opened",
		zap.String("collection", collectionName),
		zap.String("database", databaseName),
	)

	return collection
}

// CloseDB 关闭数据库连接
func CloseDB() {
	if Client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := Client.Disconnect(ctx); err != nil {
			getLogger().Error("Failed to disconnect from MongoDB", zap.Error(err))
		} else {
			getLogger().Info("MongoDB connection closed")
		}
	}
}
