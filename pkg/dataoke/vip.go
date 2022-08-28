package dataoke

// 唯品会 API

// 唯品会联盟搜索
func (c *Client) VIPGoodsSearch(params map[string]string) (string, error) {
	return c.Get("/vip/search-by-keywords", params)
}
