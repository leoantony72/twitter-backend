package services

func (t *TweetService) DeleteTweet(id, user string) error {
	err := t.repo.DeleteTweet(id, user)
	if err != nil {
		return err
	}
	return nil
}
