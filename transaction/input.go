package transaction

import "bwastartup/user"

type GetTransactionCampaignDetailInput struct {
	Id   int `uri:"id" binding:"required"`
	User user.User
}
