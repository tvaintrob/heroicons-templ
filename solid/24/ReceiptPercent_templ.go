// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ReceiptPercentIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M12 1.5C10.079 1.5 8.18374 1.61114 6.32022 1.82741C4.82283 2.00119 3.75 3.28722 3.75 4.75699V21.75C3.75 21.9989 3.87345 22.2315 4.07953 22.3711C4.28561 22.5106 4.54748 22.5388 4.77854 22.4464L8.25 21.0578L11.7215 22.4464C11.9003 22.5179 12.0997 22.5179 12.2785 22.4464L15.75 21.0578L19.2215 22.4464C19.4525 22.5388 19.7144 22.5106 19.9205 22.3711C20.1266 22.2315 20.25 21.9989 20.25 21.75V4.75699C20.25 3.28722 19.1772 2.00119 17.6798 1.82741C15.8163 1.61114 13.921 1.5 12 1.5ZM15.5303 8.78033C15.8232 8.48744 15.8232 8.01256 15.5303 7.71967C15.2374 7.42678 14.7626 7.42678 14.4697 7.71967L8.46967 13.7197C8.17678 14.0126 8.17678 14.4874 8.46967 14.7803C8.76256 15.0732 9.23744 15.0732 9.53033 14.7803L15.5303 8.78033ZM8.625 9C8.625 8.37868 9.12868 7.875 9.75 7.875C10.3713 7.875 10.875 8.37868 10.875 9C10.875 9.62132 10.3713 10.125 9.75 10.125C9.12868 10.125 8.625 9.62132 8.625 9ZM14.25 12.375C13.6287 12.375 13.125 12.8787 13.125 13.5C13.125 14.1213 13.6287 14.625 14.25 14.625C14.8713 14.625 15.375 14.1213 15.375 13.5C15.375 12.8787 14.8713 12.375 14.25 12.375Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
