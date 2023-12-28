// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func EyeDropperIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M16.0982 2.59835C17.5627 1.13388 19.9371 1.13388 21.4015 2.59835C22.866 4.06282 22.866 6.43718 21.4015 7.90165C20.9192 8.38401 20.3362 8.7081 19.7198 8.87264L18 9.33348V12C18 12.1989 17.921 12.3897 17.7803 12.5303L17.0303 13.2803C16.7374 13.5732 16.2626 13.5732 15.9697 13.2803L15 12.3107L7.06066 20.25C6.58044 20.7302 5.92913 21 5.25 21C4.9687 21 4.69891 21.1117 4.5 21.3107L3.53033 22.2803C3.23744 22.5732 2.76256 22.5732 2.46967 22.2803L1.71967 21.5303C1.42678 21.2374 1.42678 20.7626 1.71967 20.4697L2.68934 19.5C2.88825 19.3011 3 19.0313 3 18.75C3 18.0709 3.26978 17.4196 3.75 16.9393L11.6893 8.99999L10.7197 8.03031C10.4268 7.73742 10.4268 7.26255 10.7197 6.96966L11.4697 6.21966C11.6103 6.079 11.8011 5.99998 12 5.99998H14.6661L15.127 4.27976C15.2917 3.66322 15.6161 3.08048 16.0982 2.59835ZM12.75 10.0607L4.81066 18C4.61175 18.1989 4.5 18.4687 4.5 18.75C4.5 19.0682 4.44077 19.3803 4.32841 19.6716C4.61967 19.5592 4.93178 19.5 5.25 19.5C5.5313 19.5 5.80109 19.3883 6 19.1893L13.9393 11.25L12.75 10.0607Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
