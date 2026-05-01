package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/taleeus/wedding/internal/db"
	"github.com/taleeus/wedding/internal/server"
	"github.com/taleeus/wedding/web/components"
	"github.com/taleeus/wedding/web/pages"
)

var dburl = "libsql://$TURSO_DATABASE_URL?authToken=$TURSO_AUTH_TOKEN"

func main() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	})))

	err := godotenv.Load()
	if err != nil {
		slog.Info(".env not found; using environment")
	}

	dbconn, err := db.New(os.ExpandEnv(dburl))
	if err != nil {
		log.Fatal(err)
	}
	defer dbconn.Close()

	if err := db.Migrate(dbconn); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	slog.Info("🚀 starting server", "port", port)
	if err := http.ListenAndServe(":"+port, server.New(dbconn, copy)); err != nil {
		log.Fatal(err)
	}
}

var copy = pages.HomeCopy{
	HeroLabel: "Vi invitano a condividere la gioia del loro matrimonio il 17 ottobre 2026.",
	InfoChips: []pages.InfoChip{{
		Title:   "QUANDO",
		Content: "Il matrimonio inizierà alle ore 15 del 17 ottobre 2026.",
	}, {
		Title: "DOVE",
		Content: `La cerimonia e il ricevimento si terranno presso la Rocca di Montalfeo.
		<a class="underline text-exotic-skin hover:text-cherry-oak" target="_blank" href="https://maps.app.goo.gl/hccg9SgnqfpMR4Ka9">
			Località Montalfeo, 9, 27052 Montalfeo PV
		</a>`,
	}, {
		Title:   "DRESS CODE",
		Content: "Sentitevi liberi di esprimere il vostro stile con abiti eleganti nei colori che amate di più.",
	}},
	TimelineCells: []components.CellConfig{{
		Time:        "15:00",
		Description: "Inizio cerimonia",
	}, {
		Time:        "16:30",
		Description: "Aperitivo",
	}, {
		Time:        "19:00",
		Description: "Cena",
	}, {
		Time:        "21:30",
		Description: "Taglio torta",
	}, {
		Time:        "22:00",
		Description: "Festa",
	}},
	GiftDescription: `Il regalo più bello è la vostra presenza. Se desiderate
	contribuire al nostro viaggio di nozze in Giappone, potete farlo
	tramite bonifico bancario.`,
	IBAN: os.Getenv("IBAN"),
	Feedback: pages.FeedbackCopy{
		Success: "Grazie di cuore per condividere con noi un momento così importante!",
		Failure: "Si è verificato un errore :( riprova o contattaci direttamente per comunicarci la tua risposta",
	},
}
