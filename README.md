# getir-test-case-MS

# About the case
This test case is about an API. It has 2 endpoints.

One of them, fetches the datas from MongoDB; other one is return it.


Note:

Please clone this repo into your local machine.


### Endpoints:
.....


Sample Input:

    {
    "startDate": "2016-01-21",
    "endDate": "2016-03-02",
    "minCount": 2900,
    "maxCount": 3000
    }

Response;

    {
    "code":0,
    "msg":"Success",
    "records": [
              {
              "key":"TAKwGc6Jr4i8Z487",
              "createdAt":"2017-01-28T01:22:14.398Z",
              "totalCount":2800
              },
              {
              "key":"NAeQ8eX7e5TEg7oH",
              "createdAt":"2017-01-27T08:19:14.135Z",
              "totalCount":2900
              }
                ]
    }




