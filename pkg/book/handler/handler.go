package handler

import (
	"github.com/andhikasamudra/fiber-starter-pack/internal/handler"
	"github.com/andhikasamudra/fiber-starter-pack/pkg/book/service"
	"github.com/go-playground/validator/v10"
)

type Dependency struct {
	BookService service.Interface
	Validator   *validator.Validate
	Resp        handler.ResponseInterface
}

type Handler struct {
	BookService service.Interface
	Validator   *validator.Validate
	Resp        handler.ResponseInterface
}

func NewHandler(d Dependency) *Handler {
	resp := handler.NewResponse()
	return &Handler{
		BookService: d.BookService,
		Validator:   d.Validator,
		Resp:        resp,
	}
}
