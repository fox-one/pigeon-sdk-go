package id

import (
	"fmt"
	"math/rand"
	"time"

	"strconv"

	"strings"

	"github.com/gofrs/uuid"
)

var timeLayout = "2006-01-02-15-04-05"

// GenReadableUint64ID generate readable uint64 id 19(bit)
func GenReadableUint64ID() uint64 {

	now := time.Now().UTC()
	formattedNow := now.Format(timeLayout)

	idStr := strings.ReplaceAll(formattedNow, "-", "") + fmt.Sprintf("%03d", now.Nanosecond()/1000000) + fmt.Sprintf("%02d", rand.Intn(100))

	id, _ := strconv.ParseUint(idStr, 10, 64)

	return id
}

// GenUint64ID generate normal uint64 id 19(bit)
func GenUint64ID() uint64 {
	now := time.Now().UTC().UnixNano()
	ran := rand.Intn(1000)

	idNum := now + int64(ran)

	return uint64(idNum)
}

func GenTraceID() string {
	return GenUUIDString()
}

func GenUUIDString() string {
	return uuid.Must(uuid.NewV4()).String()
}

func Num2Str(id uint64) string {
	return strconv.FormatUint(id, 10)
}

func Str2Num(idStr string) uint64 {
	v, _ := strconv.ParseUint(idStr, 10, 64)
	return v
}
