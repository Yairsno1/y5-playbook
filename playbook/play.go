package playbook

//Play represents half court set play
type Play struct {
	phases      []*Phase
	name        string
	description string
}

//NewPlay constructs Play value.
func NewPlay(name, desc string, intialSetup Setting) *Play {
	rv := Play{
		phases:      make([]*Phase, 1, 10),
		name:        name,
		description: desc,
	}
	rv.phases[0] = newPhase(intialSetup)

	return &rv
}

//Completed indicates if the play is done and can't be changed.
func (play *Play) Completed() bool {
	count := play.PhaseCount()
	return play.phases[count-1].active
}

//CurrPhase is a getter for the current(last) phase of the play.
func (play Play) CurrPhase() Phase {
	iPhase := len(play.phases) - 1
	rv := play.phases[iPhase]

	return *rv
}

//Description gets the play's decription.
func (play Play) Description() string {
	return play.description
}

//Name gets the name of the play.
func (play Play) Name() string {
	return play.name
}

//PhaseCount indicates the play's number of phases.
func (play Play) PhaseCount() int {
	return len(play.phases)
}

//Add an action of O to the play according to order.
func (play *Play) Add(order int, threat ...IAction) {
	//TODO: order validation
	iPhase := len(play.phases) - 1
	phase := play.phases[iPhase]

	for _, t := range threat {
		phase.addThreat(t, order)
	}
}

//Done blocks the play to accept changes.
//Call to Rollback to unblock.
func (play *Play) Done() {
	count := play.PhaseCount()
	play.phases[count-1].active = false
}

//GetPhase gets phase number phaseNum.
func (play Play) GetPhase(phaseNum int) Phase {
	//TODO: validate index
	rv := play.phases[phaseNum-1]

	return *rv
}

//GoOn completes the current phase and creates the next phase
func (play *Play) GoOn() {
	iPhase := len(play.phases) - 1
	curr := play.phases[iPhase]
	curr.active = false

	set := curr.start
	for _, order := range curr.actions {
		for _, act := range order {
			set = act.Do(set)
		}
	}
	play.phases = append(play.phases, newPhase(set))
}

//Rollback rewinds the play to phase number phaseNum
func (play *Play) Rollback(phaseNum int) {
	count := play.PhaseCount()
	//TODO: validate index

	if phaseNum < count {
		play.phases = play.phases[:phaseNum]
	}

	count = play.PhaseCount()
	play.phases[count-1].active = true
}
