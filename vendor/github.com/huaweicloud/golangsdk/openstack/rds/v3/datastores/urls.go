package datastores

import "github.com/huaweicloud/golangsdk"

func listURL(sc *golangsdk.ServiceClient, databasename string) string {
	return sc.ServiceURL("datastores", databasename)
}
