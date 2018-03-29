package domain

import (
	"github.com/midorigreen/gt/repository"
	"github.com/spf13/viper"
)

type ListDomain struct {
	// dependency
}

func Create() ListDomain {
	token := viper.GetString("token")
	key := viper.GetString("apikey")
	board := viper.GetString("board")

	tr := repository.NewTrello(token, key, board)

	tr.

	return ListDomain{}
}
