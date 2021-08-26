package controllers

import "github.com/khafido/simple-app-go/api/middlewares"

func (s *Server) initializeRoutes() {
	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//---Users routes
	//Insert
	s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	//Select All
	s.Router.HandleFunc("/api/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	//Select Once
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	//Update
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	//Delete
	s.Router.HandleFunc("/api/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteUser))).Methods("DELETE")

	//Products routes
	//Insert
	s.Router.HandleFunc("/api/products", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CreateProduct))).Methods("POST")
	//Select All
	s.Router.HandleFunc("/api/products", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetProducts))).Methods("GET")
	//Select Once
	s.Router.HandleFunc("/api/products/{id_produk}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.GetProduct))).Methods("GET")
	//Update
	s.Router.HandleFunc("/api/products/{id_produk}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateProduct))).Methods("PUT")
	//Delete
	s.Router.HandleFunc("/api/products/{id_produk}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.DeleteProduct))).Methods("DELETE")



}
