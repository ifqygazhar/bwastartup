package transaction

import (
	"time"
)

type CampaignTransactionFormater struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Amount   int       `json:"amount"`
	CreateAt time.Time `json:"created_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormater {
	formatter := CampaignTransactionFormater{}
	formatter.Id = transaction.Id
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreateAt = transaction.CreateAt
	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormater {
	if len(transactions) == 0 {
		return []CampaignTransactionFormater{}
	}

	var transactionsFormater []CampaignTransactionFormater

	for _, transaction := range transactions {
		formater := FormatCampaignTransaction(transaction)
		transactionsFormater = append(transactionsFormater, formater)
	}
	return transactionsFormater
}

type UserTransactionFormater struct {
	Id       int              `json:"id"`
	Name     string           `json:"name"`
	Amount   int              `json:"amount"`
	Status   string           `json:"status"`
	CreateAt time.Time        `json:"created_at"`
	Campaign CampaignFormater `json:"campaign"`
}

type CampaignFormater struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormater {
	formater := UserTransactionFormater{}
	formater.Id = transaction.Id
	formater.Amount = transaction.Amount
	formater.Status = transaction.Status

	formater.CreateAt = transaction.CreateAt

	campaignFormater := CampaignFormater{}
	campaignFormater.Name = transaction.Campaign.Name
	campaignFormater.ImageURL = ""

	if len(transaction.Campaign.CampaignImages) > 0 {
		campaignFormater.ImageURL = transaction.Campaign.CampaignImages[0].FileName
	}

	formater.Campaign = campaignFormater
	return formater
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormater {
	if len(transactions) == 0 {
		return []UserTransactionFormater{}
	}

	var transactionsFormater []UserTransactionFormater

	for _, transaction := range transactions {
		formater := FormatUserTransaction(transaction)
		transactionsFormater = append(transactionsFormater, formater)
	}
	return transactionsFormater
}

type TransactionFormater struct {
	Id         int    `json:"id"`
	CampaignId int    `json:"campaign_id"`
	UserId     int    `json:"user_id"`
	Amount     int    `json:"amount"`
	Status     string `json:"status"`
	Code       string `json:"code"`
	PaymentURL string `json:"payment_url"`
}

func FormatTransaction(transaction Transaction) TransactionFormater {
	formatter := TransactionFormater{}
	formatter.Id = transaction.Id
	formatter.CampaignId = transaction.CampaignId
	formatter.UserId = transaction.User.Id
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.Code = transaction.Code
	formatter.PaymentURL = transaction.PaymentURL
	return formatter
}
