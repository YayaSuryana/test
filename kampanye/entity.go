package kampanye

import (
	"time"
	"yayasuryana/user"
)

type Kampanye struct{
	ID 					int
	UserID 				int
	Name				string
	ShortDescription 	string
	Description 		string
	Perks				string
	BackerCount			int
	GoalAmount			int
	CurrentAmount		int
	Slug				string
	CreatedAt			time.Time
	UpdatedAt			time.Time
	KampanyeImages		[]KampanyeImage
	User 				user.User
}

type KampanyeImage struct{
	ID 				int
	KampanyeID		int
	FileName		string
	IsPrimary		int
	CreatedAt		time.Time
	UpdatedAt		time.Time
}