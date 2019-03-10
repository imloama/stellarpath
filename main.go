package main
import (
	"fmt"
    "log"
    "github.com/stellar/go/keypair"
	"github.com/jinzhu/configor"
)

var Config = struct {
	
	Secret String
	
	Assets []struct{
		Code string
		Issuer string
	}
	
	Paths []string
	
	/*
	APPName string `default:"app name"`

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}

	Contacts []struct {
		Name  string
		Email string `required:"true"`
	}
	*/
}{}

func main(){
	configor.Load(&Config, "config.toml")
	fmt.Printf("config: %#v", Config)

	pair, err := keypair.Random()
    if err != nil {
        log.Fatal(err)
    }

    log.Println(pair.Seed())
    log.Println(pair.Address())
}