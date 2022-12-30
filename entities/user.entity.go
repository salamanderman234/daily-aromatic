package entity

type User struct {
	ID            uint           `json:"id"`
	Username      string         `json:"username"`
	FollowerTotal int            `json:"follower_total"`
	ReviewTotal   int            `json:"review_total"`
	Bio           string         `json:"bio"`
	Notifications []Notification `json:"notifications"`
	Reviews       []Review       `json:"reviews"`
}
