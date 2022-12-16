package calc

import (
	"fmt"

	"predmer-api/OLD/app/helper"

	"gioui.org/layout"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

func (w *WingCal) Nazad() func(gtx C) D {
	return func(gtx C) D {
		var btn layout.Dimensions
		if len(w.Putanja) > 1 {
			btnNazad := material.Button(w.UI.Tema.T, nazadDugme, w.text("NAZAD"))
			btnNazad.Background = helper.HexARGB(w.UI.Tema.Colors["Secondary"])
			btnNazad.CornerRadius = unit.Dp(0)
			for nazadDugme.Clicked() {
				if w.Element {
					w.Element = false
				} else {
					//komanda := ""
					if len(w.Putanja) == 3 {
						//komanda = "/" + fmt.Sprint(w.Roditelj)
						//podvrstaradova = fmt.Sprint(w.Roditelj)
						fmt.Println("roddddditeL111::", w.Roditelj)
					}
					if len(w.Putanja) == 4 {
						//komanda = "/" + Podvrstaradova + "/" + fmt.Sprint(w.Roditelj)
					}
					w.APIpozivIzbornik(w.Radovi.PodvrsteRadova)
					//w.LinkoviIzboraVrsteRadova = GenerisanjeLinkova(w.IzbornikRadova)
					w.GenerisanjeLinkova(w.IzbornikRadova)
					w.Putanja = w.Putanja[:len(w.Putanja)-1]
				}
			}
			btn = btnNazad.Layout(gtx)
		}
		return btn
	}
}
