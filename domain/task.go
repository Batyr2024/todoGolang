package domain

type Task struct {
	ID      int32  `db:"id" json:"id"`
	Text    string `db:"text" json:"text"`
	Checked bool   `db:"checked" json:"checked"`
}
