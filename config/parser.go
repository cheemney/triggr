package config

import "fmt"

type parser struct {
	lex *lexer
	cur token
}

func newParser(src string) *parser {
	p := &parser{lex: newLexer(src)}
	p.advance()
	return p
}

func (p *parser) advance() {
	p.cur = p.lex.next()
}

func (p *parser) expectIdent(want string) error {
	if p.cur.kind != tokIdent || p.cur.text != want {
		return fmt.Errorf("expected %q, got %q", want, p.cur.text)
	}
	p.advance()
	return nil
}

// Parse turns config source into a Program. Format:
//
//	rule NAME {
//	    trigger KIND ARG...
//	    action  KIND ARG...
//	}
func Parse(src string) (*Program, error) {
	p := newParser(src)
	prog := &Program{}
	for p.cur.kind != tokEOF {
		rule, err := p.parseRule()
		if err != nil {
			return nil, err
		}
		prog.Rules = append(prog.Rules, rule)
	}
	return prog, nil
}

func (p *parser) parseRule() (RuleDecl, error) {
	if err := p.expectIdent("rule"); err != nil {
		return RuleDecl{}, err
	}
	name := p.cur.text
	p.advance()

	if p.cur.kind != tokLBrace {
		return RuleDecl{}, fmt.Errorf("expected '{' after rule name, got %q", p.cur.text)
	}
	p.advance()

	if err := p.expectIdent("trigger"); err != nil {
		return RuleDecl{}, err
	}
	trigger, err := p.parseSpec()
	if err != nil {
		return RuleDecl{}, err
	}

	if err := p.expectIdent("action"); err != nil {
		return RuleDecl{}, err
	}
	action, err := p.parseSpec()
	if err != nil {
		return RuleDecl{}, err
	}

	if p.cur.kind != tokRBrace {
		return RuleDecl{}, fmt.Errorf("expected '}', got %q", p.cur.text)
	}
	p.advance()

	return RuleDecl{Name: name, Trigger: trigger, Action: action}, nil
}

// parseSpec reads a kind identifier followed by whatever args come
// after it, stopping at the next "trigger"/"action" keyword or the
// closing brace.
func (p *parser) parseSpec() (Spec, error) {
	kind := p.cur.text
	p.advance()

	var args []string
	for p.cur.kind != tokRBrace &&
		!(p.cur.kind == tokIdent && (p.cur.text == "trigger" || p.cur.text == "action")) {
		if p.cur.kind == tokEOF {
			return Spec{}, fmt.Errorf("unexpected end of input while parsing %q", kind)
		}
		args = append(args, p.cur.text)
		p.advance()
	}
	return Spec{Kind: kind, Args: args}, nil
}
