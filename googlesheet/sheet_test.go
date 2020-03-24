package googlesheet

import (
	"testing"
)

func TestConfiguration_Validate(t *testing.T) {
	type fields struct {
		GoogleCredentialPath string
		SpreadsheetID        string
		NoSkipHeader         int
		SheetName            string
		Columns              string
	}
	tests := []struct {
		name    string
		fields  fields
		want    bool
		wantErr bool
	}{
		{
			name: "Missing GoogleCredentialPath",
			fields: fields{
				SpreadsheetID: "ABC",
				NoSkipHeader:  5,
				SheetName:     "Sheet1",
				Columns:       "area_detail.postal_city,housing_form,price,,,,number_of_rooms,living_area,,borattavgift,driftkostnad,construction_year,street_name",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Missing SpreadsheetID",
			fields: fields{
				GoogleCredentialPath: "XYZ",
				NoSkipHeader:         5,
				SheetName:            "Sheet1",
				Columns:              "area_detail.postal_city,housing_form,price,,,,number_of_rooms,living_area,,borattavgift,driftkostnad,construction_year,street_name",
			},
			want:    false,
			wantErr: true,
		},
		{
			name: "Valid",
			fields: fields{
				GoogleCredentialPath: "XYZ",
				SpreadsheetID:        "ABC",
				NoSkipHeader:         5,
				SheetName:            "Sheet1",
				Columns:              "area_detail.postal_city,housing_form,price,,,,number_of_rooms,living_area,,borattavgift,driftkostnad,construction_year,street_name",
			},
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Configuration{
				GoogleCredentialPath: tt.fields.GoogleCredentialPath,
				SpreadsheetID:        tt.fields.SpreadsheetID,
				NoSkipHeader:         tt.fields.NoSkipHeader,
				SheetName:            tt.fields.SheetName,
				Columns:              tt.fields.Columns,
			}
			got, err := c.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Configuration.Validate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Configuration.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConfiguration_Default(t *testing.T) {
	type fields struct {
		GoogleCredentialPath string
		SpreadsheetID        string
		NoSkipHeader         int
		SheetName            string
		Columns              string
	}
	type wants struct {
		SheetName string
		Columns   string
	}
	tests := []struct {
		name   string
		fields fields
		wants  wants
	}{
		{
			name: "Has values",
			fields: fields{
				GoogleCredentialPath: "XYZ",
				SpreadsheetID:        "ABC",
				NoSkipHeader:         5,
				SheetName:            "mnn",
				Columns:              "uuu",
			},
			wants: wants{
				SheetName: "mnn",
				Columns:   "uuu",
			},
		},
		{
			name: "Miss Columns",
			fields: fields{
				GoogleCredentialPath: "XYZ",
				SpreadsheetID:        "ABC",
				NoSkipHeader:         5,
				SheetName:            "mnn",
			},
			wants: wants{
				SheetName: "mnn",
				Columns:   DefaultColumns,
			},
		},
		{
			name: "Miss SheetName",
			fields: fields{
				GoogleCredentialPath: "XYZ",
				SpreadsheetID:        "ABC",
				NoSkipHeader:         5,
				Columns:              "uuu",
			},
			wants: wants{
				SheetName: DefaultSheetName,
				Columns:   "uuu",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Configuration{
				GoogleCredentialPath: tt.fields.GoogleCredentialPath,
				SpreadsheetID:        tt.fields.SpreadsheetID,
				NoSkipHeader:         tt.fields.NoSkipHeader,
				SheetName:            tt.fields.SheetName,
				Columns:              tt.fields.Columns,
			}
			c.Default()
			if c.SheetName != tt.wants.SheetName {
				t.Errorf("Configuration.Validate().SheetName = %v, want %v", c.SheetName, tt.wants.SheetName)
			}
			if c.Columns != tt.wants.Columns {
				t.Errorf("Configuration.Validate().Columns = %v, want %v", c.Columns, tt.wants.Columns)
			}
		})
	}
}
