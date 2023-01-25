package util

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"strconv"
	"time"
)

func GetUnixTime(uid string) (int64, error) {
	hexTime := uid[2:4] + uid[6:8] + uid[10:12] + uid[14:16]
	time, err := strconv.ParseInt(hexTime, 16, 64)

	if err != nil {
		return 0, err
	}
	return time, nil

}

func GetUniqueId(dateString string) (string, error) {

	timeString, err := getTimeStampHex(dateString)

	if err != nil {
		log.Println(err)
		return "", err
	}

	idString, err := getRandHexString()
	if err != nil {
		log.Println(err)
		return "", err
	}

	randString, err := getRandHexString()

	if err != nil {
		log.Println(err)
		return "", err
	}

	finalKey := randString[:2] + timeString[:2] + idString[:2] + timeString[2:4] + randString[2:] + timeString[4:6] + idString[2:] + timeString[6:8]

	return finalKey, nil

}

func getTimeStampHex(dateString string) (string, error) {
	layout := "2006-01-02 15:04:05"

	timestamp, err := time.Parse(layout, dateString)
	if err != nil {
		return "", err
	}

	timeUnix := timestamp.Unix()

	return strconv.FormatInt(timeUnix, 16), nil

}

// func getMachineId() (string, error) {
// 	machineId, err := machineid.ID()
// 	if err != nil {
// 		return "", err
// 	}

// 	return machineId[:4], nil
// }

func getRandHexString() (string, error) {
	bytes := make([]byte, 2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
