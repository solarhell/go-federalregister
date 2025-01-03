package federalregister

import (
	"context"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/davecgh/go-spew/spew"
	"strconv"
)

func (c *Client) GetFederalRegisterDocuments(ctx context.Context, pageNumber uint64, pageSize uint64, agencySlugs []string) (*GetFederalRegisterDocumentsResponse, error) {
	link := "https://www.federalregister.gov/api/v1/documents.json"

	if pageNumber < 1 {
		pageNumber = 1
	}

	if pageSize > 1000 {
		pageSize = 1000
	}

	queryParams := map[string]string{
		"order":    "newest",
		"page":     strconv.FormatUint(pageNumber, 10),
		"per_page": strconv.FormatUint(pageSize, 10),
	}

	req := c.httpClient.R().SetContext(ctx).SetQueryParams(queryParams)

	if len(agencySlugs) > 0 {
		for _, slug := range agencySlugs {
			req.AddQueryParam("conditions[agencies][]", slug)
		}
	}

	resp, err := req.Get(link)
	if err != nil {
		return nil, err
	}

	body := resp.String()

	res := new(GetFederalRegisterDocumentsResponse)
	err = sonic.UnmarshalString(body, res)
	if err != nil {
		return nil, err
	}

	if res.Errors != nil {
		return nil, fmt.Errorf("请求出错: %s", spew.Sdump(res.Errors))
	}

	return res, nil
}

type GetFederalRegisterDocumentsResponse struct {
	Errors      map[string]string `json:"errors,omitempty"`
	Count       int               `json:"count"`
	Description string            `json:"description"`
	TotalPages  int               `json:"total_pages"`
	NextPageUrl string            `json:"next_page_url"`
	Results     []struct {
		Title                  string `json:"title"`
		Type                   string `json:"type"`
		Abstract               string `json:"abstract"`
		DocumentNumber         string `json:"document_number"`
		HtmlUrl                string `json:"html_url"`
		PdfUrl                 string `json:"pdf_url"`
		PublicInspectionPdfUrl string `json:"public_inspection_pdf_url"`
		PublicationDate        string `json:"publication_date"`
		Agencies               []struct {
			RawName  string `json:"raw_name"`
			Name     string `json:"name,omitempty"`
			Id       int    `json:"id,omitempty"`
			Url      string `json:"url,omitempty"`
			JsonUrl  string `json:"json_url,omitempty"`
			ParentId *int   `json:"parent_id,omitempty"`
			Slug     string `json:"slug,omitempty"`
		} `json:"agencies"`
		Excerpts string `json:"excerpts"`
	} `json:"results"`
}
