package api

import (
	"context"
	"time"

	"github.com/bytedance/sonic"
)

func (c *Client) GetPublicInspectionDocuments(ctx context.Context, availableOn time.Time) (*GetPublicInspectionDocumentsResponse, error) {
	link := "https://www.federalregister.gov/api/v1/public-inspection-documents.json"

	queryParams := map[string]string{
		"conditions[available_on]": availableOn.Format("2006-01-02"),
	}

	resp, err := c.httpClient.R().SetContext(ctx).
		SetQueryParams(queryParams).
		Get(link)
	if err != nil {
		return nil, err
	}

	body := resp.String()

	res := new(GetPublicInspectionDocumentsResponse)
	err = sonic.UnmarshalString(body, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type GetPublicInspectionDocumentsResponse struct {
	Count   int `json:"count"`
	Results []struct {
		Agencies []struct {
			RawName  string `json:"raw_name"`
			Name     string `json:"name"`
			Id       int    `json:"id"`
			Url      string `json:"url"`
			JsonUrl  string `json:"json_url"`
			ParentId *int   `json:"parent_id"`
			Slug     string `json:"slug"`
		} `json:"agencies"`
		AgencyLetters             []interface{} `json:"agency_letters"`
		AgencyNames               []string      `json:"agency_names"`
		DocketNumbers             []interface{} `json:"docket_numbers"`
		DocumentNumber            string        `json:"document_number"`
		EditorialNote             interface{}   `json:"editorial_note"`
		Excerpts                  interface{}   `json:"excerpts"`
		FiledAt                   time.Time     `json:"filed_at"`
		FilingType                string        `json:"filing_type"`
		HtmlUrl                   string        `json:"html_url"`
		JsonUrl                   string        `json:"json_url"`
		LastPublicInspectionIssue string        `json:"last_public_inspection_issue"`
		NumPages                  int           `json:"num_pages"`
		PageViews                 struct {
			Count       int    `json:"count"`
			LastUpdated string `json:"last_updated"`
		} `json:"page_views"`
		PdfFileName     string      `json:"pdf_file_name"`
		PdfFileSize     int         `json:"pdf_file_size"`
		PdfUpdatedAt    time.Time   `json:"pdf_updated_at"`
		PdfUrl          string      `json:"pdf_url"`
		PublicationDate string      `json:"publication_date"`
		RawTextUrl      string      `json:"raw_text_url"`
		Subject1        string      `json:"subject_1"`
		Subject2        *string     `json:"subject_2"`
		Subject3        interface{} `json:"subject_3"`
		Title           string      `json:"title"`
		TocDoc          string      `json:"toc_doc"`
		TocSubject      string      `json:"toc_subject"`
		Type            string      `json:"type"`
	} `json:"results"`
	SpecialFilingsUpdatedAt time.Time `json:"special_filings_updated_at"`
	RegularFilingsUpdatedAt time.Time `json:"regular_filings_updated_at"`
}
