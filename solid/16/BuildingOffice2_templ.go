// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func BuildingOffice2Icon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M1.75 2C1.33579 2 1 2.33579 1 2.75C1 3.16421 1.33579 3.5 1.75 3.5H2V12.5H1.75C1.33579 12.5 1 12.8358 1 13.25C1 13.6642 1.33579 14 1.75 14H3.25C3.66421 14 4 13.6642 4 13.25V11.75C4 11.3358 4.33579 11 4.75 11H6.25C6.66421 11 7 11.3358 7 11.75V13.25C7 13.6642 7.33579 14 7.75 14H8.25C8.66421 14 9 13.6642 9 13.25V3.5H9.25C9.66421 3.5 10 3.16421 10 2.75C10 2.33579 9.66421 2 9.25 2H1.75ZM3.5 5.5C3.5 5.22386 3.72386 5 4 5H4.5C4.77614 5 5 5.22386 5 5.5V6C5 6.27614 4.77614 6.5 4.5 6.5H4C3.72386 6.5 3.5 6.27614 3.5 6V5.5ZM4 7.5C3.72386 7.5 3.5 7.72386 3.5 8V8.5C3.5 8.77614 3.72386 9 4 9H4.5C4.77614 9 5 8.77614 5 8.5V8C5 7.72386 4.77614 7.5 4.5 7.5H4ZM6 5.5C6 5.22386 6.22386 5 6.5 5H7C7.27614 5 7.5 5.22386 7.5 5.5V6C7.5 6.27614 7.27614 6.5 7 6.5H6.5C6.22386 6.5 6 6.27614 6 6V5.5ZM6.5 7.5C6.22386 7.5 6 7.72386 6 8V8.5C6 8.77614 6.22386 9 6.5 9H7C7.27614 9 7.5 8.77614 7.5 8.5V8C7.5 7.72386 7.27614 7.5 7 7.5H6.5Z\" fill=\"#0F172A\"></path> <path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M11.5 6C10.9477 6 10.5 6.44772 10.5 7V13C10.5 13.5523 10.9477 14 11.5 14H14.25C14.6642 14 15 13.6642 15 13.25C15 12.8358 14.6642 12.5 14.25 12.5H14V7.5H14.25C14.6642 7.5 15 7.16421 15 6.75C15 6.33579 14.6642 6 14.25 6H11.5ZM12 7.5H12.5C12.7761 7.5 13 7.72386 13 8V8.5C13 8.77614 12.7761 9 12.5 9H12C11.7239 9 11.5 8.77614 11.5 8.5V8C11.5 7.72386 11.7239 7.5 12 7.5ZM12 10C11.7239 10 11.5 10.2239 11.5 10.5V11C11.5 11.2761 11.7239 11.5 12 11.5H12.5C12.7761 11.5 13 11.2761 13 11V10.5C13 10.2239 12.7761 10 12.5 10H12Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}