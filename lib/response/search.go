package response

/*
{
  "result" : [
    {
      "fields" : [
        {
          "name" : "date",
          "value" : "2018-01-02"
        },
        {
          "name" : "user",
          "value" : "root"
        }
      ]
    },
    {
      "fields" : [
        {
          "name" : "date",
          "value" : "2018-01-02"
        },
        {
          "name" : "user",
          "value" : "root"
        }
      ]
    }
  ]
}
*/

type Search struct {
	Result []struct {
		Fields []struct {
			Name  string
			Value string
		}
	}
}
