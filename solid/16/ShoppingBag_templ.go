// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ShoppingBagIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M5.00001 4C5.00001 2.34315 6.34316 1 8.00001 1C9.65687 1 11 2.34315 11 4V5H11.6425C12.4132 5 13.0584 5.58394 13.1351 6.35074L13.8351 13.3507C13.9234 14.2338 13.23 15 12.3425 15H3.65749C2.77006 15 2.07664 14.2338 2.16494 13.3507L2.86494 6.35075C2.94162 5.58394 3.58687 5 4.35749 5H5.00001V4ZM9.50001 4V5H6.50001V4C6.50001 3.17157 7.17159 2.5 8.00001 2.5C8.82844 2.5 9.50001 3.17157 9.50001 4ZM6.5 7.75C6.5 7.33579 6.16421 7 5.75 7C5.33579 7 5 7.33579 5 7.75V8.75C5 10.4069 6.34315 11.75 8 11.75C9.65685 11.75 11 10.4069 11 8.75V7.75C11 7.33579 10.6642 7 10.25 7C9.83579 7 9.5 7.33579 9.5 7.75V8.75C9.5 9.57843 8.82843 10.25 8 10.25C7.17157 10.25 6.5 9.57843 6.5 8.75V7.75Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
