package go2css

import (
	"fmt"
)

type (
	ModifierConfigurator struct {
		rule       *Rule
		traitRules []*Rule
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
	rule := cfg.stylesheet.includeTrait(path)
	for _, newDeclaration := range rule.declarations {
		for _, oldDeclaration := range cfg.rule.declarations {
			if newDeclaration.property == oldDeclaration.property {
				panic(
					fmt.Sprintf(
						"declaration \"%s\" already exists. old: %s | new: %s",
						newDeclaration.property,
						oldDeclaration.value,
						newDeclaration.value,
					))
			}
		}
		for _, previousTraitRule := range cfg.traitRules {
			for _, oldDeclaration := range previousTraitRule.declarations {
				if newDeclaration.property == oldDeclaration.property {
					panic(
						fmt.Sprintf(
							"declaration \"%s\" already exists. old: \"%s\" from trait \"%s\" | new: \"%s\" from trait \"%s\"",
							newDeclaration.property,
							oldDeclaration.value,
							previousTraitRule.comment,
							newDeclaration.value,
							rule.comment,
						),
					)
				}
			}
		}
	}
	rule.selectors = append(rule.selectors, cfg.rule.selectors[0])
	cfg.traitRules = append(cfg.traitRules, rule)
}
