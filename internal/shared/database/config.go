package database

import "time"

type ConfigDatabase struct {
	Username           string        `json:"username"`
	Password           string        `json:"password"`
	Name               string        `json:"name"`
	Schema             string        `json:"schema"`
	Host               string        `json:"host"`
	Port               string        `json:"port"`
	MinIdleConnections int           `json:"min_idle_connections"`
	MaxOpenConnections int           `json:"max_open_connections"`
	ConnMaxLifetime    time.Duration `json:"conn_max_lifetime"`
	DebugMode          bool          `json:"debug_mode"`
}
