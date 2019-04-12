package zhenai

import "regexp"

const cityListRe = `<a href="(http://www.zhenai.com/zhenhun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

//demo link: https://blog.csdn.net/qq_36183935/article/details/80506175
func ParseCityList(
	contents []byte) engine.ParseResult  {

		re := regexp.MustCompile(cityListRe)
		matches := re.FindAllStringSubmatch(contents, -1)

		result := engine.ParseResult{}

		limit := 10

		for _, m := range matches {
			result.Items = append(result.Items, "City "+string(m[2]))

			result.Requests = append(
				result.Requests,
				engine.Request{
					Url: string(m[1]),
					ParseFunc: ParseCity,
				})

			limit--

			if limit == 0 {
				break
			}
		}

		return result
}
