package sgtmstore

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/bwmarrin/snowflake"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"moul.io/sgtm/pkg/sgtmpb"
)

type Store interface {
	// user store
	GetUserByID(userID int64) (*sgtmpb.User, error)
	GetLastUsersList(limit int) ([]*sgtmpb.User, error)
	CreateUser(dbUser *sgtmpb.User) (*sgtmpb.User, error)
	GetUserBySlug(slug string) (*sgtmpb.User, error)
	UpdateUser(user *sgtmpb.User, updates interface{}) error
	GetUserRecentPost(userID int64) (*sgtmpb.User, error)

	// post store
	GetPostList(limit int) ([]*sgtmpb.Post, error)
	GetTrackByCID(cid string) (*sgtmpb.Post, error)
	GetTrackBySCID(scid uint64) (*sgtmpb.Post, error)
	GetLastActivities(moulID int64) ([]*sgtmpb.Post, error)
	CreatePost(post *sgtmpb.Post) error
	GetPostBySlugOrID(postSlug string) (*sgtmpb.Post, error)
	GetPostComments(postID int64) ([]*sgtmpb.Post, error)
	UpdatePost(post *sgtmpb.Post, updates interface{}) error
	GetPostListByUserID(userID int64, limit int) ([]*sgtmpb.Post, int64, error)
	CheckAndUpdatePost(post *sgtmpb.Post) error

	// counts
	GetUploadsByWeek() ([]*sgtmpb.UploadsByWeek, error)
	GetNumberOfDraftPosts() (int64, error)
	GetNumberOfUsers() (int64, error)
	GetNumberOfPostsByKind() ([]*sgtmpb.PostByKind, error)
	GetTotalDuration() (int64, error)
	GetCalendarHeatMap(authorID int64) ([]int64, error)

	// internal
	DB() *gorm.DB
}

type store struct {
	db *gorm.DB
}

func New(db *gorm.DB, sfn *snowflake.Node) (Store, error) {
	wrap, err := dbInit(db, sfn)
	if err != nil {
		return nil, fmt.Errorf("db init: %w", err)
	}
	return &store{db: wrap}, nil
}

func (s *store) DB() *gorm.DB { return s.db }

func (s *store) GetUserByID(userID int64) (*sgtmpb.User, error) {
	var user sgtmpb.User

	err := s.db.
		Where("id = ?", userID).
		First(&user).
		Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *store) GetLastUsersList(limit int) ([]*sgtmpb.User, error) {
	var users []*sgtmpb.User
	err := s.db.
		Order("created_at desc").
		Limit(limit).
		Find(&users).
		Error
	if err != nil {
		return nil, err
	}

	for _, u := range users {
		u.Filter()
	}
	return users, nil
}

func (s *store) GetPostList(limit int) ([]*sgtmpb.Post, error) {
	var p