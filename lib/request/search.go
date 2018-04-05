package request

/*
{
  "project" : "Vagrant",
  "environment" : "live",
  "fields" : [
    {
      "name" : "date",
      "start" : "2018-04-03",
      "end" : "2018-04-03"
    },
    {
      "name" : "user",
      "value" : "test"
    }
  ]
}
 */

type Search struct {
	Project string
	Environment string
	Fields []struct{
		Name string
		Value string
		Start string
		End string
	}
}