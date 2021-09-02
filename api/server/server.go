package server

func InitServer() error {
	router := NewRouter()
	if err := router.Run(":8080"); err != nil {
		return err
	}
	return nil
}
