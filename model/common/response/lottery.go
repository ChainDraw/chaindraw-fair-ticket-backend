/*
 * @author: biturd
 * @date: 2024/5/12 08:58
 * @description:
 */
package commonresp

type LotteryRecordResponse struct {
	CommonResp
	Result struct {
		RecordID string `json:"record_id"`
	} `json:"result"`
}
