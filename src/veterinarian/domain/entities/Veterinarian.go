package entities

type Veterinarian struct {
	ID       int    `json:"id_veterinario" db:"id_veterinario"`
	Name     string `json:"name_veterinario" db:"name_veterinario"`
	LastName string `json:"last_name_veterinario" db:"last_name_veterinario"`
	Phone    string `json:"phone_veterinario" db:"phone_veterinario"`
}
