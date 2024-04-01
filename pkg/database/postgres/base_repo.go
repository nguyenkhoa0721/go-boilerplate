package postgres

type BaseSqlRepo struct {
	Sql *Sql
}

func NewBaseSqlRepo(sql *Sql) *BaseSqlRepo {
	return &BaseSqlRepo{
		Sql: sql,
	}
}

func (s *BaseSqlRepo) FindOne(target interface{}, query interface{}, args ...interface{}) error {
	return s.Sql.First(target, query, args).Error
}

func (s *BaseSqlRepo) FindAll(target interface{}, query interface{}, args ...interface{}) error {
	return s.Sql.Find(target, query, args).Error
}

func (s *BaseSqlRepo) Create(target interface{}) error {
	return s.Sql.Create(target).Error
}

func (s *BaseSqlRepo) Update(target interface{}) error {
	return s.Sql.Updates(target).Error
}

func (s *BaseSqlRepo) Delete(target interface{}) error {
	return s.Sql.Delete(target).Error
}
