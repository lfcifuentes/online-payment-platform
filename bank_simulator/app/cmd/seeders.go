package cmd

import (
	"log"

	"github.com/fatih/color"
	"github.com/lfcifuentes/online-payment-platform/bank/app/internal/adapters/pgsql"
	_ "github.com/lib/pq"
	"github.com/spf13/cobra"
)

// createUserCmd represents the createUser command
var createUserCmd = &cobra.Command{
	Use:   "seeders",
	Short: "Run seeders",
	Long:  `This command will run the seeders to create the banks`,
	Run: func(cmd *cobra.Command, args []string) {

		db, err := pgsql.NewDBAdapter()
		if err != nil {
			log.Fatalf("Could not connect to the database: %v", err)
		}
		err = db.Ping()
		if err != nil {
			log.Fatalf("Could not ping the database: %v", err)
		}

		createBanks(db)
	},
}

func init() {
	rootCmd.AddCommand(createUserCmd)
}

func createBanks(db *pgsql.DBAdapter) {
	color.Green("Creating banks...")

	color.Blue("Creating Bank of Bogota")
	query := `INSERT INTO banks (name, code) VALUES ($1, $2)`
	_, err := db.DB.Exec(query, "Banco de Bogota", "bb")
	if err != nil {
		log.Fatalf("Could not create bank: %v", err)
	}

	color.Blue("Creating Bank of Occidente")
	_, err = db.DB.Exec(query, "Banco de Occidente", "bo")
	if err != nil {
		log.Fatalf("Could not create bank: %v", err)
	}

	color.Blue("Creating Bank of Republica")
	_, err = db.DB.Exec(query, "Banco de la Republica", "br")
	if err != nil {
		log.Fatalf("Could not create bank: %v", err)
	}

	color.Blue("Creating Bank of Nacion")
	_, err = db.DB.Exec(query, "Banco de la Nacion", "bn")
	if err != nil {
		log.Fatalf("Could not create bank: %v", err)
	}
}
