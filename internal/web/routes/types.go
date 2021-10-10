package routes

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
}
