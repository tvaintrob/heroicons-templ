// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ChatBubbleLeftIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M1 8.74067C1 9.72338 1.71344 10.5655 2.68911 10.6828C3.45355 10.7748 4.22402 10.8472 5 10.8994V13.25C5 13.5533 5.18273 13.8268 5.46299 13.9429C5.74324 14.059 6.06583 13.9948 6.28033 13.7803L8.7905 11.2702C8.97217 11.0885 9.21682 10.9842 9.47361 10.9758C10.7676 10.9332 12.0475 10.8347 13.311 10.6826C14.2866 10.5653 15 9.72316 15 8.74048V4.25947C15 3.27678 14.2866 2.43469 13.3109 2.3173C11.5698 2.1078 9.79728 2 7.99962 2C6.20224 2 4.43002 2.10777 2.68909 2.31721C1.71343 2.43458 1 3.27668 1 4.25938V8.74067Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}