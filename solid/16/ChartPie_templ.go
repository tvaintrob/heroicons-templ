// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ChartPieIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"currentColor\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path d=\"M13.9754 6.50077C14.0026 6.77556 13.7761 7.00007 13.5 7.00007H9.5C9.22386 7.00007 9 6.77621 9 6.50007V2.50007C9 2.22393 9.2245 1.99744 9.4993 2.02469C11.8621 2.25894 13.7411 4.138 13.9754 6.50077Z\"></path> <path d=\"M6.5007 4.02469C6.7755 3.99744 7 4.22393 7 4.50007V8.50007C7 8.77621 7.22386 9.00007 7.5 9.00007H11.5C11.7761 9.00007 12.0026 9.22457 11.9754 9.49937C11.7248 12.0264 9.59291 14.0001 7 14.0001C4.23858 14.0001 2 11.7615 2 9.00007C2 6.40716 3.9737 4.27523 6.5007 4.02469Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
