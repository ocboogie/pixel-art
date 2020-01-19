package api

import "net/http"

var (
	errInvalidCredentials = newSimpleAPIError(http.StatusBadRequest, true, "Username or password is invalid")
	errEmailAlreadyInUse  = newSimpleAPIError(http.StatusBadRequest, true, "Email already in use")
	errUnauthenticated    = newSimpleAPIError(http.StatusUnauthorized, true, "You must be logged in")
	errInvalidAvatar      = newSimpleAPIError(http.StatusBadRequest, false, "Invalid avatar")
	errAlreadyLiked       = newSimpleAPIError(http.StatusConflict, false, `Already liked that post`)
	errInvalidUserState   = newSimpleAPIError(http.StatusBadRequest, false, `Invalid user state`)
	errLikeNotFound       = newSimpleAPIError(http.StatusNotFound, false, `Like not found`)
	errPostNotFound       = newSimpleAPIError(http.StatusNotFound, false, "Post not found")
	errInvalidLimit       = newSimpleAPIError(http.StatusBadRequest, false, `The "limit" parameter must be a number`)
	errInvalidAfter       = newSimpleAPIError(http.StatusBadRequest, false, `The "after" parameter must be a iso-8601 formatted date`)
	errInvalidArt         = newSimpleAPIError(http.StatusBadRequest, false, "Art data does not meet specifications")
)
