package main

type Mail struct {
	Domain string
	Host string
	Port int
	Username string
	Password string
	Encryption string
	FromAddress string
	FromName string
}

type Message struct {
	Fromt string
	FromtName string
	To string
	Subject string
	Attachments []string
	Data any
}