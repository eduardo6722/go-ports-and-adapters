package application

type ProductService struct {
	ProductPersistence IProductPersistence
}

func (s *ProductService) Get(id string) (IProduct, error) {
	product, err := s.ProductPersistence.Get(id)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) Create(name string, price float64) (IProduct, error) {
	product := NewProduct()
	product.Name = name
	product.Price = price
	_, err := product.IsValid()
	if err != nil {
		return &Product{}, err
	}
	result, err := s.ProductPersistence.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

func (p *ProductService) Enable(product IProduct) (IProduct, error) {
	err := product.Enable()
	if err != nil {
		return &Product{}, err
	}
	result, err := p.ProductPersistence.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}

func (p *ProductService) Disable(product IProduct) (IProduct, error) {
	err := product.Disable()
	if err != nil {
		return &Product{}, err
	}
	result, err := p.ProductPersistence.Save(product)
	if err != nil {
		return &Product{}, err
	}
	return result, nil
}
