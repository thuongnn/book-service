package api

type HomeAPI struct {
	BaseAPI
}

// default home page
func (h *HomeAPI) Get() {
	h.Data["json"] = "Welcome to book service !"
	h.ServeJSON()
}
