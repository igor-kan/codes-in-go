package mod_merkle_audit_tree_d56

func hashLeaf(v uint64) uint64 {
	return (v ^ 0x9E3779B97F4A7C15) * 0xBF58476D1CE4E5B9
}

func hashPair(l, r uint64) uint64 {
	h := l ^ ((r << 13) | (r >> 51))
	return (h * 0x94D049BB133111EB) ^ (h >> 29)
}

type MerkleTree struct {
	levels [][]uint64
}

func NewMerkleTree(leaves []uint64) *MerkleTree {
	cur := make([]uint64, len(leaves))
	for i, v := range leaves {
		cur[i] = hashLeaf(v)
	}
	lvls := [][]uint64{cur}
	for len(cur) > 1 {
		var nxt []uint64
		for i := 0; i < len(cur); i += 2 {
			r := cur[i]
			if i+1 < len(cur) {
				r = cur[i+1]
			}
			nxt = append(nxt, hashPair(cur[i], r))
		}
		lvls = append(lvls, nxt)
		cur = nxt
	}
	return &MerkleTree{levels: lvls}
}

func (m *MerkleTree) Root() uint64 {
	if len(m.levels) == 0 || len(m.levels[len(m.levels)-1]) == 0 {
		return 0
	}
	return m.levels[len(m.levels)-1][0]
}

func (m *MerkleTree) Proof(idx int) []ProofStep {
	var steps []ProofStep
	pos := idx
	for l := 0; l < len(m.levels)-1; l++ {
		row := m.levels[l]
		if pos%2 == 0 {
			sib := row[pos]
			if pos+1 < len(row) {
				sib = row[pos+1]
			}
			steps = append(steps, ProofStep{Sibling: sib, IsRight: true})
		} else {
			steps = append(steps, ProofStep{Sibling: row[pos-1], IsRight: false})
		}
		pos /= 2
	}
	return steps
}

func VerifyProof(leaf uint64, steps []ProofStep, expectedRoot uint64) bool {
	cur := hashLeaf(leaf)
	for _, st := range steps {
		if st.IsRight {
			cur = hashPair(cur, st.Sibling)
		} else {
			cur = hashPair(st.Sibling, cur)
		}
	}
	return cur == expectedRoot
}
