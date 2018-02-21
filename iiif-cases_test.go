package gollico

type StructTest struct {
	ark       string
	expected  IIIFDoc
	expectErr bool
}

var iDocTest = IIIFDoc{
	ID:          "http://gallica.bnf.fr/iiif/ark:/12148/btv1b531610266/manifest.json",
	Label:       "EI-13 (1314)",
	Attribution: "Bibliothèque nationale de France",
	License:     "http://gallica.bnf.fr/html/und/conditions-dutilisation-des-contenus-de-gallica",
	Related:     "http://gallica.bnf.fr/ark:/12148/btv1b531610266",
	Description: "13/4/26, village indou [sic] au Jardin d'acclimatation [Paris, charmeurs de serpents] : [photographie de presse] / [Agence Rol]",
	IIIFMetadata: IIIFMetadata{
		Repository:     "Bibliothèque nationale de France",
		Provider:       "Bibliothèque nationale de France",
		Disseminator:   "Gallica",
		SourceImages:   "http://gallica.bnf.fr/ark:/12148/btv1b531610266",
		MetadataSource: "http://oai.bnf.fr/oai2/OAIHandler?verb=GetRecord&metadataPrefix=oai_dc&identifier=oai:bnf.fr:gallica/ark:/12148/btv1b531610266",
		Shelfmark:      "Bibliothèque nationale de France, département Estampes et photographie, EI-13 (1314)",
		Title:          "13/4/26, village indou [sic] au Jardin d'acclimatation [Paris, charmeurs de serpents] : [photographie de presse] / [Agence Rol]",
		Date:           "1926",
		Language:       []string{"fre", "français"},
		Format:         []string{"1 photogr. nég. sur verre ; 13 x 18 cm (sup.)", "image/jpeg", "Nombre total de vues :  1"},
		Creator:        "Agence Rol. Agence photographique",
		Relation:       "Notice du catalogue : http://catalogue.bnf.fr/ark:/12148/cb453946454",
		Type: map[string][]string{
			"eng": {"image", "still image", "photograph"},
			"fre": {"image fixe", "photographie"},
		},
	},
	Images: []Image{
		Image{
			ID:        "http://gallica.bnf.fr/iiif/ark:/12148/btv1b531610266/canvas/f1",
			Label:     "NP",
			Height:    6145.0,
			Width:     8500.0,
			Thumbnail: "http://gallica.bnf.fr/ark:/12148/btv1b531610266/f1.thumbnail",
		},
	},
	Thumbnail: "http://gallica.bnf.fr/ark:/12148/btv1b531610266.thumbnail",
}

var ManifestStructTest = []StructTest{
	{
		ark:       "",
		expected:  IIIFDoc{},
		expectErr: true,
	},
	{
		ark:       "btv1b531610266", // "btv1b550076223",
		expected:  iDocTest,
		expectErr: false,
	},
}
