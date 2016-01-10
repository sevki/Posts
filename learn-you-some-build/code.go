
import (
	"crypto/sha1"
	"fmt"
	"io"
	"net/http"
	"path/filepath"

	"sevki.org/build/util"
)

// START1 OMIT
// Target defines the interface that rules must implement for becoming build targets.
type Target interface {
	GetName() string
	GetDependencies() []string

	Hash() []byte
	Build(*Context) error
	Installs() map[string]string // map[dst]src
}

// END1 OMIT
// START2 OMIT
func (cl *CLib) Hash() []byte {
	h := sha1.New()
	io.WriteString(h, CCVersion)
	io.WriteString(h, cl.Name)
	util.HashFiles(h, cl.Includes)
	io.WriteString(h, "clib")
	util.HashFiles(h, []string(cl.Sources))
	util.HashStrings(h, cl.CompilerOptions)
	util.HashStrings(h, cl.LinkerOptions)
	if cl.LinkShared {
		io.WriteString(h, "shared")
	}
	if cl.LinkStatic {
		io.WriteString(h, "static")
	}
	return h.Sum(nil)
}

// END2 OMIT
// START3 OMIT
func Hash() []byte {
	res, _ := http.Get("http://www.yr.no/place/Norway/Nord-Tr%C3%B8ndelag/Stj%C3%B8rdal/Hell/")
	h := sha1.New()
	io.Copy(h, res.Body)
	return h.Sum(nil)
}

// END3 OMIT
// START4 OMIT
func (cl *CLib) Installs() map[string]string {
	exports := make(map[string]string)
	libName := fmt.Sprintf("%s.a", cl.Name)

	exports[filepath.Join("lib", libName)] = libName

	return exports
}

// END4 OMIT
