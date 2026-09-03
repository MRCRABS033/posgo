package main

import (
	"database/sql"
	"log"
	"posgo/internal/repository/sqlite"
	"posgo/internal/usecase"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./pos.db")
	if err != nil {
		log.Fatalf("Error critico al abrir la base de datos", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("No se pudo establecer conexión con SQLite: %v", err)
	}
	log.Println("Conexión con SQLite establecida exitosamente.")

	//inyeccion de dependencias: repositorios (capa de infraestructura)
	cashInRepo := sqlite.NewCashInRepository(db)
	cashOutRepo := sqlite.NewCashOutRepository(db)
	departmentRepo := sqlite.NewDeparmentRepository(db)
	permissionRepo := sqlite.NewPermissionRepository(db)
	productRepo := sqlite.NewProductRepository(db)
	ticketItemRepo := sqlite.NewTicketItemRepository(db)
	ticketRepo := sqlite.NewTicketRepository(db)
	sessionRepo := sqlite.NewSessionRepository(db)
	userRepo := sqlite.NewUserRepository(db)

	//inyeccion de dependencias: casos de uso (capa de negocio)
	// Aquí es donde inyectas los repositorios en sus respectivos casos de uso:
	cashInUseCase := usecase.NewCashInUseCase(cashInRepo)
	cashOutUseCase := usecase.NewCashOutUseCase(cashOutRepo)
	departmentUseCase := usecase.NewDepartmentUseCase(departmentRepo)
	permissionUseCase := usecase.NewPermissionUseCase(permissionRepo)
	productUseCase := usecase.NewProductUseCase(productRepo)
	ticketItemUseCase := usecase.NewTicketItemUseCase(ticketItemRepo)
	ticketUseCase := usecase.NewTicketUseCase(ticketRepo)
	sessionUseCase := usecase.NewSessionUseCase(sessionRepo)
	userUseCase := usecase.NewUserUseCase(userRepo)

	log.Println("Todos los casos de uso han sido inicializados correctamente.")

}
