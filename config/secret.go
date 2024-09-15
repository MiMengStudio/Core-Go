package config

const ADDR, PORT, USER, PASSWORD, DATABASE = "ip", "port", "user", "password", "database"

// token var
var (
	Secret     = "MiMeng"       // 加盐
	ExpireTime = 3600 * 24 * 30 // token有效期
)
