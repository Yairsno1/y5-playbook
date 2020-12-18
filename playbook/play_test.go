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
