package navigation

var (
	groupLastIdentifier = int64(0)
)

type Group struct {
	Identifier int64
	Title      string
	Priority   int64
}

func NewGroup(title string, priority int64) *Group {
	groupLastIdentifier++

	return &Group{
		Identifier: groupLastIdentifier,
		Title:      title,
		Priority:   priority,
	}
}

type Link struct {
	Title      string
	Priority   int64
	Identifier string
}

func NewLink(title string, priority int64) *Link {
	return &Link{
		Title:    title,
		Priority: priority,
	}
}
