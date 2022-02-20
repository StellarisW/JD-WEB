package backend

type Group struct{}

func (g *Group) Index() *IndexApi {
	return &insIndex
}

func (g *Group) Login() *LoginApi {
	return &insLogin
}
