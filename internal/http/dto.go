package Http

import "study/pkg/postgres/models"

type CreateDTO struct {
	Title       string `json:"Title"`
	Description string `json:"Description"`
}

type ReadDTO struct {
	Tasks []models.ReadModel
}
