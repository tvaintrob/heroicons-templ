// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func WalletIcon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 16 16\" class=\"w-4 h-4\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path d=\"M2 3.5C2 2.67157 2.67157 2 3.5 2H12.5C13.3284 2 14 2.67157 14 3.5V3.90135C13.5587 3.64609 13.0464 3.5 12.5 3.5H3.5C2.95357 3.5 2.44126 3.64609 2 3.90135V3.5Z\"></path><path d=\"M3.5 5C2.67157 5 2 5.67157 2 6.5V6.90135C2.44126 6.64609 2.95357 6.5 3.5 6.5H12.5C13.0464 6.5 13.5587 6.64609 14 6.90135V6.5C14 5.67157 13.3284 5 12.5 5H3.5Z\"></path><path d=\"M8 10C8.93374 10 9.71803 9.36012 9.93829 8.49494C10.0064 8.22733 10.2239 8 10.5 8H12.5C13.3284 8 14 8.67157 14 9.5V12.5C14 13.3284 13.3284 14 12.5 14H3.5C2.67157 14 2 13.3284 2 12.5V9.5C2 8.67157 2.67157 8 3.5 8H5.5C5.77614 8 5.99359 8.22733 6.06171 8.49494C6.28197 9.36012 7.06626 10 8 10Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
