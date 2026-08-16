package oncall

import "testing"

func TestBug1_FairnessIncludesRosterMembersWithZeroDays(t *testing.T) {
	s, err := Build(Request{
		Roster:     []string{"alice", "bob"},
		Start:      "2026-03-02",
		End:        "2026-03-02",
		StartIndex: 0,
	})
	if err != nil {
		t.Fatal(err)
	}
	want := []FairnessCount{{Engineer: "alice", Days: 1}, {Engineer: "bob", Days: 0}}
	if !equalFairness(s.Fairness, want) {
		t.Fatalf("fairness = %+v, want %+v", s.Fairness, want)
	}
}
