package rutina

import (
	"TP-2024-TSPORT/paquete/ejercicio"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAgregarRutina(tester *testing.T) {
	gestorDeEjercicios := ejercicio.NuevoGestorDeEjercicios()
	gestorDeRutinas := NuevoGestorDeRutinas(gestorDeEjercicios)
	ejercicio1 := &ejercicio.Ejercicio {
		NombreDeEjercicio: 				"Ejercicio1", 
		TiempoEnSegundosDeEjercicio:	600, 
		CaloriasDeEjercicio: 			200, 
		TipoDeEjercicio: 				"Cardio", 
		PuntosPorTipoDeEjercicio: 		10, 
		DificultadDeEjercicio: 			"Media",
	}
	gestorDeEjercicios.AgregarEjercicio(ejercicio1)
	rutina1 := &Rutina {
		NombreDeRutina: 						"Rutina1", 
		CaracteristicasIndividualesDeRutina:	[]ejercicio.Ejercicio {
			*ejercicio1,
		},
	}
	err := gestorDeRutinas.AgregarRutina(rutina1)
	assert.Nil(tester, err)
	assert.Equal(tester, 1, len(gestorDeRutinas.ListarRutinas()))
}

func TestEliminarRutina(tester *testing.T) {
	gestorDeEjercicios := ejercicio.NuevoGestorDeEjercicios()
	gestorDeRutinas := NuevoGestorDeRutinas(gestorDeEjercicios)
	ejercicio1 := &ejercicio.Ejercicio {
		NombreDeEjercicio: 				"Ejercicio1", 
		TiempoEnSegundosDeEjercicio:	600, 
		CaloriasDeEjercicio: 			200, 
		TipoDeEjercicio: 				"Cardio", 
		PuntosPorTipoDeEjercicio: 		10, 
		DificultadDeEjercicio: 			"Media",
	}
	gestorDeEjercicios.AgregarEjercicio(ejercicio1)
	rutina1 := &Rutina {
		NombreDeRutina: 						"Rutina1", 
		CaracteristicasIndividualesDeRutina:	[]ejercicio.Ejercicio {
			*ejercicio1,
		},
	}
	gestorDeRutinas.AgregarRutina(rutina1)
	err := gestorDeRutinas.EliminarRutina("Rutina1")
	assert.Nil(tester, err)
	assert.Equal(tester, 0, len(gestorDeRutinas.ListarRutinas()))
}

func TestConsultarRutina(tester *testing.T) {
	gestorDeEjercicios := ejercicio.NuevoGestorDeEjercicios()
	gestorDeRutinas := NuevoGestorDeRutinas(gestorDeEjercicios)
	ejercicio1 := &ejercicio.Ejercicio {
		NombreDeEjercicio: 				"Ejercicio1", 
		TiempoEnSegundosDeEjercicio:	600, 
		CaloriasDeEjercicio: 			200, 
		TipoDeEjercicio: 				"Cardio", 
		PuntosPorTipoDeEjercicio: 		10, 
		DificultadDeEjercicio: 			"Media",
	}
	gestorDeEjercicios.AgregarEjercicio(ejercicio1)
	rutina1 := &Rutina {
		NombreDeRutina: 						"Rutina1", 
		CaracteristicasIndividualesDeRutina:	[]ejercicio.Ejercicio {
			*ejercicio1,
		},
	}
	gestorDeRutinas.AgregarRutina(rutina1)
	rutinaConsultada, err := gestorDeRutinas.ConsultarRutina("Rutina1")
	assert.Nil(tester, err)
	assert.Equal(tester, "Rutina1", rutinaConsultada.NombreDeRutina)
}

func TestModificarRutina(tester *testing.T) {
	gestorDeEjercicios := ejercicio.NuevoGestorDeEjercicios()
	gestorDeRutinas := NuevoGestorDeRutinas(gestorDeEjercicios)
	ejercicio1 := &ejercicio.Ejercicio {
		NombreDeEjercicio: 				"Ejercicio1", 
		TiempoEnSegundosDeEjercicio:	600, 
		CaloriasDeEjercicio: 			200, 
		TipoDeEjercicio: 				"Cardio", 
		PuntosPorTipoDeEjercicio: 		10, 
		DificultadDeEjercicio: 			"Media",
	}
	ejercicio2 := &ejercicio.Ejercicio {
		NombreDeEjercicio: 				"Ejercicio2", 
		TiempoEnSegundosDeEjercicio:	300, 
		CaloriasDeEjercicio: 			100, 
		TipoDeEjercicio: 				"Fuerza", 
		PuntosPorTipoDeEjercicio: 		5, 
		DificultadDeEjercicio: 			"Alta",
	}
	gestorDeEjercicios.AgregarEjercicio(ejercicio1)
	gestorDeEjercicios.AgregarEjercicio(ejercicio2)
	rutina1 := &Rutina {
		NombreDeRutina: 						"Rutina1", 
		CaracteristicasIndividualesDeRutina:	[]ejercicio.Ejercicio {
			*ejercicio1,
		},
	}
	gestorDeRutinas.AgregarRutina(rutina1)
	nuevaRutina := &Rutina {
		NombreDeRutina: 						"Rutina1", 
		CaracteristicasIndividualesDeRutina:	[]ejercicio.Ejercicio {
			*ejercicio2,
		},
	}
	err := gestorDeRutinas.ModificarRutina("Rutina1", nuevaRutina)
	assert.Nil(tester, err)
	rutinaConsultada, _ := gestorDeRutinas.ConsultarRutina("Rutina1")
	assert.Equal(tester, 1, len(rutinaConsultada.CaracteristicasIndividualesDeRutina))
	assert.Equal(tester, "Ejercicio2", rutinaConsultada.CaracteristicasIndividualesDeRutina[0].NombreDeEjercicio)
}

func TestListarRutinas(tester *testing.T) {
	gestorDeEjercicios := ejercicio.NuevoGestorDeEjercicios()
	gestorDeRutinas := NuevoGestorDeRutinas(gestorDeEjercicios)
	ejercicio1 := &ejercicio.Ejercicio {
		NombreDeEjercicio: 				"Ejercicio1", 
		TiempoEnSegundosDeEjercicio:	600, 
		CaloriasDeEjercicio: 			200, 
		TipoDeEjercicio: 				"Cardio", 
		PuntosPorTipoDeEjercicio: 		10, 
		DificultadDeEjercicio: 			"Media",
	}
	ejercicio2 := &ejercicio.Ejercicio {
		NombreDeEjercicio: 				"Ejercicio2", 
		TiempoEnSegundosDeEjercicio:	300, 
		CaloriasDeEjercicio: 			100, 
		TipoDeEjercicio: 				"Fuerza", 
		PuntosPorTipoDeEjercicio: 		5, 
		DificultadDeEjercicio: 			"Alta",
	}
	gestorDeEjercicios.AgregarEjercicio(ejercicio1)
	gestorDeEjercicios.AgregarEjercicio(ejercicio2)
	rutina1 := &Rutina {
		NombreDeRutina: 						"Rutina1", 
		CaracteristicasIndividualesDeRutina:	[]ejercicio.Ejercicio {
			*ejercicio1,
		},
	}
	rutina2 := &Rutina {
		NombreDeRutina: 						"Rutina2", 
		CaracteristicasIndividualesDeRutina:	[]ejercicio.Ejercicio {
			*ejercicio2,
		},
	}
	gestorDeRutinas.AgregarRutina(rutina1)
	gestorDeRutinas.AgregarRutina(rutina2)
	rutinas := gestorDeRutinas.ListarRutinas()
	assert.Equal(tester, 2, len(rutinas))
	assert.Equal(tester, "Rutina1", rutinas[0].NombreDeRutina)
	assert.Equal(tester, "Rutina2", rutinas[1].NombreDeRutina)
}

func TestListarRutinasPorDificultad(tester *testing.T) {
	gestorDeEjercicios := ejercicio.NuevoGestorDeEjercicios()
	gestorDeRutinas := NuevoGestorDeRutinas(gestorDeEjercicios)
	ejercicio1 := &ejercicio.Ejercicio {
		NombreDeEjercicio: 				"Ejercicio1", 
		TiempoEnSegundosDeEjercicio:	600, 
		CaloriasDeEjercicio: 			200, 
		TipoDeEjercicio: 				"Cardio", 
		PuntosPorTipoDeEjercicio: 		10, 
		DificultadDeEjercicio: 			"Media",
	}
	ejercicio2 := &ejercicio.Ejercicio {
		NombreDeEjercicio: 				"Ejercicio2", 
		TiempoEnSegundosDeEjercicio:	300, 
		CaloriasDeEjercicio: 			100, 
		TipoDeEjercicio: 				"Fuerza", 
		PuntosPorTipoDeEjercicio: 		5, 
		DificultadDeEjercicio: 			"Alta",
	}
	gestorDeEjercicios.AgregarEjercicio(ejercicio1)
	gestorDeEjercicios.AgregarEjercicio(ejercicio2)
	rutina1 := &Rutina {
		NombreDeRutina: 						"Rutina1", 
		DificultadDeRutina: 					"Media", 
		CaracteristicasIndividualesDeRutina:	[]ejercicio.Ejercicio {
			*ejercicio1,
		},
	}
	rutina2 := &Rutina {
		NombreDeRutina: 						"Rutina2", 
		DificultadDeRutina: 					"Alta", 
		CaracteristicasIndividualesDeRutina:	[]ejercicio.Ejercicio {
			*ejercicio2,
		},
	}
	gestorDeRutinas.AgregarRutina(rutina1)
	gestorDeRutinas.AgregarRutina(rutina2)
	rutinasMedia := gestorDeRutinas.ListarRutinasPorDificultad("Media")
	rutinasAlta := gestorDeRutinas.ListarRutinasPorDificultad("Alta")
	assert.Equal(tester, 1, len(rutinasMedia))
	assert.Equal(tester, "Rutina1", rutinasMedia[0].NombreDeRutina)
	assert.Equal(tester, 1, len(rutinasAlta))
	assert.Equal(tester, "Rutina2", rutinasAlta[0].NombreDeRutina)
}