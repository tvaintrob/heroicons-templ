// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func InboxStackIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2.74193 2.75518C3.16891 2.27483 3.78091 2 4.4236 2H11.5764C12.2191 2 12.8311 2.27483 13.2581 2.75518L14.4317 4.07548C14.7978 4.48735 15 5.01924 15 5.5703V6.25C15 7.49264 13.9926 8.5 12.75 8.5H3.25C2.00736 8.5 1 7.49264 1 6.25V5.5703C1 5.01924 1.20223 4.48735 1.56833 4.07548L2.74193 2.75518ZM4.4236 3.5C4.20937 3.5 4.00537 3.59161 3.86304 3.75173L2.75347 5H4.96482C5.29917 5 5.6114 5.1671 5.79687 5.4453L6.20313 6.0547C6.3886 6.3329 6.70083 6.5 7.03518 6.5H8.96482C9.29917 6.5 9.6114 6.3329 9.79687 6.0547L10.2031 5.4453C10.3886 5.1671 10.7008 5 11.0352 5H13.2465L12.137 3.75173C11.9946 3.59161 11.7906 3.5 11.5764 3.5H4.4236Z\" fill=\"black\"></path> <path d=\"M1 10.75C1 10.3358 1.33579 10 1.75 10H4.96482C5.29917 10 5.6114 10.1671 5.79687 10.4453L6.20313 11.0547C6.3886 11.3329 6.70083 11.5 7.03518 11.5H8.96482C9.29917 11.5 9.6114 11.3329 9.79687 11.0547L10.2031 10.4453C10.3886 10.1671 10.7008 10 11.0352 10H14.25C14.6642 10 15 10.3358 15 10.75V11.75C15 12.9926 13.9926 14 12.75 14H3.25C2.00736 14 1 12.9926 1 11.75V10.75Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
