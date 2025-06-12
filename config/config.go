package config

import (
	"fmt"
	"os"
)

type Config struct {
    MySQL MySQLConfig
    Kafka KafkaConfig
    HTTP  HTTPConfig
    GRPC GRPCConfig
}

type MySQLConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
}

type KafkaConfig struct {
    Brokers []string
    Topic   string
}

type HTTPConfig struct {
    Port string
}

type GRPCConfig struct {
    Port string
}

func LoadConfig() *Config {
    return &Config{
        MySQL: MySQLConfig{
            Host:     getEnv("MYSQL_HOST", "mysql"),
            Port:     getEnv("MYSQL_PORT", "3306"),
            User:     getEnv("MYSQL_USER", "root"),
            Password: getEnv("MYSQL_PASSWORD", "password"),
            DBName:   getEnv("MYSQL_DBNAME", "myapp"),
        },
        Kafka: KafkaConfig{
            Brokers: []string{getEnv("KAFKA_BROKER", "localhost:9092")},
            Topic:   getEnv("KAFKA_TOPIC", "user-events"),
        },
        HTTP: HTTPConfig{
            Port: getEnv("HTTP_PORT", "8080"),
        },
        GRPC: GRPCConfig{
            Port: getEnv("GRPC_PORT","9090"),
        },
    }
}


func getEnv(key, defaultValue string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return defaultValue
}


// DSN builds the MySQL Data Source Name string
func (m MySQLConfig) DSN() string {
    return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        m.User, m.Password, m.Host, m.Port, m.DBName)
}
