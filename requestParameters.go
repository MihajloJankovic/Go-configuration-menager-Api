package main

<<<<<<< Updated upstream
// swagger:parameters deletePost
type DeleteRequest struct {
	// Post ID
=======
<<<<<<< HEAD
// swagger:parameters deleteConfigGroup
type DeleteRequestConfigGroup struct {
	// ConfigGroup ID
=======
// swagger:parameters deletePost
type DeleteRequest struct {
	// Post ID
>>>>>>> 50d7e5a859f24150df0067dc453311c285efaddf
>>>>>>> Stashed changes
	// in: path
	Id string `json:"id"`
 }

<<<<<<< Updated upstream
 // swagger:parameters getPostById
 type GetRequest struct {
	// Post ID
=======
<<<<<<< HEAD
 // swagger:parameters deleteConfig
 type DeleteRequestConfig struct {
	// Config ID
	// in: path
	Id string `json:"id"`

	// Config Version
	// in: path
	Version string `json:"version"`
 }

 // swagger:parameters getConfigGroupById
 type GetRequestConfigGroup struct {
	// ConfigGroup ID
=======
 // swagger:parameters getPostById
 type GetRequest struct {
	// Post ID
>>>>>>> 50d7e5a859f24150df0067dc453311c285efaddf
>>>>>>> Stashed changes
	// in: path
	Id string `json:"id"`
 }

<<<<<<< Updated upstream
 // swagger:parameters post createPost
type RequestPostBody struct {
=======
<<<<<<< HEAD
 // swagger:parameters getConfigByIdAndVersion
 type GetRequestConfig struct {
	// Config ID
	// in: path
	Id string `json:"id"`

	// Config Version
	// in: path
	Version string `json:"version"`
 }

 // swagger:parameters config createConfig
type ConfigRequestBody struct {
	// - name: body
	// version:
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/ConfigRequestBody"
	//  required: true
	Body ConfigRequestBody `json:"body"`
 }

  // swagger:parameters configGroup createConfigGroup
type ConfigGroupRequestBody struct {
=======
 // swagger:parameters post createPost
type RequestPostBody struct {
>>>>>>> 50d7e5a859f24150df0067dc453311c285efaddf
>>>>>>> Stashed changes
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
<<<<<<< Updated upstream
	//     "$ref": "#/definitions/RequestPost"
	//  required: true
	Body RequestPost `json:"body"`
=======
<<<<<<< HEAD
	//     "$ref": "#/definitions/ConfigGroupRequestBody"
	//  required: true
	Body ConfigGroupRequestBody `json:"body"`
=======
	//     "$ref": "#/definitions/RequestPost"
	//  required: true
	Body RequestPost `json:"body"`
>>>>>>> 50d7e5a859f24150df0067dc453311c285efaddf
>>>>>>> Stashed changes
 }
