// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func FolderPlusIcon() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-5 h-5\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M3.75 3C2.7835 3 2 3.7835 2 4.75V15.25C2 16.2165 2.7835 17 3.75 17H16.25C17.2165 17 18 16.2165 18 15.25V6.75C18 5.7835 17.2165 5 16.25 5H11.4142C11.3479 5 11.2843 4.97366 11.2374 4.92678L9.82322 3.51256C9.49503 3.18437 9.04992 3 8.58579 3H3.75ZM10 8C10.4142 8 10.75 8.33579 10.75 8.75V10.25H12.25C12.6642 10.25 13 10.5858 13 11C13 11.4142 12.6642 11.75 12.25 11.75H10.75V13.25C10.75 13.6642 10.4142 14 10 14C9.58579 14 9.25 13.6642 9.25 13.25V11.75H7.75C7.33579 11.75 7 11.4142 7 11C7 10.5858 7.33579 10.25 7.75 10.25H9.25V8.75C9.25 8.33579 9.58579 8 10 8Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
