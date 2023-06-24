package model

type Person struct {
	Name  string `json:"Name"`
	Age   int
	Hobby []string
}

type Developer interface {
	Code() string
	Build() string
	Test() string
	Deploy() string
}

func (p Person) Code() string {
	return "coding"
}
func (p Person) Build() string {
	return "building..."
}
func (p Person) Test() string {
	return "testing..."
}
func (p Person) Deploy() string {
	return "deploying..."
}
