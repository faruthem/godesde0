package modelos

type Vegetal struct {
	Nombre string
}

func (v *Vegetal) EstaVivo() bool               { return true }
func (v *Vegetal) ClasificacionVegetal() string { return "Planta" }
