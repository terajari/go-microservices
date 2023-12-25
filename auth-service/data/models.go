package data

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const dbTimeout = 3 * time.Second

var db *sql.DB

// NewModels initializes a new instance of the Models struct.
//
// It takes a *sql.DB parameter called dbPool and returns a Models struct.
func NewModels(dbPool *sql.DB) Models {
	db = dbPool
	return Models{
		User: User{},
	}
}

type Models struct {
	User User
}

type User struct {
	Id        int       `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"-"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// GetAll returns all the users from the database.
//
// It does not take any parameters.
// It returns a slice of pointers to User objects and an error.
func (u *User) GetAll() ([]*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "SELECT * FROM users ORDER BY last_name"
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.Password,
			&user.Active,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning")
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

// GetByEmail retrieves a user from the database by their email.
//
// It takes a string parameter email which represents the email of the user.
// It returns a pointer to a User struct and an error.
func (u *User) GetByEmail(email string) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "SELECT * FROM users WHERE email = $1"

	var user User
	row := db.QueryRowContext(ctx, query, email)

	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		log.Println("Error scanning")
		return nil, err
	}

	return &user, nil
}

// GetOne retrieves a single user by their ID.
//
// Parameters:
// - id: the ID of the user to retrieve.
//
// Returns:
// - *User: a pointer to the retrieved user.
// - error: an error if any occurred during the retrieval process.
func (u *User) GetOne(id int) (*User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := "SELECT * FROM users WHERE id = $1"

	var user User

	row := db.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Active,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		log.Print("Error scanning")
		return nil, err
	}

	return &user, nil
}

// Update updates the user information in the database.
//
// It takes no parameters.
// It returns an error if the update operation fails.
func (u *User) Update() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := "UPDATE users SET email = $1, first_name = $2, last_name = $3, user_active = $4 updated_at = $5 WHERE id = $6"

	_, err := db.ExecContext(ctx, stmt, u.Email, u.FirstName, u.LastName, u.Active, time.Now(), u.Id)
	if err != nil {
		return err
	}

	return nil
}

// Delete deletes a user from the database.
//
// It takes no parameters.
// It returns an error.
func (u *User) Delete() error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from users where id = $1`

	_, err := db.ExecContext(ctx, stmt, u.Id)
	if err != nil {
		return err
	}

	return nil
}

// DeleteByID deletes one user from the database, by ID
func (u *User) DeleteByID(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	stmt := `delete from users where id = $1`

	_, err := db.ExecContext(ctx, stmt, id)
	if err != nil {
		return err
	}

	return nil
}

// Insert inserts a new user into the database.
//
// It takes a User struct as a parameter and returns the ID of the newly inserted user and an error if any.
func (u *User) Insert(user User) (int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return 0, err
	}

	var newID int
	stmt := `insert into users (email, first_name, last_name, password, user_active, created_at, updated_at)
		values ($1, $2, $3, $4, $5, $6, $7) returning id`

	err = db.QueryRowContext(ctx, stmt,
		user.Email,
		user.FirstName,
		user.LastName,
		hashedPassword,
		user.Active,
		time.Now(),
		time.Now(),
	).Scan(&newID)

	if err != nil {
		return 0, err
	}

	return newID, nil
}

// ResetPassword resets the password for a user.
//
// It takes a password as a string parameter and updates the user's password
// in the database. The function returns an error if there is any issue with
// generating the hashed password or executing the database update statement.
// Otherwise, it returns nil.
func (u *User) ResetPassword(password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `update users set password = $1 where id = $2`
	_, err = db.ExecContext(ctx, stmt, hashedPassword, u.Id)
	if err != nil {
		return err
	}

	return nil
}

// PasswordMatches checks if the provided plain text password matches the hashed password of the user.
//
// It takes a single parameter 'plainText' of type string, which is the plain text password to be checked.
// It returns a boolean value indicating whether the passwords match or not, and an error if any occurred during the comparison.
func (u *User) PasswordMatches(plainText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}

	return true, nil
}
