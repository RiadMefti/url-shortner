package models

type URL struct {
	Id_url       string
	Original_url string
}

type URLExists struct {
	Exists bool
	IdUrl  *string 
}
