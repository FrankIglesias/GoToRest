package model

type Student struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Subjects []Subject `json:"subjects"`
}
