package config

type Log struct {
	Level     string `mapstructure:"level" json:"level" yaml:"level"`
	RootDir   string `mapstructure:"root_dir" json:"root_dir" yaml:"root_dir"`
	Filename  string `mapstructure:"filename" json:"filename" yaml:"filename"`
	Format    string `mapstructure:"format" json:"format" yaml:"format"`
	ShowLine  bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	MaxBackup int    `mapstructure:"max_backup" json:"max_backup" yaml:"max_backup"`
	MaxSize   int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"`
	MaxAge    int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`
	Compress  bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
}
