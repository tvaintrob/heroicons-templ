// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CloudArrowDownIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M4.5 13C2.567 13 1 11.433 1 9.5C1 8.06836 1.85956 6.83748 3.09075 6.29529C3.03138 6.03979 3 5.77356 3 5.5C3 3.567 4.567 2 6.5 2C7.94462 2 9.18482 2.87521 9.71929 4.12432C9.96493 4.04364 10.2274 4 10.5 4C11.8807 4 13 5.11929 13 6.5C13 6.72217 12.971 6.93756 12.9166 7.14262C14.1252 7.53 15 8.66283 15 10C15 11.6569 13.6569 13 12 13H4.5ZM10.7803 9.03033C11.0732 8.73744 11.0732 8.26256 10.7803 7.96967C10.4874 7.67678 10.0126 7.67678 9.71967 7.96967L8.75 8.93934L8.75 6.25C8.75 5.83579 8.41421 5.5 8 5.5C7.58579 5.5 7.25 5.83579 7.25 6.25L7.25 8.93934L6.28033 7.96967C5.98744 7.67678 5.51256 7.67678 5.21967 7.96967C4.92678 8.26256 4.92678 8.73744 5.21967 9.03033L7.46967 11.2803C7.76256 11.5732 8.23744 11.5732 8.53033 11.2803L10.7803 9.03033Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
