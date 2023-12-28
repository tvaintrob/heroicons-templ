// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BarsArrowDownIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2.25 4.5C2.25 4.08579 2.58579 3.75 3 3.75H17.25C17.6642 3.75 18 4.08579 18 4.5C18 4.91421 17.6642 5.25 17.25 5.25H3C2.58579 5.25 2.25 4.91421 2.25 4.5ZM2.25 9C2.25 8.58579 2.58579 8.25 3 8.25H12.75C13.1642 8.25 13.5 8.58579 13.5 9C13.5 9.41421 13.1642 9.75 12.75 9.75H3C2.58579 9.75 2.25 9.41421 2.25 9ZM17.25 8.25C17.6642 8.25 18 8.58579 18 9V19.1893L20.4697 16.7197C20.7626 16.4268 21.2374 16.4268 21.5303 16.7197C21.8232 17.0126 21.8232 17.4874 21.5303 17.7803L17.7803 21.5303C17.4874 21.8232 17.0126 21.8232 16.7197 21.5303L12.9697 17.7803C12.6768 17.4874 12.6768 17.0126 12.9697 16.7197C13.2626 16.4268 13.7374 16.4268 14.0303 16.7197L16.5 19.1893V9C16.5 8.58579 16.8358 8.25 17.25 8.25ZM2.25 13.5C2.25 13.0858 2.58579 12.75 3 12.75H12.75C13.1642 12.75 13.5 13.0858 13.5 13.5C13.5 13.9142 13.1642 14.25 12.75 14.25H3C2.58579 14.25 2.25 13.9142 2.25 13.5Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
