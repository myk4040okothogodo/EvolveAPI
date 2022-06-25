//this package contains all the database operations


package db

import (
    "github.com/jinzhu/gorm"
    // postgres blank import for gorm
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "github.com/myk4040okothogodo/EvolveAPI/pkg/types"
)


const (
   pageSize = 10
)


//ClientInterface resembles a db interface to interact with an underlying db
type ClientInterface interface {
    Ping() error
    Connect(connStr string) error
    GetUserByID(id int) *types.User
    //GetUserByEmail(email string) *types.User
    //GetUserByPhonenumber (phonenumber string) *types.User
    SetUser(user *types.User) error
    GetUsers(pageID int) *types.UserList

}

// Client is a custom db client
type Client struct {
	Client *gorm.DB
}

// Ping allows the db to be pinged.
func (c *Client) Ping() error {
	return c.Client.DB().Ping()
}


// Connect establishes a connection to the database and auto migrates the database schema
func (c *Client) Connect(connStr string) error {
	var err error
	// Create the database connection
	c.Client, err = gorm.Open(
		"postgres",
		connStr,
	)

	// End the program with an error if it could not connect to the database
	if err != nil {
		return err
	}
	c.Client.LogMode(false)
	c.autoMigrate()
	return nil
}

// autoMigrate creates the default database schema
func (c *Client) autoMigrate() {
	c.Client.AutoMigrate(&types.User{})
	
}


// GetUserByID queries a user from the database
func (c *Client) GetUserByID(id int) *types.User {
	user := &types.User{}
	c.Client.Where("id = ?", id).First(&user).Scan(user)
	return user
}


// SetUser writes a user to the database
func (c *Client) SetUser(user *types.User) error {
	// Upsert by trying to create and updating on conflict
	if err := c.Client.Create(&user).Error; err != nil {
		return c.Client.Model(&user).Where("id = ?", user.ID).Update(&user).Error
	}
	return nil
}

// GetUsers returns all users from the database
func (c *Client) GetUsers(pageID int) *types.UserList {
	users := &types.UserList{}
	c.Client.Where("id >= ?", pageID).Order("id").Limit(pageSize + 1).Find(&users.Items)
	if len(users.Items) == pageSize+1 {
		users.NextPageID = users.Items[len(users.Items)-1].ID
		users.Items = users.Items[:pageSize+1]
	}
	return users
}

