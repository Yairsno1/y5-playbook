/*
Package playbook doc
under construction */
package playbook

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

//O represents O symbol in a play
type O struct {
	//O offensive role as a number, 1 to 5
	role ONum
}

//Role getter
func (o O) Role() ONum {
	return o.role
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
