package api

type HomeAPI struct {
	BaseAPI
}

// default home page
func (h *HomeAPI) Get() {
	h.Data["json"] = "Welcome to Book API !\n" +
		"GET /books \tList all books in DB\n" +
		"POST /books \tCreate one book in DB\n" +
		"GET /books/:id \tGet one book by id\n" +
		"PUT /books/:id \tUpdate one book by id\n" +
		"DELETE /books/:id \tDelete one book by id"
	h.ServeJSON()
}
