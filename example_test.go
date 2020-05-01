package bookmarks_test

import (
	"bytes"
	"fmt"

	"github.com/aquilax/bookmarks"
)

func ExampleBookmarks_Parse() {
	const sampleFile = `<!DOCTYPE NETSCAPE-Bookmark-file-1>
	<!-- This is an automatically generated file.
		 It will be read and overwritten.
		 DO NOT EDIT! -->
	<META HTTP-EQUIV="Content-Type" CONTENT="text/html; charset=UTF-8">
	<TITLE>Bookmarks</TITLE>
	<H1>Bookmarks</H1>
	<DL><p>
		<DT><H3 ADD_DATE="1418637929" LAST_MODIFIED="1562120488" PERSONAL_TOOLBAR_FOLDER="true">Bookmarks bar</H3>
		<DL><p>
			<DT><H3 ADD_DATE="1418638085" LAST_MODIFIED="1481039402">games</H3>
			<DL><p>
				<DT><A HREF="https://www.example.com/" ADD_DATE="1381998752" extra="cool">Example</A>
			</DL><p>
		</DL><p>
	</DL><p>
	`
	list, _ := bookmarks.Parse(bytes.NewBufferString(sampleFile))
	for _, b := range list {
		fmt.Println(b.Name)
		fmt.Println(b.Href)
	}
	// Output:
	// Example
	// https://www.example.com/
}
