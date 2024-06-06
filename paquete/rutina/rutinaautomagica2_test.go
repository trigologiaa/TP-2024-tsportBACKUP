package rutina

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerarRutinaAutomagica2(t *testing.T) {
	gestorRutinas := &GestorRutinas{}
	nombre := "MiRutina"
	caloriasTotales := 150
	rutina, err := gestorRutinas.GenerarRutinaAutomagica2(nombre, caloriasTotales)
	assert.NoError(t, err, "Se esperaba nil, pero se obtuvo error")
	assert.NotNil(t, rutina, "La rutina no debería ser nil")
	assert.Equal(t, nombre, rutina.Nombre, "Nombre de la rutina incorrecto")
	assert.Equal(t, caloriasTotales, rutina.Calorias, "Calorías totales de la rutina incorrectas")
}