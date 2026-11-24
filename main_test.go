package main

import "testing"

func same(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}
	return a.Val == b.Val && same(a.Left, b.Left) && same(a.Right, b.Right)
}

func TestRoundTrip(t *testing.T) {
	trees := []*TreeNode{
		tn(1, tn(2, nil, nil), tn(3, tn(4, nil, nil), tn(5, nil, nil))),
		nil,
		tn(7, nil, nil),
		tn(1, tn(2, tn(3, nil, nil), nil), nil), // left chain
		tn(1, nil, tn(2, nil, tn(3, nil, nil))), // right chain
		tn(-5, tn(0, nil, nil), tn(-10, nil, nil)), // negatives and zero
	}
	for _, tree := range trees {
		got := Deserialize(Serialize(tree))
		if !same(got, tree) {
			t.Errorf("round trip failed for %q", Serialize(tree))
		}
	}
}

func TestSerializeFormat(t *testing.T) {
	if s := Serialize(nil); s != "#" {
		t.Errorf("Serialize(nil) = %q, want \"#\"", s)
	}
	if s := Serialize(tn(1, nil, nil)); s != "1,#,#" {
		t.Errorf("Serialize(leaf) = %q, want \"1,#,#\"", s)
	}
	// Distinct shapes with the same values must serialize differently.
	a := tn(1, tn(2, nil, nil), nil)
	b := tn(1, nil, tn(2, nil, nil))
	if Serialize(a) == Serialize(b) {
		t.Error("different shapes produced identical encodings")
	}
}
