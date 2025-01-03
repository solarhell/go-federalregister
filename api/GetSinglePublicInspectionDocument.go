package api

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"time"
)

func (c *Client) GetSinglePublicInspectionDocument(ctx context.Context, documentID string) (*GetSinglePublicInspectionDocumentResponse, error) {
	link := fmt.Sprintf("https://www.federalregister.gov/api/v1/public-inspection-documents/%s.json", documentID)

	resp, err := c.httpClient.R().SetContext(ctx).
		Get(link)
	if err != nil {
		return nil, err
	}

	body := resp.String()

	res := new(GetSinglePublicInspectionDocumentResponse)
	err = sonic.UnmarshalString(body, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type GetSinglePublicInspectionDocumentResponse struct {
	Agencies []struct {
		RawName  string `json:"raw_name"`
		Name     string `json:"name"`
		Id       int    `json:"id"`
		Url      string `json:"url"`
		JsonUrl  string `json:"json_url"`
		ParentId int    `json:"parent_id"`
		Slug     string `json:"slug"`
	} `json:"agencies"`
	AgencyLetters             []string  `json:"agency_letters"`
	AgencyNames               []string  `json:"agency_names"`
	DocketNumbers             []string  `json:"docket_numbers"`
	DocumentNumber            string    `json:"document_number"`
	EditorialNote             *string   `json:"editorial_note"`
	FiledAt                   time.Time `json:"filed_at"`
	FilingType                string    `json:"filing_type"`
	HtmlUrl                   string    `json:"html_url"`
	LastPublicInspectionIssue string    `json:"last_public_inspection_issue"`
	NumPages                  int       `json:"num_pages"`
	PageViews                 struct {
		Count       int    `json:"count"`
		LastUpdated string `json:"last_updated"`
	} `json:"page_views"`
	PdfFileName     string    `json:"pdf_file_name"`
	PdfFileSize     int       `json:"pdf_file_size"`
	PdfUpdatedAt    time.Time `json:"pdf_updated_at"`
	PdfUrl          string    `json:"pdf_url"`
	PublicationDate string    `json:"publication_date"`
	RawTextUrl      string    `json:"raw_text_url"`
	Subject1        string    `json:"subject_1"`
	Subject2        string    `json:"subject_2"`
	Subject3        *string   `json:"subject_3"`
	Title           string    `json:"title"`
	TocDoc          string    `json:"toc_doc"`
	TocSubject      string    `json:"toc_subject"`
	Type            string    `json:"type"`
}
