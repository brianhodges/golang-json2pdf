package main
import(
    "fmt"
    "encoding/json"
    "strconv"
    "net/http"
    "io/ioutil"
    "./user"
    gofpdf "github.com/jung-kurt/gofpdf"
)

func check(err error) {
    if err != nil {
        panic(err)
    }
}

func main() {
    var users []user.User
    url := "https://golang-jsonservice.herokuapp.com/users.json"

    client := http.Client{}
    req, err := http.NewRequest(http.MethodGet, url, nil)
    check(err)

    req.Header.Set("User-Agent", "BASIC_APP_USERS")
    res, getErr := client.Do(req)
    check(getErr)

    body, readErr := ioutil.ReadAll(res.Body)
    check(readErr)

    jsonErr := json.Unmarshal(body, &users)
    check(jsonErr)

    pdf := gofpdf.New("P", "mm", "A4", "")
    pdf.AddPage()
    for _, u := range users {
        //check if page has enough room for user record, else add new page
        if pdf.GetY() > 255 {
            pdf.AddPage()
        }
        
        //User ID
        pdf.SetFont("Times", "B", 16)
        pdf.Cell(25, 10, "ID:")
        pdf.SetFont("Times", "", 16)
        pdf.SetX(pdf.GetX() + 20)
        pdf.Cell(25, 10, strconv.Itoa(u.ID))
        
        //User Full Name
        pdf.SetY(pdf.GetY() + 5)
        pdf.SetFont("Times", "B", 16)
        pdf.Cell(25, 10, "Full Name:")
        pdf.SetFont("Times", "", 16)
        pdf.SetX(pdf.GetX() + 20)
        pdf.Cell(25, 10, u.Full_Name())
        
        //User Email
        pdf.SetY(pdf.GetY() + 5)
        pdf.SetFont("Times", "B", 16)
        pdf.Cell(25, 10, "Email:")
        pdf.SetFont("Times", "", 16)
        pdf.SetX(pdf.GetX() + 20)
        pdf.Cell(25, 10, u.Email)
        
        //User Permissions (Role)
        pdf.SetY(pdf.GetY() + 5)
        pdf.SetFont("Times", "B", 16)
        pdf.Cell(25, 10, "Role:")
        pdf.SetFont("Times", "", 16)
        pdf.SetX(pdf.GetX() + 20)
        pdf.Cell(25, 10, u.Role())
        
        pdf.SetY(pdf.GetY() + 10)
        
        //footer (page numbers)
        pdf.SetFooterFunc(func() {
            pdf.SetY(-15)
            pdf.SetFont("Arial", "I", 8)
            pdf.CellFormat(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()), "", 0, "C", false, 0, "")
        })
    }
    
    err = pdf.OutputFileAndClose("output.pdf")
    check(err)
}