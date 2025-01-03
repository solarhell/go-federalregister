package federalregister

import (
	"context"
	"github.com/bytedance/sonic"
)

// GetAllAgencies
func (c *Client) GetAllAgencies(ctx context.Context) ([]APIAgencyItem, error) {
	link := "https://www.federalregister.gov/api/v1/agencies"

	resp, err := c.httpClient.R().SetContext(ctx).Get(link)
	if err != nil {
		return nil, err
	}

	body := resp.String()

	res := make([]APIAgencyItem, 0)
	err = sonic.UnmarshalString(body, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type APIAgencyItem struct {
	AgencyUrl   string   `json:"agency_url"`
	ChildIds    []int    `json:"child_ids"`
	ChildSlugs  []string `json:"child_slugs"`
	Description string   `json:"description"`
	Id          int      `json:"id"`
	Logo        *struct {
		ThumbUrl  string `json:"thumb_url"`
		SmallUrl  string `json:"small_url"`
		MediumUrl string `json:"medium_url"`
	} `json:"logo"`
	Name              string `json:"name"`
	ParentId          *int   `json:"parent_id"`
	RecentArticlesUrl string `json:"recent_articles_url"`
	ShortName         string `json:"short_name"`
	Slug              string `json:"slug"`
	Url               string `json:"url"`
	JsonUrl           string `json:"json_url"`
}

type TopAgency struct {
	AgencyUrl   string `json:"agency_url"`
	Description string `json:"description"`
	Id          int    `json:"id"`
	Logo        *struct {
		ThumbUrl  string `json:"thumb_url"`
		SmallUrl  string `json:"small_url"`
		MediumUrl string `json:"medium_url"`
	} `json:"logo"`
	Name              string   `json:"name"`
	RecentArticlesUrl string   `json:"recent_articles_url"`
	ShortName         string   `json:"short_name"`
	Slug              string   `json:"slug"`
	Url               string   `json:"url"`
	JsonUrl           string   `json:"json_url"`
	Children          []Agency `json:"children"`
}

type Agency struct {
	AgencyUrl   string `json:"agency_url"`
	Description string `json:"description"`
	Id          int    `json:"id"`
	Logo        *struct {
		ThumbUrl  string `json:"thumb_url"`
		SmallUrl  string `json:"small_url"`
		MediumUrl string `json:"medium_url"`
	} `json:"logo"`
	Name              string `json:"name"`
	ParentId          *int   `json:"parent_id"`
	RecentArticlesUrl string `json:"recent_articles_url"`
	ShortName         string `json:"short_name"`
	Slug              string `json:"slug"`
	Url               string `json:"url"`
	JsonUrl           string `json:"json_url"`
}

func (c *Client) GetAllAgenciesWithChild(ctx context.Context) ([]TopAgency, error) {
	// 获取所有机构
	items, err := c.GetAllAgencies(ctx)
	if err != nil {
		return nil, err
	}

	// 创建一个map用于快速查找
	agencyMap := make(map[int]APIAgencyItem)
	for _, item := range items {
		agencyMap[item.Id] = item
	}

	// 创建结果slice，只存储顶级机构
	var result []TopAgency

	// 遍历所有机构，只处理顶级机构（ParentId为nil的）
	for _, item := range items {
		if item.ParentId == nil {
			topAgency := TopAgency{
				AgencyUrl:         item.AgencyUrl,
				Description:       item.Description,
				Id:                item.Id,
				Logo:              item.Logo,
				Name:              item.Name,
				RecentArticlesUrl: item.RecentArticlesUrl,
				ShortName:         item.ShortName,
				Slug:              item.Slug,
				Url:               item.Url,
				JsonUrl:           item.JsonUrl,
				Children:          make([]Agency, 0),
			}

			// 查找并添加子机构
			for _, childId := range item.ChildIds {
				if childItem, ok := agencyMap[childId]; ok {
					childAgency := Agency{
						AgencyUrl:         childItem.AgencyUrl,
						Description:       childItem.Description,
						Id:                childItem.Id,
						Logo:              childItem.Logo,
						Name:              childItem.Name,
						ParentId:          childItem.ParentId,
						RecentArticlesUrl: childItem.RecentArticlesUrl,
						ShortName:         childItem.ShortName,
						Slug:              childItem.Slug,
						Url:               childItem.Url,
						JsonUrl:           childItem.JsonUrl,
					}
					topAgency.Children = append(topAgency.Children, childAgency)
				}
			}

			result = append(result, topAgency)
		}
	}

	return result, nil
}
