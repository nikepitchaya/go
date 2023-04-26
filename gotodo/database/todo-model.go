package database

type Todo struct {
	id uint `json:"id" gorm:"primary_key"`
	Username string `json:"username"` //ผูกกับตัว JSON
	Title    string `json:"title"`
	Message  string `json:"message"`
}
