package limit

import (
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/kyaxcorp/go-helper/conv"
)

var hideAllObjects bool

// var showObjects bool
// var hideObjects bool

var rShowObjects []*regexp.Regexp
var rHideObjects []*regexp.Regexp

func init() {
	sHideAllObjects := os.Getenv("KYAX_OBJECTS_HIDE_ALL")
	sShowObjects := os.Getenv("KYAX_OBJECTS_SHOW")
	sHideObjects := os.Getenv("KYAX_OBJECTS_HIDE")

	hideAllObjects = conv.ParseBool(sHideAllObjects)
	// showObjects = conv.ParseBool(sShowObjects)
	// hideObjects = conv.ParseBool(sHideObjects)

	// object should be delimited by comma
	// different patterns can be used, mostly regex!

	if sShowObjects != "" {
		showObjectsSlice := strings.Split(sShowObjects, ",")
		for _, obj := range showObjectsSlice {
			if obj == "" {
				continue
			}
			regexCompiled, err := regexp.Compile(obj)
			if err != nil {
				log.Println("Error compiling regex for show objects:", err)
				// return objects
			}
			rShowObjects = append(rShowObjects, regexCompiled)
		}
	}

	if sHideObjects != "" {
		hideObjectsSlice := strings.Split(sHideObjects, ",")
		for _, obj := range hideObjectsSlice {
			if obj == "" {
				continue
			}
			regexCompiled, err := regexp.Compile(obj)
			if err != nil {
				log.Println("Error compiling regex for hide objects:", err)
				// return objects
			}
			rHideObjects = append(rHideObjects, regexCompiled)
		}
	}

}

func ObjectAllow(matchString string) bool {

	// matchString := fmt.Sprintf("%s.%s", schema, name)

	if conv.ParseBool(hideAllObjects) {
		// Check only in the showObjects
		if len(rShowObjects) > 0 {
			showObject := false
			for _, r := range rShowObjects {
				if r.MatchString(matchString) {
					showObject = true
					break
				}
			}
			if !showObject {
				return false
			}
		} else {
			return false
		}
	} else {
		// check only in the hideObjects
		if len(rHideObjects) > 0 {
			hideObject := false
			for _, r := range rHideObjects {
				if r.MatchString(matchString) {
					hideObject = true
					break
				}
			}
			if hideObject {
				return false
			}
		} else {
			return false
		}
	}
	return true
}
