package main

// swagger:parameters deleteConfig
type DeleteRequest struct {
	// Post ID
	// in: path
	Id string `json:"id"`
}

// swagger:parameters getPostById
type GetRequest struct {
	// Post ID
	// in: path
	Id string `json:"id"`
}

<<<<<<< Updated upstream
// swagger:parameters post createPost
type RequestPostBody struct {
=======
 // swagger:parameters post createConfigGroup
type RequestConfigGroupBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/RequestPost"
	//  required: true
	Body RequestPost `json:"body"`
 }

// swagger:parameters post createConfig
type RequestConfigBody struct {
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/RequestPost"
	//  required: true
	Body RequestPost `json:"body"`
 }

 // swagger:parameters post getConfigs
type RequestConfigBody struct {
>>>>>>> Stashed changes
	// - name: body
	//  in: body
	//  description: name and status
	//  schema:
	//  type: object
	//     "$ref": "#/definitions/RequestPost"
	//  required: true

}
