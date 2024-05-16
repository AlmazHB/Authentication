package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

func Init(path, name string) error {
	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	fmt.Println("Configs imported successfuly!")
	return nil
}
