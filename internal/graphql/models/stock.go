package models

import (
	"time"
)

type Strain struct {
	ID                  string    `json:"id"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	CreatedBy           string    `json:"created_by"`
	UpdatedBy           string    `json:"updated_by"`
	Summary             *string   `json:"summary"`
	EditableSummary     *string   `json:"editable_summary"`
	Depositor           *string   `json:"depositor"`
	Genes               []*string `json:"genes"`
	Dbxrefs             []*string `json:"dbxrefs"`
	Publications        []*string `json:"publications"`
	SystematicName      string    `json:"systematic_name"`
	Label               string    `json:"label"`
	Species             string    `json:"species"`
	Plasmid             *string   `json:"plasmid"`
	Parent              *string   `json:"parent"`
	Names               []*string `json:"names"`
	InStock             bool      `json:"in_stock"`
	Phenotypes          []*string `json:"phenotypes"`
	GeneticModification *string   `json:"genetic_modification"`
	MutagenesisMethod   *string   `json:"mutagenesis_method"`
	Characteristics     []*string `json:"characteristics"`
	Genotypes           []*string `json:"genotypes"`
}

type Plasmid struct {
	ID               string    `json:"id"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CreatedBy        string    `json:"created_by"`
	UpdatedBy        string    `json:"updated_by"`
	Summary          *string   `json:"summary"`
	EditableSummary  *string   `json:"editable_summary"`
	Depositor        *string   `json:"depositor"`
	Genes            []*string `json:"genes"`
	Dbxrefs          []*string `json:"dbxrefs"`
	Publications     []*string `json:"publications"`
	ImageMap         *string   `json:"image_map"`
	Sequence         *string   `json:"sequence"`
	Name             string    `json:"name"`
	InStock          bool      `json:"in_stock"`
	Keywords         []*string `json:"keywords"`
	GenbankAccession *string   `json:"genbank_accession"`
}

func (Strain) IsStock()  {}
func (Plasmid) IsStock() {}
