package dataoke

import "encoding/json"

// 淘宝API

// 商品列表
func (c *Client) GetGoodsList(params map[string]string) (string, error) {
	return c.Get("/goods/get-goods-list", params)
}

// 单品详情
func (c *Client) GetGoodsDetails(params map[string]string) (string, error) {
	return c.Get("/goods/get-goods-details", params)
}

// 超级分类
func (c *Client) GetSuperCategory(params map[string]string) (*GetSuperCategoryResp, error) {
	resp, err := c.Get("/category/get-super-category", params)
	if err != nil {
		return nil, err
	}
	var data GetSuperCategoryResp
	if err := json.Unmarshal([]byte(resp), &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 轮播图
func (c *Client) CarouseList(params map[string]string) (*CarouseListResp, error) {
	resp, err := c.Get("/goods/topic/carouse-list", params)
	if err != nil {
		return nil, err
	}
	var data CarouseListResp
	if err := json.Unmarshal([]byte(resp), &data); err != nil {
		return nil, err
	}
	return &data, nil
}

// 高效转链
func (c *Client) GetPrivilegeLink(params map[string]string) (string, error) {
	return c.Get("/tb-service/get-privilege-link", params)
}

// 超级搜索
func (c *Client) ListSuperGoods(params map[string]string) (string, error) {
	return c.Get("/goods/list-super-goods", params)
}

// 大淘客搜索
func (c *Client) GetDtkSearchGoods(params map[string]string) (string, error) {
	return c.Get("/goods/get-dtk-search-goods", params)
}

// 联盟搜索
func (c *Client) GetTbService(params map[string]string) (string, error) {
	return c.Get("/tb-service/get-tb-service", params)
}

// 联想词
func (c *Client) SearchSuggestion(params map[string]string) (string, error) {
	return c.Get("/goods/search-suggestion", params)
}

// 热搜记录
func (c *Client) GetTop100(params map[string]string) (string, error) {
	return c.Get("/category/get-top100", params)
}

// 热搜记录
func (c *Client) GetRankingList(params map[string]string) (string, error) {
	return c.Get("/goods/get-goods-details", params)
}
