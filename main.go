package main

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
)

type Sale struct {
	Product int
	Volume  int
	Date    string
}

// String реализует метод интерфейса fmt.Stringer для Sale, возвращает строковое представление объекта Sale.
// Теперь, если передать объект Sale в fmt.Println(), то выведется строка, которую вернёт эта функция.
func (s Sale) String() string {
	return fmt.Sprintf("Product: %d Volume: %d Date:%s", s.Product, s.Volume, s.Date)
}

func selectSales(client int) ([]Sale, error) {
	var sales []Sale

	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
        return nil, fmt.Errorf("ошибка подключения к БД: %w", err)
    }
    defer db.Close()

	query := "SELECT product, volume, date FROM sales WHERE client = ?"
	rows, err := db.Query(query, client)
    if err != nil {
        return nil, fmt.Errorf("ошибка выполнения запроса: %w", err)
    }
    defer rows.Close()

	for rows.Next() {
		var sale Sale
		err := rows.Scan(&sale.Product, &sale.Volume, &sale.Date)
		if err != nil {
            return nil, fmt.Errorf("ошибка при сканировании строки: %w", err)
        }
		sales = append(sales, sale)
	}

	if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("ошибка при обработке результатов: %w", err)
    }

	return sales, nil
}

func main() {
	client := 208

	sales, err := selectSales(client)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sale := range sales {
		fmt.Println(sale)
	}
}
