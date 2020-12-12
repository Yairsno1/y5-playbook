/*
Package playbook doc
under construction */
package playbook

import "fmt"

//HalfCourtPos defines half court spots
type HalfCourtPos uint8

//Enumeration of HalfCourtPos
const (
	RightGuard HalfCourtPos = iota
	TopOfTheKey
	LeftGuard
	RightWing
	RightElbow
	HighPost
	LeftElbow
	LeftWing
	RightMidPost
	InCircle
	LeftMidPost
	RightCorner
	RightShortCorner
	RightBlock
	UnderTheRim
	LeftBlock
	LeftShortCorner
	LeftCorner
)

func (spot HalfCourtPos) String() string {
	switch spot {
	case RightGuard:
		return fmt.Sprintf("(R.) Guard")
	case TopOfTheKey:
		return fmt.Sprintf("Top of the 3pt. Arc")
	case LeftGuard:
		return fmt.Sprintf("(L.) Guard")
	case RightWing:
		return fmt.Sprintf("(R.) Wing")
	case RightElbow:
		return fmt.Sprintf("(R.) Elbow")
	case HighPost:
		return fmt.Sprintf("High post")
	case LeftElbow:
		return fmt.Sprintf("(L.) Elbow")
	case LeftWing:
		return fmt.Sprintf("(L.) Wing")
	case RightMidPost:
		return fmt.Sprintf("(R.) Mid post")
	case InCircle:
		return fmt.Sprintf("Inner circle")
	case LeftMidPost:
		return fmt.Sprintf("(L.) Mid post")
	case RightCorner:
		return fmt.Sprintf("(R.) Corner")
	case RightShortCorner:
		return fmt.Sprintf("(R.) Short corner")
	case RightBlock:
		return fmt.Sprintf("(R.) Block")
	case UnderTheRim:
		return fmt.Sprintf("Under the rim")
	case LeftBlock:
		return fmt.Sprintf("(L.) Block")
	case LeftShortCorner:
		return fmt.Sprintf("(L.) Short corner")
	case LeftCorner:
		return fmt.Sprintf("(L.) Corner")
	default:
		return fmt.Sprintf("Unknown!?") //TODO: error
	}
}

//ONum defines player ("O") role number
type ONum uint8

//O numbers 1, 2, 3, 4, 5
const (
	O1 ONum = iota + 1
	O2
	O3
	O4
	O5
)

func (n ONum) String() string {
	switch n {
	case O1:
		return fmt.Sprintf("1")
	case O2:
		return fmt.Sprintf("2")
	case O3:
		return fmt.Sprintf("3")
	case O4:
		return fmt.Sprintf("4")
	case O5:
		return fmt.Sprintf("5")
	default:
		return fmt.Sprintf("Unknown!?") //TODO: error
	}
}

//Threat defines (O)ffensive player threats, moves and options.
//For example: pass, dribble, cut, screen etc.
type Threat uint8

//Enumeration of offensive player threats.
const (
	Shot Threat = iota
	Pass
	Dribble
	PenetrationDribble
	HandOff
	Cut
	CurlCut
	FlashCut
	GiveAndGo
	GiveAndGoBehind
	VCut
	LCut
	Backdoor
	UCLACut
	Move
	Screen
	PickAndRoll
	DownScreen
	BackScreen
	Flare
	ScreenAndGo
	DoubleScreen
	StaggeredScreen
)

func (t Threat) String() string {
	switch t {
	case Shot:
		return fmt.Sprintf("Shot")
	case Pass:
		return fmt.Sprintf("Pass")
	case Dribble:
		return fmt.Sprintf("Dribble")
	case PenetrationDribble:
		return fmt.Sprintf("Penetration")
	case HandOff:
		return fmt.Sprintf("Hand-off")
	case Cut:
		return fmt.Sprintf("Cut")
	case CurlCut:
		return fmt.Sprintf("Curl cut")
	case FlashCut:
		return fmt.Sprintf("Flash cut")
	case GiveAndGo:
		return fmt.Sprintf("Give&go")
	case GiveAndGoBehind:
		return fmt.Sprintf("Give&go behind")
	case VCut:
		return fmt.Sprintf("V-cut")
	case LCut:
		return fmt.Sprintf("L-cut")
	case Backdoor:
		return fmt.Sprintf("Backdoor")
	case UCLACut:
		return fmt.Sprintf("UCLA cut")
	case Move:
		return fmt.Sprintf("Move")
	case Screen:
		return fmt.Sprintf("Screen")
	case PickAndRoll:
		return fmt.Sprintf("Pick&roll")
	case DownScreen:
		return fmt.Sprintf("Down screen")
	case BackScreen:
		return fmt.Sprintf("Back screen")
	case Flare:
		return fmt.Sprintf("Flare screen")
	case ScreenAndGo:
		return fmt.Sprintf("Screen&go")
	case DoubleScreen:
		return fmt.Sprintf("Double screen")
	case StaggeredScreen:
		return fmt.Sprintf("Staggered screen")
	default:
		return fmt.Sprintf("%q", int(t))
	}
}

//O represents O symbol in a play
type O struct {
	//O offensive role as a number, 1 to 5
	role ONum
}

//Role getter
func (o O) Role() ONum {
	return o.role
}

func (o O) String() string {
	return fmt.Sprintf("O%s", o.role)
}

/*
O values declerations as FlyWeight pattern, value for each O role.
Additional variables of type O are not necessary, other properties for O
like spot, threat and so on will be defined through other types.
*/
var (
	O1Sym O = O{
		role: O1,
	}
	O2Sym O = O{
		role: O2,
	}
	O3Sym O = O{
		role: O3,
	}
	O4Sym O = O{
		role: O4,
	}
	O5Sym O = O{
		role: O5,
	}
)
