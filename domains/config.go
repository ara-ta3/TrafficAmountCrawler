package domains

import "github.com/spf13/viper"

func LoadEnv() (config *EnvConfigs, err error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	if err := v.Unmarshal(&config); err != nil {
		return nil, err
	}
	return
}

type EnvConfigs struct {
	NihonTsushinID   string `mapstructure:"NIHON_TSUSHIN_ID"`
	NihonTsushinPass string `mapstructure:"NIHON_TSUSHIN_PASS"`
	SlackToken       string `mapstructure:"SLACK_TOKEN"`
	SlackChannelID   string `mapstructure:"SLACK_CHANNEL_ID"`
}
