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
	GetUserR