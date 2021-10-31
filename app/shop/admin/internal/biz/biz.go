package biz

import (
	"github.com/google/wire"
	"github.com/lalifeier/vgo/app/shop/admin/internal/biz/ums"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(ums.ProviderSet)
