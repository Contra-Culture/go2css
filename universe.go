package go2css

import (
	"fmt"
	"strings"
)

type (
	Universe struct {
		stylesheets map[string]*Stylesheet
		traits      map[string]interface{}
		cvars       map[string]interface{}
		rvars       map[string]interface{}
	}
	UniverseConfigurator struct {
		universe *Universe
	}
)

func (cfg *UniverseConfigurator) Stylesheet(name string, configure func(*StylesheetConfigurator)) {
	stylesheet := &Stylesheet{
		name:       name,
		universe:   cfg.universe,
		rules:      []*Rule{},
		traitRules: map[int][]string{},
	}
	scfg := &StylesheetConfigurator{
		stylesheet: stylesheet,
		universe:   cfg.universe,
	}
	configure(scfg)
	cfg.universe.stylesheets[name] = stylesheet
}
func (cfg *UniverseConfigurator) Trait(path []string, setup func(*TraitConfigurator)) {
	var (
		trait     *Trait
		tsp       *TraitConfigurator
		ok        bool
		rawTraits interface{}
		nested    map[string]interface{}
		traits    = cfg.universe.traits
		name      = path[len(path)-1]
	)
	for _, chunk := range path[:len(path)-1] {
		rawTraits, ok = traits[chunk]
		if !ok {
			nested = map[string]interface{}{}
			traits[chunk] = nested
			traits = nested
			continue
		}
		traits, ok = rawTraits.(map[string]interface{})
		if !ok {
			panic("expected map[string]interface{}")
		}
	}
	trait = &Trait{
		name:         name,
		declarations: []*Declaration{},
	}
	tsp = &TraitConfigurator{
		trait: trait,
	}
	setup(tsp)
	traits[name] = trait
}
func (cfg *UniverseConfigurator) CVar(path []string, val string) {
	var (
		ok       bool
		rawCvars interface{}
		nested   map[string]interface{}

		cvars = cfg.universe.cvars
		name  = path[len(path)-1]
	)
	for _, chunk := range path[:len(path)-1] {
		rawCvars, ok = cvars[chunk]
		if !ok {
			nested = map[string]interface{}{}
			cvars[chunk] = nested
			cvars = nested
			continue
		}
		cvars, ok = rawCvars.(map[string]interface{})
		if !ok {
			panic("expected map[string]interface{}")
		}
	}
	cvars[name] = val
}
func (cfg *UniverseConfigurator) RVar(path []string, defaultVal string) {
	var (
		ok       bool
		rawRvars interface{}
		nested   map[string]interface{}

		rvars = cfg.universe.rvars
		name  = path[len(path)-1]
	)
	for _, chunk := range path[:len(path)-1] {
		rawRvars, ok = rvars[chunk]
		if !ok {
			nested = map[string]interface{}{}
			rvars[chunk] = nested
			rvars = nested
			continue
		}
		rvars, ok = rawRvars.(map[string]interface{})
		if !ok {
			panic("expected map[string]interface{}")
		}
	}
	rvars[name] = defaultVal
}

func NewUniverse(setup func(*UniverseConfigurator)) *Universe {
	var (
		universe = &Universe{
			stylesheets: map[string]*Stylesheet{},
			traits:      map[string]interface{}{},
			cvars:       map[string]interface{}{},
			rvars:       map[string]interface{}{},
		}
		ucfg = &UniverseConfigurator{
			universe: universe,
		}
	)
	setup(ucfg)
	return universe
}
func (u *Universe) Stylesheet(name string) (*Stylesheet, error) {
	stylesheet, ok := u.stylesheets[name]
	if !ok {
		return nil, fmt.Errorf("stylesheet \"%s\"", name)
	}
	return stylesheet, nil
}

func (u *Universe) trait(path []string) (*Trait, error) {
	var (
		ok        bool
		trait     *Trait
		rawTrait  interface{}
		rawTraits interface{}
		newTraits map[string]interface{}
		chunks    = path[:len(path)-1]
		name      = path[len(path)-1]
		traits    = u.traits
	)
	for _, chunk := range chunks {
		rawTraits, ok = traits[chunk]
		if !ok {
			return nil, fmt.Errorf("wrong path1 \"%s\"", strings.Join(path, "/"))
		}
		newTraits, ok = rawTraits.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("wrong path2 \"%s\"", strings.Join(path, "/"))
		}
		traits = newTraits
	}
	rawTrait, ok = traits[name]
	if !ok {
		return nil, fmt.Errorf("wrong path3 \"%s\"", strings.Join(path, "/"))
	}
	trait, ok = rawTrait.(*Trait)
	if !ok {
		return nil, fmt.Errorf("wrong path4 \"%s\"", strings.Join(path, "/"))
	}
	return trait, nil
}
func (u *Universe) cvar(path []string) (string, error) {
	var (
		ok       bool
		cvar     string
		rawCvar  interface{}
		rawCvars interface{}
		newcvars map[string]interface{}
		chunks   = path[:len(path)-1]
		name     = path[len(path)-1]
		cvars    = u.cvars
	)
	for _, chunk := range chunks {
		rawCvars, ok = cvars[chunk]
		if !ok {
			return "", fmt.Errorf("wrong path \"%s\"", strings.Join(path, "/"))
		}
		newcvars, ok = rawCvars.(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("wrong path \"%s\"", strings.Join(path, "/"))
		}
		cvars = newcvars
	}
	rawCvar, ok = cvars[name]
	if !ok {
		return "", fmt.Errorf("wrong path \"%s\"", strings.Join(path, "/"))
	}
	cvar, ok = rawCvar.(string)
	if !ok {
		return "", fmt.Errorf("wrong path \"%s\"", strings.Join(path, "/"))
	}
	return cvar, nil
}
func (u *Universe) rvar(path []string) (string, error) {
	var (
		ok       bool
		rvar     string
		rawRvar  interface{}
		rawRvars interface{}
		newRvars map[string]interface{}
		chunks   = path[:len(path)-1]
		name     = path[len(path)-1]
		rvars    = u.rvars
	)
	for _, chunk := range chunks {
		rawRvars, ok = rvars[chunk]
		if !ok {
			return "", fmt.Errorf("wrong path \"%s\"", strings.Join(path, "/"))
		}
		newRvars, ok = rawRvars.(map[string]interface{})
		if !ok {
			return "", fmt.Errorf("wrong path \"%s\"", strings.Join(path, "/"))
		}
		rvars = newRvars
	}
	rawRvar, ok = rvars[name]
	if !ok {
		return "", fmt.Errorf("wrong path \"%s\"", strings.Join(path, "/"))
	}
	rvar, ok = rawRvar.(string)
	if !ok {
		return "", fmt.Errorf("wrong path \"%s\"", strings.Join(path, "/"))
	}
	return rvar, nil
}
