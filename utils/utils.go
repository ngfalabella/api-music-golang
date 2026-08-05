package utils

import (
	"github.com/ngfalabella/api-music-golang/domain"
)

const (
	HeaderContentType = "Content-Type"
	ContentTypeJSON   = "application/json"
)

var Canciones = []domain.Cancion{
	{
		ID:       1,
		Nombre:   "Ji Ji Ji",
		Artista:  "Patricio Rey y sus Redonditos de Ricota",
		Duracion: 330,
	},
	{
		ID:       2,
		Nombre:   "Un Ángel Para Tu Soledad",
		Artista:  "Patricio Rey y sus Redonditos de Ricota",
		Duracion: 230,
	},
	{
		ID:       3,
		Nombre:   "Todo un Palo",
		Artista:  "Patricio Rey y sus Redonditos de Ricota",
		Duracion: 455,
	},
	{
		ID:       4,
		Nombre:   "Flight 956",
		Artista:  "Los Fundamentalistas del Aire Acondicionado",
		Duracion: 275,
	},
	{
		ID:       5,
		Nombre:   "Nike es la Cultura",
		Artista:  "Los Fundamentalistas del Aire Acondicionado",
		Duracion: 270,
	},
	{
		ID:       6,
		Nombre:   "El Tesoro de los Inocentes",
		Artista:  "Los Fundamentalistas del Aire Acondicionado",
		Duracion: 355,
	},
}
