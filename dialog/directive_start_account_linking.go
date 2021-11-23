package dialog

type StartAccountLinkingDirective struct{}

func (StartAccountLinkingDirective) Type() DirectiveType {
	return DirectiveStartAccountLinking
}
