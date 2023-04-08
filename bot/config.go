package bot

type Config struct {
	Token    string `yaml:"token"`
	Commands struct {
		Start        string `yaml:"start"`
		Rules        string `yaml:"rules"`
		Registration string `yaml:"registration"`
		Subscription string `yaml:"subscription"`
		Sale         string `yaml:"sale"`
	}
	InternalCommands struct {
		EnterCarBrand  string `yaml:"enterCarBrand"`
		EnterCarModel  string `yaml:"enterCarModel"`
		EnterCarEngine string `yaml:"enterCarEngine"`
	}

	ValidateData struct {
		ChannelID  int64  `yaml:"channelID"`
		ChannelUrl string `yaml:"channelUrl"`
	}
	Psql struct {
		Login              string `yaml:"login"`
		Password           string `yaml:"password"`
		SslMode            string `yaml:"ssl_mode"`
		DriverName         string `yaml:"driver_name"`
		DatabaseName       string `yaml:"database_name"`
		AttemptsConnection string `yaml:"attempts_connection"`
	}
}

func (c *Config) printCommands() string {
	return "\n" + c.Commands.Start + "\n" +
		"\n" + c.Commands.Rules +
		"\n" + c.Commands.Registration +
		"\n" + c.Commands.Subscription +
		"\n" + c.Commands.Sale

}
