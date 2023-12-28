// Code generated by templ - DO NOT EDIT.

// templ: version: 0.2.476
package solid20

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func SpeakerXMarkIcon() templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<svg width=\"20\" height=\"20\" viewBox=\"0 0 20 20\" fill=\"none\" xmlns=\"http://www.w3.org/2000/svg\"><path d=\"M10.0475 3.06153C10.3221 3.18021 10.5 3.4508 10.5 3.75002V16.25C10.5 16.5492 10.3221 16.8198 10.0475 16.9385C9.7728 17.0572 9.45387 17.0012 9.23598 16.7962L5.20257 13H2.66724C2.35725 13 2.0792 12.8093 1.96756 12.5201C1.66534 11.7372 1.5 10.887 1.5 10C1.5 9.11302 1.66534 8.26287 1.96756 7.47993C2.0792 7.19074 2.35725 7.00002 2.66724 7.00002H5.20257L9.23598 3.20387C9.45387 2.99879 9.7728 2.94286 10.0475 3.06153Z\" fill=\"#0F172A\"></path> <path d=\"M13.7803 7.2197C13.4874 6.92681 13.0126 6.92681 12.7197 7.2197C12.4268 7.51259 12.4268 7.98747 12.7197 8.28036L14.4393 10L12.7197 11.7197C12.4268 12.0126 12.4268 12.4875 12.7197 12.7804C13.0126 13.0733 13.4874 13.0733 13.7803 12.7804L15.5 11.0607L17.2197 12.7804C17.5126 13.0733 17.9874 13.0733 18.2803 12.7804C18.5732 12.4875 18.5732 12.0126 18.2803 11.7197L16.5607 10L18.2803 8.28036C18.5732 7.98747 18.5732 7.51259 18.2803 7.2197C17.9874 6.92681 17.5126 6.92681 17.2197 7.2197L15.5 8.93937L13.7803 7.2197Z\" fill=\"#0F172A\"></path></svg>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
