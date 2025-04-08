package config

type Config struct {
	URL         string `json:"db_url"`
	CurrentUser string `json:"current_user_name"`
}
