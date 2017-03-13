package model

type User struct {
	ID int64

	Name        string
	PassKey     []byte
	PassSalt    []byte
	PassVersion int
}

func (s *Store) AddUser(user *User) error {
	const addUserSQL = `
INSERT INTO users (id, username, pass_key, pass_salt, pass_version)
VALUES (?, ?, ?, ?, ?)`

	_, err := s.db.Exec(
		addUserSQL,
		user.ID,
		user.Name,
		user.PassKey,
		user.PassSalt,
		user.PassVersion,
	)

	return err
}

func (s *Store) GetUser(id int64) (*User, error) {
	const getUserSQL = `
SELECT id, username, pass_key, pass_salt, pass_version
FROM users WHERE id = ?`

	var user User

	err := s.db.QueryRow(getUserSQL, id).Scan(
		&user.ID,
		&user.Name,
		&user.PassKey,
		&user.PassSalt,
		&user.PassVersion,
	)

	return &user, err
}
