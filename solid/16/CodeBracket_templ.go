// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CodeBracketIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M4.78033 4.96977C5.07322 5.26266 5.07322 5.73753 4.78033 6.03043L2.81066 8.0001L4.78033 9.96977C5.07322 10.2627 5.07322 10.7375 4.78033 11.0304C4.48744 11.3233 4.01256 11.3233 3.71967 11.0304L1.21967 8.53043C0.926777 8.23753 0.926777 7.76266 1.21967 7.46977L3.71967 4.96977C4.01256 4.67687 4.48744 4.67687 4.78033 4.96977Z\" fill=\"black\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M11.2197 4.96977C10.9268 5.26266 10.9268 5.73753 11.2197 6.03043L13.1893 8.0001L11.2197 9.96977C10.9268 10.2627 10.9268 10.7375 11.2197 11.0304C11.5126 11.3233 11.9874 11.3233 12.2803 11.0304L14.7803 8.53043C15.0732 8.23753 15.0732 7.76266 14.7803 7.46977L12.2803 4.96977C11.9874 4.67687 11.5126 4.67687 11.2197 4.96977Z\" fill=\"black\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M8.85607 2.00763C9.26612 2.06621 9.55104 2.44611 9.49246 2.85616L7.99246 13.3562C7.93388 13.7662 7.55398 14.0511 7.14393 13.9926C6.73388 13.934 6.44896 13.5541 6.50754 13.144L8.00754 2.64403C8.06612 2.23398 8.44602 1.94905 8.85607 2.00763Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}