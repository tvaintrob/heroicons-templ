// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package outline24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BuildingOffice2Icon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-6 h-6\"><path d=\"M2.25 21H21.75M3.75 3V21M14.25 3V21M20.25 7.5V21M6.75 6.75H7.5M6.75 9.75H7.5M6.75 12.75H7.5M10.5 6.75H11.25M10.5 9.75H11.25M10.5 12.75H11.25M6.75 21V17.625C6.75 17.0037 7.25368 16.5 7.875 16.5H10.125C10.7463 16.5 11.25 17.0037 11.25 17.625V21M3 3H15M14.25 7.5H21M17.25 11.25H17.2575V11.2575H17.25V11.25ZM17.25 14.25H17.2575V14.2575H17.25V14.25ZM17.25 17.25H17.2575V17.2575H17.25V17.25Z\" stroke=\"#0F172A\" stroke-width=\"1.5\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
