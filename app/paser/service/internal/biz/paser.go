package biz

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	pb "github.com/lalifeier/vvgo-mall/api/paser/service/v1"
	"github.com/lalifeier/vvgo-mall/app/paser/service/internal/data/douyin"
)

type PaserRepo interface {
}

type PaserUseCase struct {
	repo PaserRepo
	log  *log.Helper
}

func NewPaserUseCase(repo PaserRepo, logger log.Logger) *PaserUseCase {
	return &PaserUseCase{repo: repo, log: log.NewHelper(logger)}
}

func getDomain(uri string) (string, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return "", nil
	}
	parts := strings.Split(u.Hostname(), ".")
	return parts[len(parts)-2] + "." + parts[len(parts)-1], nil
}

func (uc PaserUseCase) Paser(url string) (*pb.PaserReply, error) {
	domain, err := getDomain(url)
	if err != nil {
		return nil, err
	}
	fmt.Println(domain)
	d := douyin.New()

	return d.Parse(url)
}
