package model

import (
	"net/http"
	"time"

	"github.com/anujc4/tweeter_api/internal/app"
	"github.com/anujc4/tweeter_api/request"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type Tweet struct {
	ID        int `gorm:"primarykey"`
	UserId    int
	Content   string
	ParentTweet  int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tweets []*Tweet

func (appModel *AppModel) CreateTweet(request *request.CreateTweetRequest) (*Tweet, *app.Error) {
	tweet := Tweet{
		UserId: request.UserId,
		Content:  request.Content,
		ParentTweet: request.ParentTweet,
	}
	result := appModel.DB.Create(&tweet)

	if result.Error != nil {
		me, ok := result.Error.(*mysql.MySQLError)
		if !ok {
			return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
		}
		if me.Number == 1062 {
			return nil, app.
				NewError(result.Error).
				SetMessage("Email " + request.Content + " is already taken").
				SetCode(http.StatusBadRequest)
		}
		return nil, app.NewError(result.Error).SetCode(http.StatusBadRequest)
	}
	return &tweet, nil
}

func (appModel *AppModel) GetTweets(request *request.GetTweetsRequest) (*Tweets, *app.Error) {
	var tweets Tweets
	var where *gorm.DB = appModel.DB
	var page, pageSize int

	if request.ID != 0 {
		where = appModel.DB.Where("ID = ?", request.ID)
	} else if request.ParentTweet != 0 {
		where = appModel.DB.Where("parent_tweet = ?", request.ParentTweet)
	}else if request.UserId != 0 {
		where = appModel.DB.Where("user_id = ?",request.UserId)
	}

	if request.Page == 0 {
		page = 1
	} else {
		page = request.Page
	}

	switch {
	case request.PageSize > 100:
		pageSize = 100
	case request.PageSize <= 0:
	}
	if request.PageSize <= 0 {
		pageSize = 10
	} else {
		pageSize = request.PageSize
	}


	offset := (page - 1) * pageSize

	result := where.
		Offset(offset).
		Limit(pageSize).
		Find(&tweets)

	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}

	return &tweets, nil
}


func (appModel *AppModel) GetTweetByID(ID string) (*Tweet, *app.Error) {
	var tweet Tweet
	result := appModel.DB.First(&tweet, ID)
	if result.Error != nil {
		return nil, app.NewError(result.Error).SetCode(http.StatusNotFound)
	}
	return &tweet, nil
}





