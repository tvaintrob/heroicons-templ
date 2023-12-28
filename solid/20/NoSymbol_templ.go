// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func NoSymbolIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M5.96461 4.90395L15.096 14.0354C15.9755 12.9265 16.5 11.525 16.5 10C16.5 6.41015 13.5899 3.5 10 3.5C8.475 3.5 7.07353 4.02446 5.96461 4.90395ZM14.0354 15.096L4.90395 5.96461C4.02446 7.07353 3.5 8.475 3.5 10C3.5 13.5899 6.41015 16.5 10 16.5C11.525 16.5 12.9265 15.9755 14.0354 15.096ZM4.34315 4.34315C5.79004 2.89625 7.79107 2 10 2C14.4183 2 18 5.58172 18 10C18 12.2089 17.1037 14.21 15.6569 15.6569C14.21 17.1037 12.2089 18 10 18C5.58172 18 2 14.4183 2 10C2 7.79107 2.89625 5.79004 4.34315 4.34315Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
