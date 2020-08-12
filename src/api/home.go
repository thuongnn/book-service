package api

type HomeAPI struct {
	BaseAPI
}

// default home page
func (h *HomeAPI) Get() {
	h.ServeXML()
	h.Data["json"] = "Welcome to Book API !"
	h.ServeJSON()
}
