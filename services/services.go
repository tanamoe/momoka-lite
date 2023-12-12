package services

func Start() error {
	if err := startUpdateSlugService(); err != nil {
		return err
	}
	return nil
}
