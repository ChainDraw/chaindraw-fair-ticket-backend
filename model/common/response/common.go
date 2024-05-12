/*
 * @author: biturd
 * @date: 2024/5/12 08:53
 * @description:
 */
package commonresp

type CommonResp struct {
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Msg       string `json:"msg"`
	Reason    string `json:"reason"`
	RequestID string `json:"request_id"`
}
