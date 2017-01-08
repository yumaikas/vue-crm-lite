package database

import (
	"database/sql"
	"fmt"
)

type Contact struct {
	ID       *int64
	Name     sql.NullString
	Phone    sql.NullString
	Email    sql.NullString
	Website  sql.NullString
	Modified sql.NullString
}

type Note struct {
	ID       *int64
	Body     sql.NullString
	Modified sql.NullString
}

// CreateDB Fill out a contact, and now you'll have it'd ID field populated
func (c *Contact) CreateInDB() error {
	res, err := db.Exec(`
    Insert into Contact (
        name, phone, email, website, created_date
    ) values (
        ?, --name
        ?, --phone
        ?, --email
        ?, --website
        datetime('now')
    );
    `, c.Name, c.Phone.String, c.Email.String, c.Website.String)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	c.ID = &id
	return nil
}

// SaveDB Save updates to a contact to the databse
func (c Contact) SaveInDB() error {
	_, err := db.Exec(`
    Update  Contact (
        set name = ?, 
            phone = ?, 
            email = ?, 
            website = ?, 
            updated_date = datetime('now')
    );`,
		*c.ID, c.Name, c.Phone.String, c.Email.String, c.Website.String)
	if err != nil {
		return err
	}
	return nil
}

func getContact(id int) (Contact, error) {
	contact := Contact{}
	err := db.QueryRow(`
        Select 
            id, 
            name, 
            email, 
            website, 
            CASE updated_date 
                WHEN NULL THEN datetime(created_date)
                ELSE datetime(updated_date)
            END
     `, &id).Scan(
		&contact.ID,
		&contact.Name,
		&contact.Email,
		&contact.Website,
		&contact.Modified,
	)
	return contact, err
}

func (c Contact) getNotes() ([]Note, error) {
	if c.ID == nil {
		return nil, fmt.Errorf("cannot fetch notes for contact without ID")
	}
	rows, err := db.Query(`
    Select 
    note.id, 
    note.body,
    CASE note.updated_date
        WHEN NULL THEN datetime(note.created_date)
        ELSE datetime(note.updated_date)
    END
    from Note
    inner join contact_note on note.id = contact_note.note_id
    where contact_note.contact_id = ?
     `, *c.ID)
	if err != nil {
		return nil, err
	}
	var notes = make([]Note, 0)
	for rows.Next() {
		// var note = Note{}

	}
	return notes, nil
}
