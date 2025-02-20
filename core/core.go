package core

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ryanvillarreal/metrics/core/db"
	"github.com/ryanvillarreal/metrics/core/file"

	"github.com/spf13/cobra"
)

var (
	dir    string
	port   int
	web    bool
	dbPath string
)

var rootCmd = &cobra.Command{
	Use:     "metrics",
	Short:   "get stats for nerds",
	Version: "0.0.1",
	Args:    cobra.MaximumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		// If --web flag is set, start the web server
		if web {
			startDB()
			startWeb()
			return
		}
		//

		// Default behavior: read XML files
		err := file.ReadXMLFiles(dir)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("Reading XML files from directory: %s", dir)
	},
}

func init() {
	// Add flags
	rootCmd.PersistentFlags().StringVarP(&dir, "dir", "d", "./data", "Directory containing XML files")
	rootCmd.PersistentFlags().BoolVar(&web, "web", false, "Start the web server")
	rootCmd.PersistentFlags().IntVarP(&port, "port", "p", 42069, "Port to run the web server on")
	rootCmd.PersistentFlags().StringVar(&dbPath, "db", "metrics.db", "Path to the SQLite database file")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func startDB() {
	db.Start()
	log.Printf("Connected to SQLite database at %s", dbPath)
}

func startWeb() {
}
