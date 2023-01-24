package services

func (t *TweetService) DeleteReTweet(id, user string) error {
	err := t.repo.DeleteReTweet(id, user)
	if err != nil {
		return err
	}
	return nil
}
