package symboltable

type InstructionAddr int

type SymbolTable struct {
	forward map[InstructionAddr]InstructionAddr
	inverse map[InstructionAddr]InstructionAddr
}

func New() *SymbolTable {
	st := &SymbolTable{
		forward: make(map[InstructionAddr]InstructionAddr),
		inverse: make(map[InstructionAddr]InstructionAddr),
	}

	return st
}

func (st *SymbolTable) AddSymbolLink(left, right InstructionAddr) {
	st.forward[left] = right
	st.inverse[right] = left
}

func (st *SymbolTable) LookupFromLeft(left InstructionAddr) InstructionAddr {
	return st.forward[left]
}

func (st *SymbolTable) LookupFromRight(right InstructionAddr) InstructionAddr {
	return st.inverse[right]
}
