package api

type HomeAPI struct {
	BaseAPI
}

// default home page
func (h *HomeAPI) Get() {
	h.TplName = "index.html"
	_ = h.Render()
}
