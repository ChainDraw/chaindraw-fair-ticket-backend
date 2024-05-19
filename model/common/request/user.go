/*
 * @author: biturd
 * @date: 2024/5/8 03:19
 * @description:
 */
package commonreq

import (
	"github.com/ethereum/go-ethereum/common"
	"net/url"
)

type SiweMessage struct {
	Domain         string         `json:"domain"`
	Address        common.Address `json:"address"`
	Uri            url.URL        `json:"uri"`
	Version        string         `json:"version"`
	Statement      *string        `json:"statement"`
	Nonce          string         `json:"nonce"`
	ChainID        int            `json:"chainID"`
	IssuedAt       string         `json:"issuedAt"`
	ExpirationTime *string        `json:"expirationTime"`
	NotBefore      *string        `json:"notBefore"`
	RequestID      *string        `json:"requestID"`
	Resources      []url.URL      `json:"resources"`
}

type LoginReq struct {
}

type RegisterReq struct {
}
