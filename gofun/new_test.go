package gofun_test

import (
	"gofun/gofun"
	"testing"
)

func TestGofun(t *testing.T) {
	c := gofun.NewCat("blue", false)

	if c.Color() != "blue" {
		t.Errorf("got %v, instead of blue", c.Color())
	}
}

type House struct {
	hunter gofun.Hunter
}

func TestHunt(t *testing.T) {
	c := gofun.NewCat("blue", false)

	var h House
	h.hunter = &c

	_, err := h.hunter.Hunt()
	if err != nil {
		t.Errorf(err.Error())
	}
}
