package routes

import (
	"errors"
	"fmt"
	"github.com/p9c/cms/controller/thing"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"net/http"
	"os"
	"strconv"
)

// mapRoutes contains lots of examples of how to map things in
// Goweb.  It is in its own function so that test code can call it
// without having to run main().
func mapRoutes() {
	/*dd a pre-handler to save the referre 	*/
	goweb.MapBefore(func(c context.Context) error {
		// add a custom header
		c.HttpResponseWriter().Header().Set("X-Custom-Header", "Goweb")
		return nil
	})
	/*dd a post-handler to log somethin 	*/
	goweb.MapAfter(func(c context.Context) error {
		// TODO: log this
		return nil
	})
	/*ap the homepage.. 	*/
	goweb.Map("/", func(c context.Context) error {
		return goweb.Respond.With(c, 200, []byte("Welcome to the Goweb example app - see the terminal for instructions."))
	})
	/*ap a specific route that will redirec 	*/
	goweb.Map("GET", "people/me", func(c context.Context) error {
		hostname, _ := os.Hostname()
		return goweb.Respond.WithRedirect(c, fmt.Sprintf("/people/%s", hostname))
	})
	/*people (with optional ID 	*/
	goweb.Map("GET", "people/[id]", func(c context.Context) error {
		if c.PathParams().Has("id") {
			return goweb.API.Respond(c, 200, fmt.Sprintf("Yes, this worked and your ID is %s", c.PathParams().Get("id")), nil)
		} else {
			return goweb.API.Respond(c, 200, "Yes, this worked but you didn't specify an ID", nil)
		}
	})
	/*status-code/xxx
	Where xxx is any HTTP status code 	*/
	goweb.Map("/status-code/{code}", func(c context.Context) error {
		// get the path value as an integer
		statusCodeInt, statusCodeIntErr := strconv.Atoi(c.PathValue("code"))
		if statusCodeIntErr != nil {
			return goweb.Respond.With(c, http.StatusInternalServerError, []byte("Failed to convert 'code' into a real status code number."))
		}
		// respond with the status
		return goweb.Respond.WithStatusText(c, statusCodeInt)
	})
	// /errortest should throw a system error and be handled by the
	// DefaultHttpHandler().ErrorHandler() Handler.
	goweb.Map("/errortest", func(c context.Context) error {
		return errors.New("This is a test error!")
	})
	/*ap a RESTful controller
	(see the ThingsController for all the methods that will get
	 mapped 	*/
	thingsController := new(thing.ThingsController)
	goweb.MapController(thingsController)
	/*ap a handler for if they hit just numbers using the goweb.RegexPath
	function.

	e.g. GET /2468

	NOTE: The goweb.RegexPath is a MatcherFunc, and so comes _after_ the
	      handler 	*/
	goweb.Map(func(c context.Context) error {
		return goweb.API.RespondWithData(c, "Just a number!")
	}, goweb.RegexPath(`^[0-9]+$`))
	/*ap the static-files directory so it's exposed as /stati 	*/
	goweb.MapStatic("/static", "static-files")
	/*ap the a favico 	*/
	goweb.MapStaticFile("/favicon.ico", "static-files/favicon.ico")
	/*atch-all handler for everything that we don't understan 	*/
	goweb.Map(func(c context.Context) error {
		// just return a 404 message
		return goweb.API.Respond(c, 404, nil, []string{"File not found"})
	})

}
