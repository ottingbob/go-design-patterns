package btcomp

type Tree struct {
	LeafValue int
	Right 	  *Tree
	Left 	  *Tree
}

root := Tree{
	LeafValue: 0,
	Right: &Tree {
		LeafValue: 5,
		Right: &Tree{ 6, nil, nil },
		Left: nil,
	},
	Left: &Tree{ 4, nil, nil }
}