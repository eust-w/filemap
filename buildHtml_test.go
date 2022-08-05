package markdownFolderMap

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)
import "markdownFolderMap/static"

func Test_build(t *testing.T) {
	//temple := static.Temple
	//var payload map[string]interface{}
	payloadString := `{
	"type": "root",
	"depth": 0,
	"content": "",
	"children": [{
		"type": "paragraph",
		"depth": 1,
		"payload": {
			"lines": [1, 2]
		},
		"content": "[toc]"
	}, {
		"type": "heading",
		"depth": 1,
		"payload": {
			"lines": [3, 4]
		},
		"content": "<a href=\"https://github.com/eyssette/myMarkmap/\">什么是分布式系统</a>"
	}, {
		"type": "heading",
		"depth": 1,
		"payload": {
			"lines": [7, 8]
		},
		"content": "为什么要使用分布式系统"
	}, {
		"type": "heading",
		"depth": 1,
		"payload": {
			"lines": [15, 16]
		},
		"content": "CAP定理",
		"children": [{
			"type": "heading",
			"depth": 2,
			"payload": {
				"lines": [19, 20]
			},
			"content": "描述",
			"children": [{
				"type": "list_item",
				"depth": 3,
				"payload": {
					"lines": [23, 24]
				},
				"content": "一致性(Consistency)所有节点访问同一份最新的数据副本"
			}, {
				"type": "list_item",
				"depth": 3,
				"payload": {
					"lines": [24, 25]
				},
				"content": "可用性(Availability)每次请求都能获取到非错的响应"
			}, {
				"type": "list_item",
				"depth": 3,
				"payload": {
					"lines": [25, 26]
				},
				"content": "分区容错性(Partition tolerance)一个分布式系统中出现故障导致分裂成多个计算节点（区）；这些区应能使系统正常运行，一般来说P应该一定要满足的，弱不满足那么就违背了分布式系统初衷"
			}]
		}, {
			"type": "heading",
			"depth": 2,
			"payload": {
				"lines": [27, 28]
			},
			"content": "策略"
		}]
	}, {
		"type": "heading",
		"depth": 1,
		"payload": {
			"lines": [37, 38]
		},
		"content": "BASE理论"
	}, {
		"type": "heading",
		"depth": 1,
		"payload": {
			"lines": [47, 48]
		},
		"content": "一致性算法",
		"children": [{
			"type": "heading",
			"depth": 2,
			"payload": {
				"lines": [51, 52]
			},
			"content": "强一致性算法",
			"children": [{
				"type": "list_item",
				"depth": 3,
				"payload": {
					"lines": [69, 70]
				},
				"content": "网络可靠"
			}, {
				"type": "list_item",
				"depth": 3,
				"payload": {
					"lines": [70, 71]
				},
				"content": "延迟为零"
			}, {
				"type": "list_item",
				"depth": 3,
				"payload": {
					"lines": [71, 72]
				},
				"content": "带宽无限"
			}, {
				"type": "list_item",
				"depth": 3,
				"payload": {
					"lines": [72, 73]
				},
				"content": "网络安全"
			}, {
				"type": "list_item",
				"depth": 3,
				"payload": {
					"lines": [73, 74]
				},
				"content": "拓扑恒定"
			}, {
				"type": "list_item",
				"depth": 3,
				"payload": {
					"lines": [74, 75]
				},
				"content": "单一管理员"
			}, {
				"type": "list_item",
				"depth": 3,
				"payload": {
					"lines": [75, 76]
				},
				"content": "运输成本为零"
			}, {
				"type": "list_item",
				"depth": 3,
				"payload": {
					"lines": [76, 77]
				},
				"content": "网络为同构的"
			}]
		}]
	}],
	"payload": {}
}`
	//p, err := buildHtml(temple, map[string]string{"payload": payload})
	//if err != nil {
	//	panic(err)
	//}
	var payload map[string]interface{}
	err := json.Unmarshal([]byte(payloadString), &payload)
	if err != nil {
		panic(err)
	}
	fmt.Println(payload)
	http.HandleFunc("/te", ppp)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

func ppp(w http.ResponseWriter, r *http.Request) {
	temple := static.Temple
	payloadsss := `{
	"type": "root",
	"content": "",
	"children": [{
		"type": "paragraph",
		"content": "[toc]"
	}, {
		"type": "heading",
		"content": "<a href=\"https://github.com/eyssette/myMarkmap/\">什么是分布式系统</a>"
	}, {
		"type": "heading",
		"content": "为什么要使用分布式系统"
	}, {
		"type": "heading",
		"content": "CAP定理",
		"children": [{
			"type": "heading",
			"depth": 2,
			"content": "描述",
			"children": [{
				"type": "list_item",
				"depth": 3,
				"content": "一致性(Consistency)所有节点访问同一份最新的数据副本"
			}, {
				"type": "list_item",
				"depth": 3,
				"content": "可用性(Availability)每次请求都能获取到非错的响应"
			}, {
				"type": "list_item",
				"depth": 3,
				"content": "分区容错性(Partition tolerance)一个分布式系统中出现故障导致分裂成多个计算节点（区）；这些区应能使系统正常运行，一般来说P应该一定要满足的，弱不满足那么就违背了分布式系统初衷"
			}]
		}, {
			"type": "heading",
			"depth": 2,
			"content": "策略"
		}]
	}, {
		"type": "heading",
		"depth": 1,
		"content": "BASE理论"
	}, {
		"type": "heading",
		"depth": 1,
		"content": "一致性算法",
		"children": [{
			"type": "heading",
			"depth": 2,
			"content": "强一致性算法",
			"children": [{
				"type": "list_item",
				"depth": 3,
				"content": "网络可靠"
			}, {
				"type": "list_item",
				"depth": 3,
				"content": "延迟为零"
			}, {
				"type": "list_item",
				"depth": 3,
				"content": "带宽无限"
			}, {
				"type": "list_item",
				"depth": 3,
				"content": "网络安全"
			}, {
				"type": "list_item",
				"depth": 3,
				"content": "拓扑恒定"
			}, {
				"type": "list_item",
				"depth": 3,
				"payload": {
					"lines": [74, 75]
				},
				"content": "单一管理员"
			}, {
				"type": "list_item",
				"depth": 3,
				"payload": {
					"lines": [75, 76]
				},
				"content": "运输成本为零"
			}, {
				"type": "list_item",
				"depth": 3,
				"payload": {
					"lines": [76, 77]
				},
				"content": "网络为同构的"
			}]
		}]
	}],
	"payload": {}
}`
	var payload map[string]interface{}
	err := json.Unmarshal([]byte(payloadsss), &payload)
	if err != nil {
		panic(err)
	}
	p, err := buildHtml(temple, map[string]map[string]interface{}{"payload": payload})
	if err != nil {
		panic(err)
	}
	fmt.Println("P:", p)
	fmt.Fprintf(w, p)
}
