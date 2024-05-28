package config

type Jwt struct {
	TokenType               string `mapstructure:"token_type" json:"token_type" yaml:"token_type"`
	Secret                  string `mapstructure:"secret" json:"secret" yaml:"secret"`
	JwtExp                  int64  `mapstructure:"jwt_exp" json:"jwt_exp" yaml:"jwt_exp"`
	JwtBlacklistGracePeriod int64  `mapstructure:"jwt_blacklist_grace_period" json:"jwt_blacklist_grace_period" yaml:"jwt_blacklist_grace_period"`
	RefreshGracePeriod      int64  `mapstructure:"refresh_grace_period" json:"refresh_grace_period" yaml:"refresh_grace_period"`
}
