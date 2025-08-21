package util

import (
	"encoding/json"
	"os"
	"strings"
	"time"
)

// MarshalToJSON marshals data to JSON []byte, returns error if failed
func MarshalToJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// NowUTC returns current time in UTC
func NowUTC() time.Time {
	return time.Now().UTC()
}

// CreateDirectory create multiple directory.
func CreateDirectory(paths ...string) (err error) {
	for _, path := range paths {
		_, notExistError := os.Stat(path)
		if os.IsNotExist(notExistError) {
			if err = os.MkdirAll(path, os.ModePerm); err != nil {
				return err
			}
		}
	}
	return
}

// ExtractPhoneNumber extracts the phone number from a WhatsApp JID string like "6287850010020@s.whatsapp.net".
func ExtractPhoneNumber(jid string) string {
	parts := strings.Split(jid, "@")
	if len(parts) > 0 {
		return parts[0]
	}
	return ""
}

// EnsureWhatsAppJID checks if the input string contains '@s.whatsapp.net'.
// If not, it appends the suffix and returns the result.
func EnsureWhatsAppJID(jid string) string {
	if !strings.HasSuffix(jid, "@s.whatsapp.net") {
		return jid + "@s.whatsapp.net"
	}
	return jid
}
