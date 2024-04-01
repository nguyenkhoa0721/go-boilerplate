package config

type App struct {
	Address      string
	AdminAddress string
	AppBaseUrl   string
	Cors         string
	AuthSecret   string
}

type SocialAuth struct {
	GoogleProvider   GoogleProviderConfig
	TelegramProvider TelegramProviderConfig
}

type GoogleProviderConfig struct {
	ClientId     string
	ClientSecret string
	Scopes       string
	UserInfoUrl  string
}

type TelegramProviderConfig struct {
	Token string
}
