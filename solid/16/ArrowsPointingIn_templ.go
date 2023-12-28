// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid16

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func ArrowsPointingInIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"16\" height=\"16\" viewBox=\"0 0 16 16\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path fill-rule=\"evenodd\" clip-rule=\"evenodd\" d=\"M2.21967 2.21967C2.51256 1.92678 2.98744 1.92678 3.28033 2.21967L5.5 4.43934V2.75C5.5 2.33579 5.83579 2 6.25 2C6.66421 2 7 2.33579 7 2.75V6.25C7 6.66421 6.66421 7 6.25 7H2.75C2.33579 7 2 6.66421 2 6.25C2 5.83579 2.33579 5.5 2.75 5.5H4.43934L2.21967 3.28033C1.92678 2.98744 1.92678 2.51256 2.21967 2.21967ZM12.7197 2.21967C13.0126 1.92678 13.4874 1.92678 13.7803 2.21967C14.0732 2.51256 14.0732 2.98744 13.7803 3.28033L11.5607 5.5H13.25C13.6642 5.5 14 5.83579 14 6.25C14 6.66421 13.6642 7 13.25 7H9.75C9.33579 7 9 6.66421 9 6.25V2.75C9 2.33579 9.33579 2 9.75 2C10.1642 2 10.5 2.33579 10.5 2.75V4.43934L12.7197 2.21967ZM2.75 9H6.25C6.66421 9 7 9.33579 7 9.75V13.25C7 13.6642 6.66421 14 6.25 14C5.83579 14 5.5 13.6642 5.5 13.25V11.5607L3.28033 13.7803C2.98744 14.0732 2.51256 14.0732 2.21967 13.7803C1.92678 13.4874 1.92678 13.0126 2.21967 12.7197L4.43934 10.5H2.75C2.33579 10.5 2 10.1642 2 9.75C2 9.33579 2.33579 9 2.75 9ZM9 9.75C9 9.33579 9.33579 9 9.75 9H13.25C13.6642 9 14 9.33579 14 9.75C14 10.1642 13.6642 10.5 13.25 10.5H11.5607L13.7803 12.7197C14.0732 13.0126 14.0732 13.4874 13.7803 13.7803C13.4874 14.0732 13.0126 14.0732 12.7197 13.7803L10.5 11.5607V13.25C10.5 13.6642 10.1642 14 9.75 14C9.33579 14 9 13.6642 9 13.25V9.75Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
