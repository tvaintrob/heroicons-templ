// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M11.8896 4.11116C9.74169 1.96327 6.25928 1.96327 4.1114 4.11116C1.96352 6.25904 1.96352 9.74145 4.1114 11.8893C6.25928 14.0372 9.74169 14.0372 11.8896 11.8893C12.1825 11.5964 12.6573 11.5964 12.9502 11.8893C13.2431 12.1822 13.2431 12.6571 12.9502 12.95C10.2166 15.6837 5.78441 15.6837 3.05074 12.95C0.31707 10.2163 0.317071 5.78417 3.05074 3.0505C5.78441 0.316827 10.2166 0.316827 12.9502 3.0505C14.3169 4.41719 15.0005 6.20997 15.0005 8.00028C15.0005 9.38147 13.8802 10.4999 12.5 10.4999C11.8993 10.4999 11.348 10.288 10.9169 9.93489C10.2899 10.8782 9.21758 11.4999 8.00002 11.4999C6.06702 11.4999 4.50002 9.93288 4.50002 7.99989C4.50002 6.06689 6.06702 4.49989 8.00002 4.49989C9.93301 4.49989 11.5 6.06689 11.5 7.99989C11.5 8.55217 11.9477 8.99989 12.5 8.99989C13.0528 8.99989 13.5005 8.55208 13.5005 8.00027C13.5005 6.59153 12.9637 5.18525 11.8896 4.11116ZM10 7.99989C10 6.89532 9.10459 5.99989 8.00002 5.99989C6.89545 5.99989 6.00002 6.89532 6.00002 7.99989C6.00002 9.10445 6.89545 9.99989 8.00002 9.99989C9.10459 9.99989 10 9.10445 10 7.99989Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
