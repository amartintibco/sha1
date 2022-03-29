package sha1

import (
	"crypto/sha1"
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivInput  = "text"
	ovResult = "result"
)

// log is the default package logger
var log = logger.GetLogger("activity-queryparser")

// Sha1 is an Activity that is used to encrypt a text
// inputs : {key, text}
// outputs: encryptedText
type Sha1 struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &Sha1{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *Sha1) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *Sha1) Eval(context activity.Context) (done bool, err error) {
	text := context.GetInput(ivInput).(string)

	h := sha1.New()

	context.SetOutput(ovResult, h.Write([]byte(text)))

	return true, nil
}

func main() {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Printf("%x\n", bs)
}
