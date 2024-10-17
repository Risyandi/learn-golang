package product

const (
	GetProductSQL = `select id, name, description from public.refactor_table where (deleted_at is null and deleted_by is null);`

	GetProductByIDSQL = `select id, name, description from public.refactor_table where id = $1 and (deleted_at is null and deleted_by is null);`

	CreateProductSQL = `insert
	into
	public.refactor_table (
		id,
		"name",
		description,
		created_at
	)
	values($1, $2, $3, $4)
	returning id, "name", description, created_at`

	UpdateProductSQL = `update
	public.refactor_table
	set
		"name" = $1,
		description = $2,
		updated_at = $3
	where
	    id::text = $4
	returning id, "name", description, created_at, updated_at`

	DeleteProductSQL = `update
	public.refactor_table
	set
		deleted_at = $1,
		deleted_by = $2
	where
		id::text = $3`
)
