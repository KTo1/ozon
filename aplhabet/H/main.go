package main

import (
	"github.com/KTo1/ozon/H/model"
	"github.com/KTo1/ozon/H/pointer"
)

type Some struct {
	verif *model.VerificationStatus
}

func main() {
	vs := strfmt.UUID4( Some{verif: pointer.From(model.VerificationStatusConfirmed)}
	a := pointer.To(vs.verif)
	_ = a
}
