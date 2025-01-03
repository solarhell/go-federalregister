package api

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"time"
)

func (c *Client) GetSingleFederalRegisterDocument(ctx context.Context, documentNumber string) (*GetSingleFederalRegisterDocumentResponse, error) {
	link := fmt.Sprintf("https://www.federalregister.gov/api/v1/documents/%s.json", documentNumber)

	resp, err := c.httpClient.R().SetContext(ctx).Get(link)
	if err != nil {
		return nil, err
	}

	body := resp.String()

	res := new(GetSingleFederalRegisterDocumentResponse)
	err = sonic.UnmarshalString(body, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type GetSingleFederalRegisterDocumentResponse struct {
	Abstract string `json:"abstract"`
	Action   string `json:"action"`
	Agencies []struct {
		RawName  string `json:"raw_name"`
		Name     string `json:"name"`
		Id       int    `json:"id"`
		Url      string `json:"url"`
		JsonUrl  string `json:"json_url"`
		ParentId *int   `json:"parent_id"`
		Slug     string `json:"slug"`
	} `json:"agencies"`
	BodyHtmlUrl          string        `json:"body_html_url"`
	CfrReferences        []interface{} `json:"cfr_references"`
	Citation             string        `json:"citation"`
	CommentUrl           interface{}   `json:"comment_url"`
	CommentsCloseOn      string        `json:"comments_close_on"`
	CorrectionOf         interface{}   `json:"correction_of"`
	Corrections          []interface{} `json:"corrections"`
	Dates                string        `json:"dates"`
	DispositionNotes     interface{}   `json:"disposition_notes"`
	DocketIds            []interface{} `json:"docket_ids"`
	Dockets              []interface{} `json:"dockets"`
	DocumentNumber       string        `json:"document_number"`
	EffectiveOn          interface{}   `json:"effective_on"`
	EndPage              int           `json:"end_page"`
	ExecutiveOrderNotes  interface{}   `json:"executive_order_notes"`
	ExecutiveOrderNumber interface{}   `json:"executive_order_number"`
	Explanation          interface{}   `json:"explanation"`
	FullTextXmlUrl       string        `json:"full_text_xml_url"`
	HtmlUrl              string        `json:"html_url"`
	Images               struct {
	} `json:"images"`
	ImagesMetadata struct {
	} `json:"images_metadata"`
	JsonUrl                   string      `json:"json_url"`
	ModsUrl                   string      `json:"mods_url"`
	NotReceivedForPublication interface{} `json:"not_received_for_publication"`
	PageLength                int         `json:"page_length"`
	PageViews                 struct {
		Count       int    `json:"count"`
		LastUpdated string `json:"last_updated"`
	} `json:"page_views"`
	PdfUrl                     string      `json:"pdf_url"`
	PresidentialDocumentNumber interface{} `json:"presidential_document_number"`
	ProclamationNumber         interface{} `json:"proclamation_number"`
	PublicInspectionPdfUrl     string      `json:"public_inspection_pdf_url"`
	PublicationDate            string      `json:"publication_date"`
	RawTextUrl                 string      `json:"raw_text_url"`
	RegulationIdNumberInfo     struct {
	} `json:"regulation_id_number_info"`
	RegulationIdNumbers   []string `json:"regulation_id_numbers"`
	RegulationsDotGovInfo struct {
		CheckedRegulationsdotgovAt time.Time `json:"checked_regulationsdotgov_at"`
	} `json:"regulations_dot_gov_info"`
	RegulationsDotGovUrl *string     `json:"regulations_dot_gov_url"`
	Significant          interface{} `json:"significant"`
	SigningDate          *string     `json:"signing_date"`
	StartPage            int         `json:"start_page"`
	Subtype              *string     `json:"subtype"`
	Title                string      `json:"title"`
	TocDoc               string      `json:"toc_doc"`
	TocSubject           *string     `json:"toc_subject"`
	Topics               []string    `json:"topics"`
	Type                 string      `json:"type"`
	Volume               int         `json:"volume"`
}
