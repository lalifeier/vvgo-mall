package data

// import (
// 	"github.com/go-kratos/kratos/v2/log"
// 	"github.com/lalifeier/vvgo-mall/app/taoke/interface/internal/biz"
// )

// var _ biz.TaoBaoRepo = (*taoBaoRepo)(nil)

// type taoBaoRepo struct {
// 	data *Data
// 	log  *log.Helper
// }

// func NewTaoBaoRepo(data *Data, logger log.Logger) biz.TaoBaoRepo {
// 	return &taoBaoRepo{
// 		data: data,
// 		log:  log.NewHelper(log.With(logger, "module", "taobao-interface/data")),
// 	}
// }

// func (rp *taoBaoRepo) CarouseList(ctx context.Context, params map[string]string) ([]biz.Carouse, error) {
// 	resp, err := rp.data.dtk.CarouseList(params)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var data []biz.Carouse
// 	if err := json.Unmarshal([]byte(resp), &data); err != nil {
// 		return nil, err
// 	}
// 	return data, nil
// }
