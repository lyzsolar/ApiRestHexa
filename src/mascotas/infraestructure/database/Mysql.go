package database

import (
	"APIDOS/src/core"
	"APIDOS/src/mascotas/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_Mysql
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatal("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(mascotas entities.Mascotas) error {
	query := "INSERT INTO Mascotas (name_mascotas, breed_mascotas) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, mascotas.Name, mascotas.Breed)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]entities.Mascotas, error) {
	query := "SELECT id_mascotas, name_mascotas, breed_mascotas FROM Mascotas"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener filas: %v", err)
	}
	defer rows.Close()

	var mascota []entities.Mascotas

	for rows.Next() {
		var mascotas entities.Mascotas
		if err := rows.Scan(&mascotas.Id, &mascotas.Name, &mascotas.Breed); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}
		mascota = append(mascota, mascotas)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %v", err)
	}

	return mascota, nil
}

func (mysql *MySQL) GetById(id int) (entities.Mascotas, error) {
	query := "SELECT id_mascotas, name_mascotas, breed_mascotas FROM Mascotas WHERE id_mascotas = ?"
	rows, err := mysql.conn.FetchRows(query, id)
	if err != nil {
		return entities.Mascotas{}, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	var mascotas entities.Mascotas
	if rows.Next() {
		if err := rows.Scan(&mascotas.Id, &mascotas.Name, &mascotas.Breed); err != nil {
			return entities.Mascotas{}, fmt.Errorf("error al escanear la mascota: %v", err)
		}
		return mascotas, nil
	}

	return entities.Mascotas{}, fmt.Errorf("no se encontró el Id de la mascota %d", id)
}

func (mysql *MySQL) Update(mascotas entities.Mascotas) error {
	query := "UPDATE Mascotas SET name_mascotas = ?, breed_mascotas = ? WHERE id_mascotas = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, mascotas.Name, mascotas.Breed, mascotas.Id)
	if err != nil {
		return fmt.Errorf("error al actualizar la mascota con ID %d: %v", mascotas.Id, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ninguna mascota con Id %d para actualizar", mascotas.Id)
	}

	log.Printf("[MySQL] - Mascota con Id %d actualizada", mascotas.Id)
	return nil
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM Mascotas WHERE id_mascotas = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar la mascota con Id %d: %v", id, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ninguna mascota con Id %d para eliminar", id)
	}

	log.Printf("[MySQL] - Mascota con Id %d eliminada", id)
	return nil
}
