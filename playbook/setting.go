package playbook

type ospot struct {
	o    O
	spot HalfCourtPos
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

//CreateSetupOpen creates "Open" (1-2-2) initail half court setup
func CreateSetupOpen() Setting {
	return CreateSetting(TopOfTheKey, RightWing, LeftWing, RightCorner, LeftCorner, O1)
}
