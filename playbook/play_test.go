package playbook

import "testing"

func Test_NewPlay(t *testing.T) {
	//Arrange + Act
	play := NewPlay("2 down", "Double scrren to low post", Setting{})

	got := play.PhaseCount()
	want := 1
	if got != want {
		t.Errorf("Phase order is %q, want %q", got, want)
	}
}

func Test_Add(t *testing.T) {
	//Arrange
	play := NewPlay("2 down", "Double scrren to low post", Setting{})

	//Act
	/*o2Vcut := O2Sym.VCut(RightWing)
	o1PassToO2 := O1Sym.Pass(O2Sym)
	o4ScreenToO1 := O4Sym.BackScreen(RightElbow, O1Sym)
	o1UCLACut := O1Sym.UCLACut(RightBlock)*/
	order3 := []IAction{O4Sym.BackScreen(RightElbow, O1Sym), O1Sym.UCLACut(RightBlock)}

	play.Add(1, O2Sym.VCut(RightWing))
	play.Add(2, O1Sym.Pass(O2Sym))
	play.Add(3, order3...)

	//Assert
	got := len(play.phases[0].actions)
	want := 3
	if got != want {
		t.Errorf("Play 1st phase order is %q, want %q", got, want)
	}
}

func Test_GoOn(t *testing.T) {
	//Arrange
	setting14 := CreateSetting(TopOfTheKey, RightWing, LeftWing, RightElbow, LeftElbow, O1) //1-4
	play := NewPlay("2 down", "Double scrren to low post", setting14)

	play.Add(1, O2Sym.VCut(RightWing))
	play.Add(2, O1Sym.Pass(O2Sym))
	play.Add(3, O4Sym.BackScreen(RightElbow, O1Sym), O1Sym.UCLACut(RightBlock))
	play.Add(4, O4Sym.Move(TopOfTheKey))

	//Act
	play.GoOn()
	gotSet := play.CurrPhase().start

	//Assert
	_, got := gotSet.OSpot(O1)
	want := RightBlock
	if got != want {
		t.Errorf("1 spot is %q, want %q", got, want)
	}
	_, got = gotSet.OSpot(O2)
	want = RightWing
	if got != want {
		t.Errorf("2 spot is %q, want %q", got, want)
	}
	_, got = gotSet.OSpot(O3)
	want = LeftWing
	if got != want {
		t.Errorf("3 spot is %q, want %q", got, want)
	}
	_, got = gotSet.OSpot(O4)
	want = TopOfTheKey
	if got != want {
		t.Errorf("4 spot is %q, want %q", got, want)
	}
	_, got = gotSet.OSpot(O5)
	want = LeftElbow
	if got != want {
		t.Errorf("5 spot is %q, want %q", got, want)
	}
	gotBall := gotSet.ball
	wantBall := O2
	if gotBall != wantBall {
		t.Errorf("Ball handler is %q, want %q", gotBall, wantBall)
	}

}

func Test_Rollback(t *testing.T) {
	//Arrange
	setting14 := CreateSetting(TopOfTheKey, RightWing, LeftWing, RightElbow, LeftElbow, O1) //1-4

	play := NewPlay("2 down", "Double scrren to low post", setting14)
	play.Add(1, O2Sym.VCut(RightWing))
	play.Add(2, O1Sym.Pass(O2Sym))
	play.Add(3, O4Sym.BackScreen(RightElbow, O1Sym), O1Sym.UCLACut(RightBlock))
	play.Add(4, O4Sym.Move(TopOfTheKey))

	play.GoOn()

	//Act
	play.Rollback(1)

	//Assert
	got := play.PhaseCount()
	want := 1
	if got != want {
		t.Errorf("Play has %q phases, want %q", got, want)

	}
	gotBall := play.CurrPhase().start.ball
	wantBall := O1
	if gotBall != wantBall {
		t.Errorf("Ball handler is %q, want %q", gotBall, wantBall)
	}
	if !play.CurrPhase().active {
		t.Errorf("Current phase is not active")
	}
}

func Test_Rollback_Last(t *testing.T) {
	//Arrange
	setting14 := CreateSetting(TopOfTheKey, RightWing, LeftWing, RightElbow, LeftElbow, O1) //1-4

	play := NewPlay("2 down", "Double scrren to low post", setting14)
	play.Add(1, O2Sym.VCut(RightWing))
	play.Add(2, O1Sym.Pass(O2Sym))
	play.Add(3, O4Sym.BackScreen(RightElbow, O1Sym), O1Sym.UCLACut(RightBlock))
	play.Add(4, O4Sym.Move(TopOfTheKey))

	play.GoOn()

	//Act
	play.Rollback(2)

	//Assert
	got := play.PhaseCount()
	want := 2
	if got != want {
		t.Errorf("Play has %q phases, want %q", got, want)

	}
	gotBall := play.CurrPhase().start.ball
	wantBall := O2
	if gotBall != wantBall {
		t.Errorf("Ball handler is %q, want %q", gotBall, wantBall)
	}
}
