package go2css

import "fmt"

type (
	ModifierConfigurator struct {
		rule       *Rule
		stylesheet *Stylesheet
	}
)

func (cfg *ModifierConfigurator) D(property string, value string) {
	newDeclaration := &Declaration{
		property: property,
		value:    value,
	}
	for _, d := range cfg.rule.declarations {
		if d.property == newDeclaration.property {
			panic(fmt.Sprintf("declaration \"%s\" already exists", d.property))
		}
	}
	cfg.rule.declarations = append(cfg.rule.declarations, newDeclaration)
}
func (cfg *ModifierConfigurator) Has(path []string) {
	cfg.stylesheet.updateOrCreateTraitRule(cfg.rule.selectors[0], path)
}
