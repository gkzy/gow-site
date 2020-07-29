package md

import "strings"

type MD struct {
	URL      string
	FilePath string
	Title    string
}

//GetMDFile GetMDFile
func GetMDFile(url string) *MD {
	data := GetMdFiles()

	for _, item := range data {
		if strings.ToLower(item.URL) == strings.ToLower(url) {
			return item
		}
	}
	return nil
}

//GetMdFiles
func GetMdFiles() (data []*MD) {
	data = make([]*MD, 0)
	data = append(data, &MD{
		URL:      "/docs/start",
		FilePath: "./md/start.md",
		Title:    "快速开始",
	})
	data = append(data, &MD{
		URL:      "/docs/html",
		FilePath: "./md/html.md",
		Title:    "做一个网站&HTML模板",
	})
	data = append(data, &MD{
		URL:      "/docs/request",
		FilePath: "./md/request.md",
		Title:    "获取请求值",
	})
	data = append(data, &MD{
		URL:      "/docs/response",
		FilePath: "./md/response.md",
		Title:    "输出值",
	})
	data = append(data, &MD{
		URL:      "/docs/router",
		FilePath: "./md/router.md",
		Title:    "路由详解",
	})
	data = append(data, &MD{
		URL:      "/docs/config",
		FilePath: "./md/config.md",
		Title:    "配置文件",
	})
	data = append(data, &MD{
		URL:      "/docs/cache",
		FilePath: "./md/cache.md",
		Title:    "内存缓存",
	})
	data = append(data, &MD{
		URL:      "/docs/func",
		FilePath: "./md/func.md",
		Title:    "HTML 模板函数",
	})
	data = append(data, &MD{
		URL:      "/docs/middleware",
		FilePath: "./md/middleware.md",
		Title:    "中间件实现&权限",
	})
	data = append(data, &MD{
		URL:      "/docs/lib",
		FilePath: "./md/lib.md",
		Title:    "lib库",
	})
	return data
}
