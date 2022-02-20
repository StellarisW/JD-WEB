package frontend

type Group struct{}

func (g *Group) Base() *sBase {
	return &insBase
}

func (g *Group) Index() *sIndex {
	return &insIndex
}

func (g *Group) Auth() *sAuth {
	return &insAuth
}

func (g *Group) User() *sUser {
	return &insUser
}

func (g *Group) Product() *sProduct {
	return &insProduct
}

func (g *Group) Cart() *sCart {
	return &insCart
}

func (g *Group) Buy() *sBuy {
	return &insBuy
}
