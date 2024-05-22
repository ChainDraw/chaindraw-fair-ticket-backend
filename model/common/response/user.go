/*
 * @author: biturd
 * @date: 2024/5/8 03:19
 * @description:
 */
package commonresp

type LoginResp struct {
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Msg       string `json:"msg"`
	Reason    string `json:"reason"`
	RequestID string `json:"request_id"`
	Result    struct {
		UserID string `json:"userid"`
	}
}

type RegisterResp struct {
	Code      int    `json:"code"`
	Status    string `json:"status"`
	Msg       string `json:"msg"`
	Reason    string `json:"reason"`
	RequestID string `json:"request_id"`
	Result    struct {
		UserID string `json:"userid"`
	} `json:"result"`
}
