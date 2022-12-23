package db

import "github.com/sousapedro11/bucketeer/models"

func (db Database) GetAllItems() (*models.ItemList, error) {
	itemList := &models.ItemList{}

	rows, err := db.Conn.Query("SELECT * FROM items ORDER BY ID DESC")

	if err != nil {
		return itemList, err
	}

	for rows.Next() {
		var item models.Item

		err := rows.Scan(&item.ID, &item.Name, &item.Description, &item.CreatedAt)

		if err != nil {
			return itemList, err
		}

		itemList.Items = append(itemList.Items, item)
	}

	return itemList, nil
}
