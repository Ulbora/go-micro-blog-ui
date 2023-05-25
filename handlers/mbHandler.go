package handlers

import (
	lg "github.com/GolangToolKits/go-level-logger"
	m "github.com/Ulbora/go-micro-blog-ui/managers"
)

// MCHandler MCHandler
type MCHandler struct {
	Log         lg.Log
	Manager     m.Manager
	APIKey      string
	APIAdminKey string
	
}

// New New
func (h *MCHandler) New() Handler {
	return h
}
