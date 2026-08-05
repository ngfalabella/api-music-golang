package domain

type Cancion struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Artista  string `json:"artista"`
	Duracion int    `json:"duracion"`
}

func NewCancion(id int, nombre string, artista string, duracion int) (Cancion, error) {
	cancionCreada := Cancion{
		ID:       id,
		Nombre:   nombre,
		Artista:  artista,
		Duracion: duracion,
	}

	return cancionCreada, nil
}
