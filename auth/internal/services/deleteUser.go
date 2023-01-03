package services

func (u *userUseCase) Delete(id string) error {
	err := u.userRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
