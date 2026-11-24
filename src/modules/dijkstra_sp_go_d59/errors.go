package mod_dijkstra_sp_go_d59

import "errors"

var (
	ErrSsspInvalid   = errors.New("dijkstra_sp_go_d59: invalid argument")
	ErrSsspBounds    = errors.New("dijkstra_sp_go_d59: index out of bounds")
	ErrSsspEmpty     = errors.New("dijkstra_sp_go_d59: empty input")
	ErrSsspInvariant = errors.New("dijkstra_sp_go_d59: invariant violated")
)
