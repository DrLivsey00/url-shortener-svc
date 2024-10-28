package db

type LinkService interface {
	AddToDb(longUrl, alias string) error
	GetLongUrl(alias string) (string, error)
	GetShortUrl(longUrl string) (string, error)
}

type Repository struct {
	LinkService
}

func NewRepo(db *Db) *Repository {
	return &Repository{
		LinkService: NewLinkSrv(db.DB),
	}
}
