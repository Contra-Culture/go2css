package go2css

import "fmt"

type (
	StateConfigurator struct {
		rule       *Rule
		stylesheet *Stylesheet
		nested     map[string]*Rule
	}
)

func (cfg *StateConfigurator) Modifier(name string, setup func(*ModifierConfigurator)) {
	_, exists := cfg.nested[name]
	if exists {
		panic(fmt.Sprintf("\"%s\" already defined", name))
	}
	selector := fmt.Sprintf("%s-%s", cfg.rule.selectors[0], name)
	rule := &Rule{
		genesisID: MODIFIER_RULE_GENESIS_ID,
		selectors: []string{
			selector,
		},
		declarations: []*Declaration{},
	}
	mcfg := &ModifierConfigurator{
		rule:       rule,
		stylesheet: cfg.stylesheet,
	}
	setup(mcfg)
	cfg.nested[name] = rule
}
func (cfg *StateConfigurator) D(property string, value string) {
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
func (cfg *StateConfigurator) Has(path []string) {
	cfg.stylesheet.updateOrCreateTraitRule(cfg.rule.selectors[0], path)
}
