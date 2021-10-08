package go2css

type (
	Rule struct {
		genesisID    int
		selectors    []string
		declarations []*Declaration
	}
	Declaration struct {
		property string
		value    string
	}
	RuleConfigurator struct {
		rule *Rule
	}
)

const (
	_ = iota
	MODULE_RULE_GENESIS_ID
	COMPONENT_RULE_GENESIS_ID
	STATE_RULE_GENESIS_ID
	MODIFIER_RULE_GENESIS_ID
	TRAIT_RULE_GENESIS_ID
)
const (
	MODULE_RULE_GENESIS_TITLE    = "module"
	COMPONENT_RULE_GENESIS_TITLE = "component"
	STATE_RULE_GENESIS_TITLE     = "state"
	MODIFIER_RULE_GENESIS_TITLE  = "modifier"
	TRAIT_RULE_GENESIS_TITLE     = "trait"
)

var genesisMap = map[int]string{
	MODULE_RULE_GENESIS_ID:    MODULE_RULE_GENESIS_TITLE,
	COMPONENT_RULE_GENESIS_ID: COMPONENT_RULE_GENESIS_TITLE,
	STATE_RULE_GENESIS_ID:     STATE_RULE_GENESIS_TITLE,
	MODIFIER_RULE_GENESIS_ID:  MODIFIER_RULE_GENESIS_TITLE,
	TRAIT_RULE_GENESIS_ID:     TRAIT_RULE_GENESIS_TITLE,
}

func D(p, v string) *Declaration {
	return nil
}
func R(configure func(*RuleConfigurator)) *Rule {
	return nil
}

func (r *Rule) genesis() string {
	return genesisMap[r.genesisID]
}
