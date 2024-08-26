package cmd

import (
	"fmt"
	"log"

	"github.com/fatih/color"
	"github.com/lfcifuentes/online-payment-platform/api/app/internal/adapters/pgsql"
	"github.com/lfcifuentes/online-payment-platform/api/app/modules/auth/data/models"
	"github.com/lfcifuentes/online-payment-platform/api/app/pkg"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

// createUserCmd represents the createUser command
// createUserCmd is a command that creates a new user in the database.
// It establishes a connection to the database and pings it to ensure connectivity.
// After that, it calls the createUser function to create the user and then calls the createBanks function.
var createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Create a new user in the database",
	Long:  `This command allows you to create a new user in the database.`,
	Run: func(cmd *cobra.Command, args []string) {

		db, err := pgsql.NewDBAdapter()
		if err != nil {
			log.Fatalf("Could not connect to the database: %v", err)
		}
		err = db.Ping()
		if err != nil {
			log.Fatalf("Could not ping the database: %v", err)
		}

		createUser(cmd, db)

		createBanks(db)
	},
}

func init() {
	rootCmd.AddCommand(createUserCmd)

	createUserCmd.Flags().StringP("username", "u", "", "Username of the new user")
	createUserCmd.Flags().StringP("password", "p", "", "Password of the new user")
	createUserCmd.Flags().StringP("email", "e", "", "Email of the new user")
}

// createUser is a function that creates a new user in the database.
// It takes the username, password, and email as input parameters.
// It checks if the required fields are provided, hashes the password, and inserts the user into the database.
// If the user already exists, it displays an error message.
func createUser(cmd *cobra.Command, db *pgsql.DBAdapter) error {
	color.Green("Creating user...")
	username, _ := cmd.Flags().GetString("username")
	password, _ := cmd.Flags().GetString("password")
	email, _ := cmd.Flags().GetString("email")

	if username == "" || password == "" || email == "" {
		log.Fatal("Username, password, and email are required")
	}

	user := models.User{
		Name:     username,
		Password: password,
		Email:    email,
	}

	hash := pkg.NewPasswordHasher()
	var err error
	user.Password, err = hash.Make(user.Password)
	if err != nil {
		return err
	}
	// check if the user already exists
	query := `SELECT id FROM users WHERE email = $1`
	row := db.DB.QueryRow(query, user.Email)
	var id int
	err = row.Scan(&id)
	color.Green(fmt.Sprintf("%d", id))
	if err != nil {
		query := `INSERT INTO users (name, password, email) VALUES ($1, $2, $3)`
		_, err = db.DB.Exec(query, user.Name, user.Password, user.Email)
		if err != nil {
			return err
		}
		color.Green("User created successfully")
	} else {
		color.Red("User already exists")
	}

	return nil
}

// createBanks is a function that creates banks in the database.
// It inserts four banks with their respective names and webhook URLs.
func createBanks(db *pgsql.DBAdapter) {
	color.Green("Creating banks...")

	color.Blue("Creating Bank of Bogota")
	query := `INSERT INTO banks (name, webhook_url) VALUES ($1, $2)`
	_, err := db.DB.Exec(query, "Banco de Bogota", "http://localhost:8002/bb")
	if err != nil {
		log.Fatalf("Could not create bank: %v", err)
	}

	color.Blue("Creating Bank of Occidente")
	_, err = db.DB.Exec(query, "Banco de Occidente", "http://localhost:8002/bo")
	if err != nil {
		log.Fatalf("Could not create bank: %v", err)
	}

	color.Blue("Creating Bank of Republica")
	_, err = db.DB.Exec(query, "Banco de la Republica", "http://localhost:8002/br")
	if err != nil {
		log.Fatalf("Could not create bank: %v", err)
	}

	color.Blue("Creating Bank of Nacion")
	_, err = db.DB.Exec(query, "Banco de la Nacion", "http://localhost:8002/bn")
	if err != nil {
		log.Fatalf("Could not create bank: %v", err)
	}
}
