package badgerepository

func (r *BadgerRepository) Close() error {
	return r.db.Close()
}
