package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/jinzhu/copier"
	"github.com/lalifeier/vvgo-mall/pkg/dataoke"
)

type TaoBaoRepo interface {
	CarouseList(ctx context.Context, params map[string]string) ([]Carouse, error)
}

type TaoBaoUseCase struct {
	log *log.Helper
	dtk *dataoke.Client
}

type Carouse struct {
	TopicID    int    `json:"topicId"`
	TopicName  string `json:"topicName"`
	TopicImage string `json:"topicImage"`
	Link       string `json:"link"`
	SourceType int    `json:"sourceType"`
	ActivityID string `json:"activityId"`
}

type Subcategories struct {
	Scpic    string `json:"scpic"`
	Subcid   int    `json:"subcid"`
	Subcname string `json:"subcname"`
}

type Categories struct {
	Cid           int             `json:"cid"`
	Cname         string          `json:"cname"`
	Cpic          string          `json:"cpic"`
	Subcategories []Subcategories `json:"subcategories"`
}

func NewTaoBaoUseCase(logger log.Logger, dtk *dataoke.Client) *TaoBaoUseCase {
	return &TaoBaoUseCase{
		dtk: dtk,
		log: log.NewHelper(logger),
	}
}

func (uc *TaoBaoUseCase) CarouseList(ctx context.Context, params map[string]string) ([]*Carouse, error) {
	resp, err := uc.dtk.CarouseList(params)
	if err != nil {
		return nil, err
	}

	carouseList := make([]*Carouse, 0)
	err = copier.Copy(&carouseList, &resp.Data)
	if err != nil {
		return nil, err
	}

	return carouseList, nil
}

func (uc *TaoBaoUseCase) GetSuperCategory(ctx context.Context, params map[string]string) ([]*Categories, error) {
	resp, err := uc.dtk.GetSuperCategory(params)
	if err != nil {
		return nil, err
	}

	categoriesList := make([]*Categories, 0)
	err = copier.Copy(&categoriesList, &resp.Data)
	if err != nil {
		return nil, err
	}

	return categoriesList, nil
}
