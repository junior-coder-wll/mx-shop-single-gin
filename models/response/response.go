//@Author: wulinlin
//@Description:
//@File:  response
//@Version: 1.0.0
//@Date: 2023/03/24 16:26

package response

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // omitempty 标签，表示如果 Data 为空，则不在 JSON 数据中显示。
}
