package playbook

//IExecution defiens threat exectution by O
type IExecution interface {
	Shot() IAction
	Pass(teamate O) IAction
	Dribble(spot HalfCourtPos) IAction
	PenetrationDribble(spot HalfCourtPos) IAction
	HandOff(teamate O) IAction
	Cut(spot HalfCourtPos) IAction
	CurlCut(spot HalfCourtPos) IAction
	FlashCut(spot HalfCourtPos) IAction
	GiveAndGo(spot HalfCourtPos) IAction
	GiveAndGoBehind(spot HalfCourtPos) IAction
	VCut(spot HalfCourtPos) IAction
	LCut(spot HalfCourtPos) IAction
	Backdoor(spot HalfCourtPos) IAction
	UCLACut(spot HalfCourtPos) IAction
	Move(spot HalfCourtPos) IAction
	Screen(spot HalfCourtPos, teamate O) IAction
	PickAndRoll(spot HalfCourtPos, teamate O) IAction
	DownScreen(spot HalfCourtPos, teamate O) IAction
	BackScreen(spot HalfCourtPos, teamate O) IAction
	Flare(spot HalfCourtPos, teamate O) IAction
	ScreenAndGo(spot HalfCourtPos, teamate O) IAction
	DoubleScreen(spot HalfCourtPos, teamate O) IAction
	StaggeredScreen(spot HalfCourtPos, teamate O) IAction
}

//Shot represents O takes a shot.
func (o O) Shot() IAction {
	return ShotAction{
		o: o,
	}
}

//Pass to a teamate.
func (o O) Pass(teamate O) IAction {
	return PassAction{
		o:       o,
		threat:  Pass,
		teamate: teamate,
	}
}

//Dribble represents an O dribbles.
func (o O) Dribble(spot HalfCourtPos) IAction {
	return DribbleAction{
		o:      o,
		threat: Dribble,
		spot:   spot,
	}
}

//PenetrationDribble represents an O penetrates.
func (o O) PenetrationDribble(spot HalfCourtPos) IAction {
	return DribbleAction{
		o:      o,
		threat: PenetrationDribble,
		spot:   spot,
	}
}

//HandOff the ball to a teamate.
func (o O) HandOff(teamate O) IAction {
	return PassAction{
		o:       o,
		threat:  HandOff,
		teamate: teamate,
	}
}

//Cut represents an O cuts.
func (o O) Cut(spot HalfCourtPos) IAction {
	return CutAction{
		o:      o,
		threat: Cut,
		spot:   spot,
	}
}

//CurlCut represents an O curls off a scrren.
func (o O) CurlCut(spot HalfCourtPos) IAction {
	return CutAction{
		o:      o,
		threat: CurlCut,
		spot:   spot,
	}
}

//FlashCut represents an O executing a flash cut.
func (o O) FlashCut(spot HalfCourtPos) IAction {
	return CutAction{
		o:      o,
		threat: FlashCut,
		spot:   spot,
	}
}

//GiveAndGo represents an O executing a give&go move.
func (o O) GiveAndGo(spot HalfCourtPos) IAction {
	return CutAction{
		o:      o,
		threat: GiveAndGo,
		spot:   spot,
	}
}

//GiveAndGoBehind represents an O executing a give&go behind move.
func (o O) GiveAndGoBehind(spot HalfCourtPos) IAction {
	return CutAction{
		o:      o,
		threat: GiveAndGoBehind,
		spot:   spot,
	}
}

//VCut represents an O doing a v-cut.
func (o O) VCut(spot HalfCourtPos) IAction {
	return MoveAction{
		o:      o,
		threat: VCut,
		spot:   spot,
	}
}

//LCut represents an O doing a L-cut.
func (o O) LCut(spot HalfCourtPos) IAction {
	return MoveAction{
		o:      o,
		threat: LCut,
		spot:   spot,
	}
}

//Backdoor represents an O executing a backdoor cut.
func (o O) Backdoor(spot HalfCourtPos) IAction {
	return CutAction{
		o:      o,
		threat: Backdoor,
		spot:   spot,
	}
}

//UCLACut represents an O executing a UCLA cut off back screen at the high post.
func (o O) UCLACut(spot HalfCourtPos) IAction {
	return CutAction{
		o:      o,
		threat: UCLACut,
		spot:   spot,
	}
}

//Move represents an O moving.
func (o O) Move(spot HalfCourtPos) IAction {
	return MoveAction{
		o:      o,
		threat: Move,
		spot:   spot,
	}
}

//Screen represents an O setting a screen to a teamate at spot.
func (o O) Screen(spot HalfCourtPos, teamate O) IAction {
	return ScreenAction{
		o:       o,
		threat:  Screen,
		spot:    spot,
		teamate: teamate,
	}
}

//PickAndRoll represents an O doing pick&roll with the ball handler.
func (o O) PickAndRoll(spot HalfCourtPos, teamate O) IAction {
	return ScreenAction{
		o:       o,
		threat:  PickAndRoll,
		spot:    spot,
		teamate: teamate,
	}
}

//DownScreen represents an O setting a down screen.
func (o O) DownScreen(spot HalfCourtPos, teamate O) IAction {
	return ScreenAction{
		o:       o,
		threat:  DownScreen,
		spot:    spot,
		teamate: teamate,
	}
}

//BackScreen represents an O setting a back screen.
func (o O) BackScreen(spot HalfCourtPos, teamate O) IAction {
	return ScreenAction{
		o:       o,
		threat:  BackScreen,
		spot:    spot,
		teamate: teamate,
	}
}

//Flare represents an O setting a flare screen.
func (o O) Flare(spot HalfCourtPos, teamate O) IAction {
	return ScreenAction{
		o:       o,
		threat:  Flare,
		spot:    spot,
		teamate: teamate,
	}
}

//ScreenAndGo represents an O starts setting a screen and then cut to the rim.
func (o O) ScreenAndGo(spot HalfCourtPos, teamate O) IAction {
	return ScreenAction{
		o:       o,
		threat:  ScreenAndGo,
		spot:    spot,
		teamate: teamate,
	}
}

//DoubleScreen represents an O takes part in setting a double screen.
func (o O) DoubleScreen(spot HalfCourtPos, teamate O) IAction {
	return ScreenAction{
		o:       o,
		threat:  DoubleScreen,
		spot:    spot,
		teamate: teamate,
	}
}

//StaggeredScreen represents an O takes part in setting a staggered screen.
func (o O) StaggeredScreen(spot HalfCourtPos, teamate O) IAction {
	return ScreenAction{
		o:       o,
		threat:  StaggeredScreen,
		spot:    spot,
		teamate: teamate,
	}

}
