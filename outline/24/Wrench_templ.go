// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package outline24

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func WrenchIcon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 24 24\" class=\"w-6 h-6\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path d=\"M21.75 6.75C21.75 9.23528 19.7353 11.25 17.25 11.25C17.1206 11.25 16.9925 11.2445 16.8659 11.2338C15.7904 11.1429 14.6016 11.3052 13.9155 12.1383L6.76432 20.8219C6.28037 21.4096 5.55897 21.75 4.79769 21.75C3.39064 21.75 2.25 20.6094 2.25 19.2023C2.25 18.441 2.59044 17.7196 3.1781 17.2357L11.8617 10.0845C12.6948 9.39841 12.8571 8.20956 12.7662 7.13411C12.7555 7.00749 12.75 6.87938 12.75 6.75C12.75 4.26472 14.7647 2.25 17.25 2.25C17.9103 2.25 18.5375 2.39223 19.1024 2.64774L15.8262 5.92397C16.0823 7.03963 16.9605 7.91785 18.0762 8.17397L21.3524 4.89779C21.6078 5.46268 21.75 6.08973 21.75 6.75Z\" stroke=\"#0F172A\" stroke-width=\"1.5\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path><path d=\"M4.86723 19.125H4.87473V19.1325H4.86723V19.125Z\" stroke=\"#0F172A\" stroke-width=\"1.5\" stroke-linecap=\"round\" stroke-linejoin=\"round\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
