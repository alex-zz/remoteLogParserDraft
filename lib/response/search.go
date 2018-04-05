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
  ],
  "error" : {
    "code" : 100,
    "title" : "data not found"
  }
}
*/

type Search struct {
	Result []struct {
		Fields []struct {
			Name  string
			Value string
		}
	}
	Error struct{
		Code int
		Title string
	}
}
