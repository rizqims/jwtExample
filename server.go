package main

import (
	"apilaundry/config"
	"apilaundry/controller"
	"apilaundry/middleware"
	"apilaundry/repository"
	"apilaundry/service"
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	bS      service.BillService
	cS      service.CustomerService
	pS      service.ProductService
	uS      service.UserService
	jS      service.JwtService
	aM      middleware.AuthMiddleware
	engine  *gin.Engine
	PortApp string
}

func (s *Server) initiateRoute() {
	routerGroup := s.engine.Group("/api/v1")
	controller.NewBillController(s.bS, routerGroup, s.aM).Route()
	controller.NewProductController(s.pS, routerGroup, s.aM).Route()
	controller.NewUserController(s.uS, routerGroup).Route()
}

func (s *Server) Start() {
	s.initiateRoute()
	s.engine.Run(s.PortApp)
}

func NewServer() *Server {
	co, err := config.NewConfig()

	urlConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", co.Host, co.Port, co.User, co.Password, co.Name)

	db, err := sql.Open("postgres", urlConnection)
	if err != nil {
		log.Fatal(err)
	}
	portApp := co.AppPort
	billRepo := repository.NewBillRepository(db)
	custRepo := repository.NewCustomerRepository(db)
	productRepo := repository.NewProductRepository(db)
	userRepo := repository.NewUserRepository(db)

	jwtService := service.NewJwtService(co.SecurityConfig)
	custService := service.NewCustomerService(custRepo)
	userService := service.NewUserService(userRepo, jwtService)
	productService := service.NewProductService(productRepo)
	billService := service.NewBillService(billRepo, userService, productService, custService)

	authMiddleware := middleware.NewAuthMiddleware(jwtService)

	return &Server{
		PortApp: portApp,
		bS:      billService,
		cS:      custService,
		pS:      productService,
		uS:      userService,
		jS:      jwtService,
		aM:      authMiddleware,
		engine:  gin.Default(),
	}
}
