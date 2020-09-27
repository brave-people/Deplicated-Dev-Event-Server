package requests

type Event struct {
	ID          uint   `json:"id"`
	Description string `json:"nickname"`
	Tages       []Tag  `json:"tages"`
}
