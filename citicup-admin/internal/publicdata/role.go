package publicdata

type Role int

const (
	Running Role = iota
	User                   // value --> 1
	Investor               // value --> 2
)

func (this Role) String() string {
	switch this {
	case Running:
		return "Running"
	case User:
		return "User"
	case Investor:
		return "Investor"
	default:
		return "Unknow"
	}
}
