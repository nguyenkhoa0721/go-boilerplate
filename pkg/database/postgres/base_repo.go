package postgres

type BaseSqlRepo struct {
	sql *Sql
}

func NewBaseSqlRepo(sql *Sql) *BaseSqlRepo {
	return &BaseSqlRepo{
		sql: sql,
	}
}

func (s *BaseSqlRepo) FindOne(target interface{}, query interface{}, args ...interface{}) error {
	return s.sql.First(target, query, args).Error
}

func (s *BaseSqlRepo) FindAll(target interface{}, query interface{}, args ...interface{}) error {
	return s.sql.Find(target, query, args).Error
}

func (s *BaseSqlRepo) Create(target interface{}) error {
	return s.sql.Create(target).Error
}

func (s *BaseSqlRepo) Update(target interface{}) error {
	return s.sql.Updates(target).Error
}
