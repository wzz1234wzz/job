package mode

type favorite struct {
	Sports []string
	Music []string
	LuckyNumber int
}

type information struct {
	Name string
	Age  int
	Alise []string
	Image string
	Public bool
}

type YamlConfig struct {
	TimeStamp string
	Author string
	PassWd string
	Information information
	Favorite favorite;
}

type Config struct {
	v  *viper.Viper;
}