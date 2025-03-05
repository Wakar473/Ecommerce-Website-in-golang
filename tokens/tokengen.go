package token
 import(
	jwt "github.com/dgerijalva/jwt-go"
	"os"
	"github.com/form3tech-oss/jwt-go"
 )

 type SignedDetails struct{
	Email string
	First_Name string
	Last_Name string
	Uid string
	jwt.StandardClaims
 }
var UserData *mongo.UserCollection = database.UserData(database.Client, "Users")


 var SECRET_KEY = os.Getenv("SECRET_KEY")

 func TokenGenerator(email string, firstname string, lastname string, uid string 0)(signedtoken string, signedrefreshtoken string, err errorf){

	claims := &SignedDetails{
		Email: email,
		First_Name: firstname,
		Last_Name: lastname,
		Uid: uid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Ass(time.Hour * time.Duration(24)).Unix.(),
		},
	}

	refreshclaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * Time.Duration(168)).Unix(),
		},
	}
		token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))

		if err != nil {
			return "", "", err
		}

		refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS384, refreshclaims).SignedString([]byte(SECRET_KEY))
		if err != nil {
			log.panic(err)
			return
		}
		return token, refreshtoken, err

 }

 func ValidateToken(){

 }

 func UpdateAllTokens(){

 }