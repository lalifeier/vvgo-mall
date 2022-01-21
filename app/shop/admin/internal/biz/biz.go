package biz

import (
	"github.com/google/wire"
	"github.com/lalifeier/vvgo-mall/app/shop/admin/internal/biz/sys"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(sys.ProviderSet)
