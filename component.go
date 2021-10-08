package go2css

import "fmt"

type (
	ComponentConfigurator struct {
		rule       *Rule
		stylesheet *Stylesheet
		nested     map[string]*Rule
		states     []string
		modifiers  []string
	}
)

func (cfg *ComponentConfigurator) State(name string, setup func(*StateConfigurator)) {
	_, exists := cfg.nested[name]
	if exists {
		panic(fmt.Sprintf("\"%s\" already defined", name))
	}
	selector := fmt.Sprintf("%s-%s", cfg.rule.selectors[0], name)
	rule := &Rule{
		genesisID: STATE_RULE_GENESIS_ID,
		selectors: []string{
			selector,
		},
		declarations: []*Declaration{},
	}
	scfg := &StateConfigurator{
		rule:       rule,
		stylesheet: cfg.stylesheet,
		nested:     map[string]*Rule{},
	}
	setup(scfg)
	cfg.nested[name] = rule
	cfg.states = append(cfg.states, name)
}
func (cfg *ComponentConfigurator) Modifier(name string, setup func(*ModifierConfigurator)) {
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
	cfg.modifiers = append(cfg.modifiers, name)
}
func (cfg *ComponentConfigurator) D(property string, value string) {
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
func (cfg *ComponentConfigurator) Has(path []string) {
	cfg.stylesheet.updateOrCreateTraitRule(cfg.rule.selectors[0], path)
}
