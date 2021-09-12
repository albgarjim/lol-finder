package input

type TestObj struct {
	Title       string             `bson:"title"`
}

type ContactMail struct {
	Sender string `bson: "sender"`
	Subject string `bson: "subject"`
	Content string `bson: "content"`
	Name string `bson: "name"`
}