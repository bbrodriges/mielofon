package dialog

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

var _ Request = (*PurchaseConfirmationRequest)(nil)

type PurchaseConfirmationRequest struct {
	PurchaseRequestId string          `json:"purchase_request_id"`
	PurchaseToken     string          `json:"purchase_token"`
	OrderId           string          `json:"order_id"`
	PurchaseTimestamp int64           `json:"purchase_timestamp"`
	PurchasePayload   json.RawMessage `json:"purchase_payload"`
	SignedData        string          `json:"signed_data"`
	Signature         string          `json:"signature"`
}

func (PurchaseConfirmationRequest) Type() RequestType {
	return TypePurchaseConfirmation
}

// Verify verifies purchase confirmation request signature
func (p *PurchaseConfirmationRequest) Verify(pubkey *rsa.PublicKey) error {
	bsig, err := base64.RawStdEncoding.DecodeString(p.Signature)
	if err != nil {
		return fmt.Errorf("cannot decode given signature from base64: %w", err)
	}

	hash := sha256.New()
	_, err = hash.Write([]byte(p.SignedData))
	if err != nil {
		return fmt.Errorf("cannot hash signed data: %w", err)
	}
	sum := hash.Sum(nil)

	err = rsa.VerifyPSS(pubkey, crypto.SHA256, sum, bsig, nil)
	if err != nil {
		return fmt.Errorf("could not verify signature: %w", err)
	}
	return nil
}
