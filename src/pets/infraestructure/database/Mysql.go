package database

import (
	"APIDOS/src/core"
	"APIDOS/src/pets/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() (*MySQL, error) {
	conn := core.GetDBPool()
	if conn.Err != "" {
		return nil, fmt.Errorf("error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}, nil
}

func (mysql *MySQL) Save(mascota entities.Mascotas) (entities.Mascotas, error) {
	query := "INSERT INTO Mascota (name_mascota, breed_mascota) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, mascota.Name, mascota.Breed)
	if err != nil {
		return entities.Mascotas{}, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
		mascota.ID = int(rowsAffected)
	}

	return mascota, nil
}

func (mysql *MySQL) GetAll() ([]entities.Mascotas, error) {
	query := "SELECT id_mascota, name_mascota, breed_mascota FROM Mascota"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener filas: %v", err)
	}
	defer rows.Close()

	var mascotas []entities.Mascotas

	for rows.Next() {
		var mascota entities.Mascotas
		if err := rows.Scan(&mascota.ID, &mascota.Name, &mascota.Breed); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}
		mascotas = append(mascotas, mascota)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %v", err)
	}

	return mascotas, nil
}

func (mysql *MySQL) Update(mascota entities.Mascotas) error {
	query := "UPDATE Mascota SET name_mascota = ?, breed_mascota = ? WHERE id_mascota = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, mascota.Name, mascota.Breed, mascota.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar la mascota con ID %d: %v", mascota.ID, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ninguna mascota con ID %d para actualizar", mascota.ID)
	}

	log.Printf("[MySQL] - Mascota con ID %d actualizada", mascota.ID)
	return nil
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM Mascota WHERE id_mascota = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar la mascota con ID %d: %v", id, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ninguna mascota con ID %d para eliminar", id)
	}

	log.Printf("[MySQL] - Mascota con ID %d eliminada", id)
	return nil
}
