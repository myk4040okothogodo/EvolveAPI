package types

import (
     "time"
    "net/http"
    "github.com/go-chi/render"

)


// User is one instace of a user object
type User struct {
  // The Unique id of this user
  ID          int      `gorm:"primaryKey" json:"id" example:"1"`
  // The email address associated with this user
  Email       string   `json:"email" example:"johndoe@gamil.com"`
  // The phone number associated with this user
  Phonenumber string   `json:"phonenumber" example:"0717256998"`
  // The address associated with this user
  Address     string   `json:"address" example:"johns-street 27"`
  // The date of birth of this user
  Birthday    time.Time `json:"birthday" example:"2021-08-19 00:39:56.191"`
  // the creation date and time of this user on the system
  Createdat   time.Time `gorm:"autoCreateTime" json:"createdat" example:"2021-08-19 00:39:56.191"`
} // @name User


// Render implements the github.com/go-chi/render.Render interfac
func (u *User) Render (w http.ResponseWriter, r *http.Request) error {
    return nil
}


//Bind implements the github.com/go-chi/render.Render interface
func (u *User) Bind (r *http.Request) error {
    return nil
}


//UserList contains a list of users
type UserList struct {
    // A list of users
    Items []*User  `json:"items"`
    // The id to query the next page
    NextPageID int `json:"next_page_id, omitempty" example:"10"`
    // The page size of the query
    //PageSize int  `json: "page_size,omitempty" example: "5"`
} // @name UserList


//Render implements the github.com/go-chi/render.Render interface
func (u *UserList) Render(w http.ResponseWriter, r *http.Request) error {
    return nil
}


//ErrResponse renderer type for handling all sorts of errors
type ErrResponse struct {
    Err             error    `json:"-"`  // low-level runtime error
    HTTPStatusCode  int      `json:"-"`  //http response status code

    StatusText      string   `json: "status" example:"Resource not found."`  // user-level status message
    AppCode         int64    `json: "code,omitempty" example:"404"`          // application-specific error code
    ErrorText       string   `json: "error,omitempty" example:"The requested resource was not found on the server"` // application-level error message, for debugging
} // @name ErrorResponse


//Render implements the github.com/go-chi/render.Renderer interface for ErrResponse
func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
    render.Status(r, e.HTTPStatusCode)
    return nil
}


//ErrInvalidRequest returns a structured http response for invalid requests
func ErrInvalidRequest(err error) render.Renderer{
    return &ErrResponse{
        Err:                    err,
        HTTPStatusCode:         http.StatusBadRequest,
        StatusText:             "Invalid request.",
        ErrorText:              err.Error(),
    }
}


//ErrRender returns a structured http response in case of rendering error
func ErrRender(err error) render.Renderer{
    return &ErrResponse {
        Err:                      err,
        HTTPStatusCode:           http.StatusUnprocessableEntity,
        StatusText:               "Error rendering response",
        ErrorText:                err.Error(),
    }
}

//ErrNotFound returns a structured http response if a resource couldnt be found
func ErrNotFound() render.Renderer {
    return &ErrResponse{
        HTTPStatusCode: http.StatusNotFound,
        StatusText:     "Resource not found.",
    }
}
