package tree_test

import (
	"testing"

	"github.com/Fotkurz/snippets/datastructures/go/pkg/tree"
)

func Test_NewNode(t *testing.T) {
	n := tree.NewNode(nil)
	if n.Parent != nil {
		t.Errorf("Expected parent to be nil, but it isn't")
	}
}

func Test_IsLeaf(t *testing.T) {
	tests := []struct {
		name string
		node *tree.Node
		want bool
	}{
		{
			"is leaf",
			tree.NewNode(nil),
			true,
		},
		{
			"is not leaf",
			func() *tree.Node {
				n := tree.NewNode(nil)
				n.Left = tree.NewNode(n)
				return n
			}(),
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isLeaf := tt.node.IsLeaf()
			if tt.want != isLeaf {
				t.Errorf("Expected=(%v), but got=(%v)", tt.want, isLeaf)
			}
		})
	}
}

func Test_Root(t *testing.T) {
	t.Run("return itself if its root", func(t *testing.T) {
		n := tree.NewNode(nil)
		root := n.Root()

		if root != n {
			t.Errorf("Expected=(%v), but got=(%v)", n, root)
		}
	})

	t.Run("return parent for 1 level", func(t *testing.T) {
		root := tree.NewNode(nil)
		root.Left = tree.NewNode(root)
		n := root.Left
		if n.Root() != root {
			t.Errorf("Expected=(%v), but got=(%v)", root, n.Root())
		}
	})

	t.Run("return parent for 10 levels", func(t *testing.T) {
		n := tree.NewNode(nil)

		pointer := n
		for i := 0; i < 10; i++ {
			nN := tree.NewNode(n)

			n.Left = nN
			pointer = nN
		}

		if pointer.Root() != n {
			t.Errorf("Expected=(%v), but got=(%v)", n, pointer.Root())
		}
	})
}

// func Test_Height(t *testing.T) {
// 	tests := []struct {
// 		name string
// 		arg  *tree.Node
// 		want uint
// 	}{
// 		{
// 			"return 0 for nil tree",
// 			nil,
// 			0,
// 		},
// 		{
// 			"return 1",
// 			tree.NewNode(nil),
// 			1,
// 		},
// 		{
// 			"return 2",
// 			func() *tree.Node {
// 				n := tree.NewNode(nil)
// 				n.Left = tree.NewNode(n)
// 				return n.Left
// 			}(),
// 			2,
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got := tree.Height(tt.arg)
// 			if tt.want != got {
// 				t.Errorf("Expected=(%v), but got=(%v)", tt.want, got)
// 			}
// 		})
// 	}
// }
