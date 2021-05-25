package cache

//
//type ModMailEntryService struct {
//	*lru.Cache
//	infidel.ModMailEntryService
//}
//
//func NewModMailEntryService(size int, ms infidel.ModMailEntryService) *ModMailEntryService {
//	l, err := lru.New(size)
//	if err != nil {
//		log.Printf("Error creating new cache")
//		return nil
//	}
//	return &ModMailEntryService{l, ms}
//}
//
//func (m *ModMailEntryService) AddModMailEntry(errCh chan error, ModMailEntry infidel.ModMailEntry) {
//	m.ModMailEntryService.AddModMailEntry(errCh, ModMailEntry)
//	m.Add(ModMailEntry.ID, ModMailEntry)
//	log.Printf("Added ModMailEntry '%s' to cache", ModMailEntry.ID)
//}
//
//func (m *ModMailEntryService) GetModMailEntry(mCh chan *infidel.ModMailEntry, errCh chan error, ModMailEntryID string) {
//	ModMailEntry, ok := m.Get(ModMailEntryID)
//	if !ok {
//		m.ModMailEntryService.GetModMailEntry(mCh, errCh, ModMailEntryID)
//		m.Add(ModMailEntryID, <-mCh)
//		return
//	}
//	mem := ModMailEntry.(infidel.ModMailEntry)
//	log.Printf("Retrieved ModMailEntry '%s' from cache", ModMailEntryID)
//	errCh<-nil
//	mCh<-&mem
//}
//
//func (m *ModMailEntryService) UpdateModMailEntry(errCh chan error, ModMailEntry *infidel.ModMailEntry) {
//	m.Remove(ModMailEntry.ID)
//	m.Add(ModMailEntry.ID, *ModMailEntry)
//	m.ModMailEntryService.UpdateModMailEntry(errCh, ModMailEntry)
//}
//
//func (m *ModMailEntryService) RemoveModMailEntry(errCh chan error, ModMailEntryID string) {
//	m.Remove(ModMailEntryID)
//	m.ModMailEntryService.RemoveModMailEntry(errCh, ModMailEntryID)
//}
