/*
 * @author: asta
 * @date: 2024/5/12 08:46
 * @description:
 */
package commonreq

type ConcertReViewReq struct {
	ConcertID string `json:"concert_id"` // 提交的演唱会id
	Pass      bool   `json:"pass"`       // 是否通过
}
