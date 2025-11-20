package sqlite

type InvoicesRepo struct{}

func NewPostRepo(db *sql.DB, logfile *os.File) *PostRepo {
	dblogger := &DBlogger{DB: db, logfile: logfile}
	return &PostRepo{DBlogger: dblogger}
}
