package databases

import "project2/config"

func InputData() error {
	var query = `insert into fasilitas (id, nama_fasilitas) value
	('1', 'TV'),
	('2', 'AC'),
	('3', 'Wifi'),
	('4', 'Heating'),
	('5', 'Include Sarapan'),
	('6', 'PlayStation'),
	('7', 'Virtual Reality')`
	tx := config.DB.Exec(query)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
