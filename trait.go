package go2css

import "fmt"

type (
	Trait struct {
		name         string
		declarations []*Declaration
	}
	TraitConfigurator struct {
		trait    *Trait
		universe *Universe
	}
)

func (cfg *TraitConfigurator) D(property string, value string) {
	for _, declaration := range cfg.trait.declarations {
		if declaration.property == property {
			panic(fmt.Sprintf("declaration \"%s\" already exists", property))
		}
	}
	d := &Declaration{
		property: property,
		value:    value,
	}
	cfg.trait.declarations = append(cfg.trait.declarations, d)
}
