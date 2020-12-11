package playbook

//Action represents threats and movements by O
type Action struct {
	o       O
	threat  Threat
	teamate O
	spot    HalfCourtPos
}

//Phase represents phase of a play
type Phase struct {
	start   Setting
	actions [][]IAction
	active  bool
}

func newPhase(setting Setting) *Phase {
	return &Phase{
		start:   setting,
		actions: make([][]IAction, 0, 5),
		active:  true,
	}
}

func (p *Phase) addThreat(threat IAction, order int) {
	//TODO: active validation
	orderCount := len(p.actions)
	//TODO: validate order

	i := order - 1
	if order == orderCount {
		p.actions[i] = append(p.actions[i], threat)
	} else if order > orderCount {
		p.actions = append(p.actions, make([]IAction, 1, 10))
		p.actions[i][0] = threat
	}
}

func (p *Phase) completed() {
	p.active = false
}

//IAction defines the Os actions during a phase of a play
type IAction interface {
	Do(set Setting) Setting
}

//CutAction represents cutting
type CutAction Action

//Do creates the new setting as reault of O cutting.
func (cutter CutAction) Do(set Setting) Setting {
	rv := copySetting(set)

	currSpot := rv.os[cutter.o.role]
	rv.os[cutter.o.role] = oMoved(currSpot, cutter.spot)

	return rv
}

//DribbleAction represents dribble O
type DribbleAction Action

//Do creates the new setting as reault of O dribbling.
func (dribble DribbleAction) Do(set Setting) Setting {
	rv := copySetting(set)

	currSpot := rv.os[dribble.o.role]
	rv.os[dribble.o.role] = oMoved(currSpot, dribble.spot)

	return rv
}

//MoveAction represents O moving
type MoveAction Action

//Do creates the new setting as reault of O moving.
func (move MoveAction) Do(set Setting) Setting {
	rv := copySetting(set)

	currSpot := rv.os[move.o.role]
	rv.os[move.o.role] = oMoved(currSpot, move.spot)

	return rv
}

//PassAction represents passing the ball
type PassAction Action

//Do creates the new setting as reault of a pass.
//Since pass is not about movement, Do() actually creates the same setting but
//with different ball handler.
func (pass PassAction) Do(set Setting) Setting {
	//TODO: validate the ball handler, passer<->set.ball
	rv := copySetting(set)
	rv.ball = pass.teamate.role

	return rv
}

//ScreenAction represents off the ball screen
type ScreenAction Action

//Do creates the new setting as reault of a screen.
func (screener ScreenAction) Do(set Setting) Setting {
	rv := copySetting(set)

	currSpot := rv.os[screener.o.role]
	rv.os[screener.o.role] = oMoved(currSpot, screener.spot)

	return rv
}

//ShotAction represents O shooting
type ShotAction Action

//Do creates the new setting as reault of a shot.
//Since shooting is not about movement, Do() actually creates the same setting.
func (shot ShotAction) Do(set Setting) Setting {
	//TODO: validate the shooter, shooter<->set.ball
	return copySetting(set)
}
