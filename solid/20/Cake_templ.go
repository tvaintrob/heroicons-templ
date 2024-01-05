// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.513
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func CakeIcon(attrs templ.Attributes) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg xmlns=\"http://www.w3.org/2000/svg\" fill=\"currentColor\" viewBox=\"0 0 20 20\" class=\"w-5 h-5\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.RenderAttributes(ctx, templ_7745c5c3_Buffer, attrs)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("><path d=\"M6.74999 0.97937L5.86612 1.86326C5.37796 2.35142 5.37796 3.14288 5.86612 3.63103C6.35427 4.11919 7.14573 4.11919 7.63388 3.63103C8.12204 3.14288 8.12204 2.35142 7.63388 1.86326L6.74999 0.97937Z\"></path><path d=\"M13.25 0.97937L12.3661 1.86326C11.878 2.35142 11.878 3.14288 12.3661 3.63103C12.8543 4.11919 13.6457 4.11919 14.1339 3.63103C14.622 3.14288 14.622 2.35142 14.1339 1.86326L13.25 0.97937Z\"></path><path d=\"M9.99999 0.97937L10.8839 1.86326C11.372 2.35142 11.372 3.14288 10.8839 3.63103C10.3957 4.11919 9.60427 4.11919 9.11612 3.63103C8.62796 3.14288 8.62796 2.35142 9.11612 1.86326L9.99999 0.97937Z\"></path><path d=\"M7.5 5.75003C7.5 5.33582 7.16421 5.00003 6.75 5.00003C6.33579 5.00003 6 5.33582 6 5.75003V6.21382C4.82122 6.5185 4 7.60435 4 8.83573V8.93025C4.10036 8.90963 4.20229 8.89232 4.30573 8.87847C6.16935 8.62879 8.07023 8.50003 10 8.50003C11.9298 8.50003 13.8306 8.62879 15.6943 8.87847C15.7977 8.89232 15.8996 8.90963 16 8.93025V8.83573C16 7.60435 15.1788 6.5185 14 6.21382V5.75003C14 5.33582 13.6642 5.00003 13.25 5.00003C12.8358 5.00003 12.5 5.33582 12.5 5.75003V6.06832C11.9195 6.03653 11.3361 6.01574 10.75 6.00616V5.75003C10.75 5.33582 10.4142 5.00003 10 5.00003C9.58579 5.00003 9.25 5.33582 9.25 5.75003V6.00616C8.66391 6.01574 8.08051 6.03653 7.5 6.06832V5.75003Z\"></path><path d=\"M4.50491 10.3652C6.30269 10.1243 8.13702 10 10 10C11.863 10 13.6973 10.1243 15.4951 10.3652C16.9666 10.5623 18 11.8379 18 13.2789V13.9722C17.4297 13.9722 16.8594 13.8412 16.3354 13.5792C14.8652 12.8441 13.1348 12.8441 11.6646 13.5792C10.6167 14.1032 9.38329 14.1032 8.33541 13.5792C6.86524 12.8441 5.13476 12.8441 3.66459 13.5792C3.14065 13.8412 2.57032 13.9722 2 13.9722V13.2789C2 11.8379 3.03337 10.5623 4.50491 10.3652Z\"></path><path d=\"M15.6646 14.9209C16.3997 15.2884 17.1998 15.4722 18 15.4722V16.5C18 17.3285 17.3284 18 16.5 18H3.5C2.67157 18 2 17.3285 2 16.5V15.4722C2.80016 15.4722 3.60032 15.2884 4.33541 14.9209C5.38329 14.3969 6.61671 14.3969 7.66459 14.9209C9.13476 15.6559 10.8652 15.6559 12.3354 14.9209C13.3833 14.3969 14.6167 14.3969 15.6646 14.9209Z\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
