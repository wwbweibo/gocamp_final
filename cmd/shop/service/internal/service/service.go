package service

import (
	"eshop/api/shop/service/v1"
	"eshop/cmd/shop/service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewShopInterface)

type ShopInterface struct {
	v1.UnimplementedShopInterfaceServer

	uc *biz.UserUseCase
	cc *biz.CatalogUseCase

	log *log.Helper
}

func NewShopInterface(uc *biz.UserUseCase, cc *biz.CatalogUseCase, logger log.Logger) *ShopInterface {
	return &ShopInterface{
		log: log.NewHelper(log.With(logger, "module", "service/interface")),
		uc:  uc,
		cc:  cc,
	}
}
