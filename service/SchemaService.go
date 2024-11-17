package service

func (h *ServiceDB) Createtables() error {
	return h.RepoDB.Createtables()
}
