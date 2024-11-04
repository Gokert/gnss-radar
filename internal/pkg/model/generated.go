// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

// Выходные параметры для проверки авторизации
type AuthcheckInput struct {
	//  Пусто
	Empty *string `json:"_empty,omitempty"`
}

// Выходные параметры для проверки авторизации
type AuthcheckOutput struct {
	//  Информация о юзере
	UserInfo *User `json:"userInfo"`
}

// Мутации связанные с авторизацией
type AuthorizationMutations struct {
	//  Регистрация
	Signup *SignupOutput `json:"signup"`
	//  Авторизация
	Signin *SigninOutput `json:"signin"`
	//  Выход
	Logout *LogoutOutput `json:"logout,omitempty"`
}

type CoordsInput struct {
	//  Координата X
	X string `json:"x"`
	//  Координата Y
	Y string `json:"y"`
	//  Координата Z
	Z string `json:"z"`
}

type CoordsResults struct {
	//  Координата X
	X string `json:"x"`
	//  Координата Y
	Y string `json:"y"`
	//  Координата Z
	Z string `json:"z"`
}

type Gnss struct {
	ID string `json:"Id"`
	//  id спутника
	SatelliteID string `json:"SatelliteId"`
	//  Имя спутника
	SatelliteName string `json:"SatelliteName"`
	//  Координаты спутника
	Coordinates *CoordsResults `json:"Coordinates"`
}

type GNSSFilter struct {
	//  Фильтр по индетификаторам
	Coordinates *CoordsInput `json:"Coordinates"`
}

type GNSSPagination struct {
	//  Загруженные элементы
	Items []*Gnss `json:"items,omitempty"`
}

// Выходные параметры для выхода
type LogoutInput struct {
	//  Пусто
	Empty *string `json:"_empty,omitempty"`
}

// Выходные параметры для выхода
type LogoutOutput struct {
	//  Пусто
	Empty *string `json:"_empty,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

// Входные параметры для авторизации
type SigninInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Выходные параметры для авторизации
type SigninOutput struct {
	//  Информация о юзере
	UserInfo *User `json:"userInfo"`
}

// Входные параметры для регистрации
type SignupInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

// Выходные параметры для регистрации
type SignupOutput struct {
	//  Информация о юзере
	UserInfo *User `json:"userInfo"`
}

// Бизнес ошибки
type Error string

const (
	//  Уже cуществует
	ErrorAlreadyExists Error = "ALREADY_EXISTS"
	//  Не авторизован
	ErrorNotAuthorized Error = "NOT_AUTHORIZED"
	//  Не найден
	ErrorNotFound Error = "NOT_FOUND"
	//  Нет прав
	ErrorPermissionDenied Error = "PERMISSION_DENIED"
	//  Ошибка сервера
	ErrorInternalError Error = "INTERNAL_ERROR"
	//  Ошибка запроса
	ErrorBadRequest Error = "BAD_REQUEST"
)

var AllError = []Error{
	ErrorAlreadyExists,
	ErrorNotAuthorized,
	ErrorNotFound,
	ErrorPermissionDenied,
	ErrorInternalError,
	ErrorBadRequest,
}

func (e Error) IsValid() bool {
	switch e {
	case ErrorAlreadyExists, ErrorNotAuthorized, ErrorNotFound, ErrorPermissionDenied, ErrorInternalError, ErrorBadRequest:
		return true
	}
	return false
}

func (e Error) String() string {
	return string(e)
}

func (e *Error) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Error(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Error", str)
	}
	return nil
}

func (e Error) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
