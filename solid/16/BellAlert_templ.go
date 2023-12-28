// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BellAlertIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M3.5996 1.69988C3.84841 1.36872 3.78165 0.89856 3.45048 0.649752C3.11932 0.400945 2.64917 0.467705 2.40036 0.798866C1.89988 1.465 1.51511 2.22396 1.27662 3.04572C1.16117 3.44352 1.39006 3.85959 1.78786 3.97504C2.18566 4.09048 2.60173 3.8616 2.71718 3.4638C2.90412 2.81965 3.20598 2.22379 3.5996 1.69988Z\" fill=\"#0F172A\"></path> <path d=\"M13.5996 0.798866C13.3508 0.467705 12.8806 0.400945 12.5495 0.649752C12.2183 0.89856 12.1516 1.36872 12.4004 1.69988C12.794 2.22379 13.0958 2.81965 13.2828 3.4638C13.3982 3.8616 13.8143 4.09048 14.2121 3.97504C14.6099 3.85959 14.8388 3.44352 14.7233 3.04572C14.4848 2.22396 14.1001 1.465 13.5996 0.798866Z\" fill=\"#0F172A\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M8 1C10.2091 1 12 2.79086 12 5V7.37868C12 7.7765 12.158 8.15804 12.4393 8.43934L13.7071 9.70711C13.8946 9.89464 14 10.149 14 10.4142V11C14 11.5523 13.5523 12 13 12H11C11 13.6569 9.65685 15 8 15C6.34315 15 5 13.6569 5 12H3C2.44772 12 2 11.5523 2 11V10.4142C2 10.149 2.10536 9.89464 2.29289 9.70711L3.56066 8.43934C3.84197 8.15804 4 7.7765 4 7.37868V5C4 2.79086 5.79086 1 8 1ZM8 13.5C7.17157 13.5 6.5 12.8284 6.5 12H9.5C9.5 12.8284 8.82843 13.5 8 13.5Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}