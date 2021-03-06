package abstractfactory

type nike struct {
}

func (n *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 21,
		},
	}
}

func (n *nike) makeShirt() iShirt {
	return &nikeShirt{
		shirt: shirt{
			logo: "nike",
			size: 21,
		},
	}
}
