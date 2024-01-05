// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func InboxIcon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 20 20\" class=\"w-5 h-5\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M1 11.2708C1 11.0237 1.03331 10.7777 1.09902 10.5395L2.622 5.0187C2.95099 3.82611 4.03584 3 5.27298 3H14.727C15.9642 3 17.049 3.8261 17.378 5.0187L18.901 10.5395C18.9667 10.7777 19 11.0237 19 11.2708V15C19 16.1046 18.1046 17 17 17H3C1.89543 17 1 16.1046 1 15V11.2708ZM4.06799 5.41759C4.21753 4.8755 4.71065 4.5 5.27298 4.5H14.727C15.2894 4.5 15.7825 4.8755 15.932 5.41759L17.455 10.9384C17.4606 10.9588 17.4657 10.9794 17.4703 11H14C13.6471 11 13.3203 11.186 13.1401 11.4895L12.5339 12.5105C12.3537 12.814 12.0269 13 11.674 13H8.23607C7.8573 13 7.51103 12.786 7.34164 12.4472L6.89443 11.5528C6.72504 11.214 6.37877 11 6 11H2.52969C2.53427 10.9794 2.53938 10.9588 2.54501 10.9384L4.06799 5.41759Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
