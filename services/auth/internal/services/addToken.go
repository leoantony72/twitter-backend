package services

func (u *userUseCase) AddToken(username, token string) error {
	err := u.userRepo.AddToken(username, token)
	if err != nil {
		return err
	}
	return nil
}
