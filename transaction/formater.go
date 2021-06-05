package transaction

import "time"

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
