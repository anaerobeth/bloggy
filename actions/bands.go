package actions

import (
	"github.com/anaerobeth/bloggy/models"
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Band)
// DB Table: Plural (bands)
// Resource: Plural (Bands)
// Path: Plural (/bands)
// View Template Folder: Plural (/templates/bands/)

// BandsResource is the resource for the band model
type BandsResource struct {
	buffalo.Resource
}

// List gets all Bands. This function is mapped to the path
// GET /bands
func (v BandsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	bands := &models.Bands{}
	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=2".
	q := tx.PaginateFromParams(c.Params())
	// You can order your list here. Just change
	err := q.All(bands)
	// to:
	// err := q.Order("create_at desc").All(bands)
	if err != nil {
		return errors.WithStack(err)
	}
	// Make Bands available inside the html template
	c.Set("bands", bands)
	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)
	return c.Render(200, r.HTML("bands/index.html"))
}

// Show gets the data for one Band. This function is mapped to
// the path GET /bands/{band_id}
func (v BandsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty Band
	band := &models.Band{}
	// To find the Band the parameter band_id is used.
	err := tx.Find(band, c.Param("band_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Make band available inside the html template
	c.Set("band", band)
	return c.Render(200, r.HTML("bands/show.html"))
}

// New renders the formular for creating a new Band.
// This function is mapped to the path GET /bands/new
func (v BandsResource) New(c buffalo.Context) error {
	// Make band available inside the html template
	c.Set("band", &models.Band{})
	return c.Render(200, r.HTML("bands/new.html"))
}

// Create adds a Band to the DB. This function is mapped to the
// path POST /bands
func (v BandsResource) Create(c buffalo.Context) error {
	// Allocate an empty Band
	band := &models.Band{}
	// Bind band to the html form elements
	err := c.Bind(band)
	if err != nil {
		return errors.WithStack(err)
	}
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(band)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		// Make band available inside the html template
		c.Set("band", band)
		// Make the errors available inside the html template
		c.Set("errors", verrs)
		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("bands/new.html"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "Band was created successfully")
	// and redirect to the bands index page
	return c.Redirect(302, "/bands/%s", band.ID)
}

// Edit renders a edit formular for a band. This function is
// mapped to the path GET /bands/{band_id}/edit
func (v BandsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty Band
	band := &models.Band{}
	err := tx.Find(band, c.Param("band_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Make band available inside the html template
	c.Set("band", band)
	return c.Render(200, r.HTML("bands/edit.html"))
}

// Update changes a band in the DB. This function is mapped to
// the path PUT /bands/{band_id}
func (v BandsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty Band
	band := &models.Band{}
	err := tx.Find(band, c.Param("band_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	// Bind Band to the html form elements
	err = c.Bind(band)
	if err != nil {
		return errors.WithStack(err)
	}
	verrs, err := tx.ValidateAndUpdate(band)
	if err != nil {
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		// Make band available inside the html template
		c.Set("band", band)
		// Make the errors available inside the html template
		c.Set("errors", verrs)
		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("bands/edit.html"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "Band was updated successfully")
	// and redirect to the bands index page
	return c.Redirect(302, "/bands/%s", band.ID)
}

// Destroy deletes a band from the DB. This function is mapped
// to the path DELETE /bands/{band_id}
func (v BandsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	// Allocate an empty Band
	band := &models.Band{}
	// To find the Band the parameter band_id is used.
	err := tx.Find(band, c.Param("band_id"))
	if err != nil {
		return errors.WithStack(err)
	}
	err = tx.Destroy(band)
	if err != nil {
		return errors.WithStack(err)
	}
	// If there are no errors set a flash message
	c.Flash().Add("success", "Band was destroyed successfully")
	// Redirect to the bands index page
	return c.Redirect(302, "/bands")
}
