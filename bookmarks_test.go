package bookmarks

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

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

func Test_parse(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []Bookmark
		wantErr bool
	}{
		{
			name: "Parses sample file",
			args: args{
				r: bytes.NewBufferString(sampleFile),
			},
			want: []Bookmark{
				{
					Name:    "Example",
					Href:    "https://www.example.com/",
					AddDate: "1381998752",
					Attributes: []Attribute{
						{"extra", "cool"},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
