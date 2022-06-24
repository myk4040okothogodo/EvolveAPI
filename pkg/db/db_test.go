package  db

import (
     "time"
    "fmt"
    "log"
    "os"
    "testing"
    "github.com/ory/dockertest/v3"
    "github.com/stretchr/testify/assert"
    "github.com/myk4040okothogodo/EvolveAPI/pkg/types"
)


var (
    testClient = &Client{}
    testUser = types.User{
        Email:  "johndoe@gmail.com",
        Phonenumber: "+254717256998",
        Address : "Banda-street 27",
        Birthday: '2021-08-19 00:39:56.191',
        Createdat: time.Now(),
    }

)

func TestMain(m *testing.M) {
    //uses a sensible default on windows (tcp/http) and linux (socket)
    pool, err := dockertest.NewPool("")
    if err != nil {
        log.Fatalf("Couldnt connect to docker : %s", err)
    }

    //pulls an image, creates a container based on it and runs it
    r, err := pool.Run("Postgres","13-alpine", []string{"POSTGRES_PASSWORD=secret"})
    if err != nil {
        log.Fatalf("Couldnt start the resource: %s", err)
    }

    //exponantial backoff-retry, because an application in the container might not be ready to acsept the connections yet
    if err := pool.Retry(func() error {
        if err := testClient.Connect(fmt.Sprintf("host=localhost port=%s user=postgres dbname=postgres password=secret sslmode=disable", r.GetPort("5432/tcp"))); err != nil {
            return err
        }
        return testClient.Ping()
    }); err != nil {
        log.Fataf("couldnt connect to docker: %s", err)
    }

    code := m.Run()

    //You cant defer this because os.Exit doesnt care for defer
    if err := pool.Purge(r); err != nil {
        log.Fatelf("Couldnt purge resource: %s", err)
    }

    os.Exit(code)

}

func TestClient_Users(t *testing.T){
    testClient.Client.DropTable(&types.User{})
    testClient.autoMigrate()
    first := testUser
    err := testClient.SetUser(&first)
    assert.NoError(t, err)
    assert.Equal(t, 1, first.ID)

    second := testUser
    err = testClient.SetUser(&second)
    assert.NoError(t, err)
    assert.Equal(t, 2, second.ID)

    update := first 
    update.Phonenumber = "+254717000000"
    err = testClient.SetUser(&update)
    assert.NoError(t, err)

    got := testClient.GetUserByID(1)
    assert.Equal(t, testUser.Email, got.Name, "")
    assert.Equal(t, "+254717000000", got.Phonenumber, "")

    got =  testClient.GetUserByID(2)
    assert.Equal(t, testUser.Email, got.Email, "")
    assert.Equal(t, "+254717256998", got.Phonenumber, "")
}


func TestClient_PaginateUsers(t *testing.T){
    testClient.Client.DropTable(&types.User{})
    testClient.autoMigrate()
    for i := 0; i < pageSize+2; i++{
        user := testUser
        _ = testClient.SetUser(&user)
    }

    got := testClient.GetUsers(0)
    assert.Equal(t, 10, len(got.Items))
    assert.Equal(t, 11, got.NextPageID)

    got = testClient.GetUsers(11)
    assert.Equal(t, 2, len(got.Items))
    assert.Equal(t, 0, got.NextPageID)
}
