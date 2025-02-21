package entities

type Mascotas struct {
	ID    int    `json:"id_mascota" db:"id_mascota"`
	Name  string `json:"name_mascota" db:"name_mascota"`
	Breed string `json:"breed_mascota" db:"breed_mascota"`
}
