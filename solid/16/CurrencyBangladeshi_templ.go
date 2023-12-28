// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CurrencyBangladeshiIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M15 8C15 11.866 11.866 15 8 15C4.13401 15 1 11.866 1 8C1 4.13401 4.13401 1 8 1C11.866 1 15 4.13401 15 8ZM5.25 4.70729C4.98617 4.80033 4.68069 4.74135 4.46967 4.53033C4.17678 4.23743 4.17678 3.76256 4.46967 3.46967C5.31117 2.62817 6.75 3.22415 6.75 4.41421V6H12.25C12.6642 6 13 6.33579 13 6.75C13 7.16421 12.6642 7.5 12.25 7.5H6.75V10.5977C6.75 11.1468 7.0454 11.4336 7.29476 11.4685C7.44315 11.4892 7.5951 11.5 7.75 11.5C8.67002 11.5 9.50155 11.1176 10.0933 10.5022H9.75C9.33579 10.5022 9 10.1664 9 9.75222C9 9.338 9.33579 9.00222 9.75 9.00222H11.4583C11.7078 9.00222 11.941 9.12628 12.0804 9.33321C12.2198 9.54014 12.2471 9.80284 12.1534 10.034C11.4486 11.772 9.74362 13 7.75 13C7.52539 13 7.304 12.9844 7.08697 12.954C5.82127 12.777 5.25 11.5784 5.25 10.5977V7.5H3.75C3.33579 7.5 3 7.16421 3 6.75C3 6.33579 3.33579 6 3.75 6H5.25V4.70729Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
