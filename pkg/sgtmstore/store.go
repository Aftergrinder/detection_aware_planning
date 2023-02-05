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
	GetUserByID