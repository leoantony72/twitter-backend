package repositories

import "github.com/leoantony72/twitter-backend/timeline/internal/model"

type TempUserFollows struct {
	Followers []string `redis:"followers"`
}

func (t *TimelineRepo) GetFollowers(username string) []string {
	user := model.User_followers{}
	TempUserFollows := TempUserFollows{}
	followers_rows, _ := t.db.Model(&user).Select("follower").Where("followee=?", username).Rows()
	for followers_rows.Next() {
		t.db.ScanRows(followers_rows, &TempUserFollows.Followers)
	}
	return TempUserFollows.Followers
}
