package model

type Student struct {
	Name     string    `json:"name"`
	Subjects []Subject `json:"subjects"`
}
