package frontend

type Group struct{}

func (g *Group) Base() *BaseApi {
	return &insBase
}

func (g *Group) Index() *IndexApi {
	return &insIndex
}

func (g *Group) Auth() *AuthApi {
	return &insAuth
}

func (g *Group) User() *UserApi {
	return &insUser
}

func (g *Group) Product() *ProductApi {
	return &insProduct
}

func (g *Group) Cart() *CartApi {
	return &insCart
}

func (g *Group) Buy() *BuyApi {
	return &insBuy
}

func (g *Group) Pay() *PayApi {
	return &insPay
}
