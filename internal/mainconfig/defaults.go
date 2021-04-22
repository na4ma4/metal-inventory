package mainconfig

import "github.com/spf13/viper"

// ConfigInit is the common config initialisation for the commands.
func ConfigInit() {
	viper.SetConfigName("mic")
	viper.SetConfigType("toml")
	viper.AddConfigPath("./artifacts")
	viper.AddConfigPath("./test")
	viper.AddConfigPath("$HOME/.mic")
	viper.AddConfigPath("/etc/mic")
	viper.AddConfigPath("/usr/local/etc")
	viper.AddConfigPath("/usr/local/mic/etc")
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath("/run/secrets")
	viper.AddConfigPath(".")

	_ = viper.ReadInConfig()
}
