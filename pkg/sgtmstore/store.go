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
	CheckAndUpdatePost(post *sgtmpb