package entity

type Notification struct {
	User    User   `json:"user"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
