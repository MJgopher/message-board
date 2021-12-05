package dao

import (
	"message-board-demo/model"
)

func InsertPost(post model.Post) error {
	_, err := DB.Exec("INSERT INTO post(username, txt, post_time, update_time) "+"values(?, ?, ?, ?);", post.Username, post.Txt, post.PostTime, post.UpdateTime)
	return err
}

func SelectPosts() ([]model.Post, error) {
	var posts []model.Post
	rows, err := DB.Query("SELECT id, username, txt, post_time, update_time FROM post")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	for rows.Next() {
		var post model.Post

		err = rows.Scan(&post.Id, &post.Username, &post.Txt, &post.PostTime, &post.UpdateTime)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
func DeletePost(post model.Post) error {
	_, err := DB.Exec("delete from  post where  id=?", post.Id)
	return err
}
