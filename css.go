package css

import (
	"github.com/Contra-Culture/treg"
)

var RULE_SPECS = map[string]RuleSpec{}

type (
	Rules    []*Rule
	RuleSpec struct {
		name         string
		val          *ValueSpec
		subRuleSpecs []*RuleSpec
	}
	ValueSpec struct {
		enums []string
		units []string
	}
	Rule struct {
		prop     string
		val      *Value
		subRules []*Rule
	}
	Value struct {
		val  string
		unit string
	}
	Selector []string
)

func R(prop string, value string) (r *Rule, err error) {

}

func Lib() *treg.Reg {
	return treg.New(treg.CheckNotNil)
}
