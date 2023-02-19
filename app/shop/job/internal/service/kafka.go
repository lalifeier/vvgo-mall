package service

import (
	"context"
	"fmt"

	v1 "github.com/lalifeier/vvgo-mall/gen/api/go/shop/job/v1"
	"github.com/tx7do/kratos-transport/broker"
)

func (s *ShopJobService) Foo(ctx context.Context, topic string, headers broker.Headers, msg *[]v1.Foo) error {
	fmt.Println("Foo() Topic: ", topic)

	return nil
}
