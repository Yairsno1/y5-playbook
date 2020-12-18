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

//Add an action of O to the play according to order.
func (play *Play) Add(order int, threat ...IAction) {
	//TODO: order validation
	iPhase := len(play.phases) - 1
	phase := play.phases[iPhase]

	for _, t := range threat {
		phase.addThreat(t, order)
	}
}

//CurrPhase is a getter for the current(last) phase of the play.
func (play Play) CurrPhase() Phase {
	iPhase := len(play.phases) - 1
	rv := play.phases[iPhase]

	return *rv
}

//PhaseCount indicates the play's number of phases.
func (play Play) PhaseCount() int {
	return len(play.phases)
}
