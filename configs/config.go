package configs

import "github.com/spf13/viper"

type config struct {
	DB  DBConfig
	API ServerAPI
}

var cfg *config

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
	Driver   string
}

type ServerAPI struct {
	Port              string
	BaseUrl           string
	KeyToken          string
	KeyPlataformCode  string
	DispatcherCNPJ    string
	DispatcherZipcode int32
	RecipientType     int
	RecipientCountry  string
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func Load() error {

	cfg = new(config)

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
		Driver:   viper.GetString("database.driver"),
	}

	cfg.API = ServerAPI{
		Port:              viper.GetString("api.port"),
		BaseUrl:           viper.GetString("api.base_url_freterapido"),
		KeyToken:          viper.GetString("api.key_token"),
		KeyPlataformCode:  viper.GetString("api.key_platformcode"),
		DispatcherCNPJ:    viper.GetString("api.dispatcher_cnpj"),
		DispatcherZipcode: viper.GetInt32("api.dispatcher_zipcode"),
		RecipientType:     viper.GetInt("api.recipient_type"),
		RecipientCountry:  viper.GetString("api.recipient_country"),
	}

	return nil
}

func GetDB() DBConfig {
	return cfg.DB
}

func GetServer() ServerAPI {
	return cfg.API
}
