package domain

import (
	"errors"
)

type errMsg struct {
	tmp string
}

type errObj struct {
	UserNotLoggedIn error
	Uint64Cast      error
	FinishSession   error

	NoUser         error
	InternalServer error

	InvalidEmail    error
	InvalidUsername error
	InvalidPassword error

	EmptyField         error
	UnmatchedPasswords error
	BadPassword        error

	BadInput error

	AlreadyIn   error
	EmailExists error

	ParseId error
	SmallDb error

	InvalidRating      error
	InvalidId          error
	InvalidCommentType error

	BadGenre error

	InvalidTitle error
	PlaylistExist error
}

type err struct {
	ErrMsg errMsg
	ErrObj errObj
}

var Err = err{
	errMsg{
		tmp: "lolkek",
	},
	errObj{
		UserNotLoggedIn: errors.New("User not logged in"),
		Uint64Cast:      errors.New("Id uint64 cast error"),
		FinishSession:   errors.New("Passed through if on FinishSession"),

		NoUser:         errors.New("No user found"),
		InternalServer: errors.New("Internal server error"),

		InvalidEmail:    errors.New("Invalid email"),
		InvalidUsername: errors.New("Invalid username"),
		InvalidPassword: errors.New("Invalid password"),

		EmptyField:         errors.New("Empty field"),
		UnmatchedPasswords: errors.New("Unmatched passwords"),
		BadPassword:        errors.New("Wrong password"),

		BadInput: errors.New("Bad input"),

		AlreadyIn:   errors.New("User is already logged in"),
		EmailExists: errors.New("Email not unique"),

		ParseId: errors.New("Parse Id error"),
		SmallDb: errors.New("Sorry, our database is too small yet"),

		InvalidRating:      errors.New("Invalid value of `rating`"),
		InvalidId:          errors.New("Invalid id in database request"),
		InvalidCommentType: errors.New("Invalid comment type (expected {1, 2, 3} default:2)"),

		BadGenre: errors.New("Genre request gives empty response from db"),

		InvalidTitle: errors.New("Invalid title"),
		PlaylistExist: errors.New("Playlist with this title already exist"),
	},
}
