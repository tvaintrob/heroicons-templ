// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BellIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\" class=\"w-5 h-5\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M9.99997 2C6.68626 2 3.99997 4.68629 3.99997 8C3.99997 9.88663 3.54624 11.6651 2.7426 13.2343C2.63591 13.4426 2.6326 13.6888 2.73365 13.9C2.83469 14.1111 3.02851 14.2629 3.25769 14.3105C4.32537 14.5322 5.41181 14.7023 6.51426 14.818C6.67494 16.602 8.17421 18 10 18C11.8258 18 13.3251 16.602 13.4857 14.818C14.5882 14.7023 15.6746 14.5322 16.7422 14.3105C16.9714 14.2629 17.1652 14.1111 17.2663 13.9C17.3673 13.6888 17.364 13.4426 17.2573 13.2343C16.4537 11.6651 16 9.88663 16 8C16 4.68629 13.3137 2 9.99997 2ZM8.0493 14.9433C8.69477 14.9809 9.34517 15 9.99997 15C10.6548 15 11.3052 14.9809 11.9507 14.9433C11.749 15.8345 10.9522 16.5 10 16.5C9.04777 16.5 8.25097 15.8345 8.0493 14.9433Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
