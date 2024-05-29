/*
 * @author: biturd
 * @date: 2024/5/12 08:46
 * @description:
 */
package commonresp

type TicketType struct {
	TicketType           int64   `json:"ticket_type"`             // 门票种类唯一键
	TypeName             string  `json:"type_name"`               // 类型名称
	Num                  int     `json:"num"`                     //数量
	TicketImg            string  `json:"ticket_img"`              // 票图片
	Trade                bool    `json:"trade"`                   //是否交易
	Price                float64 `json:"price"`                   // 价格
	MaxQuantityPerWallet int64   `json:"max_quantity_per_wallet"` // 单个钱包最大购买数量
	CreateAt             string  `json:"create_at"`               // 创建时间
	UpdateAt             string  `json:"update_at"`               // 更新时间
}
type Concert struct {
	ConcertID        string       `json:"concert_id"`         // 演唱会ID
	ConcertName      string       `json:"concert_name"`       // 演唱会名称
	Address          string       `json:"address"`            //演唱会地点
	ConcertImg       string       `json:"concert_img"`        // 演唱会图片
	ConcertDate      string       `json:"concert_date"`       // 演唱会日期
	LotteryStartDate string       `json:"lottery_start_date"` //彩票开始日期
	LotteryEndDate   string       `json:"lottery_end_date"`   //彩票开始日期
	ReviewStatus     int          `json:"review_status"`      //审核状态
	CancelReason     string       `json:"cancel_reason"`      //取消原因
	ConcertStatus    int          `json:"concert_status"`     // 演唱会状态
	Remark           string       `json:"remark"`             //演唱会备注
	Status           string       `json:"status"`             // 演唱会状态
	TicketTypes      []TicketType `json:"ticket_types"`       // 门票类型列表
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
