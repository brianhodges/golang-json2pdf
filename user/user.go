package user
import "strings"

type User struct {
    ID int `json:"id"`
    First_Name string `json:"first_name"`
    Last_Name  string `json:"last_name"`
    Email string `json:"email"`
    Role_ID int `json:"role_id"`
}

func (u User) Full_Name() string {
    if (u.First_Name != "") {
        return strings.Title(u.First_Name) + " " + strings.Title(u.Last_Name)
    } else {
        return strings.Title(u.Last_Name)
    }
}

func (u User) Role() string {
    if (u.Role_ID == 1) {
        return "Admin"
    } else {
        return "User"
    }
}