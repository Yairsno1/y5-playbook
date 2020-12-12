package playbook

import (
	"fmt"
	"testing"
)

func Test_O_String(t *testing.T) {
	//Arrange + Act + Assert
	got := fmt.Sprintf("%s", O3Sym)
	want := "O3"

	if got != want {
		t.Errorf("string(O3) == %q, want %q", got, want)
	}
}

func Test_Threat_GiveAndGo_String(t *testing.T) {
	//Arrange + Act + Assert
	got := fmt.Sprintf("%s", GiveAndGo)
	want := "Give&go"

	if got != want {
		t.Errorf("string(GiveAndGo) == %q, want %q", got, want)
	}
}

func Test_HalfCourtPos_LeftCorner_String(t *testing.T) {
	//Arrange + Act + Assert
	got := fmt.Sprintf("%s", LeftCorner)
	want := "(L.) Corner"

	if got != want {
		t.Errorf("string(GiveAndGo) == %q, want %q", got, want)
	}
}
