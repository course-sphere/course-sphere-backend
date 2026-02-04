package http

import "github.com/go-fuego/fuego"

func (h *Handler) Ping(c fuego.ContextNoBody) (string, error) {
	return "pong", nil
}
