// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func IdentificationIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"24\" height=\"24\" viewBox=\"0 0 24 24\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-6 h-6\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M4.5 3.75C2.84315 3.75 1.5 5.09315 1.5 6.75V17.25C1.5 18.9069 2.84315 20.25 4.5 20.25H19.5C21.1569 20.25 22.5 18.9069 22.5 17.25V6.75C22.5 5.09315 21.1569 3.75 19.5 3.75H4.5ZM8.625 6.75C7.38236 6.75 6.375 7.75736 6.375 9C6.375 10.2426 7.38236 11.25 8.625 11.25C9.86764 11.25 10.875 10.2426 10.875 9C10.875 7.75736 9.86764 6.75 8.625 6.75ZM4.75191 15.4528C5.3309 13.8765 6.84542 12.75 8.62496 12.75C10.4045 12.75 11.919 13.8765 12.498 15.4528C12.6271 15.8043 12.4771 16.1972 12.1466 16.3733C11.0958 16.9331 9.89627 17.25 8.62496 17.25C7.35364 17.25 6.15413 16.9331 5.10331 16.3733C4.77278 16.1972 4.62279 15.8043 4.75191 15.4528ZM15 8.25C14.5858 8.25 14.25 8.58579 14.25 9C14.25 9.41421 14.5858 9.75 15 9.75H18.75C19.1642 9.75 19.5 9.41421 19.5 9C19.5 8.58579 19.1642 8.25 18.75 8.25H15ZM14.25 12C14.25 11.5858 14.5858 11.25 15 11.25H18.75C19.1642 11.25 19.5 11.5858 19.5 12C19.5 12.4142 19.1642 12.75 18.75 12.75H15C14.5858 12.75 14.25 12.4142 14.25 12ZM15 14.25C14.5858 14.25 14.25 14.5858 14.25 15C14.25 15.4142 14.5858 15.75 15 15.75H18.75C19.1642 15.75 19.5 15.4142 19.5 15C19.5 14.5858 19.1642 14.25 18.75 14.25H15Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
