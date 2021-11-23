package dialog

type ConfirmPurchaseDirective struct{}

func (ConfirmPurchaseDirective) Type() DirectiveType {
	return DirectiveConfirmPurchase
}
