package iko

import (
	"github.com/kittycash/kittiverse/src/kitty"
	"github.com/skycoin/skycoin/src/cipher"
	"github.com/skycoin/skycoin/src/cipher/encoder"
)

/*
	<<< KITTY DETAILS >>>
	>>> Used by multiple services, provides off-chain details for kitties and IKO.
*/

// type Kitty struct {
// 	ID        KittyID `json:"kitty_id"`    // Identifier for kitty.
// 	Name      string  `json:"name"`        // Name of kitty.
// 	Desc      string  `json:"description"` // Description of kitty.
// 	Breed     string  `json:"breed"`       // Kitty breed.
// 	Legendary bool    `json:"legendary"`   // Whether kitty is legendary.
//
// 	PriceBTC int64 `json:"price_btc"` // Price of kitty in BTC.
// 	PriceSKY int64 `json:"price_sky"` // Price of kitty in SKY.
//
// 	BoxOpen bool `json:"box_open"` // Whether box is open.
//
// 	Created   int64  `json:"created,omitempty"`    // Timestamp that the kitty box began existing.
// 	BirthDate int64  `json:"birth_date,omitempty"` // Timestamp of box opening (after box opening).
// 	KittyDNA  string `json:"kitty_dna,omitempty"`  // Hex representation of kitty DNA (after box opening).
//
// 	BoxImgURL   string `json:"box_image_url,omitempty"`   // Box image TransformURL.
// 	KittyImgURL string `json:"kitty_image_url,omitempty"` // Kitty image TransformURL (after box opening).
// }

/*
	<<< KITTY ENTRY >>>
	>>> The way a kitty is entered in the kitty-api.
*/

// TODO For now replacing Kitty with Kitty.KittyInfo to mock things
// KittyEntry should be totally replaced with Kitty
type KittyEntry struct {
	kitty.KittyInfo
	Sig         string `json:"sig"`               // Signature should be verified with
	Reservation string `json:"reservation"`       // Whether kitty is reserved or not.
	Address     string `json:"address,omitempty"` // The address in which the kitty resides in.
}

// func KittyEntryFromJson(raw []byte) (*KittyEntry, error) {
// 	out := new(KittyEntry)
// 	err := json.Unmarshal(raw, out)
// 	return out, err
// }
//
// func (e *KittyEntry) Json() []byte {
// 	raw, _ := json.Marshal(e)
// 	return raw
// }
//
// func (e *KittyEntry) Sign(sk cipher.SecKey) {
// 	e.Sig = e.KittyInfo.Sign(sk).Hex()
// }
//
// func (e *KittyEntry) Verify(pk cipher.PubKey) error {
// 	sig, err := cipher.SigFromHex(e.Sig)
// 	if err != nil {
// 		return err
// 	}
// 	return e.KittyInfo.Verify(pk, sig)
// }

/*
	<<< KITTY STATE >>>
	>>> The state of a kitty as represented when the IKO Chain is compiled.
*/

type KittyState struct {
	Address      cipher.Address
	Transactions TxHashes
}

func (s KittyState) Serialize() []byte {
	return encoder.Serialize(s)
}

/*
	<<< ADDRESS STATE >>>
	>>> The state of an address as represented when the IKO Chain is compiled.
*/

type AddressState struct {
	Kitties      kitty.KittyIDs
	Transactions TxHashes
}

func NewAddressState() *AddressState {
	return &AddressState{
		Kitties:      make(kitty.KittyIDs, 0),
		Transactions: make(TxHashes, 0),
	}
}

func (a AddressState) Serialize() []byte {
	return encoder.Serialize(a)
}
