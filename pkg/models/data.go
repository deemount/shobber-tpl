package models

type Data struct {
	PageTitle       string
	PageAuthor      string
	PageDescription string
	NavBarLinks     []Link
	OG              []OpenGraph
	Widgets         []Widget
}
