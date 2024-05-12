/*
 * @author: biturd
 * @date: 2024/5/12 08:46
 * @description:
 */
package commonreq

type LotteryRecordReq struct {
	ConcertID     string `json:"concert_id"`     // 提交的演唱会id
	TypeName      string `json:"type_name"`      // 票种类型
	TicketType    string `json:"ticket_type"`    // 票种类型
	Quantity      int    `json:"quantity"`       // 登记抽票数量
	WalletAddress string `json:"wallet_address"` // 抽票地址
}
