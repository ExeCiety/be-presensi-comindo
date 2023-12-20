package utils

func GetFullAddress() string {
	return GetEnvValue("APP_URL", "")
}
