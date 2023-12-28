// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ChatBubbleBottomCenterIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-4 h-4\"><path d=\"M1 8.74053C1 9.72321 1.71341 10.5653 2.68906 10.6827C3.5937 10.7915 4.50678 10.8729 5.42746 10.926C5.78969 10.9469 6.11506 11.1572 6.27733 11.4817L7.32918 13.5854C7.45622 13.8395 7.71592 14 8 14C8.28408 14 8.54378 13.8395 8.67082 13.5854L9.72265 11.4817C9.88492 11.1572 10.2103 10.9469 10.5725 10.9261C11.4932 10.873 12.4063 10.7916 13.3109 10.6828C14.2866 10.5654 15 9.72332 15 8.74062V4.25938C15 3.27668 14.2866 2.43458 13.3109 2.3172C11.57 2.10777 9.79777 2 8.00039 2C6.20273 2 4.43025 2.10781 2.68906 2.3173C1.71341 2.43469 1 3.27679 1 4.25947V8.74053Z\" fill=\"black\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
