// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func GiftIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-5 h-5\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M14.0002 6C14.314 5.58217 14.5 5.0628 14.5 4.5C14.5 3.11929 13.3807 2 12 2C11.1822 2 10.4561 2.39267 10 2.99976C9.54389 2.39267 8.8178 2 8 2C6.61929 2 5.5 3.11929 5.5 4.5C5.5 5.0628 5.68597 5.58217 5.99982 6H3.25C2.55964 6 2 6.55964 2 7.25V7.75C2 8.44036 2.55964 9 3.25 9H9.25V6H10.75V9H16.75C17.4404 9 18 8.44036 18 7.75V7.25C18 6.55964 17.4404 6 16.75 6H14.0002ZM13 4.5C13 5.05228 12.5523 5.5 12 5.5H11L11 4.5C11 3.94772 11.4477 3.5 12 3.5C12.5523 3.5 13 3.94772 13 4.5ZM7 4.5C7 5.05228 7.44772 5.5 8 5.5H9V4.5C9 3.94772 8.55228 3.5 8 3.5C7.44772 3.5 7 3.94772 7 4.5Z\" fill=\"#0F172A\"></path> <path d=\"M9.25 10.5H3V15.25C3 16.7688 4.23122 18 5.75 18H9.25V10.5Z\" fill=\"#0F172A\"></path> <path d=\"M10.75 18V10.5H17V15.25C17 16.7688 15.7688 18 14.25 18H10.75Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
