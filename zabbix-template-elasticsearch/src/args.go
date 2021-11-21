/**
 * @Author: Einic <einicyeo AT gmail.com>
 * @Version: 1.0.0
 * @Date: 2018/11/21 17:05
 * @BLOG:  https://www.infvie.com
 * @Description:
 */
package main

import "os"

func obtainCAPath() string {
	caPath := os.Getenv("ZBX_ES_CA_PATH")

	if len(caPath) == 0 {
		return "None"
	}

	return caPath
}

func obtainESDSN() string {
	dsn := os.Getenv("ZBX_ES_DSN")

	if len(dsn) == 0 {
		return "http://127.0.0.1:9200"
	}

	return dsn
}
