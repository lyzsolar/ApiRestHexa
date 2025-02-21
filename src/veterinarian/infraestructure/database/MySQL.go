package database

import (
	"APIDOS/src/core"
	"APIDOS/src/veterinarian/domain/entities"
	"fmt"
	"log"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(veterinarian entities.Veterinarian) error {
	query := "INSERT INTO Veterinario (name_veterinario, last_name_veterinario, phone_veterinario) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, veterinarian.Name, veterinarian.LastName, veterinarian.Phone)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]entities.Veterinarian, error) {
	query := "SELECT id_veterinario, name_veterinario, last_name_veterinario, phone_veterinario FROM Veterinario"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener filas: %v", err)
	}
	defer rows.Close()

	var veterinarians []entities.Veterinarian

	for rows.Next() {
		var veterinarian entities.Veterinarian
		if err := rows.Scan(&veterinarian.ID, &veterinarian.Name, &veterinarian.LastName, &veterinarian.Phone); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}
		veterinarians = append(veterinarians, veterinarian)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %v", err)
	}

	return veterinarians, nil
}

func (mysql *MySQL) Update(veterinarian entities.Veterinarian) error {
	query := "UPDATE Veterinario SET name_veterinario = ?, last_name_veterinario = ?, phone_veterinario = ? WHERE id_veterinario = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, veterinarian.Name, veterinarian.LastName, veterinarian.Phone, veterinarian.ID)
	if err != nil {
		return fmt.Errorf("error al actualizar el veterinario con ID %d: %v", veterinarian.ID, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ningún veterinario con ID %d para actualizar", veterinarian.ID)
	}

	log.Printf("[MySQL] - Veterinario con ID %d actualizado", veterinarian.ID)
	return nil
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM Veterinario WHERE id_veterinario = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar al veterinario con ID %d: %v", id, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ningún veterinario con ID %d para eliminar", id)
	}

	log.Printf("[MySQL] - Veterinario con ID %d eliminado", id)
	return nil
}
