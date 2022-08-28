package dataoke

// 拼多多API

// 拼多多联盟搜索
func (c *Client) PddGoodsSearch(params map[string]string) (string, error) {
	return c.Get("/dels/pdd/goods/search", params)
}
