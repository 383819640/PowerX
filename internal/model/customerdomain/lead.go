package customerdomain

import "PowerX/internal/model"

type Lead struct {
	//Inviter *Customer

	model.Model
	Name        string
	Mobile      string `gorm:"unique"`
	Email       string
	InviterID   int64
	Source      string
	Status      int8
	IsActivated bool
	ExternalId
}

const LeadUniqueId = "mobile"