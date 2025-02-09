package database

import (
	"APIDOS/src/core"
	"APIDOS/src/veterinario/domain/entities"
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

func (mysql *MySQL) Save(veterinario entities.Veterinario) error {
	query := "INSERT INTO Veterinario (name_veterinario, lastName_veterinario, phone_veterinario) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, veterinario.Name, veterinario.LastName, veterinario.Phone)
	if err != nil {
		return fmt.Errorf("error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Filas afectadas: %d", rowsAffected)
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]entities.Veterinario, error) {
	query := "SELECT id_veterinario, name_veterinario, lastName_veterinario, phone_veterinario FROM Veterinario"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener filas: %v", err)
	}
	defer rows.Close()

	var veterinarios []entities.Veterinario

	for rows.Next() {
		var veterinario entities.Veterinario
		if err := rows.Scan(&veterinario.Id, &veterinario.Name, &veterinario.LastName, &veterinario.Phone); err != nil {
			return nil, fmt.Errorf("error al escanear la fila: %v", err)
		}
		veterinarios = append(veterinarios, veterinario)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterando sobre las filas: %v", err)
	}

	return veterinarios, nil
}

func (mysql *MySQL) GetById(id int) (entities.Veterinario, error) {
	query := "SELECT id_veterinario, name_veterinario, lastName_veterinario, phone_veterinario FROM Veterinario WHERE id_veterinario = ?"
	rows, err := mysql.conn.FetchRows(query, id)
	if err != nil {
		return entities.Veterinario{}, fmt.Errorf("error al ejecutar la consulta: %v", err)
	}
	defer rows.Close()

	var veterinario entities.Veterinario
	if rows.Next() {
		if err := rows.Scan(&veterinario.Id, &veterinario.Name, &veterinario.LastName, &veterinario.Phone); err != nil {
			return entities.Veterinario{}, fmt.Errorf("error al escanear datos del medico: %v", err)
		}
		return veterinario, nil
	}

	return entities.Veterinario{}, fmt.Errorf("no se encontró el Id del medico %d", id)
}

func (mysql *MySQL) Update(veterinario entities.Veterinario) error {
	query := "UPDATE Mascotas SET name_veterinario = ?, lastName_veterinario = ?, phone_veterinario WHERE id_veterinario = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, veterinario.Name, veterinario.LastName, veterinario.Phone, veterinario.Id)
	if err != nil {
		return fmt.Errorf("error al actualizar los datos del medico con ID %d: %v", veterinario.Id, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ninguna medico con ese Id %d para actualizar", veterinario.Id)
	}

	log.Printf("[MySQL] - Medico con Id %d actualizado", veterinario.Id)
	return nil
}

func (mysql *MySQL) Delete(id int) error {
	query := "DELETE FROM veterinario WHERE id_veterinario = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminarlos datos del medico  con Id %d: %v", id, err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no se encontró ningún medico con ese Id %d para eliminar", id)
	}

	log.Printf("[MySQL] - Medico con Id %d eliminado", id)
	return nil
}
