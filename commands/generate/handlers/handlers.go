// Package handlers is used by go generate for analizing
// controller package's files and generation of handlers.
package handlers

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/anonx/sunplate/internal/action"
	"github.com/anonx/sunplate/internal/command"
	"github.com/anonx/sunplate/internal/generation"
	"github.com/anonx/sunplate/internal/method"
	"github.com/anonx/sunplate/internal/path"
	"github.com/anonx/sunplate/log"
)

// Start is an entry point of the generate handlers command.
func Start(params command.Data) {
	inputDir := params.Default("--input", "./controllers")
	outputDir := params.Default("--output", "./assets/handlers")
	outPkg := params.Default("--package", "handlers")

	// Clean the out directory.
	log.Trace.Printf(`Removing "%s" directory if already exists...`, outputDir)
	err := os.RemoveAll(outputDir)
	log.AssertNil(err)

	// Start processing of controllers.
	ps := packages{}
	absImport := path.AbsoluteImport(inputDir)
	absImportOut := path.AbsoluteImport(outputDir)
	log.Trace.Printf(`Processing "%s" package...`, absImport)
	ps.processPackage(absImport)

	// Start generation of handler packages.
	t := generation.NewType(
		"", filepath.Join(path.SunplateDir("commands", "generate", "handlers"), "./handlers.go.template"),
	)
	t.Extension = ".go" // Save generated files as a .go source.

	// Iterate through all available packages and generate handlers for them.
	log.Trace.Printf(`Starting generation of "%s" package...`, outPkg)
	for imp := range ps {
		// Check whether current package is the main one
		// and should be stored at the root directory or it is a subpackage.
		//
		// I.e. if --input is "./controllers" and --output is "./assets/handlers",
		// we are saving processed "./controllers" package to "./assets/handlers"
		// and some "github.com/anonx/smth" to "./assets/handlers/github.com/anonx/smth".
		out := outputDir
		if imp != absImport {
			out = filepath.Join(out, imp)
		}
		t.CreateDir(out)

		// Iterate over all available controllers, generate handlers package on
		// every of them.
		for name := range ps[imp] {
			// Find parent controllers of this controller.
			cs := []parent{}
			for i, p := range ps[imp][name].Parents {
				// Make sure it is a controller rather than just some embedded struct.
				check := p.Import
				if check == "" { // Embedded parent is a local structure.
					check = absImport
				}
				if _, ok := ps[check]; !ok { // Such package is not in the list of scanned ones.
					continue
				}
				if _, ok := ps[check][p.Name]; !ok { // There is no such controller.
					continue
				}

				// It is a valid parent controller, add it to the list.
				cs = append(cs, parent{
					ID:     i,
					Import: p.Import,
					Name:   p.Name,
				})
			}

			// Initialize parameters and generate a package.
			t.Package = strings.ToLower(name)
			t.Context = map[string]interface{}{
				"after":     action.MethodAfter,
				"before":    action.MethodBefore,
				"initially": method.InitiallyName,
				"finally":   method.FinallyName,

				"controller":   ps[imp][name],
				"import":       imp,
				"input":        inputDir,
				"name":         name,
				"outputImport": absImportOut,
				"output":       outputDir,
				"package":      outPkg,
				"parents":      cs,

				"actionImport":    action.InterfaceImport,
				"actionInterface": action.Interface,
				"strconv":         action.StrconvContext,
			}
			t.Generate()
		}
	}
}
