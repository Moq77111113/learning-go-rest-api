package database

import "moq.com/test/cmd/models"

func (d *DB) Create(p *models.Post) error {
	res, err := d.db.Exec(insertSchema, p.Title, p.Content, p.Author)
	if err != nil {
		return err
	}
	res.LastInsertId()
	return err
}

func (d *DB) Get() ([]*models.Post, error) {
	var posts []*models.Post
	err := d.db.Select(&posts, "SELECT * FROM posts")
	if err != nil {
		return posts, err
}
	return posts, nil
}
