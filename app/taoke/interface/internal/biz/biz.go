package biz

import (
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewTaoBaoUseCase, NewJDUseCase, NewPddUseCase, NewVIPUseCase, NewTikTokUseCase)
