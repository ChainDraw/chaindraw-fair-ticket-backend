/*
 * @author: ASWLaunchs
 * @date: 2024/5/20 06:07
 * @description:
 */
package commonreq

type ConcertAddReq struct {
	ConcertID        string   `json:"concert_id"`         // 演唱会ID
	ConcertName      string   `json:"concert_name"`       // 演唱会名称
	Address          string   `json:"address"`            // 演唱会地址
	Remark           string   `json:"remark"`             // 演唱会备注信息
	LotteryStartDate string   `json:"lottery_start_date"` // 抽奖开始时间
	LotteryEndDate   string   `json:"lottery_end_date"`   // 抽奖结束时间
	ConcertDate      string   `json:"concert_date"`       // 演唱会日期
	ConcertImg       string   `json:"concert_img"`        // 演唱会图片
	ConcertStatus    int      `json:"concert_status"`     // 演唱会状态（0: 未开始, 1: 已过期, 2: 已取消）
	ReviewStatus     int      `json:"review_status"`      // 审核状态（0: 未审核, 1: 审核通过, 2: 审核失败）
	Tickets          []Ticket `json:"tickets"`            // 演唱会门票列表
}

type Ticket struct {
	TicketID             string `json:"ticket_id"`               // 门票ID
	TicketType           string `json:"ticket_type"`             // 门票类型唯一键（对应抵押品合约和抽选合约）
	TypeName             string `json:"type_name"`               // 门票类型名称
	Num                  int    `json:"num"`                     // 门票数量
	Price                string `json:"price"`                   // 门票价格
	TicketImg            string `json:"ticket_img"`              // 门票图片
	MaxQuantityPerWallet int    `json:"max_quantity_per_wallet"` // 每个钱包最大可购买数量
	Trade                bool   `json:"trade"`                   // 是否可以二手交易
}
