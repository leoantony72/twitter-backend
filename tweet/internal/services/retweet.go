package services

func (t *TweetService) ReTweet(id,user string) error {
	err := t.repo.ReTweet(id, user)
	if err != nil {
		return err
	}
	return nil
}
