package services

type APIResponse struct {
	Status  string      `json:"status"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
}

type Task struct {
	ID          int    `json:"id,omitempty"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Priority    string `json:"priority" binding:"required"`
}

type PatchTask struct {
	ID          int    `json:"id"`
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Priority    string `json:"priority,omitempty"`
}