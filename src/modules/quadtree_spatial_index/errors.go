package mod_quadtree_spatial_index

import "errors"

var (
	ErrQuadtreeInvalid   = errors.New("quadtree: invalid argument")
	ErrQuadtreeBounds    = errors.New("quadtree: index out of bounds")
	ErrQuadtreeEmpty     = errors.New("quadtree: empty container")
	ErrQuadtreeInvariant = errors.New("quadtree: invariant violated")
)
