// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func AtSymbolIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M5.40405 14.5962C2.86564 12.0578 2.86564 7.94221 5.40405 5.40381C7.94246 2.8654 12.058 2.8654 14.5964 5.40381C15.8658 6.67316 16.5002 8.33534 16.5002 10C16.5002 10.6904 15.9404 11.25 15.25 11.25C14.5597 11.25 14 10.6904 14 10C14 7.79086 12.2092 6 10 6C7.79087 6 6.00001 7.79086 6.00001 10C6.00001 12.2091 7.79087 14 10 14C11.4554 14 12.7292 13.2228 13.429 12.0607C13.9141 12.4897 14.5516 12.75 15.25 12.75C16.7614 12.75 17.9881 11.5307 17.9999 10.022C18.0001 10.0147 18.0002 10.0074 18.0002 10C18.0002 7.95378 17.219 5.9051 15.6571 4.34315C12.5329 1.21895 7.46758 1.21895 4.34339 4.34315C1.2192 7.46734 1.2192 12.5327 4.34339 15.6569C7.46758 18.781 12.5329 18.781 15.6571 15.6569C15.95 15.364 15.95 14.8891 15.6571 14.5962C15.3642 14.3033 14.8893 14.3033 14.5964 14.5962C12.058 17.1346 7.94246 17.1346 5.40405 14.5962ZM10 7.5C8.6193 7.5 7.50001 8.61929 7.50001 10C7.50001 11.3807 8.6193 12.5 10 12.5C11.3807 12.5 12.5 11.3807 12.5 10C12.5 8.61929 11.3807 7.5 10 7.5Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
