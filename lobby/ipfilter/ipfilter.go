package ipfilter

import (
	"encoding/binary"
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/lrita/cmap"
)

var _EXISTS struct{}

var _IPFILTER cmap.Map[uint32, struct{}]

// check if the client IP is blocked in _IPFILTER
func BlockIP() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, ok := _IPFILTER.Load(binary.BigEndian.Uint32([]byte(c.ClientIP())))
		if ok {
			c.AbortWithStatus(http.StatusForbidden)

			return
		}
	}
}

var _BANNED_BEGINS = []string{
	"/manager",
	"/php",
	"/sql",
	"/web",
	"/pma",
	"/cgi-bin",
	"/?XDEBUG_SESSION_START",
	"/_",
	"/vendor",
	"/dbadmin",
	"/my",
	"/admin",
	"/db",
}

var _BANNED_STRINGS = []string{
	".cgi",
	"wget",
	"curl",
	"sh",
	"php",
	"..",
}

// init the module lowercosing and sorting the _BANNED_BEGINS so they can be searched faster
func Init() {
	// lowecase the array
	for i := 0; i < len(_BANNED_BEGINS); i++ {
		_BANNED_BEGINS[i] = strings.ToLower(_BANNED_BEGINS[i])
	}
	// sort the array
	sort.Strings(_BANNED_BEGINS)
}

// check if the any of the emlements of the slice matches the beginning of item.
// it expects a sorted slice
func beginsWith(prefixes []string, item string) bool {

	NumPrefixes := len(prefixes)
	item = strings.ToLower(item)

	// for each prefix in the slice
	for i := 0; i < NumPrefixes; i++ {

		// if the element in the slice is smaller than the item
		// and the item is bigger we know we will not find it anymore
		if len(item) >= len(prefixes[i]) && item[:len(prefixes[i])] < prefixes[i] {
			return false
		}

		// if it has the prefix we found it
		if strings.HasPrefix(item, prefixes[i]) {
			return true
		}
	}

	// if we have finished the slice we have not found it
	return false
}

// check if any of th elements in the slice are part of the item
func contains(slice []string, item string) bool {

	length := len(slice)

	for i := 0; i < length; i++ {
		if strings.Contains(item, slice[i]) {
			return true
		}

	}
	return false
}

func BanIfNeeded(c *gin.Context) {
	if beginsWith(_BANNED_BEGINS, c.Request.URL.Path) || contains(_BANNED_STRINGS, c.Request.URL.Path) {
		_IPFILTER.Store(binary.BigEndian.Uint32([]byte(c.ClientIP())), _EXISTS)

		c.AbortWithStatus(http.StatusForbidden)

		return
	}

	c.Data(http.StatusNotFound, gin.MIMEHTML, []byte("404 page not found"))
}
