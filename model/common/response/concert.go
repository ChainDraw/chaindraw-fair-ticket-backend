/*
 * @author: biturd
 * @date: 2024/5/12 08:46
 * @description:
 */
package commonresp

type TicketType struct {
	TicketType           string `json:"ticket_type"`             // 门票种类唯一键
	TypeName             string `json:"type_name"`               // 类型名称
	Price                string `json:"price"`                   // 价格
	MaxQuantityPerWallet int    `json:"max_quantity_per_wallet"` // 单个钱包最大购买数量
}
type Concert struct {
	ConcertID     string       `json:"concert_id"`     // 演唱会ID
	ConcertName   string       `json:"concert_name"`   // 演唱会名称
	ConcertImg    string       `json:"concert_img"`    // 演唱会图片
	ConcertDate   string       `json:"concert_date"`   // 演唱会日期
	ConcertStatus int          `json:"concert_status"` // 演唱会状态
	TicketTypes   []TicketType `json:"ticket_types"`   // 门票类型列表
	Status        string       `json:"status"`         // 演唱会状态
}

type ConcertListResponse struct {
	CommonResp
	Result []Concert `json:"result"`
}

type ConcertCancellationResponse struct {
	Code      int          `json:"code"`
	Status    string       `json:"status"`
	Msg       string       `json:"msg"`
	Reason    string       `json:"reason"`
	RequestID string       `json:"request_id"`
	Result    CancelResult `json:"result"`
}

type CancelResult struct {
	ConcertID string `json:"concert_id"`
}
