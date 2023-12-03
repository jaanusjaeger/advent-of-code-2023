package matrix

type Index struct {
	I, J int
}

func (i Index) Empty() bool {
	return i == Index{}
}

type Indexes []Index
