package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/masato25/my104job_board/models"
	"github.com/pkg/errors"
)

func JobQuery(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)
	jobs := models.Jobs{}
	// Default values are "page=1" and "per_page=20".
	q := tx.Where("cat = ? and sal_low >= ?", c.Param("cat"), c.Param("salary")).PaginateFromParams(c.Params())
	// You can order your list here. Just change
	err := q.All(&jobs)
	// to:
	// err := q.Order("created_at desc").All(jobs)
	if err != nil {
		return errors.WithStack(err)
	}
	return c.Render(
		http.StatusOK,
		r.JSON(map[string]interface{}{"status": "ok", "data": jobs, "pagging": q.Paginator}),
	)
}
