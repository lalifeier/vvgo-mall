package dataoke

// 京东API

// 京东联盟搜索
func (c *Client) JdGoodsSearch(params map[string]string) (string, error) {
	return c.Get("/dels/jd/goods/search", params)
}
