package engines

import (
	"carimurah/entity"
	"carimurah/postgres"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

func (e *Engine) DoSearchQuery(keyword string) {

	url := "https://gql.tokopedia.com/graphql/SearchProductQueryV4"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`[
		{
			"operationName": "SearchProductQueryV4",
			"variables": {
				"params": "device=desktop&navsource=home&navsource=home&ob=23&page=1&q=%s&related=true&rows=60&safe_search=false&scheme=https&shipping=&source=search&srp_component_id=02.01.00.00&srp_page_id=&srp_page_title=&st=product&start=0&topads_bucket=true&unique_id=934f945f07ffa2ad12ec751ad04bf5cf&user_addressId=3553458&user_cityId=253&user_districtId=3546&user_id=2733931&user_lat=-7.8903395&user_long=110.3280655&user_postCode=55711&user_warehouseId=0&variants="
			},
			"query": "query SearchProductQueryV4($params: String!) { ace_search_product_v4(params: $params) { data { products {  id  name originalPrice imageUrl price priceRange rating ratingAverage url } } } }"
		}
	]`, keyword))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", "_abck=81EF3A7EDB46C546A37999B44766F156~-1~YAAQPuUcuMw9k2GCAQAAYnNFZAhm5zYR7LVm1g2QYbhw81gmc7zCI8U2Rq9A6nvxNYXX70uPrrqjwIoKuxBhkEoAKMbNHibgXcHy5t69LJ0Qmx6cdVNYEe7peDkBHZ3NCFpeMgmY0HtwWcuZA7Vg976LofzUge0KloFCw2Euic1qitfEhT6qIDz78SSWlGvpRpOLP42+jBecOeHlNFdCrIQzhtQbQWP6hxcgsrocxF77ZXbE0R/DTeq/YOMZ4Ih6XlYi0uuKH+1OUGadTlnVrfydYwb6G0oPj8/Q/0pSr3/rTSJ2xWbx0eHoaHUbSKh9tpPHeA9jrn75OjHPqt1gb8BPwyFpqKSdn6+qGYMOLiWnFGQDZL3GpR7jsNA=~-1~-1~-1; bm_sz=6758DA9A8677EAF00240EF507D7601D9~YAAQPuUcuM09k2GCAQAAYnNFZBDUhM3PkLxLWXLXHJ/bCsYCAl7HtG5W0LuTpKH6hwPo9ctwI5G+NaWjIGBQR1v9luEwcyrddTlPuo+vLu4WGRK8qL6XOLNkXhbTyKw1YbNOXwbZob2wpWiq0yeBdw3OWEZz9i11Qsyq9dnmqc1dSyXN/X7mGWs6ypmuVtCzDPSd9qPGfdh7tFbAppqivi4DfTKoAT27BRvNOknet2Jg3nQ8yxNy5IT97lWZh37Vkwi3IRGtccQkxX7NmHULXUUNuBuBEhTIw15EjrZOBiOKSFNFhU8=~3289145~4340038")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var response []entity.TokopediaResponseParent

	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, item := range response[0].Data.AceSearchProductV4.Data.Products {
		// Convert string to float64
		ratingAvg, err := strconv.ParseFloat(item.RatingAverage, 2)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if item.Price == "" || item.OriginalPrice == "" {
			continue
		}

		// Remove unused character from Price key
		reg := regexp.MustCompile("[^0-9]+")
		price := reg.ReplaceAllString(item.Price, "")
		originalPrice := reg.ReplaceAllString(item.OriginalPrice, "")

		product := postgres.Product{
			ID:            uuid.New(),
			ExtID:         fmt.Sprint(item.ID),
			Name:          item.Name,
			Price:         price,
			ImageUrl:      item.ImageUrl,
			Url:           item.Url,
			CreatedAt:     time.Now(),
			OriginalPrice: originalPrice,
			Rating:        item.Rating,
			RatingAverage: ratingAvg,
		}

		err = e.db.InsertItemToProduct(product)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

}
