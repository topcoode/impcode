package main

import (
	"fmt"
	"net/url"
)

// creating a fictitious (fake)URL
const myurl string = "https://www.mongodb.com/docs/manual/reference/operator/aggregation/"

func main() {
	//creating URls=========================================handling URL
	// net package

	// fmt.Println(myurl)
	// parsing the URL
	// /*In parsing, code is taken from the preprocessor,
	//  broken into smaller pieces and analyzed so other software can understand it.
	result, _ := url.Parse(myurl) //................
	fmt.Println(result.Scheme)    //....................
	// https
	fmt.Println(result.Host) //........................
	// www.google.com
	fmt.Println(result.Path) //.....................
	// /search
	fmt.Println(result.Port()) //............................
	// q=fictitious+meaning&rlz=1C1JJTC_enIN1024IN1024&ei=QaGIY_CqK_vZseMP7_m-cA&oq=fictitious+&
	// fmt.Println(result.RawQuery)//...........................
	/*gs_lcp=Cgxnd3Mtd2l6LXNlcnAQAxgBMggIABCABBCxAzIICAAQgAQQsQMyCAgAEIAEELEDMgUIABCABDIFCAAQgAQyBQgAEIAEMgUIABCABDIFCAAQgAQyBQgAEIAEMgUIABCABDoMCAAQ6gIQtAIQQxgBOhYIABDqAhC0AhCKAxC3AxDUAxDlAhgBOhMIABCPARDqAhC0AhCMAxDlAhgCOhMILhCPARDqAhC0AhCMAxDlAhgCOgQIABBDOgoIABCxAxCDARBDOggILhCDARCxAzoLCAAQgAQQsQMQgwE6CAgAELEDEIMBOgoILhDHARDRAxBDOgUIABCRAjoLCC4QgwEQsQMQgAQ6DgguEIAEEMcBENEDENQCOggILhCABBCxAzoFCC4QsQM6DQguEIAEEMcBENEDEAo6EAguEIAEELEDEIMBELEDEAo6CgguEIAEENQCEAo6CggAEIAEELEDEAo6BwgAEIAEEAo6DQgAEIAEELEDEIMBEAo6BwguEIAEEA06BwgAEIAEEA06CQgAEIAEEA0QCjoKCAAQgAQQsQMQDToNCAAQgAQQsQMQRhD5AUoECEEYAEoECEYYAVAAWPlGYKFaaAFwAXgAgAHXAYgB-RKSAQYwLjE0LjKYAQCgAQGwARTAAQHaAQQIARgH2gEGCAIQARgK&sclient=gws-wiz-serp*/
	name := result.Query()
	fmt.Println(name)
	fmt.Println(name["manual"])
	for _, value := range name {
		fmt.Println(value)
	}
	// &
	// creating URL by using key:value pair
	gange := &url.URL{
		// without using & output is {https   lco.dev /tutcss user=hitesh false false   }
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	// anotherURl := gange.String()
	fmt.Println(gange)
	// https://lco.dev/tutcss

}
