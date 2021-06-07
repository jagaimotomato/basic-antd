package tools

import (
	"encoding/base64"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"golang.org/x/crypto/scrypt"
	"log"
	"math/rand"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

// 生成加盐值
func Salt() string {
	return strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100000000))
}

// 验证密码
func Verify(pass, loginPass, salt string) bool {
	dk, err := scrypt.Key([]byte(loginPass), []byte(salt), 1<<15, 8, 1, 32)
	if err != nil {
		return false
	}
	// fmt.Println(pass, loginPass, salt, "***************", base64.StdEncoding.EncodeToString(dk), base64.StdEncoding.EncodeToString(dk) == pass)
	if base64.StdEncoding.EncodeToString(dk) != pass {
		return false
	}
	return true
}

// 生成加盐密码
func GeneratePassword(password string) (saltPass, salt string, err error) {
	salt = Salt()
	dk, err := scrypt.Key([]byte(password), []byte(salt), 1<<15, 8, 1, 32)
	if err != nil {
		return
	}
	saltPass = base64.StdEncoding.EncodeToString(dk)
	return
}

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func StringGroupToIntGroup(str []string) (intGroup []int) {
	for _, v := range str {
		i, _ := strconv.Atoi(v)
		intGroup = append(intGroup, i)
	}
	return
}

func StrContains(target string, strArr []string) bool {
	sort.Strings(strArr)
	index := sort.SearchStrings(strArr, target)
	if index < len(strArr) && strArr[index] == target {
		return true
	}
	return false
}

func IdStrToIdsGroup(key string, c *gin.Context) (Ids []int) {
	ids := strings.Split(c.Param(key), ",")
	for i := 0; i < len(ids); i++ {
		Id, _ := StringToInt(ids[i])
		Ids = append(Ids, Id)
	}
	return
}

func StrToStrGroup(key string, c *gin.Context) (str []string) {
	str = strings.Split(c.Request.FormValue(key), ",")
	return
}

func IntToString(i int) string {
	return strconv.Itoa(i)
}

func UintToString(u uint) string {
	return strconv.Itoa(int(u))
}

func Int64ToString(i int64) string {
	return strconv.FormatInt(i, 10)
}

func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func HasError(err error, msg string, code int) {
	if err != nil {
		statusCode := 200
		if msg == "" {
			msg = err.Error()
		}
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%v error: %#v", file, line, err)
		panic("CustomError#" + strconv.Itoa(statusCode) + "#" + msg)
	}
}

// 生成http请求的requestId
func GenerateMsgIdFromContext(c *gin.Context) string {
	var msgId string
	data, ok := c.Get("msgID")
	if !ok {
		msgId = uuid.New().String()
		c.Set("msgID", msgId)
		return msgId
	}
	msgId = cast.ToString(data)
	return msgId
}

func StructToJsonStr(i interface{}) (string, error) {
	if j, err := json.Marshal(i); err != nil {
		return string(j), err
	} else {
		return "", err
	}
}
