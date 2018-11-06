package create

import (
	"crypto/md5"
	"fmt"
	"io"
	"strconv"
	"time"
)
func CreateToken() string {

	cruTime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(cruTime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}
