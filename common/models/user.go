package models

type User struct {
	ID           uint     `json:"id"`
	Username     string   `json:"username"`
	Age          uint8    `json:"age"`
	ListOfSeries []Series `json:"list_of_series,omitempty"`
	TimeSpent    uint     `json:"time_spent,omitempty"`
}

func (user *User) CalculateTotalTimeSpent() uint {
	for _, series := range user.ListOfSeries {
		user.TimeSpent += series.CalculateTotalDuration()
	}

	return user.TimeSpent
}
