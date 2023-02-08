package play

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/CameronGorrie/sc"
	"github.com/pkg/errors"
)

// App is a command line app.
type App struct {
	// Options
	list  bool
	sound string

	c  *sc.Client
	m  map[string]*sc.Synthdef
	mu sync.RWMutex
}

// New creates a new app with some options already added:
//     -l        Lists the synthdefs for the app.
//     -s SOUND  Plays a sound.
func NewApp(fs *flag.FlagSet) *App {
	app := &App{
		m: map[string]*sc.Synthdef{},
	}

	fs.BoolVar(&app.list, "l", false, "list sounds")
	fs.StringVar(&app.sound, "s", "", "play a sound")

	return app
}

// Add adds a synthdef to the app.
// Returns an error if there is already a synthdef with the provided name.
func (app *App) Add(name string, f sc.UgenFunc) error {
	app.mu.Lock()

	if _, exists := app.m[name]; exists {
		app.mu.Unlock()
		return errors.New("sound is already defined")
	}

	app.m[name] = sc.NewSynthdef(name, f)
	app.mu.Unlock()

	return nil
}

// List prints a list of the synthdefs.
func (app *App) List() {
	app.mu.RLock()

	for name := range app.m {
		fmt.Println(name)
	}

	app.mu.RUnlock()
}

// Play plays a sound.
// params should be key-value pairs formatted as "key=value"
func (app *App) Play(sound string, params []string) error {
	app.mu.RLock()

	def, exists := app.m[sound]
	if !exists {
		app.mu.RUnlock()
		return errors.Errorf("unrecognized sound: " + sound)
	}

	app.mu.RUnlock()

	ctls := map[string]float32{}
	for _, param := range params {
		a := strings.Split(param, "=")
		if len(a) < 2 {
			return errors.Errorf("could not parse key=value from " + param)
		}

		fv, err := strconv.ParseFloat(a[1], 32)
		if err != nil {
			return errors.Wrap(err, "parsing control value")
		}

		ctls[a[0]] = float32(fv)
	}

	if err := app.c.SendDef(def); err != nil {
		return errors.Errorf("could not send def to server: ", err)
	}

	if _, err := app.c.Synth(def.Name, app.c.NextSynthID(), sc.AddToTail, sc.DefaultGroupID, ctls); err != nil {
		return errors.Wrap(err, "playing synthdef")
	}

	return nil
}

// Run runs the app.
// args should be the non-flag command line arguments.
func (app *App) Run(args []string) error {
	if app.list {
		app.List()

		return nil
	}

	return app.Play(app.sound, args)
}
