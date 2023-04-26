package dto

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type CreateBookRequest struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

func (r *CreateBookRequest) Validate() error {
	var emptyFields []string
	if r.Title == "" {
		emptyFields = append(emptyFields, "title")
	}
	if r.Author == "" {
		emptyFields = append(emptyFields, "author")
	}

	if len(emptyFields) > 0 {
		return errors.New(fmt.Sprintf("field is required %s", emptyFields))
	}

	return nil
}

type GetBookResponse struct {
	GUID   uuid.UUID `json:"guid"`
	Title  string    `json:"title"`
	Author string    `json:"author"`
}
