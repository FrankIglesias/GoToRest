package model

type Student struct {
	ID       int       `json:"id"`
	Name     string    `json:"name"`
	Subjects []Subject `json:"subjects"`
}
