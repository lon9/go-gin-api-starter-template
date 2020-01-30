package server

import "{{ .ProjectPath }}/config"

// Init initialize server
func Init() error {
	c := config.GetConfig()
	r, err := NewRouter()
	if err != nil {
		return err
	}
	r.Run(":"+c.GetString("server.port"))
	return nil
}
