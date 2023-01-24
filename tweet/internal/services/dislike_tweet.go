package services

func (t *TweetService) DislikeTweet(id, user string) error {
	err := t.repo.DislikeTweet(id, user)
	if err != nil {
		return err
	}
	return nil
}
