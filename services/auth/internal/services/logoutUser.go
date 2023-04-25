package services

func (u *userUseCase) Logout(id string) error {
	err := u.userRepo.Logout(id)
	if err != nil {
		return err
	}
	return nil
}
