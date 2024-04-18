package validators

type Task struct {
	Title string `json:"title" binding:"required"`
	Note  string `json:"note"`
}
