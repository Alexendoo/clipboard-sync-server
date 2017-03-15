package model

type User struct {
	ID int64

	Name        string
	PassKey     []byte
	PassSalt    []byte
	PassVersion int
}

const addUserSQL = `
INSERT INTO users (id, username, pass_key, pass_salt, pass_version)
VALUES (?, ?, ?, ?, ?)`

func (s *Store) AddUser(user *User) error {
	_, err := s.stmts.addUser.Exec(
		user.ID,
		user.Name,
		user.PassKey,
		user.PassSalt,
		user.PassVersion,
	)

	return err
}

const getUserSQL = `
SELECT id, username, pass_key, pass_salt, pass_version
FROM users WHERE id = ?`

func (s *Store) GetUser(id int64) (*User, error) {
	var user User

	err := s.stmts.getUser.QueryRow(id).Scan(
		&user.ID,
		&user.Name,
		&user.PassKey,
		&user.PassSalt,
		&user.PassVersion,
	)

	return &user, err
}
