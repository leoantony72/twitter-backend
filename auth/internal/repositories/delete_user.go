package repositories

func (u *UserPostgresRepo) Delete(id string) error {
	result := u.db.Delete(id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
