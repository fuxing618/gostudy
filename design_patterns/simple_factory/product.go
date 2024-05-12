package simple_factory

type produceType int
const (
	p1 produceType = iota
	p2
)

type Product interface {
	setName(name string)
	getName() string
}

type Product1 struct {
	name string
}

func (p *Product1) setName(name string) {
	p.name = name
}

func (p *Product1) getName() string {
	return p.name
}

type Product2 struct {
	name string
}

func (p *Product2) setName(name string) {
	p.name = name
}
func (p *Product2) getName() string{
	return p.name
}

type Factory struct {

}

func (f *Factory) createProduct(produceType produceType) Product {
	if produceType == p1 {
		return &Product1{}
	} else if produceType == p2 {
		return &Product1{}
	}
	return nil
}
