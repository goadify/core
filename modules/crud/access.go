package crud

const (
	CreateRuleName = "create"
	ReadRuleName   = "read"
	UpdateRuleName = "update"
	DeleteRuleName = "delete"
)

type AccessRule interface {
	RuleName() string
}

type AccessRead struct{}

func (a *AccessRead) RuleName() string {
	return ReadRuleName
}

type AccessCreate struct{}

func (a *AccessCreate) RuleName() string {
	return CreateRuleName
}

type AccessUpdate struct{}

func (a *AccessUpdate) RuleName() string {
	return UpdateRuleName
}

type AccessDelete struct{}

func (a *AccessDelete) RuleName() string {
	return DeleteRuleName
}
