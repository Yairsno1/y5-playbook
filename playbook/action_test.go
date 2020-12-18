package playbook

import "testing"

func Test_Do_Pass(t *testing.T) {
	//Arrange
	set1 := CreateSetting(TopOfTheKey, RightWing, LeftWing, RightElbow, LeftElbow, O1)
	pass := PassAction{
		o:       O1Sym,
		threat:  Pass,
		teamate: O2Sym,
	}

	//Act
	set2 := pass.Do(set1)

	//Assert
	want := O2
	got := set2.ball
	if got != want {
		t.Errorf("Ball handler is %q, should be %q", got, want)
	}
}

func Test_Do_Screen(t *testing.T) {
	//Arrange
	set1 := CreateSetting(TopOfTheKey, RightWing, LeftWing, RightElbow, LeftElbow, O1)
	screen := ScreenAction{
		o:       O4Sym,
		threat:  PickAndRoll,
		teamate: O1Sym,
		spot:    TopOfTheKey,
	}

	//Act
	set2 := screen.Do(set1)

	//Assert
	want := TopOfTheKey
	got := set2.os[O4].spot
	if got != want {
		t.Errorf("4 screens at %q, should be at %q", got, want)
	}
}

func Test_Do_Cut(t *testing.T) {
	//Arrange
	set1 := CreateSetting(TopOfTheKey, RightWing, LeftWing, RightElbow, LeftElbow, O2)
	cut := CutAction{
		o:      O1Sym,
		threat: UCLACut,
		spot:   RightBlock,
	}

	//Act
	set2 := cut.Do(set1)

	//Assert
	want := RightBlock
	got := set2.os[O1].spot
	if got != want {
		t.Errorf("1 cuts to %q, should be to %q", got, want)
	}
}

func Test_Do_Dribble(t *testing.T) {
	//Arrange
	set1 := CreateSetting(TopOfTheKey, RightWing, LeftWing, RightElbow, LeftElbow, O1)
	dribble := DribbleAction{
		o:      O1Sym,
		threat: Dribble,
		spot:   RightWing,
	}

	//Act
	set2 := dribble.Do(set1)

	//Assert
	want := RightWing
	got := set2.os[O1].spot
	if got != want {
		t.Errorf("1 dribbles to %q, should be to %q", got, want)
	}
}

func Test_Do_Move(t *testing.T) {
	//Arrange
	set1 := CreateSetting(TopOfTheKey, RightWing, LeftWing, RightElbow, LeftElbow, O1)
	move := MoveAction{
		o:      O2Sym,
		threat: Move,
		spot:   RightCorner,
	}

	//Act
	set2 := move.Do(set1)

	//Assert
	want := RightCorner
	got := set2.os[O2].spot
	if got != want {
		t.Errorf("2 moves to the %q, should be move to %q", got, want)
	}
}

func Test_Do_Shot(t *testing.T) {
	//Arrange
	set1 := CreateSetting(TopOfTheKey, RightWing, LeftWing, HighPost, HighPost, O1)
	shot := ShotAction{
		o:      O1Sym,
		threat: Shot,
	}

	//Act
	set2 := shot.Do(set1)

	//Assert
	want := TopOfTheKey
	got := set2.os[O1].spot
	if got != want {
		t.Errorf("1 shoots from %q, should be shoot from %q", got, want)
	}
}

func Test_IExecution_Pass(t *testing.T) {
	//Arrange + Act
	pass := O1Sym.Pass(O2Sym)

	//Assert
	want := O2Sym
	got := pass.(PassAction).teamate
	if got != want {
		t.Errorf("Ball handler is %q, should be %q", got, want)
	}
}

func Test_IExecution_VCut(t *testing.T) {
	//Arrange + Act
	want := RightWing
	vcut := O1Sym.VCut(want)

	//Assert
	got := vcut.(MoveAction).spot
	if got != want {
		t.Errorf("O1 did a v-cut to %q, should be to %q", got, want)
	}
}

func Test_IExecution_DownScreen(t *testing.T) {
	//Arrange + Act
	want := LeftBlock
	downScreen := O4Sym.DownScreen(want, O2Sym)

	//Assert
	got := downScreen.(ScreenAction).spot
	if got != want {
		t.Errorf("O4 set a down screen at %q, should be at %q", got, want)
	}
}
