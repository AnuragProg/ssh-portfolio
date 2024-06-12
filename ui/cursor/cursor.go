package cursor

type Cursor string

const (
	BlockCursor Cursor = "█"
	EmptyCursor Cursor = ""
)

type CursorUpdateMsg struct{}

