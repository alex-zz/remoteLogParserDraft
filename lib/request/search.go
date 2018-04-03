package request

/*
{
  "project" : "GN",
  "environment" : "live",
  "criteria" : [
    {
      "field" : "date",
      "start" : "2018-04-03",
      "end" : "2018-04-03"
    },
    {
      "field" : "user",
      "value" : "test"
    }
  ]
}
 */

type Search struct {
	Project string
	Environment string
	Fields []struct{
		Field string
		Value string
		Start string
		End string
	}
}