package message


// Indicates movement from one page to another
type NavMsg struct {
	From int
	To   int
}
