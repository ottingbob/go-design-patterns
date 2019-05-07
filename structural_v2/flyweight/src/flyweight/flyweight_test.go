package flyweight

import (
	"fmt"
	"testing"
	"unsafe"

	f "flyweight"
)

func TestTeamFlyweightFactory_GetTeam(t *testing.T) {
	factory := f.NewTeamFactory()

	teamA1 := factory.GetTeam(f.TEAM_A)
	if teamA1 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}

	teamA2 := factory.GetTeam(f.TEAM_A)
	if teamA2 == nil {
		t.Error("The pointer to the TEAM_A was nil")
	}

	if teamA1 != teamA2 {
		t.Error("TEAM_A objects weren't the same")
	}

	if factory.GetNumberOfObjects() != 1 {
		t.Errorf("The number of objects created was not 1: %d\n", factory.GetNumberOfObjects())
	}
}

func Test_HighVolume(t *testing.T) {
	factory := f.NewTeamFactory()

	teams := make([]*f.Team, 500000*2)
	for i := 0; i < 500000; i++ {
		teams[i] = factory.GetTeam(f.TEAM_A)
	}

	for i := 500000; i < 2*500000; i++ {
		teams[i] = factory.GetTeam(f.TEAM_B)
	}

	if factory.GetNumberOfObjects() != 2 {
		t.Errorf("The number of objects created was not 2: %d\n", factory.GetNumberOfObjects())
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("Pointer %d points to %p and is located in %p\n", i, teams[i], &teams[i])
	}
	fmt.Printf("Size of s: %d\nSize of [500k*2]s: %d\n", unsafe.Sizeof(teams), unsafe.Sizeof([500000*2]*f.Team{}))
}