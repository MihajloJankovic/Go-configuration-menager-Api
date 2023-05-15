package main

<<<<<<< Updated upstream
=======
<<<<<<< HEAD
// swagger:response ResponseConfig
type ResponseConfig struct {
	// Id of the config
   	// in: string
	Id string `json:"id"`
	
	// Version of the config
	// in: string
	Version string `json:"version"`
}

// swagger:response ResponseConfigGroup
type ResponseConfigGroup struct {
	// Id of the configGroup
   	// in: string
	Id string `json:"id"`
=======
>>>>>>> Stashed changes
// swagger:response ResponsePost
type ResponsePost struct {
	// Id of the config
	// in: string
	Id      string            `json:"id"`

	// Version of the config
	// in: string
	Version string            `json:"version"`

<<<<<<< Updated upstream
=======
>>>>>>> 50d7e5a859f24150df0067dc453311c285efaddf
>>>>>>> Stashed changes
}

// swagger:response ErrorResponse
type ErrorResponse struct {
	// Error status code
	// in: int64
	Status int64 `json:"status"`
	// Message of the error
	// in: string
	Message string `json:"message"`
 }

 // swagger:response NoContentResponse
<<<<<<< Updated upstream
type NoContentResponse struct {}
 
=======
<<<<<<< HEAD
type NoContentResponse struct {}
=======
type NoContentResponse struct {}
 
>>>>>>> 50d7e5a859f24150df0067dc453311c285efaddf
>>>>>>> Stashed changes
