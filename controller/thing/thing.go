package thing

import (
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"net/http"
)

// Thing is just a thing
type Thing struct {
	Id   string
	Text string
}

// ThingsController is the RESTful MVC controller for Things.
type ThingsController struct {
	// Things holds the things... obviously, you would never do this
	// in the real world - you'd be reading from some kind of datastore.
	Things []*Thing
}

// Before gets called before any other method.
func (r *ThingsController) Before(ctx context.Context) error {

	// set a Things specific header
	ctx.HttpResponseWriter().Header().Set("X-Things-Controller", "true")

	return nil

}

func (r *ThingsController) Create(ctx context.Context) error {

	data, dataErr := ctx.RequestData()

	if dataErr != nil {
		return goweb.API.RespondWithError(ctx, http.StatusInternalServerError, dataErr.Error())
	}

	dataMap := data.(map[string]interface{})

	thing := new(Thing)
	thing.Id = dataMap["Id"].(string)
	thing.Text = dataMap["Text"].(string)

	r.Things = append(r.Things, thing)

	return goweb.Respond.WithStatus(ctx, http.StatusCreated)

}

func (r *ThingsController) ReadMany(ctx context.Context) error {
	if r.Things == nil {
		r.Things = make([]*Thing, 0)
	}
	return goweb.API.RespondWithData(ctx, r.Things)
}

func (r *ThingsController) Read(id string, ctx context.Context) error {
	for _, thing := range r.Things {
		if thing.Id == id {
			return goweb.API.RespondWithData(ctx, thing)
		}
	}
	return goweb.Respond.WithStatus(ctx, http.StatusNotFound)
}

func (r *ThingsController) DeleteMany(ctx context.Context) error {
	r.Things = make([]*Thing, 0)
	return goweb.Respond.WithOK(ctx)
}

func (r *ThingsController) Delete(id string, ctx context.Context) error {
	newThings := make([]*Thing, 0)
	for _, thing := range r.Things {
		if thing.Id != id {
			newThings = append(newThings, thing)
		}
	}
	r.Things = newThings
	return goweb.Respond.WithOK(ctx)
}
