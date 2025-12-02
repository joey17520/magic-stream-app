package config

import (
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

// Config 应用配置结构体
type Config struct {
	// 服务器配置
	ServerPort string `env:"PORT" envDefault:"8088"`
	GinMode    string `env:"GIN_MODE" envDefault:"release"`
	LogLevel   string `env:"LOG_LEVEL" envDefault:"info"`
	Timezone   string `env:"TZ" envDefault:"Asia/Shanghai"`

	// 数据库配置
	MongoDBURI   string `env:"MONGODB_URI,required"`
	DatabaseName string `env:"DATABASE_NAME" envDefault:"magicstream"`

	// 安全配置
	SecretKey        string `env:"SECRET_KEY,required"`
	SecretRefreshKey string `env:"SECRET_REFRESH_KEY,required"`

	// CORS配置
	AllowedOrigins []string `env:"ALLOWED_ORIGINS" envDefault:"http://localhost:5173,http://localhost:80" envSeparator:","`

	// AI服务配置
	DeepSeekAPIKey     string `env:"DEEPSEEK_API_KEY"`
	BasePromptTemplate string `env:"BASE_PROMPT_TEMPLATE" envDefault:"You are a sentiment analysis assistant. Classify the following movie review into one of these sentiment categories: {rankings}. Only respond with the category name. Review:"`

	// 业务配置
	RecommendedMovieLimit int `env:"RECOMMENDED_MOVIE_LIMIT" envDefault:"5"`
}

// LoadConfig 加载配置
func LoadConfig(logger *zap.Logger) *Config {
	// 尝试加载.env文件（仅用于开发环境）
	// 在生产环境中，应该通过环境变量设置
	_ = godotenv.Load()

	config := &Config{
		// 服务器配置
		ServerPort: getEnv("PORT", "8088"),
		GinMode:    getEnv("GIN_MODE", "release"),
		LogLevel:   getEnv("LOG_LEVEL", "info"),
		Timezone:   getEnv("TZ", "Asia/Shanghai"),

		// 数据库配置
		MongoDBURI:   getEnv("MONGODB_URI", ""),
		DatabaseName: getEnv("DATABASE_NAME", "magicstream"),

		// 安全配置
		SecretKey:        getEnv("SECRET_KEY", ""),
		SecretRefreshKey: getEnv("SECRET_REFRESH_KEY", ""),

		// AI服务配置
		DeepSeekAPIKey:     getEnv("DEEPSEEK_API_KEY", ""),
		BasePromptTemplate: getEnv("BASE_PROMPT_TEMPLATE", "You are a sentiment analysis assistant. Classify the following movie review into one of these sentiment categories: {rankings}. Only respond with the category name. Review:"),

		// 业务配置
		RecommendedMovieLimit: getEnvAsInt("RECOMMENDED_MOVIE_LIMIT", 5),
	}

	// 处理CORS配置
	allowedOrigins := getEnv("ALLOWED_ORIGINS", "http://localhost:5173,http://localhost:80")
	config.AllowedOrigins = strings.Split(allowedOrigins, ",")
	for i := range config.AllowedOrigins {
		config.AllowedOrigins[i] = strings.TrimSpace(config.AllowedOrigins[i])
	}

	// 验证必需配置
	config.validate(logger)

	return config
}

// getEnv 获取环境变量，提供默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvAsInt 获取环境变量作为整数
func getEnvAsInt(key string, defaultValue int) int {
	strValue := getEnv(key, "")
	if strValue == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(strValue)
	if err != nil {
		return defaultValue
	}
	return value
}

// validate 验证配置
func (c *Config) validate(logger *zap.Logger) {
	// 必需配置验证
	requiredConfigs := map[string]string{
		"MONGODB_URI":        c.MongoDBURI,
		"SECRET_KEY":         c.SecretKey,
		"SECRET_REFRESH_KEY": c.SecretRefreshKey,
	}

	for key, value := range requiredConfigs {
		if value == "" {
			logger.Fatal("Required configuration is missing", zap.String("config_key", key))
		}
	}

	// 配置合理性检查
	if c.RecommendedMovieLimit <= 0 || c.RecommendedMovieLimit > 100 {
		logger.Warn("Recommended movie limit is out of reasonable range, using default",
			zap.Int("provided", c.RecommendedMovieLimit),
			zap.Int("default", 5),
		)
		c.RecommendedMovieLimit = 5
	}

	// 记录配置摘要（敏感信息不记录）
	logger.Info("Configuration loaded",
		zap.String("server_port", c.ServerPort),
		zap.String("gin_mode", c.GinMode),
		zap.String("log_level", c.LogLevel),
		zap.String("database", c.DatabaseName),
		zap.Strings("allowed_origins", c.AllowedOrigins),
		zap.Int("recommended_movie_limit", c.RecommendedMovieLimit),
		zap.Bool("deepseek_configured", c.DeepSeekAPIKey != ""),
	)
}
