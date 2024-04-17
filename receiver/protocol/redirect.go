package protocol

// / 64
type SiteRedirect struct {
	Port uint16
	Site [50]byte
	Keep [12]byte
}
