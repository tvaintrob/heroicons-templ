// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CheckCircleIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M8 15C11.866 15 15 11.866 15 8C15 4.13401 11.866 1 8 1C4.13401 1 1 4.13401 1 8C1 11.866 4.13401 15 8 15ZM11.8435 6.20859C12.0967 5.88082 12.0363 5.40981 11.7086 5.15654C11.3808 4.90327 10.9098 4.96365 10.6565 5.29141L6.95615 10.0801L5.30747 8.24828C5.03038 7.94039 4.55616 7.91543 4.24828 8.19253C3.94039 8.46962 3.91544 8.94384 4.19253 9.25172L6.44253 11.7517C6.59132 11.917 6.80582 12.0078 7.02809 11.9995C7.25036 11.9911 7.45746 11.8846 7.59346 11.7086L11.8435 6.20859Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
