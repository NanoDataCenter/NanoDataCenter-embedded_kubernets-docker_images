import (
	"lacima.com/Patterns/secrets"
)

func Init_Secret_Support(site_data map[string]interface{}) {
	secrets.Init_file_handler(site_data)
}

func Decode_access_key(key string) string {
	return secrets.Get_Secret("eto", key)
}