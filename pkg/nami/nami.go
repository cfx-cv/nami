package nami

type Nami struct {
	store Store
}

func NewNami(store Store) *Nami {
	return &Nami{store: store}
}
