// // package main

// // import (
// // 	"log"
// // 	"net/http"

// // 	"github.com/gorilla/mux"

// // 	"github.com/maryamjamal7/smart-light-city/adapters/api"
// // 	"github.com/maryamjamal7/smart-light-city/adapters/storage"
// // 	"github.com/maryamjamal7/smart-light-city/domain/model"
// // 	"github.com/maryamjamal7/smart-light-city/domain/service"
// // )

// // func main() {
// // 	// 1. Connect to database
// // 	db, err := storage.ConnectPostgres()
// // 	if err != nil {
// // 		log.Fatal("❌ Failed to connect to DB:", err)
// // 	}

// // 	// 2. Migrate models
// // 	err = db.AutoMigrate(&model.Area{}, &model.Lumiere{}, &model.Command{})
// // 	if err != nil {
// // 		log.Fatal("❌ Auto-migration failed:", err)
// // 	}
// // 	log.Println("✅ Database schema migrated")

// // 	// 3. Initialize repositories and services

// // 	// Area
// // 	areaRepo := storage.NewAreaRepository(db)
// // 	areaService := service.NewAreaService(areaRepo)

// // 	// Lumiere
// // 	lumiereRepo := storage.NewLumiereRepository(db)
// // 	lumiereService := service.NewLumiereService(lumiereRepo)

// // 	// 4. Setup router
// // 	r := mux.NewRouter()

// // 	// 5. Register routes
// // 	api.RegisterAreaRoutes(r, areaService)
// // 	api.RegisterLumiereRoutes(r, lumiereService)

// //		// 6. Start server
// //		log.Println("🚀 Server running on http://localhost:8080")
// //		log.Fatal(http.ListenAndServe(":8080", r))
// //	}
// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// 	"github.com/maryamjamal7/smart-light-city/adapters/mqtt"

// 	"github.com/maryamjamal7/smart-light-city/adapters/api"
// 	"github.com/maryamjamal7/smart-light-city/adapters/storage"
// 	"github.com/maryamjamal7/smart-light-city/domain/model"
// 	"github.com/maryamjamal7/smart-light-city/domain/service"
// )

// func main() {
// 	// 1. Connect to database
// 	db, err := storage.ConnectPostgres()
// 	if err != nil {
// 		log.Fatal(" Failed to connect to DB:", err)
// 	}
// 	mqttPub, err := mqtt.NewMQTTPublisher("tcp://localhost:1883", "smart-light-city")
// 	if err != nil {
// 	log.Fatal("MQTT connection failed:", err)
// 	}

// 	// 2. Auto migrate models
// 	err = db.AutoMigrate(&model.Area{}, &model.Lumiere{}, &model.Command{})
// 	if err != nil {
// 		log.Fatal(" Auto-migration failed:", err)
// 	}
// 	log.Println("✅ Database schema migrated")

// 	// 3. Initialize Repositories
// 	areaRepo := storage.NewAreaRepository(db)
// 	lumiereRepo := storage.NewLumiereRepository(db)
// 	commandRepo := storage.NewCommandRepository(db)

// 	// 4. Initialize Services
// 	areaService := service.NewAreaService(areaRepo)
// 	lumiereService := service.NewLumiereService(lumiereRepo)
// 	commandService := service.NewCommandService(commandRepo)
// 	cityManager := service.NewCityManager(areaService, lumiereService, commandService)

// 	// 5. Setup Router
// 	r := mux.NewRouter()
// 	api.RegisterAreaRoutes(r, areaService)
// 	api.RegisterLumiereRoutes(r, lumiereService)
// 	api.RegisterCommandRoutes(r, commandService)

// 	// City-level control: /city/poweroff, /city/dim
// 	r.HandleFunc("/city/poweroff", api.HandleCityPowerOff(cityManager)).Methods("POST")
// 	r.HandleFunc("/city/dim", api.HandleCityScheduleDim(cityManager)).Methods("POST")

//		// 6. Start Server
//		log.Println("🚀 Server running on http://localhost:8080")
//		log.Fatal(http.ListenAndServe(":8080", r))
//	}
package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/maryamjamal7/smart-light-city/adapters/api"
	"github.com/maryamjamal7/smart-light-city/adapters/mqtt"
	"github.com/maryamjamal7/smart-light-city/adapters/storage"

	"github.com/maryamjamal7/smart-light-city/domain/model"
	"github.com/maryamjamal7/smart-light-city/domain/service"
)

func main() {
	// 1. Connect to database
	db, err := storage.ConnectPostgres()
	if err != nil {
		log.Fatal("❌ Failed to connect to DB:", err)
	}

	// 2. Connect to MQTT broker
	mqttPub, err := mqtt.NewMQTTPublisher("tcp://localhost:1883", "smart-light-city")
	if err != nil {
		log.Fatal("❌ MQTT connection failed:", err)
	}
	log.Println("✅ Connected to MQTT broker")

	// 3. Auto migrate models
	err = db.AutoMigrate(&model.Area{}, &model.Lumiere{}, &model.Command{})
	if err != nil {
		log.Fatal("❌ Auto-migration failed:", err)
	}
	log.Println("✅ Database schema migrated")

	// 4. Initialize Repositories
	areaRepo := storage.NewAreaRepository(db)
	lumiereRepo := storage.NewLumiereRepository(db)
	commandRepo := storage.NewCommandRepository(db)

	// 5. Initialize Services (pass mqtt publisher to command service)
	areaService := service.NewAreaService(areaRepo)
	lumiereService := service.NewLumiereService(lumiereRepo)
	commandService := service.NewCommandService(commandRepo, mqttPub)
	cityManager := service.NewCityManager(areaService, lumiereService, commandService)

	// // 6. Start Dkron listener in background (runs scheduled commands)
	// scheduler := scheduler.NewDkronScheduler(commandService)
	// scheduler.Start()

	// 7. Setup HTTP Router
	r := mux.NewRouter()
	api.RegisterAreaRoutes(r, areaService)
	api.RegisterLumiereRoutes(r, lumiereService)
	api.RegisterCommandRoutes(r, commandService)

	// 8. City-wide control endpoints
	r.HandleFunc("/city/poweroff", api.HandleCityPowerOff(cityManager)).Methods("POST")
	r.HandleFunc("/city/dim", api.HandleCityScheduleDim(cityManager)).Methods("POST")

	// 9. Start HTTP server
	log.Println("🚀 Server running on http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", r))

}
