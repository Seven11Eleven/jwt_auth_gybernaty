package domain

import "errors"

var (
	//Возвращается, когда заголовок статьи не соотвествует требованиям
	//http status code - 400 Bad Request
	ErrInvalidTitle = errors.New("title must be between 4 and 100 characters and contain only letters")

	//Возвращается, когда контент статьи не соотвествует требованиям
	//http status code - 400 Bad Request
	ErrInvalidContent = errors.New("content must contain only letters")

	//Возвращается, когда юзернейм не соотвествует требованиям (состоит не из только латинских букв)
	//http status code - 400 Bad Request
	ErrInvalidUsername = errors.New("username must contain only latin letters")

	//Возвращается, когда пытаются создать аккаунт с ником, который уже существует
	//http status code - 409 Conflict
	ErrUsernameExists = errors.New("username already exists")

	//Возвращается, когда автор не был найден в базе данных
	//http status code - 404 Not Found
	ErrAuthorNotFound = errors.New("author not found")

	//http status code - 404 Not Found
	ErrArticleNotFound = errors.New("article not found")

	// http status code - 401 Unathorized Request https://www.webnots.com/4xx-http-status-codes/
	ErrInvalidCredentials = errors.New("invalid credentials")
	// http status code - 401 Unathorized Request
	ErrInvalidToken = errors.New("invalid token")

	ErrTokenCreation = errors.New("token creation error")
)
