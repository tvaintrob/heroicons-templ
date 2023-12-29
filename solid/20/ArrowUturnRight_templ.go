// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowUturnRightIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-5 h-5\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M12.2075 2.23214C11.9215 2.53177 11.9325 3.00651 12.2321 3.29252L16.3781 7.25H6.375C3.40647 7.25 1 9.65647 1 12.625C1 15.5935 3.40647 18 6.375 18H9.25C9.66421 18 10 17.6642 10 17.25C10 16.8358 9.66421 16.5 9.25 16.5H6.375C4.2349 16.5 2.5 14.7651 2.5 12.625C2.5 10.4849 4.2349 8.75 6.375 8.75H16.3781L12.2321 12.7075C11.9325 12.9935 11.9215 13.4682 12.2075 13.7679C12.4935 14.0675 12.9682 14.0785 13.2679 13.7925L18.7679 8.54252C18.9161 8.401 19 8.20496 19 8C19 7.79504 18.9161 7.59901 18.7679 7.45748L13.2679 2.20748C12.9682 1.92148 12.4935 1.93252 12.2075 2.23214Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
