/*
 * @author: ASWLaunchs
 * @date: 2024/5/20 06:07
 * @description:
 */
package commonresp

type CommitResp struct {
	Code      int         `json:"code"`       // 返回状态码
	Status    string      `json:"status"`     // 返回状态
	Msg       string      `json:"msg"`        // 返回消息
	Reason    string      `json:"reason"`     // 错误原因
	RequestID string      `json:"request_id"` // 请求ID
	Result    *ResultData `json:"result"`     // 返回结果
}

type ResultData struct {
	ConcertID string `json:"concert_id"` // 演唱会ID
}
