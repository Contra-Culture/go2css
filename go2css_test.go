package go2css_test

import (
	. "github.com/Contra-Culture/go2css"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("go2css", func() {
	Describe("universe", func() {
		Describe("NewUniverse()", func() {
			It("returns lib", func() {
				u := NewUniverse(func(u *UniverseConfigurator) {
					u.Trait([]string{"global", "link"}, func(t *TraitConfigurator) {
						t.D("color", "#0033ff")
						t.D("text-decoration", "underline")
					})
					u.Trait([]string{"global", "widget"}, func(t *TraitConfigurator) {
						t.D("background-color", "#e5e6e7")
						t.D("border-width", "1px")
						t.D("border-color", "#a5a6a7")
						t.D("border-style", "solid")
						t.D("padding", "2em")
					})
					u.Stylesheet("test", func(s *StylesheetConfigurator) {
						s.Module("header", func(m *ModuleConfigurator) {
							m.Has([]string{"global", "widget"})
							m.Component("title", func(c *ComponentConfigurator) {
								c.D("font-size", "3em")
								c.D("font-weight", "600")
							})
						})
						s.Module("nav", func(m *ModuleConfigurator) {
							m.D("border-width", "1px")
							m.D("border-color", "#000000")
							m.D("border-style", "solid")
							m.D("padding", "1em")
							m.Component("link", func(c *ComponentConfigurator) {
								c.Has([]string{"global", "link"})
							})
						})
					})
				})
				Expect(u).NotTo(BeNil())
				s, err := u.Stylesheet("test")
				Expect(err).NotTo(HaveOccurred())
				Expect(s).NotTo(BeNil())
				Expect(s.Compile()).To(Equal("\n/* trait */\n.header {\n\tbackground-color: #e5e6e7;\n\tborder-width: 1px;\n\tborder-color: #a5a6a7;\n\tborder-style: solid;\n\tpadding: 2em;\n\t}\n/* component */\n.header-title {\n\tfont-size: 3em;\n\tfont-weight: 600;\n\t}\n/* trait */\n.nav-link {\n\tcolor: #0033ff;\n\ttext-decoration: underline;\n\t}\n/* module */\n.nav {\n\tborder-width: 1px;\n\tborder-color: #000000;\n\tborder-style: solid;\n\tpadding: 1em;\n\t}"))
			})
		})
	})
})
