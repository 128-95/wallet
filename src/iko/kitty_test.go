package iko

import (
	"testing"

	"github.com/kittycash/kittiverse/src/kitty"
)

func TestKittyIDs_Sort(t *testing.T) {
	ids := kitty.KittyIDs{
		kitty.KittyID(65),
		kitty.KittyID(2),
		kitty.KittyID(20),
		kitty.KittyID(23),
		kitty.KittyID(12),
		kitty.KittyID(3),
		kitty.KittyID(94),
		kitty.KittyID(24),
	}
	ids.Sort()
	t.Log(ids)
}
