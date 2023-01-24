package services

func (t *TweetService) LikeTweet(id, user string) error {
	err := t.repo.LikeTweet(id, user)
	if err != nil {
		return err
	}
	return nil
}
