// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func FolderArrowDownIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2 4.75C2 3.7835 2.7835 3 3.75 3H8.58579C9.04992 3 9.49503 3.18437 9.82322 3.51256L11.2374 4.92678C11.2843 4.97366 11.3479 5 11.4142 5H16.25C17.2165 5 18 5.7835 18 6.75V15.25C18 16.2165 17.2165 17 16.25 17H3.75C2.7835 17 2 16.2165 2 15.25V4.75ZM10.75 8.75C10.75 8.33579 10.4142 8 10 8C9.58579 8 9.25 8.33579 9.25 8.75V11.2955L8.30747 10.2483C8.03038 9.94039 7.55616 9.91543 7.24828 10.1925C6.94039 10.4696 6.91544 10.9438 7.19253 11.2517L9.44253 13.7517C9.58476 13.9098 9.78738 14 10 14C10.2126 14 10.4152 13.9098 10.5575 13.7517L12.8075 11.2517C13.0846 10.9438 13.0596 10.4696 12.7517 10.1925C12.4438 9.91543 11.9696 9.94039 11.6925 10.2483L10.75 11.2955V8.75Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
