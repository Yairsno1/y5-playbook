package playbook

type ospot struct {
	o    O
	spot HalfCourtPos
}

func oMoved(o ospot, to HalfCourtPos) ospot {
	return ospot{
		o:    o.o,
		spot: to,
	}
}

//Setting represents the setting of the Os on the court during the play
type Setting struct {
	os   map[ONum]ospot
	ball ONum
}

//Ball gets the ball handler
func (s Setting) Ball() O {
	return s.os[s.ball].o
}

//OSpot finds in a setting the O and its spot on the court for a given role number
func (s Setting) OSpot(role ONum) (O, HalfCourtPos) {
	rv := s.os[role]

	return rv.o, rv.spot
}

//CreateSetting sets the Os on the half court
func CreateSetting(O1spot HalfCourtPos, O2spot HalfCourtPos, O3spot HalfCourtPos, O4spot HalfCourtPos, O5spot HalfCourtPos, ballHandler ONum) Setting {
	rv := Setting{
		os:   make(map[ONum]ospot),
		ball: ballHandler,
	}

	rv.os[O1] = ospot{
		o:    O1Sym,
		spot: O1spot,
	}
	rv.os[O2] = ospot{
		o:    O2Sym,
		spot: O2spot,
	}
	rv.os[O3] = ospot{
		o:    O3Sym,
		spot: O3spot,
	}
	rv.os[O4] = ospot{
		o:    O4Sym,
		spot: O4spot,
	}
	rv.os[O5] = ospot{
		o:    O5Sym,
		spot: O5spot,
	}

	return rv
}

//CreateSetup122 creates 1-2-2 (left and right blocks are filled) initial half court setup
func CreateSetup122() Setting {
	return CreateSetting(TopOfTheKey, RightWing, LeftWing, RightBlock, LeftBlock, O1)
}

//CreateSetup131 creates 1-3-1 initial half court setup
func CreateSetup131() Setting {
	return CreateSetting(TopOfTheKey, RightWing, LeftWing, RightBlock, HighPost, O1)
}

//CreateSetup14High creates high 1-4 initial half court setup
func CreateSetup14High() Setting {
	return CreateSetting(TopOfTheKey, RightWing, LeftWing, RightElbow, LeftElbow, O1)
}

//CreateSetup212 creates 2-1-2 initial half court setup
func CreateSetup212() Setting {
	return CreateSetting(RightGuard, LeftGuard, RightCorner, LeftCorner, HighPost, O1)
}

//CreateSetup23 creates 2-3 initial half court setup
func CreateSetup23() Setting {
	return CreateSetting(RightGuard, LeftGuard, RightWing, LeftWing, HighPost, O1)
}

//CreateSetup4Out creates 4-out 1-in  initial half court setup
func CreateSetup4Out() Setting {
	return CreateSetting(RightGuard, LeftGuard, RightCorner, LeftCorner, RightBlock, O1)
}

//CreateSetupDoubleElbow creates double elbow initial half court setup
func CreateSetupDoubleElbow() Setting {
	return CreateSetting(TopOfTheKey, RightBlock, LeftBlock, RightElbow, LeftElbow, O1)
}

//CreateSetupHighDoubleStack creates high double stack initial half court setup
func CreateSetupHighDoubleStack() Setting {
	return CreateSetting(TopOfTheKey, RightElbow, LeftElbow, RightElbow, LeftElbow, O1)
}

//CreateSetupLowDoubleStack creates low double stack initial half court setup
func CreateSetupLowDoubleStack() Setting {
	return CreateSetting(TopOfTheKey, RightBlock, LeftBlock, RightBlock, LeftBlock, O1)
}

//CreateSetupOpen creates "Open" (1-2-2) initial half court setup
func CreateSetupOpen() Setting {
	return CreateSetting(TopOfTheKey, RightWing, LeftWing, RightCorner, LeftCorner, O1)
}

func copySetting(set Setting) Setting {
	return CreateSetting(set.os[O1].spot, set.os[O2].spot, set.os[O3].spot, set.os[O4].spot, set.os[O5].spot, set.ball)
}
