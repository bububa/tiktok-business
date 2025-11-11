package enum

// RepresentativeDocumentType 法人证件类型。
type RepresentativeDocumentType string

const (
	// RepresentativeDocumentType_ID_CARD ：居民身份证。
	RepresentativeDocumentType_ID_CARD RepresentativeDocumentType = "ID_CARD"
	// RepresentativeDocumentType_PASSPORT：护照。
	RepresentativeDocumentType_PASSPORT RepresentativeDocumentType = "PASSPORT"
	// RepresentativeDocumentType_HK_MACAO_EXIT_ENTRY_PERMIT ：往来港澳通行证，又称港澳通行证。
	RepresentativeDocumentType_HK_MACAO_EXIT_ENTRY_PERMIT RepresentativeDocumentType = "HK_MACAO_EXIT_ENTRY_PERMIT"
	// RepresentativeDocumentType_TAIWAN_MAINLAND_TRAVEL_PERMIT ：台湾居民来往大陆通行证，又称台胞证。
	RepresentativeDocumentType_TAIWAN_MAINLAND_TRAVEL_PERMIT RepresentativeDocumentType = "TAIWAN_MAINLAND_TRAVEL_PERMIT"
	// RepresentativeDocumentType_HK_MACAO_MAINLAND_TRAVEL_PERMIT：港澳居民来往内地通行证，又称回乡证。
	RepresentativeDocumentType_HK_MACAO_MAINLAND_TRAVEL_PERMIT RepresentativeDocumentType = "HK_MACAO_MAINLAND_TRAVEL_PERMIT"
)
