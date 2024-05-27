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

type LotteryListResponse struct {
	CommonResp
	Result struct {
		LotteryList []string `json:"lottery_list"`
	} `json:"result"`
}

type SoldTicket struct {
	ConcertName string `json:"concert_name"`
	TypeName    string `json:"type_name"`
	Seller      string `json:"seller"`
	TokenID     string `json:"token_id"`
}
type TicketListResponse struct {
	CommonResp
	Result []SoldTicket `json:"result"`
}
