package biz

import (
	"github.com/google/wire"
	"github.com/lalifeier/vvgo/app/shop/admin/internal/biz/sys"
	"github.com/lalifeier/vvgo/app/shop/admin/internal/biz/ums"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(ums.ProviderSet, sys.ProviderSet)
