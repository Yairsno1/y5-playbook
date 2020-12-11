package playbook

import "testing"

func Test_ctr(t *testing.T) {
	phase := newPhase(Setting{})

	got := len(phase.actions)
	want := 0
	if got != want {
		t.Errorf("actions.count == %q, want %q", got, want)
	}
	if !phase.active {
		t.Errorf("Phase not active")
	}
}

func Test_addThreat_no_actions(t *testing.T) {
	//Arrange
	phase := newPhase(Setting{})

	//Act
	action := mockAction{}
	phase.addThreat(action, 1)

	//Assert
	got := len(phase.actions)
	want := 1
	if got != want {
		t.Errorf("actions.count == %q, want %q", got, want)
	}

	got = len(phase.actions[0])
	want = 1
	if got != want {
		t.Errorf("actions.0.count == %q, want %q", got, want)
	}
}

func Test_addThreat_same_order(t *testing.T) {
	//Arrange
	phase := newPhase(Setting{})

	//Act
	action := mockAction{}
	phase.addThreat(action, 1)
	phase.addThreat(action, 1)

	//Assert
	got := len(phase.actions)
	want := 1
	if got != want {
		t.Errorf("actions.count == %q, want %q", got, want)
	}

	got = len(phase.actions[0])
	want = 2
	if got != want {
		t.Errorf("actions.0.count == %q, want %q", got, want)
	}
}

func Test_addThreat_add_order(t *testing.T) {
	//Arrange
	phase := newPhase(Setting{})

	action := mockAction{}
	phase.addThreat(action, 1)
	phase.addThreat(action, 1)

	//Act
	phase.addThreat(action, 2)

	//Assert
	got := len(phase.actions)
	want := 2
	if got != want {
		t.Errorf("actions.count == %q, want %q", got, want)
	}

	got = len(phase.actions[1])
	want = 1
	if got != want {
		t.Errorf("actions.0.count == %q, want %q", got, want)
	}
}

type mockAction struct {
}

func (act mockAction) Do(set Setting) Setting {
	return Setting{}
}
