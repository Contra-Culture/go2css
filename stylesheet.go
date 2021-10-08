package go2css

import "strings"

type (
	Stylesheet struct {
		name       string
		universe   *Universe
		rules      []*Rule
		traitRules map[int][]string
	}
	StylesheetConfigurator struct {
		stylesheet *Stylesheet
		universe   *Universe
	}
)

func (cfg *StylesheetConfigurator) Module(name string, setup func(*ModuleConfigurator)) {
	var (
		rule = &Rule{
			genesisID: MODULE_RULE_GENESIS_ID,
			selectors: []string{
				"." + name,
			},
			declarations: []*Declaration{},
		}
		mcfg = &ModuleConfigurator{
			stylesheet: cfg.stylesheet,
			rule:       rule,
			nested:     map[string]*Rule{},
			components: []string{},
			states:     []string{},
			modifiers:  []string{},
		}
	)
	setup(mcfg)
	cfg.stylesheet.rules = append(cfg.stylesheet.rules, rule)
	for _, rule := range mcfg.nested {
		cfg.stylesheet.rules = append(cfg.stylesheet.rules, rule)
	}
}
func (s *Stylesheet) Compile() string {
	var sb strings.Builder
	for _, rule := range s.rules {
		if len(rule.declarations) == 0 {
			continue
		}
		sb.WriteString("\n/* ")
		sb.WriteString(rule.genesis())
		if len(rule.comment) > 0 {
			sb.WriteRune(' ')
			sb.WriteString(rule.comment)
		}
		sb.WriteString(" */\n")
		selectors := strings.Join(rule.selectors, ",\n")
		sb.WriteString(selectors)
		sb.WriteString(" {")
		for _, declaration := range rule.declarations {
			sb.WriteString("\n\t")
			sb.WriteString(declaration.property)
			sb.WriteString(": ")
			sb.WriteString(declaration.value)
			sb.WriteRune(';')
		}
		sb.WriteString("\n\t}")
	}
	return sb.String()
}
func (s *Stylesheet) traitRule(path []string) (*Rule, bool) {
outer:
	for i, tpath := range s.traitRules {
		if len(tpath) != len(path) {
			continue
		}
		for j, chunk := range tpath {
			if chunk != path[j] {
				continue outer
			}
		}
		return s.rules[i], true
	}
	return nil, false
}
func (s *Stylesheet) updateOrCreateTraitRule(selector string, path []string) {
	rule, ok := s.traitRule(path)
	if ok {
		rule.selectors = append(rule.selectors, selector)
		return
	}
	trait, err := s.universe.trait(path)
	if err != nil {
		panic(err)
	}
	rule = &Rule{
		genesisID: TRAIT_RULE_GENESIS_ID,
		selectors: []string{
			selector,
		},
		declarations: trait.declarations,
		comment:      strings.Join(path, "/"),
	}
	s.rules = append(s.rules, rule)
	s.traitRules[len(s.rules)-1] = path
}
